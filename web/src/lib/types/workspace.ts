// Workspace feature contract — mirrors internal/workspace/dto on the Go backend.

// Lifecycle state — server-validated enum; new rooms default to `prepare`.
export type WorkspaceStatus = 'prepare' | 'active' | 'archive';

// owner_id is NOT sent: the backend resolves it from the JWT claims.
export interface CreateWorkspacePayload {
	name: string;
	description: string;
}

// Backend PUT /workspaces/:id — name required, description optional.
export interface UpdateWorkspacePayload {
	name: string;
	description: string;
}

export interface WorkspaceData {
	id: string;
	owner_id: string;
	name: string;
	slug: string;
	description: string;
	status: WorkspaceStatus;
	// ISO-8601 strings over the wire (Go time.Time marshals to RFC3339).
	created_at: string;
	updated_at: string;
}
