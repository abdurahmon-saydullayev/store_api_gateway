package handlers

// import (
// 	"Projects/store/store_api_gateway/api/http"
// 	"Projects/store/store_api_gateway/genproto/order_service"
// 	"Projects/store/store_api_gateway/pkg/util"
// 	"context"

// 	"github.com/gin-gonic/gin"
// )

// // CreateOrder godoc
// // @ID create_order
// // @Router /order [POST]
// // @Summary Create Order
// // @Description  Create Order
// // @Tags Order
// // @Accept json
// // @Produce json
// // @Param profile body order_service.CreateOrder true "CreateOrderRequestBody"
// // @Success 200 {object} http.Response{data=order_service.Order} "GetOrderBody"
// // @Response 400 {object} http.Response{data=string} "Invalid Argument"
// // @Failure 500 {object} http.Response{data=string} "Server Error"
// func (h *Handler) CreateOrder(c *gin.Context) {
// 	var order order_service.CreateOrder

// 	err := c.ShouldBindJSON(&order)
// 	if err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}

// 	resp, err := h.services.OrderService().Create(
// 		c.Request.Context(),
// 		&order,
// 	)
// 	if err != nil {
// 		h.handleResponse(c, http.GRPCError, err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.Created, resp)
// }

// // GetOrderByID godoc
// // @ID get_order_by_id
// // @Router /order/{id} [GET]
// // @Summary Get Order By ID
// // @Description Get Order By ID
// // @Tags Order
// // @Accept json
// // @Produce json
// // @Param id path string true "id"
// // @Success 200 {object} http.Response{data=order_service.Order} "OrderBody"
// // @Response 400 {object} http.Response{data=string} "Invalid Argument"
// // @Failure 500 {object} http.Response{data=string} "Server Error"
// func (h *Handler) GetOrderByID(c *gin.Context) {

// 	orderId := c.Param("id")

// 	if !util.IsValidUUID(orderId) {
// 		h.handleResponse(c, http.InvalidArgument, "order id is an invalid uuid")
// 		return
// 	}

// 	resp, err := h.services.OrderService().GetById(
// 		context.Background(),
// 		&order_service.OrderPrimaryKey{
// 			Id: orderId,
// 		},
// 	)
// 	if err != nil {
// 		h.handleResponse(c, http.GRPCError, err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.OK, resp)
// }
