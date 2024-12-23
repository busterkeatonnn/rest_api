package route

import (
	"net/http"
	"strconv"

	"tspo_final/internal/middleware"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"
	services "tspo_final/internal/service"

	"github.com/gin-gonic/gin"
)

var userRepository *repository.UserRepository

// @Summary		Создать юзера
// @Description	Creates a users
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			user	body		models.User	true	"User details"
// @IsResponseSuccess		201		{object}	dto.InfoResponse
// @Failure		400		{object}	dto.InfoResponse
// @Router			/users/create [post]
func createUser(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to user")
		return
	}

	code := http.StatusOK

	response := services.CreateUser(&user, *userRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary		Получить инфу по всем заказам
// @Description	Returns list of users
// @Tags			users
// @Produce		json
// @IsResponseSuccess		200	{object}	dto.InfoResponse
// @Failure		400	{object}	dto.InfoResponse
// @Router			/users/ [get]
func getUsers(context *gin.Context) {
	code := http.StatusOK

	response := services.FindAllUsers(*userRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary		Получить инфу по одному конкретному юзеру
// @Description	Returns one user
// @Tags			users
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @IsResponseSuccess		200	{object}	dto.InfoResponse
// @Failure		400	{object}	dto.InfoResponse
// @Router			/users/show/{id} [get]
func getUser(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.FindOneUserById(id, *userRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary		Обновить инфу по одному конкретному юзеру
// @Description	Updates a users
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id		path		string		true	"User ID"
// @Param			user	body		models.User	true	"User details"
// @IsResponseSuccess		201		{object}	dto.InfoResponse
// @Failure		400		{object}	dto.InfoResponse
// @Router			/users/update/{id} [put]
func updateUser(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, "Can't bind JSON to user")
		return
	}

	code := http.StatusOK

	response := services.UpdateUserById(id, &user, *userRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

// @Summary		Удалить юзера
// @Description	Deletes a users
// @Tags			users
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @IsResponseSuccess		201	{object}	dto.InfoResponse
// @Failure		400	{object}	dto.InfoResponse
// @Router			/users/delete/{id} [delete]
func deleteUser(context *gin.Context) {
	id_, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	id := uint(id_)

	code := http.StatusOK

	response := services.DeleteOneUserById(id, *userRepository)

	if !response.IsResponseSuccess {
		code = http.StatusBadRequest
	}

	context.JSON(code, response)
}

func SetupUsersRoutes(userRepository_ *repository.UserRepository, route *gin.Engine) {
	userRepository = userRepository_
	route.Group("/").Use(middleware.AuthMiddleware())
	{
		route.POST("/users/create", createUser)
		route.GET("/users/", getUsers)
		route.GET("/users/show/:id", getUser)
		route.PUT("/users/update/:id", updateUser)
		route.DELETE("/users/delete/:id", deleteUser)
	}
}
