package service

//LoginInformation is ...
type LoginInformation struct {
	id       int
	Username string
	password string
}

//LoginService is ...
type LoginService interface {
	LoginUser(Username string, password string) bool
}

func loginService() LoginService {
	return &LoginInformation{

		Username: "dinesh",
		password: "nopasword",
	}
}

//LoginUser is ...
func (L *LoginInformation) LoginUser(Username string, password string) bool {
	return L.Username == Username && L.password == password
}
