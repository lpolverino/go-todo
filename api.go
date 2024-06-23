package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type APIServer struct {
	ListenAddr string
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		ListenAddr: listenAddr,
	}
}

func (a *APIServer) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	CreateUsersGroup(e)
	CreateTodoGroup(e)

	e.Logger.Fatal(e.Start(":" + a.ListenAddr))
}

func CreateUsersGroup(e *echo.Echo) {
	g := e.Group("/users")
	g.POST("/log-in", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return echo.ErrBadRequest
		}

		fmt.Printf("the user %s : %s", u.Name, u.Password)

		if u.Name != "jon" || u.Password != "shhh" {
			return echo.ErrUnauthorized
		}

		claims := &jwtCustomClaims{
			"Jon Snow",
			true,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})

	g.POST("/sing-up", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return echo.ErrBadRequest
		}
		claims := &jwtCustomClaims{
			u.Name,
			false,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	})
}

func CreateTodoGroup(e *echo.Echo) {
	g := e.Group("/todos")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte("secret"),
	}

	g.Use(echojwt.WithConfig(config))

	g.GET("/", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)
		name := claims.Name
		return c.String(http.StatusOK, "Hello"+name)
	})
}
