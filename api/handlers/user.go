package handlers

import (
	"Projects/store/store_api_gateway/api/http"
	"Projects/store/store_api_gateway/genproto/user_service"
	"Projects/store/store_api_gateway/models"
	"Projects/store/store_api_gateway/pkg/helper"
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
// @Param profile body user_service.CreateUser true "CreateUserRequestBody"
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
func (h *Handler) GetUserByID(c *gin.Context) {
	userId := c.Param("id")

	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
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

// GetUserList godoc
// @ID get_user_list
// @Router /user [GET]
// @Summary Get User List
// @Description Get User List
// @Tags User
// @Accept json
// @Produce json
// @Param offset query integer false "offset"
// @Param limit query integer false "limit"
// @Param search query string false "search"
// @Success 200 {object} http.Response{data=user_service.GetListUserResponse} "GetAllUserResponseBody"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) GetUserList(c *gin.Context) {

	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserService().GetList(
		context.Background(),
		&user_service.GetListUserRequest{
			Limit:  int64(limit),
			Offset: int64(offset),
			Search: c.Query("search"),
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// @ID update_user
// @Router /user/{id} [PUT]
// @Summary Update User
// @Description Update User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body user_service.UpdateUser true "UpdateUserRequestBody"
// @Success 200 {object} http.Response{data=user_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdateUser(c *gin.Context) {

	var user user_service.UpdateUser

	user.Id = c.Param("id")

	if !util.IsValidUUID(user.Id) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().Update(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// PatchCUser godoc
// @ID patch_user
// @Router /user/{id} [PATCH]
// @Summary Patch User
// @Description Patch User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param profile body models.UpdatePatch true "UpdatePatchRequestBody"
// @Success 200 {object} http.Response{data=user_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) UpdatePatchUser(c *gin.Context) {

	var updatePatchUser models.UpdatePatch

	err := c.ShouldBindJSON(&updatePatchUser)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	updatePatchUser.ID = c.Param("id")

	if !util.IsValidUUID(updatePatchUser.ID) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	structData, err := helper.ConvertMapToStruct(updatePatchUser.Data)
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	resp, err := h.services.UserService().UpdatePatch(
		c.Request.Context(),
		&user_service.UpdatePatchUser{
			Id:     updatePatchUser.ID,
			Fields: structData,
		},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// DeleteUser godoc
// @ID delete_user
// @Router /user/{id} [DELETE]
// @Summary Delete User
// @Description Delete User
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} http.Response{data=object{}} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) DeleteUser(c *gin.Context) {

	userId := c.Param("id")

	if !util.IsValidUUID(userId) {
		h.handleResponse(c, http.InvalidArgument, "user id is an invalid uuid")
		return
	}

	resp, err := h.services.UserService().Delete(
		c.Request.Context(),
		&user_service.UserPrimaryKey{Id: userId},
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.NoContent, resp)
}
