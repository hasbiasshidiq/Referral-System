package token

import (
	"io/ioutil"
	"log"

	"github.com/dgrijalva/jwt-go"
	auth_pb "github.com/hasbiasshidiq/auth-stub-3"
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
