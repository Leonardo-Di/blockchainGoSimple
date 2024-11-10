package main

type User struct {
	Username string
	Password string
}

var users = map[string]string{
	"user1": "password1", // Example user
}

// Authenticate checks if the user credentials are valid
func Authenticate(username, password string) bool {
	if pass, ok := users[username]; ok {
		return pass == password
	}
	return false
}
