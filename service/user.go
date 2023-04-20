package service

import (
	"domo1/util/common"
	"domo1/util/dto"
	"domo1/util/model"
	"domo1/util/response"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginService(login dto.LoginDto) response.ResponseStruct {
	res := response.NewResponse()
	var user model.User
	DB := common.GetDB()
	DB.Where("username = ?", login.Username).First(&user)
	if user.ID == 0 {
		res.HttpStatus = http.StatusBadRequest
		res.Code = response.FailCode
		res.Msg = response.UserNotExist
		return res
	}
	//判断密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.NameOrPasswordError
		return res
	}
	return res
}

func RegisterService(user dto.RegisterDto) response.ResponseStruct {
	res := response.NewResponse()
	DB := common.GetDB()

	//用户是否存在
	if isUserExist(DB, user.Username) {
		res.HttpStatus = http.StatusUnprocessableEntity
		res.Code = response.CheckFailCode
		res.Msg = response.EmailRegistered
		return res
	}

	//创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		res.HttpStatus = http.StatusInternalServerError
		res.Code = response.ServerErrorCode
		res.Msg = response.SystemError
		//记录日志
		// util.LogError("hashed password " + err.Error())
		return res
	}

	newUser := model.User{
		Username: user.Username,
		Password: string(hashedPassword),
	}
	DB.Create(&newUser)
	return res
}

func isUserExist(DB *gorm.DB, s string) bool {
	var user model.User
	DB.Where("username = ?", s).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
