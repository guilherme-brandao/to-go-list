package services

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "foobar",
		authorizedPassword: "123456",
	}
}

func (services *loginService) Login(username string, password string) bool {
	return services.authorizedUsername == username &&
		services.authorizedPassword == password
}
