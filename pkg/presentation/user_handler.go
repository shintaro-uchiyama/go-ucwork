package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
)

type UserHandler struct {
	userApplicationService UserApplicationServiceInterface
}

func NewUserHandler(userApplicationService UserApplicationServiceInterface) *UserHandler {
	return &UserHandler{
		userApplicationService: userApplicationService,
	}
}

func (h UserHandler) Create(c *gin.Context) {
	var userCreateRequest UserCreateRequest
	if err := c.ShouldBind(&userCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, NewErrorResponse(err))
		return
	}

	requestUser := domain.NewUser(userCreateRequest.Email, userCreateRequest.Password)
	err := h.userApplicationService.Create(*requestUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h UserHandler) List(c *gin.Context) {
	users, err := h.userApplicationService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, NewErrorResponse(err))
		return
	}
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, NewUserListResponse(users))
		return
	}
	c.JSON(http.StatusOK, NewUserListResponse(users))
}
