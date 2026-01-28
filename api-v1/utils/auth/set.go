package auth

var secretKey []byte

func SetSecretJWT(secret string) {
	secretKey = []byte(secret)
}
