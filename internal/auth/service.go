package auth

import (
	"context"
	"database/sql"
	"time"

	"creditoreal-crm/internal/auth/password"
	"creditoreal-crm/internal/domain"
	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type Store interface {
	GetUserByEmail(context.Context, string) (db.User, error)
	GetUserByID(context.Context, uuid.UUID) (db.User, error)
	ListUserTenants(context.Context, uuid.UUID) ([]db.TenantUser, error)
	CreateRefreshToken(context.Context, db.CreateRefreshTokenParams) (db.RefreshToken, error)
	GetRefreshTokenByHash(context.Context, string) (db.RefreshToken, error)
	RevokeRefreshToken(context.Context, string) error
}

type Service struct {
	store           Store
	tokens          *TokenManager
	accessDuration  time.Duration
	refreshDuration time.Duration
}

type LoginRequest struct {
	Email    string  `json:"email"`
	Senha    string  `json:"senha"`
	TenantID *string `json:"tenant_id"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

func NewService(store Store, tokens *TokenManager, accessDuration, refreshDuration time.Duration) *Service {
	return &Service{store: store, tokens: tokens, accessDuration: accessDuration, refreshDuration: refreshDuration}
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (TokenResponse, error) {
	user, err := s.store.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return TokenResponse{}, domain.ErrInvalidCredentials
		}
		return TokenResponse{}, err
	}
	if user.Status != "ativo" || !password.Verify(req.Senha, user.SenhaHash) {
		return TokenResponse{}, domain.ErrInvalidCredentials
	}

	tenantID, err := s.resolveTenant(ctx, user.ID, req.TenantID)
	if err != nil {
		return TokenResponse{}, err
	}
	return s.issue(ctx, user.ID, tenantID)
}

func (s *Service) Refresh(ctx context.Context, req RefreshRequest) (TokenResponse, error) {
	hash := HashRefreshToken(req.RefreshToken)
	current, err := s.store.GetRefreshTokenByHash(ctx, hash)
	if err != nil {
		if err == sql.ErrNoRows {
			return TokenResponse{}, domain.ErrUnauthorized
		}
		return TokenResponse{}, err
	}
	if current.RevokedAt.Valid || current.ExpiresAt.Before(time.Now()) {
		return TokenResponse{}, domain.ErrUnauthorized
	}
	_ = s.store.RevokeRefreshToken(ctx, hash)

	tenantID, err := s.resolveTenant(ctx, current.UserID, nil)
	if err != nil {
		return TokenResponse{}, err
	}
	return s.issue(ctx, current.UserID, tenantID)
}

func (s *Service) issue(ctx context.Context, userID uuid.UUID, tenantID *uuid.UUID) (TokenResponse, error) {
	access, err := s.tokens.AccessToken(userID, tenantID, s.accessDuration)
	if err != nil {
		return TokenResponse{}, err
	}
	refresh, refreshHash, err := NewRefreshToken()
	if err != nil {
		return TokenResponse{}, err
	}
	if _, err := s.store.CreateRefreshToken(ctx, db.CreateRefreshTokenParams{
		UserID:    userID,
		TokenHash: refreshHash,
		ExpiresAt: time.Now().Add(s.refreshDuration),
	}); err != nil {
		return TokenResponse{}, err
	}

	return TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		TokenType:    "bearer",
		ExpiresIn:    int64(s.accessDuration.Seconds()),
	}, nil
}

func (s *Service) resolveTenant(ctx context.Context, userID uuid.UUID, raw *string) (*uuid.UUID, error) {
	links, err := s.store.ListUserTenants(ctx, userID)
	if err != nil {
		return nil, err
	}
	if raw == nil || *raw == "" {
		for _, link := range links {
			if link.Status == "ativo" {
				id := link.TenantID
				return &id, nil
			}
		}
		return nil, nil
	}

	wanted, err := uuid.Parse(*raw)
	if err != nil {
		return nil, domain.ErrInvalidInput
	}
	for _, link := range links {
		if link.TenantID == wanted && link.Status == "ativo" {
			id := link.TenantID
			return &id, nil
		}
	}
	return nil, domain.ErrUnauthorized
}
