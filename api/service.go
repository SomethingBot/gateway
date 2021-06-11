package api

type ServiceNode struct {
	Address string `json:"address"`
}

type Service struct {
	Name string `json:"name"`
	//Endpoints of format {"endpoints":{"192.168.88.4", "192.168.88.2}}
	Commands []Command `json:"commands"`
}

//func (service Service) AddNode(address string) error {
//
//}

func (service Service) ContainsCommand(command string) bool {
	for cmd := range service.Commands {
		if service.Commands[cmd].Name == command {
			return true
		}
	}
	return false
}
