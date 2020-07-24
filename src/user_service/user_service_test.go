package user_service

import (
	"errors"
	"github.com/SND1231/user-service/db"
	"github.com/SND1231/user-service/model"
	pb "github.com/SND1231/user-service/proto"
	"github.com/dgrijalva/jwt-go"
	"os"
	"testing"
)

const (
	Name     = "テスト"
	Email    = "test@test.com"
	PhotoUrl = "https://test"
	Password = "abcd1234"

	UserId  = 9999
	PostId  = 99999
	Content = "美味しかった"
)

func TestCheckGetUsersRequestSuccess(t *testing.T) {
	request := pb.GetUsersRequest{Limit: 1, Offset: 0, Id: 0, Name: ""}
	err := CheckGetUsersRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckGetUsersRequestLimitError(t *testing.T) {
	request := pb.GetUsersRequest{Limit: 0, Offset: 1, Id: 0, Name: ""}
	err := CheckGetUsersRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckLoginRequestSuccess(t *testing.T) {
	request := pb.LoginRequest{Email: Email, Password: Password}
	err := CheckLoginUserRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckLoginRequestEmailError(t *testing.T) {
	request := pb.LoginRequest{Email: "", Password: Password}
	err := CheckLoginUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckLoginRequestPasswordError(t *testing.T) {
	request := pb.LoginRequest{Email: Email, Password: ""}
	err := CheckLoginUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateUserRequestSuccess(t *testing.T) {
	request := pb.CreateUserRequest{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	err := CheckCreateUserRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckCreateUserRequestNameError(t *testing.T) {
	request := pb.CreateUserRequest{Name: "", Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	err := CheckCreateUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateUserRequestEmailError(t *testing.T) {
	request := pb.CreateUserRequest{Name: Name, Email: "",
		PhotoUrl: PhotoUrl, Password: Password}
	err := CheckCreateUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateUserRequestPasswordError(t *testing.T) {
	request := pb.CreateUserRequest{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: ""}
	err := CheckCreateUserRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestUserExistsByIdSuccess(t *testing.T) {
	err := UserExistsById("diff@test.com", 0)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想：", "正常終了")
	}
}

func TestUserExistsByIdExistsError(t *testing.T) {
	_ = CreateUserForTest()
	err := UserExistsById(Email, 0)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	InitUserTable()
}

func TestCheckUserExistsSuccess(t *testing.T) {
	err := CheckUserExists("diff@test.com")

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想：", "正常終了")
	}
}

func TestCheckUserExistsExistsError(t *testing.T) {
	_ = CreateUserForTest()
	err := CheckUserExists(Email)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
	InitUserTable()
}

func TestCheckUpdateUserRequestSuccess(t *testing.T) {
	request := pb.UpdateUserRequest{Name: Name, PhotoUrl: PhotoUrl}
	err := CheckUpdateUserRequest(request)

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckUpdateUserRequestNameError(t *testing.T) {
	request := pb.UpdateUserRequest{Name: "", PhotoUrl: PhotoUrl}
	err := CheckUpdateUserRequest(request)

	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCreateTokenSuccess(t *testing.T) {
	userId := CreateUserForTest()
	user := GetUserForTest(userId)
	tokenCreated, errCreated := CreateToken(user)

	if errCreated != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	token, err := jwt.Parse(tokenCreated, func(token *jwt.Token) (interface{}, error) {
		//HMAC(共通),ECDSA,RSA(公開)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("alg error")
		}
		secret := os.Getenv("SECRET_KEY")
		//keyを返す
		return []byte(secret), nil
	})

	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}

	if !token.Valid {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckCreateCommentRequestSuccess(t *testing.T) {
	request := pb.CreateCommentRequest{UserId: UserId, PostId: PostId,
		Content: Content}
	err := CheckCreateCommentRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckCreateCommentRequestUserIdError(t *testing.T) {
	request := pb.CreateCommentRequest{UserId: 0, PostId: PostId,
		Content: Content}
	err := CheckCreateCommentRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateCommentRequestPostIdError(t *testing.T) {
	request := pb.CreateCommentRequest{UserId: UserId, PostId: 0,
		Content: Content}
	err := CheckCreateCommentRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckCreateCommentRequestContentError(t *testing.T) {
	request := pb.CreateCommentRequest{UserId: UserId, PostId: PostId,
		Content: ""}
	err := CheckCreateCommentRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckGetCommentsRequestSuccess(t *testing.T) {
	request := pb.GetCommentsRequest{Limit: 1, Offset: 0, PostId: PostId}
	err := CheckGetCommentsRequest(request)
	if err != nil {
		t.Error("\n実際： ", "エラー", "\n理想： ", "正常終了")
	}
}

func TestCheckGetCommentsRequestPostIdError(t *testing.T) {
	request := pb.GetCommentsRequest{Limit: 1, Offset: 1, PostId: 0}
	err := CheckGetCommentsRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func TestCheckGetCommentsRequestLimitError(t *testing.T) {
	request := pb.GetCommentsRequest{Limit: 0, Offset: 1, PostId: PostId}
	err := CheckGetCommentsRequest(request)
	if err == nil {
		t.Error("\n実際： ", "正常終了", "\n理想： ", "エラー")
	}
}

func CreateUserForTest() int32 {
	user := model.User{Name: Name, Email: Email,
		PhotoUrl: PhotoUrl, Password: Password}
	db := db.Connection()
	defer db.Close()
	db.Create(&user)

	return user.ID
}

func GetUserForTest(id int32) model.User {
	var user model.User
	db := db.Connection()
	defer db.Close()
	db.Find(&user, id)

	return user
}

func InitUserTable() {
	db := db.Connection()
	defer db.Close()

	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM users")
}
