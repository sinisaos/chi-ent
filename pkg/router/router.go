package router

import (
	"github.com/sinisaos/chi-ent/pkg/database"
	"github.com/sinisaos/chi-ent/pkg/handler"
	"github.com/sinisaos/chi-ent/pkg/middleware"
	"github.com/sinisaos/chi-ent/pkg/service"

	"github.com/go-chi/chi/v5"
)

// Setup api routes
func SetupRoutes(app chi.Router) {
	// DB client
	client := database.DbConnection()

	// Services
	questionService := service.NewQuestionService(client)
	answerService := service.NewAnswerService(client)
	userService := service.NewUserService(client)
	authService := service.NewAuthService(client)
	tagService := service.NewTagService(client)

	// Handlers
	questionHandler := handler.NewQuestionHandler(*questionService)
	answerHandler := handler.NewAnswerHandler(*answerService)
	userHandler := handler.NewUserHandler(*userService)
	authHandler := handler.NewAuthHandler(*authService)
	tagHandler := handler.NewTagHandler(*tagService)

	// Index route
	app.Get("/", handler.Index)

	// Auth route
	app.Post("/login", authHandler.LoginHandler)

	// User routes
	app.Group(func(r chi.Router) {
		app.Get("/users", userHandler.GetAllUsersHandler)
		app.Get("/users/{id}", userHandler.GetUserHandler)
		app.Get("/users/{id}/answers", userHandler.GetUserAnswersHandler)
		app.Get("/users/{id}/questions", userHandler.GetUserQuestionsHandler)
		app.Post("/users", userHandler.CreateUserHandler)
		app.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Patch("/users/{id}", userHandler.UpdateUserHandler)
			r.Delete("/users/{id}", userHandler.DeleteUserHandler)
		})
	})

	// Question routes
	app.Group(func(r chi.Router) {
		app.Get("/questions", questionHandler.GetAllQuestionsHandler)
		app.Get("/questions/{id}", questionHandler.GetQuestionHandler)
		app.Get("/questions/{id}/answers", questionHandler.GetQuestionAnswersHandler)
		app.Get("/questions/{id}/author", questionHandler.GetQuestionAuthorHandler)
		app.Get("/questions/{id}/tags", questionHandler.GetQuestionTagsHandler)
		app.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Post("/questions", questionHandler.CreateQuestionHandler)
			r.Patch("/questions/{id}", questionHandler.UpdateQuestionHandler)
			r.Delete("/questions/{id}", questionHandler.DeleteQuestionHandler)
		})
	})

	// Answer routes
	app.Group(func(r chi.Router) {
		app.Get("/answers", answerHandler.GetAllAnswersHandler)
		app.Get("/answers/{id}", answerHandler.GetAnswerHandler)
		app.Get("/answers/{id}/author", answerHandler.GetAnswerAuthorHandler)
		app.Get("/answers/{id}/question", answerHandler.GetAnswerQuestionHandler)
		app.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Post("/answers", answerHandler.CreateAnswerHandler)
			r.Patch("/answers/{id}", answerHandler.UpdateAnswerHandler)
			r.Delete("/answers/{id}", answerHandler.DeleteAnswerHandler)
		})
	})

	// Tag routes
	app.Group(func(r chi.Router) {
		app.Get("/tags", tagHandler.GetAllTagsHandler)
		app.Get("/tags/{id}", tagHandler.GetTagHandler)
		app.Get("/tags/{id}/question", tagHandler.GetTagQuestionHandler)
		app.Group(func(r chi.Router) {
			r.Use(middleware.AuthMiddleware)
			r.Patch("/tags/{id}", tagHandler.UpdateTagHandler)
			r.Delete("/tags/{id}", tagHandler.DeleteTagHandler)
		})
	})
}
