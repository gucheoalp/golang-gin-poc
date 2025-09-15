package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/gucheoalp/golang-gin-poc/entity"
	"github.com/gucheoalp/golang-gin-poc/service"
	"github.com/gucheoalp/golang-gin-poc/validators"
)

type VideoController interface {
	FindAll() []entity.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	// Register custom validator with Gin's validator instance
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	}
	
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Video  {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		// Return a proper error response to the client
		ctx.JSON(400, gin.H{"error": err.Error()})
		return err
	}
	
	c.service.Save(video)
	return nil
}