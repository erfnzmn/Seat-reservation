package waitinglist

type Service interface {
	Join(showID uint, userName, userPhone string) (*WaitingList, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Join(showID uint, userName, userPhone string) (*WaitingList, error) {
	w := &WaitingList{
		ShowID:    showID,
		SeatID:    0,
		UserName:  userName,
		UserPhone: userPhone,
		Status:    WaitingListStatusWaiting,
	}
	if err := s.repo.Add(w); err != nil {
		return nil, err
	}
	return w, nil
}
