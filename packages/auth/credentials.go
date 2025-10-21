package auth

func LoginWithCredentials(username string, password string) bool {
	// Dummy implementation for illustration purposes
	if username == "admin" && password == "password" {
		return true
	}
	return false

}
