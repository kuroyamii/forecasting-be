package pingService

type pingService struct {
}

func (ps pingService) GetPing() string {
	return "ping"
}

func NewPingService() pingService {
	return pingService{}
}
