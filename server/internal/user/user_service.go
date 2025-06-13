package user

import (
	"context"
	"github.com/Gurshan94/chatapp/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)


type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedpassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedpassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       r.ID,
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil
}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, util.MyJWTClaims{
		ID:       u.ID,
		Username: u.Username,
		Email: u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &LoginUserRes{
		AccessToken: ss,
		Username:    u.Username,
		ID:          u.ID,
		Email: u.Email,
	}, nil
}

func(s* service) GetUserByID(c context.Context, userID int64) (*User, error){
	return s.Repository.GetUserByID(c, userID)
}

