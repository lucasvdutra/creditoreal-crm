package user

import (
	"context"
	"database/sql"
	"testing"

	db "creditoreal-crm/pkg/database/queries"

	"github.com/google/uuid"
)

type fakeUserStore struct {
	createdUser       db.CreateUserParams
	createdTenantUser db.CreateTenantUserParams
}

func (f *fakeUserStore) CreateUser(_ context.Context, arg db.CreateUserParams) (db.User, error) {
	f.createdUser = arg
	return db.User{ID: uuid.New(), Nome: arg.Nome, Email: arg.Lower, Status: arg.Status}, nil
}
func (f *fakeUserStore) GetUserByID(context.Context, uuid.UUID) (db.User, error) {
	return db.User{}, nil
}
func (f *fakeUserStore) ListUsers(context.Context, db.ListUsersParams) ([]db.User, error) {
	return nil, nil
}
func (f *fakeUserStore) CountUsers(context.Context, sql.NullString) (int64, error) {
	return 0, nil
}
func (f *fakeUserStore) CreateTenantUser(_ context.Context, arg db.CreateTenantUserParams) (db.TenantUser, error) {
	f.createdTenantUser = arg
	return db.TenantUser{ID: uuid.New(), TenantID: arg.TenantID, UserID: arg.UserID, Status: arg.Status}, nil
}
func (f *fakeUserStore) ListTenantUsers(context.Context, db.ListTenantUsersParams) ([]db.TenantUser, error) {
	return nil, nil
}
func (f *fakeUserStore) ListUserTenants(context.Context, uuid.UUID) ([]db.TenantUser, error) {
	return nil, nil
}

func TestCreateUserRequiresStrongEnoughPassword(t *testing.T) {
	t.Parallel()

	service := NewService(&fakeUserStore{})
	_, err := service.Create(context.Background(), CreateRequest{
		Nome:  "Lucas",
		Email: "lucas@example.com",
		Senha: "curta",
	})
	if err == nil {
		t.Fatal("expected validation error")
	}
}

func TestCreateTenantUserAllowsSameUserInTenant(t *testing.T) {
	t.Parallel()

	store := &fakeUserStore{}
	service := NewService(store)
	tenantID := uuid.New()
	userID := uuid.New()

	item, err := service.CreateTenantUser(context.Background(), CreateTenantUserRequest{
		TenantID: tenantID.String(),
		UserID:   userID.String(),
		Status:   "ativo",
	})
	if err != nil {
		t.Fatalf("create tenant user: %v", err)
	}
	if item.TenantID != tenantID.String() || item.UserID != userID.String() {
		t.Fatalf("unexpected tenant user response: %+v", item)
	}
}
