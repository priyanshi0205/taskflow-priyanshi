package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	authInputMapper "backedn_go/input_mapper/auth"
	"backedn_go/internal/shared/response"
	authOutputMapper "backedn_go/output_mapper/auth"
	authUsecase "backedn_go/usecase/auth"
)

type controller struct {
	usecase authUsecase.Usecase
}

func New(usecase authUsecase.Usecase) Controller {
	return &controller{usecase: usecase}
}

func (ctl *controller) Register(c *gin.Context) {
	input, appErr := authInputMapper.MapRegister(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map register input")
		return
	}
	result, appErr := ctl.usecase.Register(input)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to register user")
		return
	}

	c.JSON(http.StatusOK, authOutputMapper.MapRegister(result))
}

func (ctl *controller) Login(c *gin.Context) {
	input, appErr := authInputMapper.MapLogin(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map login input")
		return
	}

	result, appErr := ctl.usecase.Login(input)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to login user")
		return
	}

	c.JSON(http.StatusOK, authOutputMapper.MapLogin(result))
}
