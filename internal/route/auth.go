package route

import (
	"net/http"
	"strings"
	"time"

	"tspo_final/internal/auth"
	"tspo_final/internal/models"
	"tspo_final/internal/repository"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var userRepositoryAuth *repository.UserRepository

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// @Summary  Регистрация пользователя
// @Tags   auth
// @Description Регистрация пользователя по логину и паролю
// @ID    auth-sign-up
// @Accept   json
// @Produce  json
// @Param   input body Credentials true "Логин, пароль, роль"
// @Router   /auth/sign_up [post]
func signup(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	res := userRepositoryAuth.FindOneByUserName(creds.Username)
	if res.Error == nil {
		c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
		return
	}

	role := "user" // дефолтная роль

	if creds.Role != "" {
		role = creds.Role
	}

	userRepository.Save(&models.User{
		Name:         creds.Username,
		HashPassword: creds.Password,
		UserRole:     role,
	})

	c.JSON(http.StatusCreated, gin.H{"message": "user signed up successfully"})
}

// SignIn docs
//
// @Summary  Вход пользователей
// @Tags   auth
// @Description Вход для всех пользователей по логину и паролю
// @ID    auth-sign-in
// @Accept   json
// @Produce  json
// @Param   input body Credentials true "Логин и пароль"
// @Router   /auth/login [post]
func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	res := userRepositoryAuth.FindOneByUserName(creds.Username)
	if res.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	user := (res.Result).(*models.User)
	if user.HashPassword != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username, user.UserRole)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func generateToken(username string, role string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &auth.Claims{
		Username: username,
		Role:     role, // Включаем роль в токен
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(auth.JwtKey)
}

// @Summary  Обновление токенов
// @Tags   auth
// @Description Обновление токенов
// @ID    auth-refresh
// @Accept   json
// @Produce  json
// @Router   /auth/refresh [post]
func refresh(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.Abort()
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.Abort()
		return
	}

	if len(headerParts[1]) == 0 {
		c.Abort()
		return
	}
	accessToken := headerParts[1]

	claims := &auth.Claims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	newToken, err := generateToken(claims.Username, claims.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}

func SetupAuthRoutes(userRepository_ *repository.UserRepository, route *gin.Engine) {
	userRepositoryAuth = userRepository_
	route.POST("/auth/sign_up", signup)
	route.POST("/auth/login", login)
	route.POST("/auth/refresh", refresh)
}
