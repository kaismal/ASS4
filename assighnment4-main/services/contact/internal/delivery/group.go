package delivery

import (
	"architecture_go/services/contact/internal/repository/controller"
	"architecture_go/services/contact/internal/repository/storage/postgres"
	use_case "architecture_go/services/contact/internal/use-case"
)

func (r *registry) NewGroupController() controller.Group {
	g := use_case.NewGroupUseCase(
		postgres.NewGroupRepository(r.db),
		postgres.NewDBRepository(r.db),
	)
	return controller.NewGroupsController(g)
}
