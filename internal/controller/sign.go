package controller

import (
	"github.com/kataras/iris/v12"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/model"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/repository"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/service"
	"gitlab.feedtoken.tech/xdance_group/xdance/internal/util"
	"gitlab.feedtoken.tech/xdance_group/xdance/pkg/logger"
	"gorm.io/gorm"
)

type SignController struct {
	SignService service.ISignService
}

func (s *SignController) PostSignIn(ctx iris.Context) interface{} {
	var req model.SignRequestModel
	err := ctx.ReadJSON(&req)
	if err != nil {
		return err
	}
	s.SignService.Save()

	return req
}

func SignUp(db *gorm.DB, req model.SignRequestModel) interface{} {
	logger.Info("request.........")

	user := &repository.User{}
	tx := db.Model(&repository.User{}).Where(&repository.User{Email: req.Email}).First(user)
	if err := tx.Error; err != nil {
		return util.FailureResponse(err)
	}
	if user == nil {
		//新注册
	} else {
		//已注册
	}
	return util.SuccessResponse("ok")
}
