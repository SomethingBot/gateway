package api

type Gateway struct {
	Address string
	Service Service
}

func NewGateway(address string, service Service) (gateway *Gateway) {
	gateway = &Gateway{
		Address: address,
		Service: service,
	}
	return
}

func (gateway *Gateway) Dial() (err error) {
	return nil
}

func (gateway *Gateway) run() {

}

func (gateway *Gateway) RegisterService(service Service) (err error) {
	return nil
}
