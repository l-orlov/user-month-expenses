package handler

import (
	"net/http"

	//"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/l-orlov/user-month-expenses/internal/config"
	"github.com/l-orlov/user-month-expenses/internal/service"
	"github.com/sirupsen/logrus"
)

type (
	Handler struct {
		cfg *config.Config
		log *logrus.Logger
		svc *service.Service
	}
)

func New(
	cfg *config.Config, log *logrus.Logger, svc *service.Service,
) *Handler {
	c := &Handler{
		cfg: cfg,
		log: log,
		svc: svc,
	}

	return c
}

func (h *Handler) InitRoutes() http.Handler {
	router := gin.New()

	//router.Use(
	//	// for static files
	//	static.Serve("/", static.LocalFile("./static", true)),
	//)

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.CreateUser)
			users.GET("/:id", h.GetUserByID)
			users.GET("/", h.GetAllUsers)
			users.GET("/with-params", h.GetUsersWithParameters)
			users.PUT("/", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}

		expenses := api.Group("/expenses")
		{
			expenses.POST("/", h.CreateUserExpense)
			expenses.GET("/:id", h.GetUserExpenseByUserID)
			expenses.GET("/", h.GetAllUserExpenses)
			expenses.GET("/with-params", h.GetUserExpensesWithParameters)
			expenses.GET("/by-categories", h.GetUserExpensesByCategories)
			expenses.PUT("/", h.UpdateUserExpense)
			expenses.DELETE("/:id", h.DeleteUserExpense)
		}
	}

	return CORS(router)
}
