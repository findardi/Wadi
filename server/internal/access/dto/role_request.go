package dto

type CreateWorkspaceRoleRequest struct {
	WorkspaceID string   `json:"-"`
	Permission  []string `json:"permissions"`
	Name        string   `json:"name" validate:"required"`
	IsSystem    bool     `json:"is_system"`
}
