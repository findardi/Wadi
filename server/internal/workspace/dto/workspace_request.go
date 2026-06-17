package dto

type WorkspaceCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	OwnerID     string `json:"-"`
}

type WorkspaceUpdateStatusRequest struct {
	ID     string `json:"-"`
	Status string `json:"status" validate:"required"`
}

type WorkspaceUpdateRequest struct {
	ID          string `json:"-"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
