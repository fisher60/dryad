package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/fisher60/dryad/internal/config"
	"github.com/fisher60/dryad/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/ydb-platform/ydb-go-sdk/v3/log"
)

func uuidToStr(uuid pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
}

func listUsers() []string {
	users, err := database.Engine.ListDryadUsers(context.Background())
	if err != nil {
		log.Error(err)
	}

	var out []string

	for _, v := range users {
		out = append(out, uuidToStr(v.AbandonauthUuid))
	}

	return out
}

func main() {
	db_config := config.DatabaseConfig{Host: "localhost", Port: 5432, User: "postgres", Password: "postgres", DbName: "postgres"}
	database.InitializeDatabse(db_config)

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{
			"message": "pong",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]any{"users": listUsers()})
	})

	http.ListenAndServe(":8000", router)
}
