package app_service

import (
	"github.com/SND1231/user_service/db"
	"github.com/SND1231/user_service/model"
	pb "github.com/SND1231/user_service/proto"
	"github.com/SND1231/user_service/user_service"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUsers(request pb.GetUsersRequest) ([]*pb.User, error) {
	var users []model.User
	var userList []*pb.User

	err := user_service.CheckGetUsersRequest(request)
	if err != nil {
		return userList, err
	}

	limit := request.Limit
	offset := limit * (request.Offset - 1)
	id := request.Id

	db := db.Connection()
	defer db.Close()
	db.Where("id <> ?", id).Limit(limit).Offset(offset).
		Find(&users).Scan(&userList)

	return userList, nil
}

func GetUser(id int32) (pb.User, error) {
	var user model.User
	var user_param pb.User

	db := db.Connection()
	defer db.Close()
	db.Find(&user, id).Scan(&user_param)

	return user_param, nil
}

func LoginUser(request pb.LoginRequest) (int32, string, error) {
	var user model.User
	var token string

	err := user_service.CheckLoginUserRequest(request)
	if err != nil {
		return -1, "", err
	}

	db := db.Connection()
	defer db.Close()
	db.Where("email = ?", request.Email).First(&user)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return -1, "", status.New(codes.InvalidArgument, "無効なEmail または　無効なパスワードです").Err()
	}

	token, err = user_service.CreateToken(user)
	if err != nil {
		return -1, "", status.New(codes.Unknown, "作成失敗").Err()
	}
	return user.ID, token, nil
}

func CreateUser(request pb.CreateUserRequest) (int32, string, error) {
	var hash []byte
	err := user_service.CheckCreateUserRequest(request)
	if err != nil {
		return -1, "", err
	}

	hash, err = bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return -1, "", err
	}

	password := string(hash)
	user_param := model.User{Name: request.Name, Email: request.Email,
		PhotoUrl: request.PhotoUrl, Password: password}

	err = user_service.CheckUserExists(request.Email)
	if err != nil {
		return -1, "", err
	}

	db := db.Connection()
	defer db.Close()
	db.Create(&user_param)
	if db.NewRecord(user_param) == false {
		token, err := user_service.CreateToken(user_param)
		return user_param.ID, token, err
	}
	return -1, "", status.New(codes.Unknown, "作成失敗").Err()
}

func UpdateUser(request pb.UpdateUserRequest) (int32, error) {
	err := user_service.CheckUpdateUserRequest(request)
	if err != nil {
		return -1, err
	}

	var id = request.Id
	err = user_service.CheckUserExistsForUpdate(request.Email, int(id))
	if err != nil {
		return -1, err
	}

	user_param := model.User{Name: request.Name, PhotoUrl: request.PhotoUrl}

	db := db.Connection()
	defer db.Close()
	user := model.User{}
	db.Find(&user, id)

	db.Model(&user).UpdateColumns(user_param)
	return id, nil

}
