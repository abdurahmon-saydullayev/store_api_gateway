package handlers

import (
	"Projects/store/store_api_gateway/api/http"
	"Projects/store/store_api_gateway/config"
	"Projects/store/store_api_gateway/grpc/client"
	"Projects/store/store_api_gateway/pkg/logger"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

type Handler struct {
	cfg      config.Config
	log      logger.LoggerI
	services client.ServiceManagerI
}

func NewHandler(cfg config.Config, log logger.LoggerI, svcs client.ServiceManagerI) Handler {
	return Handler{
		cfg:      cfg,
		log:      log,
		services: svcs,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, data)
}

func ProtoToStruct(s interface{}, p protoiface.MessageV1) error {
	var jm jsonpb.Marshaler

	jm.EmitDefaults = true
	jm.OrigName = true

	ms, err := jm.MarshalToString(p)

	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(ms), &s)

	return err
}

func (h *Handler) handleErrorResponse(c *gin.Context, code int, message string, err interface{}) {
	h.log.Error(message, logger.Int("code", code), logger.Any("error", err))
	c.JSON(code, ResponseModel{
		Code:    code,
		Message: message,
		Error:   err,
	})
}

func (h *Handler) handleSuccessResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, ResponseModel{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
