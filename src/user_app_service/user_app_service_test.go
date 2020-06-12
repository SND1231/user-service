package app_service

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

func TestGetUsers(t *testing.T){
	CreateUserForTest()
	request := pb.GetUsersRequest{Limit:1, Offset:0, Id:0, Name:""}
	users, count := GetUsers(request)

	if count != 1 {
		t.Error("\n実際： ", count, "\n理想： ", 1)
	}
}

func CreateUserForTest(){
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

func get_user_id() int {
	var count int
	db := db.Connection()
	db.Table("users").Count(&count)

	return count
}