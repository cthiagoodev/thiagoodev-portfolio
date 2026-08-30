package usecases

type DeleteProjectUseCase interface {
	Execute(uuid string) error
}
