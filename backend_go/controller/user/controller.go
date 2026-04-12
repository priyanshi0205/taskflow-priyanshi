package user

import (
	"net/http"

	"github.com/gin-gonic/gin"

	userinputmapper "backedn_go/input_mapper/user"
	"backedn_go/internal/shared/response"
	useroutputmapper "backedn_go/output_mapper/user"
	userusecase "backedn_go/usecase/user"
)

type controller struct {
	usecase userusecase.Usecase
}

func New(usecase userusecase.Usecase) Controller {
	return &controller{usecase: usecase}
}

func (ctl *controller) List(c *gin.Context) {
	requesterID, appErr := userinputmapper.MapRequesterUserID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map requester user id for user list")
		return
	}

	users, appErr := ctl.usecase.ListExceptUser(requesterID)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to list users", "requester_user_id", requesterID)
		return
	}

	c.JSON(http.StatusOK, useroutputmapper.MapList(users))
}
