package main

import (
	"database/sql"
	"log"
	"net/http"
	// "strconv"

	"github.com/Go11Group/at_lesson/lesson35/handler"
	"github.com/Go11Group/at_lesson/lesson35/storage/postgres"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "your_connection_string")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()

	problemStorage := postgres.CreateProblemStorage(db)
	userStorage := postgres.NewUserStorage(db)
	solvedProblemStorage := postgres.NewSolvedProblemStorage(db)

	router.POST("/problems", handler.CreateProblemHandler(problemStorage))
	router.GET("/problems/:id", handler.GetProblemHandler(problemStorage))
	router.PUT("/problems/:id", handler.UpdateProblemHandler(problemStorage))
	router.DELETE("/problems/:id", handler.DeleteProblemHandler(problemStorage))

	router.POST("/users", handler.CreateUserHandler(userStorage))
	router.GET("/users/:id", handler.GetUserHandler(userStorage))
	router.PUT("/users/:id", handler.UpdateUserHandler(userStorage))
	router.DELETE("/users/:id", handler.DeleteUserHandler(userStorage))

	router.POST("/solved_problems", handler.CreateSolvedProblemHandler(solvedProblemStorage))
	router.GET("/solved_problems/:user_id/:problem_id", handler.GetSolvedProblemHandler(solvedProblemStorage))
	router.DELETE("/solved_problems/:user_id/:problem_id", handler.DeleteSolvedProblemHandler(solvedProblemStorage))

	log.Fatal(http.ListenAndServe(":8080", router))
}
