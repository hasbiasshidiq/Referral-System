package grpcdriver

import (
	"context"
	"log"
	"time"

	auth_pb "github.com/hasbiasshidiq/auth-stub-5"

	"referral-server/config"
	entity "referral-server/entity"

	"google.golang.org/grpc"
)

//Token grpc client
type TokenGRPC struct{}

//NewTokenGRPC create new repository
func NewTokenGRPC() *TokenGRPC {
	return &TokenGRPC{}
}

//Create will send create token request to grpc server
func (r *TokenGRPC) Create(e *entity.Token) (AccessToken string, err error) {
	conn, err := grpc.Dial(config.TOKEN_GRPC_URL, grpc.WithTimeout(5*time.Second), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}

	defer conn.Close()

	a := auth_pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := a.CreateToken(ctx, &auth_pb.CreateTokenRequest{
		Issuer:       e.Issuer,
		Iat:          e.Iat.Format("2006-01-02 15:04:05"),
		Exp:          e.Exp.Format("2006-01-02 15:04:05"),
		ReferralLink: e.ReferralLink,
		Role:         e.Role,
	})
	if err != nil {
		log.Printf("Could not create Token :%v", err)
		return
	}

	AccessToken = resp.AccessToken

	log.Printf("Token Successfully Created")

	return
}

//Introspect will validate access token and return some claims like role and referral link
func (r *TokenGRPC) Introspect(AccessToken string) (Role, ReferralLink string, err error) {
	conn, err := grpc.Dial(config.TOKEN_GRPC_URL, grpc.WithTimeout(5*time.Second), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}

	defer conn.Close()

	a := auth_pb.NewAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := a.IntrospectToken(ctx, &auth_pb.IntrospectTokenRequest{
		AccessToken: AccessToken,
	})
	if err != nil {
		log.Printf("Could not introspect Token :%v", err)
		return
	}

	if resp.StatusCode == auth_pb.AuthStatusCode_EXPIRATE_TOKEN {
		err = entity.ErrExpirateToken
		return
	}

	if resp.StatusCode == auth_pb.AuthStatusCode_INVALID_TOKEN {
		err = entity.ErrInvalidToken
		return
	}

	ReferralLink = resp.ReferralLink
	Role = resp.Role

	return
}
