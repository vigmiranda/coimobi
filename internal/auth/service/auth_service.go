package service

func Authenticate(email, password string) bool {
	// Simples, só pra exemplo
	return email == "admin@coimobi.com" && password == "123456"
}
