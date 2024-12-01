package routes

import (
	"workout-tracker/handlers"
	"workout-tracker/middlewares"
	"workout-tracker/services"

	"github.com/gorilla/mux"
)

type AppRouter struct {
	Router             *mux.Router
	UserService        services.UserService
	WorkoutService     services.WorkoutPlanService
	ExerciseService    services.ExerciseService
	AuthHandler        handlers.AuthHandlers
	UserHandler        handlers.UserHandler
	WorkoutPlanHandler handlers.WorkoutPlanHandler
	ExerciseHandler    handlers.ExerciseHandler
}

func NewAppRouter(userService services.UserService, workoutService services.WorkoutPlanService, exerciseService services.ExerciseService) *AppRouter {
	// Initialize handlers
	authHandler := handlers.NewAuthHandler(userService)
	userHandler := handlers.NewUserHandler(userService)
	workoutPlanHandler := handlers.NewWorkoutServiceHandler(workoutService, userService)
	exerciseHandler := handlers.NewExerciseHandler(exerciseService)

	// Create a new AppRouter instance
	appRouter := &AppRouter{
		Router:             mux.NewRouter(),
		UserService:        userService,
		WorkoutService:     workoutService,
		ExerciseService:    exerciseService,
		AuthHandler:        authHandler,
		UserHandler:        userHandler,
		WorkoutPlanHandler: workoutPlanHandler,
		ExerciseHandler:    exerciseHandler,
	}

	// Set up routes
	appRouter.setupRoutes()

	return appRouter
}

func (app *AppRouter) setupRoutes() {
	// Public routes
	app.Router.HandleFunc("/signup", app.AuthHandler.SignUp).Methods("POST")
	app.Router.HandleFunc("/login", app.AuthHandler.Login).Methods("POST")
	app.Router.HandleFunc("/logout", app.AuthHandler.Logout).Methods("POST")

	// Auth Middleware routes
	userRouter := app.Router.PathPrefix("/user").Subrouter()
	userRouter.Use(middlewares.JwtMiddleware)
	userRouter.HandleFunc("/{id}", app.UserHandler.GetUser).Methods("GET")
	userRouter.HandleFunc("/{id}", app.UserHandler.DeleteUser).Methods("DELETE")
	userRouter.HandleFunc("/{id}", app.UserHandler.UpdateUser).Methods("POST")

	workoutPlanRouter := app.Router.PathPrefix("/workout_plan").Subrouter()
	workoutPlanRouter.Use(middlewares.JwtMiddleware)
	workoutPlanRouter.HandleFunc("/{id}", app.WorkoutPlanHandler.CreateWorkoutPlanForUser).Methods("POST")

	userWorkoutRouter := app.Router.PathPrefix("/user/{id}/workout_plan").Subrouter()
	userWorkoutRouter.HandleFunc("", app.WorkoutPlanHandler.GetAllWorkoutPlansForUser).Methods("GET")
	userWorkoutRouter.HandleFunc("/{wp_id}", app.WorkoutPlanHandler.GetWorkoutPlanForUser).Methods("GET")
	userWorkoutRouter.HandleFunc("/{wp_id}", app.WorkoutPlanHandler.UpdateWorkoutPlanForUser).Methods("POST")
	userWorkoutRouter.HandleFunc("/{wp_id}", app.WorkoutPlanHandler.RemoveWorkoutPlanForUser).Methods("DELETE")
	userWorkoutRouter.HandleFunc("/{wp_id}/status", app.WorkoutPlanHandler.UpdateWorkoutPlanStatusForUser).Methods("POST")

	exerciseRouter := userWorkoutRouter.PathPrefix("/{wp_id}/exercise").Subrouter()
	exerciseRouter.Use(middlewares.JwtMiddleware)
	exerciseRouter.HandleFunc("", app.ExerciseHandler.CreateExercise).Methods("POST")
	exerciseRouter.HandleFunc("/{ex_id}", app.ExerciseHandler.UpdateExercise).Methods("POST")
	exerciseRouter.HandleFunc("/{ex_id}", app.ExerciseHandler.DeleteExercise).Methods("DELETE")
}
