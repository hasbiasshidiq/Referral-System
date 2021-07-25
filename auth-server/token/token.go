package token

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	auth_pb "github.com/hasbiasshidiq/auth-stub-5"
	"golang.org/x/net/context"
)

// Server interface for our service methods
type Server struct {
	auth_pb.UnimplementedAuthServer
}

// CreateToken create jwt token
func (s *Server) CreateToken(ctx context.Context, req *auth_pb.CreateTokenRequest) (resp *auth_pb.CreateTokenResponse, err error) {

	prvKey, err := ioutil.ReadFile("cert/jwtRS256.key")
	if err != nil {
		log.Println(err.Error())

		resp = &auth_pb.CreateTokenResponse{
			StatusCode: auth_pb.AuthStatusCode_INTERNAL_SERVER_ERROR,
		}

		return
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(prvKey)
	if err != nil {
		log.Println(err.Error())

		resp = &auth_pb.CreateTokenResponse{
			StatusCode: auth_pb.AuthStatusCode_INTERNAL_SERVER_ERROR,
		}

		return
	}

	claims := MakeClaims(req)

	tok, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		log.Println(err.Error())

		resp = &auth_pb.CreateTokenResponse{
			StatusCode: auth_pb.AuthStatusCode_INTERNAL_SERVER_ERROR,
		}

		return
	}

	resp = &auth_pb.CreateTokenResponse{
		StatusCode:  auth_pb.AuthStatusCode_SUCCESS,
		AccessToken: tok,
	}

	log.Println("Create JWT Token Success")

	return
}

// IntrospectToken introspect jwt token
func (s *Server) IntrospectToken(ctx context.Context, req *auth_pb.IntrospectTokenRequest) (resp *auth_pb.IntrospectTokenResponse, err error) {

	// read public key from .key.pub file
	pubKey, err := ioutil.ReadFile("cert/jwtRS256.key.pub")
	if err != nil {
		log.Println("Introspect- Readfile err : ", err.Error())

		return
	}

	// validate token based on public key and extract claims
	claims, validated := Validate(pubKey, req.AccessToken)
	if !validated {

		resp = &auth_pb.IntrospectTokenResponse{
			StatusCode: auth_pb.AuthStatusCode_INVALID_TOKEN,
		}

		return
	}

	// check issuer
	issuer := fmt.Sprintf("%v", claims["iss"])
	if issuer != "Referral-System" {

		log.Println("issuer : ", issuer)

		log.Println("issuer is not match")

		resp = &auth_pb.IntrospectTokenResponse{
			StatusCode: auth_pb.AuthStatusCode_INVALID_TOKEN,
		}
		return
	}

	exp := fmt.Sprintf("%v", claims["exp"])

	// parsing expiration time (string) into time format
	val, err := time.Parse("2006-01-02 15:04:05", exp)
	if err != nil {
		log.Println("Introspect Time Parse err : ", err.Error())

		return
	}

	// if token has already expirated
	if val.Before(time.Now()) {
		log.Println("Expirate Token")

		resp = &auth_pb.IntrospectTokenResponse{
			StatusCode: auth_pb.AuthStatusCode_EXPIRATE_TOKEN,
		}
		return
	}

	// return successful response
	resp = &auth_pb.IntrospectTokenResponse{
		StatusCode:   auth_pb.AuthStatusCode_SUCCESS,
		ReferralLink: fmt.Sprintf("%v", claims["referral_link"]),
		Role:         fmt.Sprintf("%v", claims["role"]),
	}

	log.Println("Introspection Success")

	return
}

// MakeClaims will parsing request to claims
func MakeClaims(req *auth_pb.CreateTokenRequest) jwt.MapClaims {
	claims := make(jwt.MapClaims)

	claims["iss"] = req.Issuer

	claims["iat"] = req.Iat
	claims["exp"] = req.Exp

	claims["referral_link"] = req.ReferralLink
	claims["role"] = req.Role

	return claims
}

func Validate(publicKey []byte, token string) (claims jwt.MapClaims, validated bool) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		log.Printf("err validate: parse key: %w", err)

		return nil, false
	}

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		log.Printf("validate: %w", err)
		return nil, false
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		log.Printf("validate: invalid")
		return nil, false
	}

	return claims, true
}
