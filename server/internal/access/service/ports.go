package service

import (
	"context"

	accessdb "github.com/findardi/Wadi/server/internal/access/repository/sqlc"
)

type AccessRepository interface {
	AddMember(ctx context.Context, arg accessdb.AddMemberParams) (accessdb.WorkspaceMember, error)
	InsertRole(ctx context.Context, arg accessdb.InsertRoleParams) (accessdb.WorkspaceRole, error)

	ExecTx(ctx context.Context, fn func(q *accessdb.Queries) error) error
}

type MailService interface {
	Send(ctx context.Context, to, subject, body string) error
}
