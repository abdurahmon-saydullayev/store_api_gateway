package handlers

import (
	"Projects/store/store_api_gateway/api/http"
	"Projects/store/store_api_gateway/genproto/user_service"
	"Projects/store/store_api_gateway/pkg/helper"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateOTP godoc
// @ID create_otp
// @Router /check [POST]
// @Summary Create OTP
// @Description  Create OTP
// @Tags OTP
// @Accept json
// @Produce json
// @Param profile body user_service.CreateOTP true "CreateOTPRequestBody"
// @Success 200 {object} http.Response{data=object{}} "OTP"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) CreateUserOTP(c *gin.Context) {
	var request *user_service.CreateOTP

	err := c.ShouldBindJSON(&request)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.UserService().CreateUserOTP(
		c.Request.Context(),
		request,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// VerifyOTP godoc
// @ID verify_otp
// @Router /check [GET]
// @Summary Verify OTP
// @Description Verify OTP
// @Tags OTP
// @Accept json
// @Produce json
// @Param otp_code query string true "otp_code"
// @Param phone_number query string true "phone_number"
// @Success 200 {object} http.Response{data=object{}} "OTP"
// @Response 400 {object} http.Response{data=string} "Invalid Argument"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) VerifyUserOTP(c *gin.Context) {
	code := c.Query("otp_code")
	phoneNumber := c.Query("phone_number")

	_, err := h.services.UserService().VerifyUserOTP(
		c.Request.Context(),
		&user_service.VerifyOTP{
			Code:        code,
			PhoneNumber: phoneNumber,
		},
	)
	if err != nil {
		if err.Error() == "rpc error: code = InvalidArgument desc = no rows in result set" {
			h.handleResponse(c, http.BadRequest, errors.New("incorrect code").Error())
			return
		}
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	// check exist
	user, err := h.services.UserService().Check(
		c.Request.Context(),
		&user_service.UserPhoneNumberReq{
			PhoneNumber: phoneNumber,
		},
	)
	// doesn't exist
	if err != nil {
		if err.Error() == "rpc error: code = InvalidArgument desc = no rows in result set" {
			h.handleResponse(c, http.BadRequest, errors.New("should register").Error())
			return
		}
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	// exists
	data := map[string]interface{}{
		"id": user.Id,
	}
	token, err := helper.GenerateJWT(data, time.Minute*10, h.cfg.SecretKey)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, token)
}
