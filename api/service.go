package api

type Service struct {
	Name     string    `json:"name"`
	Commands []Command `json:"commands"`
}

func NewService(name string, commands []Command) (service Service) {
	return Service{
		Name:     name,
		Commands: commands,
	}
}

func (service Service) ContainsCommand(command string) bool {
	for cmd := range service.Commands {
		if service.Commands[cmd].Name == command {
			return true
		}
	}
	return false
}
