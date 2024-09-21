package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mddsilva/gobet/internal/auth/domain/dao"
	"github.com/mddsilva/gobet/internal/auth/domain/dao/dto"
	"github.com/mddsilva/gobet/internal/auth/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Signup(c *gin.Context)
	Signin(c *gin.Context)
}

type AuthServiceImpl struct {
	userRepository repository.UserRepository
}

func (u AuthServiceImpl) Signup(c *gin.Context) {
	var request dto.SignupDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input data",
			"details": err.Error(),
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error while hashing the password",
			"details": err.Error(),
		})
		return
	}

	data, err := u.userRepository.Save(&dao.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hash),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error while saving user",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data":    data,
	})
}

func (u AuthServiceImpl) Signin(c *gin.Context) {
	var request dto.SigninDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid input data",
			"details": err.Error(),
		})
		return
	}

	user, err := u.userRepository.FindByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"details": err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credentials invalid",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Credentials invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func AuthServiceInit(userRepository repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepository: userRepository,
	}
}
