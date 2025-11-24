package halls

type Service interface{
	GetAllHalls() ([]Hall, error)
    GetHallByID(id uint) (*Hall, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service{
	return &service{repo: repo}
}

func (s *service) GetAllHalls() ([]Hall, error){
	return s.repo.GetAll()
}

func (s *service) GetHallByID(id uint) (*Hall, error) {
    return s.repo.GetByID(id)
}