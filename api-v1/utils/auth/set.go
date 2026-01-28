package auth

var SecretKey []byte

func SetSecretJWT(secret string) {
	SecretKey = []byte(secret)
}
