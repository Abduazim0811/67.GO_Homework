package userservice

import (
	"User/internal/method"
	"User/userpb"
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	userpb.UnimplementedUserserviceServer
	db *sql.DB
}

func NewUserService(db *sql.DB)*UserService{
	return &UserService{db: db}
}

func (u *UserService) CreateUser(ctx context.Context, req *userpb.Request)(*userpb.Responce, error){
	user, err := method.StoreNewUser(u.db, req)
	if err!=nil{
		log.Println(err)
		return nil, err
	}

	return &userpb.Responce{Id: user.Id}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *userpb.Responce)(*userpb.User, error){
	user, err := method.GetbyIdUser(u.db, req)
	if err!=nil{
		log.Println(err)
		return nil, err
	}
	
	return &userpb.User{
		Id: user.Id,
		Name: user.Name,
		Email: user.Email,
		Age: user.Age,
	}, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *userpb.User)(*userpb.UserRequest, error){
	err := method.UpdateUser(u.db, req)
	if err!=nil{
		log.Println(err)
		return nil ,err
	}

	return &userpb.UserRequest{Message: "Update user"}, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *userpb.Responce)(*userpb.UserRequest, error){
	err := method.DeleteUser(u.db, req)
	if err!=nil{
		log.Println(err)
		return nil,err
	}

	return &userpb.UserRequest{Message: "Deleted user"}, nil
}