package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

var ErrInvalidToken = errors.New("token invalido")

type Claims struct {
	Subject  string `json:"sub"`
	TenantID string `json:"tenant_id,omitempty"`
	Type     string `json:"typ"`
	IssuedAt int64  `json:"iat"`
	Expires  int64  `json:"exp"`
}

type TokenManager struct {
	secret []byte
}

func NewTokenManager(secret string) *TokenManager {
	return &TokenManager{secret: []byte(secret)}
}

func (m *TokenManager) AccessToken(userID uuid.UUID, tenantID *uuid.UUID, ttl time.Duration) (string, error) {
	tenant := ""
	if tenantID != nil {
		tenant = tenantID.String()
	}
	return m.sign(Claims{
		Subject:  userID.String(),
		TenantID: tenant,
		Type:     "access",
		IssuedAt: time.Now().Unix(),
		Expires:  time.Now().Add(ttl).Unix(),
	})
}

func (m *TokenManager) VerifyAccess(token string) (Claims, error) {
	claims, err := m.verify(token)
	if err != nil {
		return Claims{}, err
	}
	if claims.Type != "access" || claims.Expires < time.Now().Unix() {
		return Claims{}, ErrInvalidToken
	}
	return claims, nil
}

func NewRefreshToken() (plain string, hash string, err error) {
	var bytes [32]byte
	if _, err := rand.Read(bytes[:]); err != nil {
		return "", "", err
	}
	plain = base64.RawURLEncoding.EncodeToString(bytes[:])
	sum := sha256.Sum256([]byte(plain))
	return plain, hex.EncodeToString(sum[:]), nil
}

func HashRefreshToken(token string) string {
	sum := sha256.Sum256([]byte(token))
	return hex.EncodeToString(sum[:])
}

func (m *TokenManager) sign(claims Claims) (string, error) {
	header := map[string]string{"alg": "HS256", "typ": "JWT"}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	unsigned := base64.RawURLEncoding.EncodeToString(headerJSON) + "." + base64.RawURLEncoding.EncodeToString(claimsJSON)
	signature := m.signature(unsigned)
	return unsigned + "." + signature, nil
}

func (m *TokenManager) verify(token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return Claims{}, ErrInvalidToken
	}

	unsigned := parts[0] + "." + parts[1]
	expected := m.signature(unsigned)
	if subtleCompare(expected, parts[2]) == false {
		return Claims{}, ErrInvalidToken
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return Claims{}, ErrInvalidToken
	}
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return Claims{}, ErrInvalidToken
	}
	return claims, nil
}

func (m *TokenManager) signature(unsigned string) string {
	mac := hmac.New(sha256.New, m.secret)
	mac.Write([]byte(unsigned))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func subtleCompare(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := range a {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
