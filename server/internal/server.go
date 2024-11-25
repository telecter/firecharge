package firecharge

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MinecraftServer struct {
	Version string `json:"version"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	ID      string `json:"id"`
}

var servers = []MinecraftServer{
	{
		Version: "1.21.3",
		Type:    "PaperMC",
		Name:    "Snap'd",
	},
}

func Start() {
	e := echo.New()
	e.HideBanner = true
	e.Static("/", "./client/dist")
	e.GET("/api/servers", func(c echo.Context) error {
		return c.JSON(200, servers)
	})

	serverPaths := e.Group("/api/server/:id")
	serverPaths.GET("/", func(c echo.Context) error {
		return c.JSON(200, servers)
	})
	e.POST("/start", func(c echo.Context) error {
		return c.JSON(200, servers)
	})
	e.POST("/stop", func(c echo.Context) error {
		return c.JSON(200, servers)
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:*",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type", "Authorization", "X-Requested-With", "Accept",
		},
	}))
	e.Start(":3000")
}
