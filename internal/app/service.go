package app

type Service struct {
	repository *Repository
}

func InitService(repository *Repository) *Service {
	return &Service{repository}
}
