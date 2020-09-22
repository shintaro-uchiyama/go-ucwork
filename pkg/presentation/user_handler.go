package presentation

import (
	"github.com/gin-gonic/gin"
	"github.com/shintaro-uchiyama/go-ucwork/pkg/domain"
	"github.com/sirupsen/logrus"
	"net/http"
)

type UserHandler struct {
	userApplicationService UserApplicationServiceInterface
}

func NewUserHandler(service UserApplicationServiceInterface) *UserHandler {
	return &UserHandler{
		userApplicationService: service,
	}
}

func (h UserHandler) Create(c *gin.Context) {
	logrus.Info("start create user")
	var userCreateRequest UserCreateRequest
	if err := c.ShouldBind(&userCreateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestUser := domain.NewUser(userCreateRequest.Email, userCreateRequest.Password)
	_, err := h.userApplicationService.Regist(*requestUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
	logrus.Info("end create user")
}

func (h UserHandler) List(c *gin.Context) {
	logrus.Info("start list user")
	users, err := h.userApplicationService.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, NewUserListResponse(users))
	logrus.Info("end list user")
}
