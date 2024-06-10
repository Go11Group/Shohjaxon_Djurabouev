// handler.go
package handler

import (
    // "encoding/json"
    "net/http"
    "strconv"

    "github.com/Go11Group/at_lesson/lesson35/model"
    "github.com/Go11Group/at_lesson/lesson35/storage/postgres"
    "github.com/gin-gonic/gin"
)

func CreateProblemHandler(storage *postgres.ProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        var problem model.Problem
        if err := c.ShouldBindJSON(&problem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := storage.Create(&problem); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, problem)
    }
}

func GetProblemHandler(storage *postgres.ProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        problem, err := storage.Get(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, problem)
    }
}

func UpdateProblemHandler(storage *postgres.ProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        var problem model.Problem
        if err := c.ShouldBindJSON(&problem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := storage.Update(&problem); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, problem)
    }
}

func DeleteProblemHandler(storage *postgres.ProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        if err := storage.Delete(id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    }
}

func CreateUserHandler(storage *postgres.UserStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user model.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := storage.Create(&user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, user)
    }
}

func GetUserHandler(storage *postgres.UserStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        user, err := storage.Get(id)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, user)
    }
}

func UpdateUserHandler(storage *postgres.UserStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        var user model.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := storage.Update(&user); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, user)
    }
}

func DeleteUserHandler(storage *postgres.UserStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
            return
        }
        if err := storage.Delete(id); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    }
}

func CreateSolvedProblemHandler(storage *postgres.SolvedProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        var solvedProblem model.SolvedProblem
        if err := c.ShouldBindJSON(&solvedProblem); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err := storage.Create(&solvedProblem); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, solvedProblem)
    }
}
func GetSolvedProblemHandler(storage *postgres.SolvedProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, err := strconv.Atoi(c.Param("user_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
            return
        }
        problemID, err := strconv.Atoi(c.Param("problem_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Problem ID"})
            return
        }
        solvedProblem, err := storage.Get(userID, problemID)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, solvedProblem)
    }
}

func DeleteSolvedProblemHandler(storage *postgres.SolvedProblemStorage) gin.HandlerFunc {
    return func(c *gin.Context) {
        userID, err := strconv.Atoi(c.Param("user_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
            return
        }
        problemID, err := strconv.Atoi(c.Param("problem_id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Problem ID"})
            return
        }
        if err := storage.Delete(userID, problemID); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        c.Status(http.StatusNoContent)
    }
}
