package clients

type CronosGateway struct {
	baseURL string
}

func NewCronosGateway(baseURL string) *CronosGateway {
	return &CronosGateway{
		baseURL: baseURL,
	}
}
