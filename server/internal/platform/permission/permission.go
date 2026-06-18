package permission

const (
	// Workspace
	PermWorkspaceView   = "workspace:view"
	PermWorkspaceEdit   = "workspace:edit"
	PermWorkspaceDelete = "workspace:delete"

	// Member
	PermMemberView   = "member:view"
	PermMemberAdd    = "member:add"
	PermMemberEdit   = "member:edit"
	PermMemberDelete = "member:delete"

	// Role
	PermRoleView   = "role:view"
	PermRoleCreate = "role:create"
	PermRoleEdit   = "role:edit"
	PermRoleDelete = "role:delete"
	PermRoleAssign = "role:assign"

	// Group
	PermGroupView   = "group:view"
	PermGroupCreate = "group:create"
	PermGroupEdit   = "group:edit"
	PermGroupDelete = "group:delete"
	PermGroupAssign = "group:assign"

	// Folder
	PermFolderView   = "folder:view"
	PermFolderCreate = "folder:create"
	PermFolderEdit   = "folder:edit"
	PermFolderDelete = "folder:delete"

	// Document
	PermDocumentView     = "document:view"
	PermDocumentUpload   = "document:upload"
	PermDocumentDownload = "document:download"
	PermDocumentEdit     = "document:edit"
	PermDocumentDelete   = "document:delete"
)

// All adalah daftar lengkap permission valid (untuk validasi mapping role).
var All = []string{
	PermWorkspaceView, PermWorkspaceEdit, PermWorkspaceDelete,
	PermMemberView, PermMemberAdd, PermMemberEdit, PermMemberDelete,
	PermRoleView, PermRoleCreate, PermRoleEdit, PermRoleDelete, PermRoleAssign,
	PermGroupView, PermGroupCreate, PermGroupEdit, PermGroupDelete, PermGroupAssign,
	PermFolderView, PermFolderCreate, PermFolderEdit, PermFolderDelete,
	PermDocumentView, PermDocumentUpload, PermDocumentDownload, PermDocumentEdit, PermDocumentDelete,
}

var set = func() map[string]struct{} {
	m := make(map[string]struct{}, len(All))
	for _, p := range All {
		m[p] = struct{}{}
	}
	return m
}()

// IsValid memeriksa apakah sebuah permission dikenali aplikasi.
func IsValid(p string) bool {
	_, ok := set[p]
	return ok
}

// Nama system role default (di-seed saat workspace dibuat).
const (
	RoleOwner       = "owner"
	RoleAdmin       = "admin"
	RoleContributor = "contributor"
	RoleViewer      = "viewer"
	RoleGuest       = "guest"
)

// SystemRole adalah cetakan role bawaan beserta permission-nya.
type SystemRole struct {
	Name        string
	Permissions []string
}

// GetOwner: kontrol penuh atas seluruh resource.
func GetOwner() []string {
	return append([]string{}, All...)
}

// GetAdmin: kelola member/role/group/konten, tapi tidak bisa hapus workspace.
func GetAdmin() []string {
	return []string{
		PermWorkspaceView, PermWorkspaceEdit,
		PermMemberView, PermMemberAdd, PermMemberEdit, PermMemberDelete,
		PermRoleView, PermRoleCreate, PermRoleEdit, PermRoleDelete, PermRoleAssign,
		PermGroupView, PermGroupCreate, PermGroupEdit, PermGroupDelete, PermGroupAssign,
		PermFolderView, PermFolderCreate, PermFolderEdit, PermFolderDelete,
		PermDocumentView, PermDocumentUpload, PermDocumentDownload, PermDocumentEdit, PermDocumentDelete,
	}
}

// GetContributor: kelola folder & dokumen, tanpa akses kelola member/role/group.
func GetContributor() []string {
	return []string{
		PermWorkspaceView,
		PermMemberView,
		PermGroupView,
		PermFolderView, PermFolderCreate, PermFolderEdit, PermFolderDelete,
		PermDocumentView, PermDocumentUpload, PermDocumentDownload, PermDocumentEdit, PermDocumentDelete,
	}
}

// GetViewer: hanya baca dan boleh download dokumen.
func GetViewer() []string {
	return []string{
		PermWorkspaceView,
		PermFolderView,
		PermDocumentView, PermDocumentDownload,
	}
}

// GetGuest: hanya baca, tanpa download.
func GetGuest() []string {
	return []string{
		PermWorkspaceView,
		PermFolderView,
		PermDocumentView,
	}
}

// DefaultSystemRoles mengembalikan seluruh role bawaan untuk di-seed ke workspace baru.
func DefaultSystemRoles() []SystemRole {
	return []SystemRole{
		{Name: RoleOwner, Permissions: GetOwner()},
		{Name: RoleAdmin, Permissions: GetAdmin()},
		{Name: RoleContributor, Permissions: GetContributor()},
		{Name: RoleViewer, Permissions: GetViewer()},
		{Name: RoleGuest, Permissions: GetGuest()},
	}
}
