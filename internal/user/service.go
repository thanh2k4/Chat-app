package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thanh2k4/Chat-app/internal/user/infras/postgres"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"time"
)

type UserServer struct {
	gen.UnimplementedUserServiceServer
	Postgres *postgres.Queries
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.Postgres.CreateUser(ctx, postgres.CreateUserParams{
		ID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		FullName: req.Name,
		Email:    pgtype.Text{String: req.Email, Valid: true},
		Phone:    pgtype.Text{String: req.Phone, Valid: true},
		Status:   req.Status,
		Avatar:   pgtype.Text{String: req.Avatar, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return &gen.User{
		Id:     req.Id,
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		Status: req.Status,
		Avatar: req.Avatar,
	}, nil

}

func (s *UserServer) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	user, err := s.Postgres.GetUserByUserID(ctx, pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil {
		return nil, err
	}
	return &gen.User{
		Id:     user.ID.String(),
		Email:  user.Email.String,
		Name:   user.FullName,
		Phone:  user.Phone.String,
		Status: user.Status,
		Avatar: user.Avatar.String,
	}, nil

}

func (s *UserServer) UpdateUser(ctx context.Context, req *gen.UpdateUserRequest) (*gen.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	userID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}
	_, err = s.Postgres.UpdateUserByID(ctx, postgres.UpdateUserByIDParams{
		ID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		FullName: req.Name,
		Email:    pgtype.Text{String: req.Email, Valid: true},
		Phone:    pgtype.Text{String: req.Phone, Valid: true},
		Status:   req.Status,
		Avatar:   pgtype.Text{String: req.Avatar, Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &gen.User{
		Id:     req.Id,
		Name:   req.Name,
		Email:  req.Email,
		Phone:  req.Phone,
		Status: req.Status,
		Avatar: req.Avatar,
	}, nil

}
