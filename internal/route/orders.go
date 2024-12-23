package route

import (
	"net/http"
	"strconv"

	"tspo_final/internal/middleware"
	"tspo_final/internal/models"
	repository "tspo_final/internal/repository"
	services "tspo_final/internal/service"

	"github.com/gin-gonic/gin"
)

var orderRepository *repository.OrderRepository

// @Summary			Создать заказ
// @Description		Creates a orders
// @Tags				orders
// @Accept				json
// @Produce			json
// @Param				order	body		models.Order	true	"Order details"
// @IsResponseSuccess	201					{object}		dto.InfoResponse
// @Failure			400		{object}	dto.InfoResponse
// @Router				/orders/create [post]
func createOrder(context *gin.Context) {
	var order models.Order

	err := context.ShouldBindJSON(&order)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to order")
		return
	}

	code := http.StatusOK

	response := services.CreateOrder(&order, *orderRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary			Получить все заказы
// @Description		Returns list of orders
// @Tags				orders
// @Produce			json
// @IsResponseSuccess	200	{object}	dto.InfoResponse
// @Failure			400	{object}	dto.InfoResponse
// @Router				/orders/ [get]
func getOrders(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllOrders(*orderRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary			Получить инфу по одному конкретному заказу
// @Description		Returns one order
// @Tags				orders
// @Produce			json
// @Param				id	path		string	true	"Order ID"
// @IsResponseSuccess	200	{object}	dto.InfoResponse
// @Failure			400	{object}	dto.InfoResponse
// @Router				/orders/show/{id} [get]
func getOrder(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneOrderById(id, *orderRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary			Обновить инфу по одному конкретному заказу
// @Description		Updates a orders
// @Tags				orders
// @Accept				json
// @Produce			json
// @Param				id		path		string			true	"Order ID"
// @Param				order	body		models.Order	true	"Order details"
// @IsResponseSuccess	201					{object}		dto.InfoResponse
// @Failure			400		{object}	dto.InfoResponse
// @Router				/orders/update/{id} [put]
func updateOrder(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var order models.Order

	err := context.ShouldBindJSON(&order)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to order")
		return
	}

	code := http.StatusOK

	response := services.UpdateOrderById(id, &order, *orderRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary			Удалить заказ
// @Description		Deletes a orders
// @Tags				orders
// @Accept				json
// @Produce			json
// @Param				id	path		string	true	"Order ID"
// @IsResponseSuccess	201	{object}	dto.InfoResponse
// @Failure			400	{object}	dto.InfoResponse
// @Router				/orders/delete/{id} [delete]
func deleteOrder(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneOrderById(id, *orderRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupOrdersRoutes(orderRepository_ *repository.OrderRepository, route *gin.Engine) {
	orderRepository = orderRepository_

	authorized := route.Group("/orders").Use(middleware.AuthMiddleware())

	authorized.POST("/create", createOrder)
	authorized.GET("/", getOrders)
	authorized.GET("/show/:id", getOrder)
	authorized.PUT("/update/:id", updateOrder)
	authorized.DELETE("/delete/:id", deleteOrder)
}
