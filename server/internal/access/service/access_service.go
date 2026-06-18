package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/findardi/Wadi/server/internal/access/dto"
	accessdb "github.com/findardi/Wadi/server/internal/access/repository/sqlc"
	"github.com/findardi/Wadi/server/internal/platform/permission"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrRoleNameTaken = errors.New("role name already taken")
)

type AccessService struct {
	repo AccessRepository
	mail MailService
}

func NewAccessService(repo AccessRepository, mail MailService) *AccessService {
	return &AccessService{
		repo: repo,
		mail: mail,
	}
}

func uuidString(u pgtype.UUID) string {
	v, err := u.Value()
	if err != nil || v == nil {
		return ""
	}
	s, _ := v.(string)
	return s
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func isUniqueViolation(err error, constraint string) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" && pgErr.ConstraintName == constraint
	}
	return false
}

func (s *AccessService) InsertRole(ctx context.Context, req dto.CreateWorkspaceRoleRequest) (dto.WorkspaceRoleResponse, error) {
	var uid pgtype.UUID
	if err := uid.Scan(req.WorkspaceID); err != nil {
		return dto.WorkspaceRoleResponse{}, fmt.Errorf("parse owner id: %w", err)
	}

	role, err := s.repo.InsertRole(ctx, accessdb.InsertRoleParams{
		WorkspaceID: uid,
		Name:        req.Name,
		Permissions: req.Permission,
		IsSystem:    req.IsSystem,
	})

	if isUniqueViolation(err, "workspace_roles_name_key") {
		return dto.WorkspaceRoleResponse{}, ErrRoleNameTaken
	}
	if err != nil {
		return dto.WorkspaceRoleResponse{}, fmt.Errorf("insert role: %w", err)
	}

	return dto.WorkspaceRoleResponse{
		ID:          uuidString(role.ID),
		WorkspaceID: uuidString(role.WorkspaceID),
		Name:        role.Name,
		Permissions: role.Permissions,
		IsSystem:    role.IsSystem,
		CreatedAt:   role.CreatedAt.Time,
		UpdatedAt:   role.UpdatedAt.Time,
	}, nil
}

func (s *AccessService) SeedSystemRoles(ctx context.Context, tx pgx.Tx, workspaceID pgtype.UUID) error {
	q := accessdb.New(tx)

	for _, r := range permission.DefaultSystemRoles() {
		if _, err := q.InsertRole(ctx, accessdb.InsertRoleParams{
			WorkspaceID: workspaceID,
			Name:        r.Name,
			Permissions: r.Permissions,
			IsSystem:    true,
		}); err != nil {
			return fmt.Errorf("seed role %s: %w", r.Name, err)
		}
	}

	return nil
}
