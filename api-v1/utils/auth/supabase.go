package auth

import (
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/golang-jwt/jwt/v5"
)

var (
	supabaseSecret []byte
	supabaseIssuer string
	jwksMutex      sync.Mutex
	jwksByURL      = map[string]*keyfunc.JWKS{}
)

func SetSupabaseSecret(secret string) {
	supabaseSecret = []byte(strings.TrimSpace(secret))
}

func SetSupabaseIssuer(issuer string) {
	supabaseIssuer = issuerNormalize(issuer)
}

type SupabaseClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func SupabaseTokenVerify(tokenString string) (email, subject string, err error) {
	claims := &SupabaseClaims{}
	parser := jwt.NewParser()

	tokenUnverified, _, parseErr := parser.ParseUnverified(tokenString, claims)
	if parseErr != nil {
		return "", "", parseErr
	}

	issuer := issuerNormalize(claims.Issuer)
	if supabaseIssuer != "" && issuer != supabaseIssuer {
		return "", "", errors.New("supabase token issuer mismatch")
	}

	algorithm, _ := tokenUnverified.Header["alg"].(string)
	if algorithm == "" {
		return "", "", errors.New("missing jwt algorithm header")
	}

	var tokenVerified *jwt.Token

	switch {
	case strings.HasPrefix(algorithm, "HS"):
		if len(supabaseSecret) == 0 {
			return "", "", errors.New("supabase secret is not set")
		}
		tokenVerified, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return supabaseSecret, nil
		}, jwt.WithValidMethods([]string{algorithm}))
	default:
		if issuer == "" {
			return "", "", errors.New("missing iss claim")
		}
		jwksURL := strings.TrimRight(issuer, "/") + "/.well-known/jwks.json"
		jwks, jwksFetchErr := jwksGet(jwksURL)
		if jwksFetchErr != nil {
			return "", "", jwksFetchErr
		}
		tokenVerified, err = jwt.ParseWithClaims(tokenString, claims, jwks.Keyfunc, jwt.WithValidMethods([]string{algorithm}))
	}

	if err != nil || tokenVerified == nil || !tokenVerified.Valid {
		return "", "", err
	}

	email = strings.ToLower(strings.TrimSpace(claims.Email))
	if email == "" {
		return "", "", errors.New("missing email claim")
	}

	subject = strings.TrimSpace(claims.Subject)

	return email, subject, nil
}

func jwksGet(jwksURL string) (*keyfunc.JWKS, error) {
	jwksMutex.Lock()
	cached := jwksByURL[jwksURL]
	jwksMutex.Unlock()
	if cached != nil {
		return cached, nil
	}
	jwks, fetchErr := keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval:   10 * time.Minute,
		RefreshRateLimit:  1 * time.Minute,
		RefreshTimeout:    5 * time.Second,
		RefreshUnknownKID: true,
	})
	if fetchErr != nil {
		return nil, fetchErr
	}
	jwksMutex.Lock()
	jwksByURL[jwksURL] = jwks
	jwksMutex.Unlock()
	return jwks, nil
}

func issuerNormalize(issuer string) string {
	issuer = strings.TrimSpace(issuer)
	if issuer == "" {
		return ""
	}
	return strings.TrimRight(issuer, "/")
}
