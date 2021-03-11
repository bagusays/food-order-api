package rest

import (
	"food-order-api/internal/delivery"
	"food-order-api/internal/delivery/rest/handler"
)

func RegisterRoute(container *delivery.Container) {
	g := container.EchoServer.Group("/api")

	g.GET("/health", handler.Check(container.HealthCheckUsecase))

	g.GET("/menu-categories", handler.FetchCategories(container.MenusUsecase))
	g.PUT("/menu-categories", handler.CreateCategory(container.MenusUsecase))
	g.PATCH("/menu-categories", handler.UpdateCategory(container.MenusUsecase))
	g.DELETE("/menu-categories", handler.DeleteCategory(container.MenusUsecase))

	g.GET("/additionals", handler.FetchAdditionals(container.MenusUsecase))
	g.PUT("/additionals", handler.CreateAdditional(container.MenusUsecase))
	g.PATCH("/additionals", handler.UpdateAdditional(container.MenusUsecase))
	g.DELETE("/additionals", handler.DeleteAdditional(container.MenusUsecase))

	g.GET("/menus", handler.FetchMenus(container.MenusUsecase))
	g.PUT("/menus", handler.CreateMenu(container.MenusUsecase))
	g.PATCH("/menus", handler.UpdateMenu(container.MenusUsecase))
	g.DELETE("/menus", handler.DeleteMenu(container.MenusUsecase))

	g.GET("/eligible-additional-menu", handler.FetchEligibleAdditionalMenu(container.MenusUsecase))
	g.PUT("/eligible-additional-menu", handler.CreateEligibleAdditionalMenu(container.MenusUsecase))
	g.PATCH("/eligible-additional-menu", handler.UpdateEligibleAdditionalMenu(container.MenusUsecase))
	g.DELETE("/eligible-additional-menu", handler.DeleteEligibleAdditionalMenu(container.MenusUsecase))

	g.GET("/orders/:userID", handler.FetchAllOrders(container.OrdersUsecase))
	g.GET("/orders/:userID/:orderID", handler.FetchOrder(container.OrdersUsecase))
}
