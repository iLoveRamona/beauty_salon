package controllers

import (
	"beauty_salon_bd/models"
	"beauty_salon_bd/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) RegisterHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Регистрация",
	})
}

func (c *AuthController) RegisterPostHandler(ctx *gin.Context) {
	user := &models.User{
		Name:     ctx.PostForm("name"),
		Phone:    ctx.PostForm("phone"),
		Password: ctx.PostForm("password"),
	}

	err := c.authService.Register(user)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusFound, "/login")
}

func (c *AuthController) LoginHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Вход",
	})
}

func (c *AuthController) LoginPostHandler(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	user, err := c.authService.Login(phone, password)
	if err != nil {
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", user.Phone)
	session.Save()

	ctx.Redirect(http.StatusFound, "/profile")
}

func (c *AuthController) ProfileHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	phone := session.Get("user")
	if phone == nil {
		ctx.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := c.authService.GetUserProfile(phone.(string))
	if err != nil {
		ctx.Redirect(http.StatusFound, "/logout")
		return
	}

	ctx.HTML(http.StatusOK, "profile.html", gin.H{
		"title": "Профиль",
		"name":  user.Name,
		"phone": user.Phone,
	})
}

func (c *AuthController) LogoutHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("user")
	session.Save()
	ctx.Redirect(http.StatusFound, "/login")
}
