package app_service

import (
	"errors"
	"github.com/SND1231/user-service/db"
	"github.com/SND1231/user-service/model"
	pb "github.com/SND1231/user-service/proto"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

const (
	Name     = "テスト"
	Email    = "test@test.com"
	PhotoUrl = "https://test"
	Password = "abcd12341231"

	UpdateName     = "テスト2"
	UpdatePhotoUrl = "https://test2"

	UserId  = int32(999991)
	PostId  = int32(999991)
	Content = "美味しかった"

	UserId2  = int32(999992)
	Content2 = "美味しかった"
)

// ユーザ取得 正常
func TestGetUsersSuccess(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	request := pb.GetUsersRequest{Limit: 1, Offset: 0, Id: 0, Name: ""}
	users, err := GetUsers(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, Name, users[0].Name, "The two words should be the same.")
	assert.Equal(t, Email, users[0].Email, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, users[0].PhotoUrl, "The two words should be the same.")
}

// ユーザ取得 エラー
func TestGetUsersError(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	request := pb.GetUsersRequest{Limit: 0, Offset: 0, Id: 0, Name: ""}
	_, err := GetUsers(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

// ユーザ詳細 正常
func TestGetUserSuccess(t *testing.T) {
	InitUserTable()
	userId := CreateUserForTest()
	user, err := GetUser(userId)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, Name, user.Name, "The two words should be the same.")
	assert.Equal(t, Email, user.Email, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, user.PhotoUrl, "The two words should be the same.")
}

// ログインユーザ 正常
func TestLoginUserSuccess(t *testing.T) {
	InitUserTable()
	userId := CreateUserForTest()
	request := pb.LoginRequest{Email: Email, Password: Password}

	id, token, err := LoginUser(request)

	if err != nil {
		t.Error("\n実際： ", err, "\n理想： ", "正常終了")
	}

	if !CheckToken(token) {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
	assert.Equal(t, userId, id, "The two words should be the same.")
}

// ログインユーザ リクエストエラー
func TestLoginUserRequestError(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	request := pb.LoginRequest{Email: Email, Password: ""}

	_, _, err := LoginUser(request)

	if err == nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// ログインユーザ 不正なパスワードエラー
func TestLoginUserPasswordError(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	request := pb.LoginRequest{Email: Email, Password: "1231"}

	_, _, err := LoginUser(request)

	if err == nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// ユーザ作成 正常
func TestCreateUserSuccess(t *testing.T) {
	InitUserTable()

	userId := CreateUserForTest() + 1
	request := pb.CreateUserRequest{Name: Name, Email: "create@test.com",
		PhotoUrl: PhotoUrl, Password: Password}

	id, token, err := CreateUser(request)

	if err != nil {
		t.Error("\n実際： ", err, "\n理想： ", "正常終了")
	}
	if !CheckToken(token) {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	user := GetUserForTest(id)
	assert.Equal(t, userId, id, "The two words should be the same.")
	assert.Equal(t, Name, user.Name, "The two words should be the same.")
	assert.Equal(t, "create@test.com", user.Email, "The two words should be the same.")
	assert.Equal(t, PhotoUrl, user.PhotoUrl, "The two words should be the same.")

}

// ユーザ作成 リクエストエラー
func TestCreateUserRequestError(t *testing.T) {
	request := pb.CreateUserRequest{Name: "", Email: "create@test.com",
		PhotoUrl: PhotoUrl, Password: Password}

	_, _, err := CreateUser(request)

	if err == nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// ユーザ作成 メールアドレスエラー
func TestCreateUserEmailError(t *testing.T) {
	InitUserTable()
	CreateUserForTest()
	request := pb.CreateUserRequest{Name: "", Email: "create@test.com",
		PhotoUrl: PhotoUrl, Password: Password}

	_, _, err := CreateUser(request)

	if err == nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// ユーザ更新 正常
func TestUpdateUserSuccess(t *testing.T) {
	InitUserTable()
	userId := CreateUserForTest()

	request := pb.UpdateUserRequest{Id: userId, Name: UpdateName,
		PhotoUrl: UpdatePhotoUrl}
	id, err := UpdateUser(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	user := GetUserForTest(id)

	assert.Equal(t, UpdateName, user.Name, "The two words should be the same.")
	assert.Equal(t, UpdatePhotoUrl, user.PhotoUrl, "The two words should be the same.")
}

// ユーザ更新 リクエストエラー
func TestUpdateUserRequestError(t *testing.T) {

	request := pb.UpdateUserRequest{Id: 0, Name: "",
		PhotoUrl: UpdatePhotoUrl}
	_, err := UpdateUser(request)
	if err == nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// コメント作成 正常
func TestCreateCommentSuccess(t *testing.T) {
	InitUserTable()
	userId := CreateUserForTest()

	request := pb.CreateCommentRequest{UserId: userId, PostId: PostId,
		Content: Content}
	id, err := CreateComment(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	comment := GetCommentByIdForTest(id)

	assert.Equal(t, userId, comment.UserId, "The two words should be the same.")
	assert.Equal(t, PostId, comment.PostId, "The two words should be the same.")
	assert.Equal(t, Content, comment.Content, "The two words should be the same.")
}

// コメント作成 リクエストエラー
func TestCreateCommentRequestError(t *testing.T) {
	InitUserTable()
	userId := CreateUserForTest()

	request := pb.CreateCommentRequest{UserId: userId, PostId: PostId,
		Content: ""}
	_, err := CreateComment(request)

	if err == nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

// コメント一覧取得　正常
func TestGetCommentsSuccess(t *testing.T) {
	InitUserTable()

	comment1 := model.Comment{UserId: UserId, PostId: PostId, Content: Content}
	comment2 := model.Comment{UserId: UserId2, PostId: PostId, Content: Content2}
	comment3 := model.Comment{UserId: UserId2, PostId: int32(32), Content: "あああ"}
	CreateCommentForTest(comment1)
	CreateCommentForTest(comment2)
	CreateCommentForTest(comment3)

	request := pb.GetCommentsRequest{PostId: PostId, Limit: 3, Offset: 1}
	commentList, count, err := GetComments(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	assert.Equal(t, int32(2), count, "The two words should be the same.")

	assert.Equal(t, UserId2, commentList[0].UserId, "The two words should be the same.")
	assert.Equal(t, Content2, commentList[0].Content, "The two words should be the same.")

	assert.Equal(t, UserId, commentList[1].UserId, "The two words should be the same.")
	assert.Equal(t, Content, commentList[1].Content, "The two words should be the same.")

}

// コメント一覧取得　エラー
func TestGetCommentsError(t *testing.T) {
	InitUserTable()

	request := pb.GetCommentsRequest{PostId: PostId, Limit: 0, Offset: 1}
	_, _, err := GetComments(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func CreateUserForTest() int32 {
	hash, _ := bcrypt.GenerateFromPassword([]byte(Password), 10)
	password := string(hash)
	user := model.User{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: password}
	db := db.Connection()
	defer db.Close()
	db.Create(&user)

	return user.ID
}

func CreateCommentForTest(comment model.Comment) {
	db := db.Connection()
	defer db.Close()
	db.Create(&comment)
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

func GetUserForTest(id int32) model.User {
	var user model.User
	db := db.Connection()
	defer db.Close()
	db.Find(&user, id)

	return user
}

func GetCommentByIdForTest(id int32) model.Comment {
	var comment model.Comment
	db := db.Connection()
	defer db.Close()
	db.Find(&comment, id)

	return comment
}

func InitUserTable() {
	db := db.Connection()
	defer db.Close()

	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM users")
}

func CheckToken(tokenCreated string) bool {
	token, err := jwt.Parse(tokenCreated, func(token *jwt.Token) (interface{}, error) {
		//HMAC(共通),ECDSA,RSA(公開)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("alg error")
		}
		//keyを返す
		return []byte("secret"), nil
	})

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	return true
}
