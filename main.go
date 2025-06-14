package main

import (
	"beauty_salon_bd/config"
	"beauty_salon_bd/controllers"
	"beauty_salon_bd/repositories"
	"beauty_salon_bd/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Инициализация базы данных
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Инициализация репозиториев, сервисов и контроллеров
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()

	// Настройка сессий
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Загрузка шаблонов
	r.LoadHTMLGlob("templates/*")

	// Статические файлы
	r.Static("/static", "./static")

	// Маршруты
	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user") != nil {
			c.Redirect(http.StatusFound, "/profile")
			return
		}
		c.Redirect(http.StatusFound, "/login")
	})

	r.GET("/register", authController.RegisterHandler)
	r.POST("/register", authController.RegisterPostHandler)
	r.GET("/login", authController.LoginHandler)
	r.POST("/login", authController.LoginPostHandler)
	r.GET("/profile", authController.ProfileHandler)
	r.GET("/logout", authController.LogoutHandler)

	r.Run(":5000")
}
