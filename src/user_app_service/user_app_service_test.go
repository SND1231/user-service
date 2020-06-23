package app_service

import (
	"github.com/SND1231/user-service/db"
	"github.com/SND1231/user-service/model"
	pb "github.com/SND1231/user-service/proto"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	Name     = "テスト"
	Email    = "test@test.com"
	PhotoUrl = "https://test"
	Password = "abcd1234"
)

func TestGetUsers(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	request := pb.GetUsersRequest{Limit: 1, Offset: 0, Id: 0, Name: ""}
	users, err := GetUsers(request)

	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	assert.Equal(t, Name, users[0].Name, "The two words should be the same.")
	assert.Equal(t, Email, users[0].Email, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, users[0].PhotoUrl, "The two words should be the same.")
}

func TestGetUser(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	id := GetUserID()
	user, err := GetUser(id)

	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	assert.Equal(t, Name, user.Name, "The two words should be the same.")
	assert.Equal(t, Email, user.Email, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, user.PhotoUrl, "The two words should be the same.")
}

func TestCreateUser(t *testing.T) {
	InitUserTable()

	userId := GetUserID() + 1
	request := pb.CreateUserRequest{Name: Name, Email: "create@test.com",
		PhotoUrl: PhotoUrl, Password: Password}

	id, token, err := CreateUser(request)
	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", err)
	}
	tokenTest := GetTokenForTest("create@test.com")
	assert.Equal(t, userId, id, "The two words should be the same.")
	assert.Equal(t, tokenTest, token, "The two words should be the same.")
}

func TestUpdateUser(t *testing.T) {
	InitUserTable()
	CreateUserForTest()

	id := GetUserID()
	request := pb.UpdateUserRequest{Id: id, Name: Name, PhotoUrl: "https://update"}
	id, err := UpdateUser(request)
	if err != nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", err)
	}

	user := GetUserById(id)

	assert.Equal(t, Name, user.Name, "The two words should be the same.")
	assert.Equal(t, "https://update", user.PhotoUrl, "The two words should be the same.")
}

func CreateUserForTest() {
	userParam := model.User{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	db := db.Connection()
	defer db.Close()
	db.Create(&userParam)
}

func InitUserTable() {
	db := db.Connection()
	var u model.User
	db.Delete(&u)
}

func GetUserID() int32 {
	var count int32
	db := db.Connection()
	db.Table("users").Count(&count)

	return count
}

func GetUserById(id int32) model.User {
	var user model.User

	db := db.Connection()
	defer db.Close()
	db.Find(&user, id)

	return user
}

func GetTokenForTest(email string) string {
	secret := "secret"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iss":   "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
	})

	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}
