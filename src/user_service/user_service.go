package user_service

import (
	"fmt"
	"github.com/SND1231/user-service/db"
	"github.com/SND1231/user-service/model"
	pb "github.com/SND1231/user-service/proto"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
)

//errorList [], err_msg string
func CreateError(code codes.Code, errorList []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "エラー発生")
	// add error message detail
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: errorList,
		},
	)
	// unexpected error
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %+v", err))
	}

	// return error
	return st.Err()
}

func CreateBadRequestFieldViolation(feild string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       feild,
		Description: desc,
	}
}

func CheckGetUsersRequest(request pb.GetUsersRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Limit == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("Limit", "値が設定されていません"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func CheckLoginUserRequest(request pb.LoginRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Email == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("Email", "必須です"))
	}
	if request.Password == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("Password", "必須です"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func CheckCreateUserRequest(request pb.CreateUserRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Name == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("名前", "必須です"))
	}
	if request.Email == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("Email", "必須です"))
	}
	if request.Password == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("パスワード", "必須です"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func UserExistsById(email string, id int32) error {
	var user model.User
	db := db.Connection()
	defer db.Close()

	db.Where("email = ? AND id <> ?", email, id).First(&user)
	log.Println(user.ID)
	if user.ID != 0 {
		return status.New(codes.AlreadyExists, "設定したEmailのユーザが他に存在します").Err()
	}
	return nil
}

func CheckUserExists(email string) error {
	return UserExistsById(email, 0)
}

func CheckUpdateUserRequest(request pb.UpdateUserRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.Name == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("名前", "必須です"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func CreateToken(user model.User) (string, error) {
	var err error

	// 鍵となる文字列
	secret := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss":   "__init__", // JWT の発行者が入る(文字列(__init__)は任意)
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func CheckCreateCommentRequest(request pb.CreateCommentRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation

	if request.UserId == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("UserID", "値が設定されていません"))
	}
	if request.PostId == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("PostId", "値が設定されていません"))
	}
	if request.Content == "" {
		errorList = append(errorList, CreateBadRequestFieldViolation("Content", "値が設定されていません"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}

func CheckGetCommentsRequest(request pb.GetCommentsRequest) error {
	var errorList []*errdetails.BadRequest_FieldViolation
	if request.PostId == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("PostId", "値が設定されていません"))
	}
	if request.Limit == 0 {
		errorList = append(errorList, CreateBadRequestFieldViolation("Limit", "値が設定されていません"))
	}

	if len(errorList) > 0 {
		return CreateError(codes.InvalidArgument, errorList)
	} else {
		return nil
	}
}
