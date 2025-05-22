package dashboard

type UseCase interface {
	GetDashboardData() (*DashboardData, error)
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo: repo}
}

func (u *usecase) GetDashboardData() (*DashboardData, error) {
	return u.repo.FetchDashboardData()
}
