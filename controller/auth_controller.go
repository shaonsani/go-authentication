package controller

import (
 "fmt"
 "go-authentication/data/request"
 "go-authentication/helper"
 "go-authentication/model"
 "net/http"

 "gorm.io/gorm"

 "github.com/gin-gonic/gin"
 "github.com/go-playground/validator/v10"
)

type AuthController struct {
 Db       *gorm.DB
 Validate *validator.Validate
}

func NewAuthControllerImpl(Db *gorm.DB, validate *validator.Validate) *AuthController {
 return &AuthController{Db: Db, Validate: validate}
}


// Register handles user registration.
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   user  body  models.User  true  "User data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]

func (c AuthController) Register(ctx *gin.Context) {
	var reqBody request.RegisterRequest
	if err := ctx.BindJSON(&reqBody); err != nil {
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 return
	}
   
	if err := c.Validate.Struct(reqBody); err != nil {
	 validationErrors := err.(validator.ValidationErrors)
	 errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	 return
	}
   
	var existingUser model.User
	result := c.Db.Where("email = ?", reqBody.Email).First(&existingUser)
	if result.RowsAffected > 0 {
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
	 return
	}
   
	password, err := helper.EncryptPassword(reqBody.Password)
   
	if err != nil {
	 ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
	 return
	}
   
	newUser := model.User{
	 Name:     reqBody.Name,
	 Email:    reqBody.Email,
	 Password: password,
	}
   
	if err := c.Db.Create(&newUser).Error; err != nil {
	 ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
	 return
	}
   
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
   }

   func (c AuthController) Login(ctx *gin.Context) {
	var reqBody request.LoginRequest
	if err := ctx.BindJSON(&reqBody); err != nil {
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	 return
	}
   
	if err := c.Validate.Struct(reqBody); err != nil {
	 validationErrors := err.(validator.ValidationErrors)
	 errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": errorMessage})
	 return
	}
   
	var existingUser model.User
	result := c.Db.Where("email = ?", reqBody.Email).First(&existingUser)
	if result.RowsAffected < 1 {
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
	 return
	}
   
	valid := helper.ComparePassword(reqBody.Password, existingUser.Password)
   
	if valid != true {
	 ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password invalid"})
	 return
	}
   
	token, err := helper.CreateToken(existingUser.Email)
   
	if err != nil {
	 ctx.JSON(http.StatusInternalServerError, gin.H{"error": "JWT Error"})
	 return
	}
   
	ctx.JSON(http.StatusOK, gin.H{"access_token": token})
   }