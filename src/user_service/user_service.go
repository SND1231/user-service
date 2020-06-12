package user_service

import (
	"github.com/SND1231/user-service/db"
	"github.com/SND1231/user-service/model"
	pb "github.com/SND1231/user-service/proto"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

//error_list [], err_msg string
func CreateError(code codes.Code, error_list []*errdetails.BadRequest_FieldViolation) error {
	st := status.New(codes.InvalidArgument, "エラー発生")
	// add error message detail
	st, err := st.WithDetails(
		&errdetails.BadRequest{
			FieldViolations: error_list,
		},
	)
	// unexpected error
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %+v", err))
	}

	// return error
	return st.Err()
}

func CreateBadRequest_FieldViolation(feild string, desc string) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       feild,
		Description: desc,
	}
}

func CheckGetUsersRequest(request pb.GetUsersRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Limit == 0 {
		error_list = append(error_list, CreateBadRequest_FieldViolation("Limit", "値が設定されていません"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckLoginUserRequest(request pb.LoginRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Email == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("Email", "必須です"))
	}
	if request.Password == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("Password", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CheckCreateUserRequest(request pb.CreateUserRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Name == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("名前", "必須です"))
	}
	if request.Email == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("Email", "必須です"))
	}
	if request.Password == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("パスワード", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func UserExistsById(email string, id int) error {
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

func CheckUserExistsForUpdate(email string, id int) error {
	return UserExistsById(email, id)
}

func CheckUpdateUserRequest(request pb.UpdateUserRequest) error {
	var error_list []*errdetails.BadRequest_FieldViolation
	if request.Name == "" {
		error_list = append(error_list, CreateBadRequest_FieldViolation("名前", "必須です"))
	}

	if len(error_list) > 0 {
		return CreateError(codes.InvalidArgument, error_list)
	} else {
		return nil
	}
}

func CreateToken(user model.User) (string, error) {
	var err error

	// 鍵となる文字列(後で変更する)
	secret := "secret"

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
