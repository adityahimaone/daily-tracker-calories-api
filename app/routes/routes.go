package routes

import (
	"daily-tracker-calories/app/presenter/calories"
	"daily-tracker-calories/app/presenter/foods"
	"daily-tracker-calories/app/presenter/histories"
	"daily-tracker-calories/app/presenter/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	UserHandler      users.Presenter
	JWTMiddleware    middleware.JWTConfig
	CalorieHandler   calories.Presenter
	FoodHandler      foods.Presenter
	HistoriesHandler histories.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")
	//users endpoint
	group.POST("/users/register", handler.UserHandler.RegisterUser)
	group.PUT("/users/edit", handler.UserHandler.EditUser, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.POST("/users/login", handler.UserHandler.LoginUser)
	group.GET("/users/:id", handler.UserHandler.FindByID)
	group.POST("/users/avatars", handler.UserHandler.UploadAvatar, middleware.JWTWithConfig(handler.JWTMiddleware))

	//calories endpoint
	group.POST("/calorie/count", handler.CalorieHandler.CountCalorie)
	group.POST("/calorie/save", handler.CalorieHandler.SaveCalorie, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.GET("/calorie/user", handler.CalorieHandler.GetCalorieByUserID, middleware.JWTWithConfig(handler.JWTMiddleware))

	//foods endpoint
	group.POST("/food/save", handler.FoodHandler.SaveFood)
	group.GET("/food/", handler.FoodHandler.GetFoodByName)
	group.GET("/food", handler.FoodHandler.GetAllFood)
	group.GET("/food/:id", handler.FoodHandler.GetFoodByID)
	group.DELETE("/food/delete/:id", handler.FoodHandler.DeleteFood)
	group.PUT("/food/edit/:id", handler.FoodHandler.EditFood)

	//histories
	group.POST("/histories/create", handler.HistoriesHandler.CreateHistory, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.GET("/histories/user", handler.HistoriesHandler.GetAllHistoriesByUserID, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.GET("/histories/stat", handler.HistoriesHandler.UserStat, middleware.JWTWithConfig(handler.JWTMiddleware))
}
