package rbac

import (
	"time"

	db "creditoreal-crm/pkg/database/queries"
)

type CreateRoleRequest struct {
	TenantID string `json:"tenant_id"`
	Nome     string `json:"nome"`
	Codigo   string `json:"codigo"`
	Status   string `json:"status"`
}

type AddPermissionRequest struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

type AssignRoleRequest struct {
	TenantID string `json:"tenant_id"`
	UserID   string `json:"user_id"`
	RoleID   string `json:"role_id"`
}

type RoleResponse struct {
	ID        string    `json:"id"`
	TenantID  string    `json:"tenant_id"`
	Nome      string    `json:"nome"`
	Codigo    string    `json:"codigo"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PermissionResponse struct {
	ID        string    `json:"id"`
	Codigo    string    `json:"codigo"`
	Descricao string    `json:"descricao"`
	CreatedAt time.Time `json:"created_at"`
}

func roleResponse(item db.Role) RoleResponse {
	return RoleResponse{
		ID:        item.ID.String(),
		TenantID:  item.TenantID.String(),
		Nome:      item.Nome,
		Codigo:    item.Codigo,
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}

func permissionResponse(item db.Permission) PermissionResponse {
	return PermissionResponse{
		ID:        item.ID.String(),
		Codigo:    item.Codigo,
		Descricao: item.Descricao,
		CreatedAt: item.CreatedAt,
	}
}
