package handlers

import (
	"Projects/store/store_api_gateway/api/http"
	"Projects/store/store_api_gateway/genproto/user_service"
	"Projects/store/store_api_gateway/pkg/util"
	"context"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @ID create_user
// @Router /user [POST]
// @Summary Create User
// @Description  Create User
// @Tags User
// @Accept json
// @Produce json
// @Param profile body user_service.CreateUserRequest true "CreateUserRequestBody"
// @Success 200 {object} http.Response{data=user_service.User} "GetUserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUser(c *gin.Context) {
	var user user_service.CreateUser

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().Create(
		c.Request.Context(),
		&user,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// GetUserByID godoc
// @ID get_user_by_id
// @Router /user/{id} [GET]
// @Summary Get User By ID
// @Description Get User By ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=user_service.User} "UserBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserById(c *gin.Context) {

	userId := c.Param("id")

	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "order id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().GetByID(
		context.Background(),
		&user_service.UserPrimaryKey{
			Id: userId,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
