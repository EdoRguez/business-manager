package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
	auth "github.com/EdoRguez/business-manager/auth-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/auth-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/jwt_manager"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/password_hash"
)

type AuthService struct {
	Repo *repo.UserRepo
	Jwt  jwt_manager.JWTWrapper
	auth.UnimplementedAuthServiceServer
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	fmt.Println("Auth Service :  Register")
	fmt.Println("Auth Service :  Register - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	_, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  Register - ERROR")
		fmt.Println(err.Error())
		return &auth.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	createUserParams := db.CreateUserParams{
		CompanyID:    req.CompanyId,
		RoleID:       req.RoleId,
		Email:        req.Email,
		PasswordHash: password_hash.HashPassword(req.Password),
	}

	_, err = s.Repo.CreateUser(ctx, createUserParams)
	if err != nil {
		fmt.Println("Auth Service :  Register - ERROR")
		fmt.Println(err.Error())
		return &auth.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  Register - SUCCESS")
	return &auth.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	fmt.Println("Auth Service :  Login")
	fmt.Println("Auth Service :  Login - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  Login - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Email or password incorrect"
		}

		return &auth.LoginResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	match := password_hash.CheckPasswordHash(req.Password, u.PasswordHash)

	if !match {
		return &auth.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Email or password incorrect",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(u.ID, u.Email, "")

	fmt.Println("Auth Service :  Login - SUCCESS")
	return &auth.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
