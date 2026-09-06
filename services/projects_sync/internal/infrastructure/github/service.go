package github

type GithubService interface {
	FetchRepositories() ([]Project, error)
}

type GithubServiceImpl struct{}

func (s *GithubServiceImpl) FetchRepositories() ([]Project, error) {
	panic("implement me")
}
