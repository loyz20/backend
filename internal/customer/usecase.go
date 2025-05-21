package customer

type UseCase interface {
	GetAll() ([]Customer, error)
	GetByID(id uint) (*Customer, error)
	Create(c *Customer) error
	Update(id uint, c *Customer) error
	Delete(id uint) error
}

type usecase struct {
	repo Repository
}

func NewUseCase(repo Repository) UseCase {
	return &usecase{repo}
}

func (u *usecase) GetAll() ([]Customer, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetByID(id uint) (*Customer, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) Create(c *Customer) error {
	return u.repo.Create(c)
}

func (u *usecase) Update(id uint, c *Customer) error {
	return u.repo.Update(id, c)
}

func (u *usecase) Delete(id uint) error {
	return u.repo.Delete(id)
}
