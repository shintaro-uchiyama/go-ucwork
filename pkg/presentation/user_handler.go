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
		jsonError(c, http.StatusBadRequest, err)
		return
	}

	requestUser := domain.NewUser(userCreateRequest.Email, userCreateRequest.Password)
	err := h.userApplicationService.Create(*requestUser)
	if err != nil {
		jsonError(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (h UserHandler) List(c *gin.Context) {
	users, err := h.userApplicationService.List()
	if err != nil {
		jsonError(c, http.StatusInternalServerError, err)
		return
	}
	if len(users) == 0 {
		jsonError(c, http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, NewUserListResponse(users))
}
