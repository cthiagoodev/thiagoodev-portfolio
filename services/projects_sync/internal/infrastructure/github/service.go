package github

type Service interface {
	FetchRepositories() ([]Project, error)
}

type ServiceImpl struct{}

func (s *ServiceImpl) FetchRepositories() ([]Project, error) {
	panic("implement me")
}
