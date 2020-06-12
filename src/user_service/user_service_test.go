package user_service

import (
	pb "github.com/SND1231/user-service/proto"
	"testing"
	"github.com/SND1231/user-service/db"
	"github.com/SND1231/user-service/model"
)

const (
	Name = "テスト"
	Email = "test@test.com"
	PhotoUrl = "https://test"
	Password = "abcd1234"
)

func TestCheckGetUsersRequestSuccess(t *testing.T) {
	request := pb.GetUsersRequest{Limit:1, Offset:0, Id:0, Name:""}
	err := CheckGetUsersRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckGetUsersRequestLimitError(t *testing.T) {
	request := pb.GetUsersRequest{Limit:0, Offset:1, Id:0, Name:""}
	err := CheckGetUsersRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckLoginRequestSuccess(t *testing.T) {
	request := pb.LoginRequest{Email:Email, Password: Password}
	err := CheckLoginUserRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckLoginRequestEmailError(t *testing.T) {
	request := pb.LoginRequest{Email:"", Password: Password}
	err := CheckLoginUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckLoginRequestPasswordError(t *testing.T) {
	request := pb.LoginRequest{Email:Email, Password: ""}
	err := CheckLoginUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateUserRequestSuccess(t *testing.T){
	request := pb.CreateUserRequest{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	err := CheckCreateUserRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckCreateUserRequestNameError(t *testing.T){
	request := pb.CreateUserRequest{Name: "", Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	err := CheckCreateUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateUserRequestEmailError(t *testing.T){
	request := pb.CreateUserRequest{Name: Name, Email: "",
		PhotoUrl: PhotoUrl, Password: Password}
	err := CheckCreateUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateUserRequestPasswordError(t *testing.T){
	request := pb.CreateUserRequest{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: ""}
	err := CheckCreateUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestUserExistsByIdSuccess(t *testing.T){
	err := UserExistsById("diff@test.com", 0)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想：", "正常終了")
	}
}

func TestUserExistsByIdExistsError(t *testing.T){
	CreateUser()
	err := UserExistsById(Email, 0)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	InitUserTable()
}

func CreateUser(){
	user_param := model.User{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	db := db.Connection()
	defer db.Close()
	db.Create(&user_param)
}

func InitUserTable(){
	db := db.Connection()
	var u model.User
	db.Delete(&u)
}