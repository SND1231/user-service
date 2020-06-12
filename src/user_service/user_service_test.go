package user_service

import (
	pb "github.com/SND1231/user_service/proto"
	"testing"
	"github.com/SND1231/user_service/db"
	"github.com/SND1231/user_service/model"
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

	t.Log("TestCheckGetUsersRequestOffsetError終了")
}

func TestCheckGetUsersRequestLimitError(t *testing.T) {
	request := pb.GetUsersRequest{Limit:0, Offset:1, Id:0, Name:""}
	err := CheckGetUsersRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}

	t.Log("TestCheckGetUsersRequestLimitError終了")
}

func TestCheckCreateUserRequestSuccess(t *testing.T) {
	CreateUser()
	request := pb.LoginRequest{Email:Email, Password: Password}
	err := CheckLoginUserRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	t.Log("TestCheckCreateUserRequestSuccess終了")
}

func TestCheckCreateUserRequestEmailError(t *testing.T) {
	request := pb.LoginRequest{Email:"", Password: Password}
	err := CheckLoginUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}

	t.Log("TestCheckCreateUserRequestEmailError終了")
}

func TestCheckCreateUserRequestPasswordError(t *testing.T) {
	request := pb.LoginRequest{Email:Email, Password: ""}
	err := CheckLoginUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}

	initUserTable()
	t.Log("TestCheckCreateUserRequestPasswordError終了")
}



func CreateUser(){
	user_param := model.User{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	db := db.Connection()
	defer db.Close()
	db.Create(&user_param)
}

func initUserTable(){
	db := db.Connection()
	var u model.User
	db.Delete(&u)
}