package seats

type Service interface{
	GetSeatsByHall(hallID uint) ([]Seat, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service{
	return &service{repo: repo}
}

func (s *service) GetSeatsByHall(hallID uint) ([]Seat, error){
	return s.repo.GetByHallID(hallID)
}