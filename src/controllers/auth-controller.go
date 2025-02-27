package controllers

import (
	"net/http"
	"smart-task-backend/src/db/models"
	"smart-task-backend/src/dto"
	"smart-task-backend/src/services"
	"smart-task-backend/src/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(context *gin.Context)
	Login(context *gin.Context)
	VerifyToken(context *gin.Context)
	RefreshToken(context *gin.Context)
	GetUsers(context *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JwtService
	logger      *utils.Logger
}

type LoginResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

func NewAuthController(authService services.AuthService, jwtService services.JwtService, logger *utils.Logger) *authController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
		logger:      logger,
	}
}

func (this_ *authController) Login(context *gin.Context) {
	var loginDto dto.Login
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}

	isValidCredential, userId := this_.authService.VerifyCredential(loginDto)

	user := this_.authService.FindUserByEmail(loginDto.Email)

	if isValidCredential {
		tokenPair := this_.jwtService.GenerateTokenPair(userId)

		context.JSON(http.StatusOK, utils.GetResponse(http.StatusOK, "Login successful...", &LoginResponse{
			AccessToken:  tokenPair["access_token"],
			RefreshToken: tokenPair["refresh_token"],
			User:         user,
		}))
		return
	}

	this_.logger.Error("Invalid credentials")
	context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, "Invalid credentials", nil))
}

func (this_ *authController) Register(context *gin.Context) {
	var userDto dto.User

	err := context.ShouldBindJSON(&userDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}

	existingUser := this_.authService.FindUserByEmail(userDto.Email)

	if existingUser.ID != 0 {
		context.JSON(
			http.StatusConflict,
			utils.GetResponse(
				http.StatusConflict,
				"A user with this email already exists",
				nil,
			),
		)

		this_.logger.Error("A user with this email already exists")

		return
	}

	result, user := this_.authService.Register(userDto)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, result.Error.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}

	tokenPair := this_.jwtService.GenerateTokenPair(user.ID)

	context.JSON(
		http.StatusCreated,
		utils.GetResponse(
			http.StatusCreated,
			"A user has been successfully created",
			&LoginResponse{
				User:         user,
				AccessToken:  tokenPair["access_token"],
				RefreshToken: tokenPair["refresh_token"],
			}),
	)
}

func (this_ *authController) VerifyToken(context *gin.Context) {
	tokenDto := dto.Token{}

	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}

	token, _ := utils.ValidateToken(tokenDto.Token)

	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, "Invalid Token", nil))
		this_.logger.Error("Invalid Token")
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(http.StatusOK, "The token is valid", gin.H{
		"is_valid": true,
	}))
}

func (this_ *authController) RefreshToken(context *gin.Context) {
	tokenDto := dto.Token{}
	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}
	token, err := utils.ValidateToken(tokenDto.Token)
	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		context.JSON(
			http.StatusOK,
			utils.GetResponse(
				http.StatusOK,
				"Token pair ready",
				this_.jwtService.GenerateTokenPair(claims["user_id"])),
		)
	} else {
		this_.logger.Error("Failed to claim a token")

		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			utils.GetResponse(
				http.StatusBadRequest,
				"Failed to claim a token",
				nil,
			),
		)
	}
}

func (this_ *authController) GetUsers(context *gin.Context) {
	users := this_.authService.GetUsers()

	context.JSON(http.StatusOK, utils.GetResponse(http.StatusOK, "Users fetched...", users))
}
