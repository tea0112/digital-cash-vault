package permissions

type PermissionServiceImpl struct {
	repo PermissionRepository
}

func NewPermissionService(repo PermissionRepository) *PermissionServiceImpl {
	return &PermissionServiceImpl{
		repo: repo,
	}
}
