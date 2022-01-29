package main

import (
	"database/sql"
	"fmt"
	"github.com/AntonioTrupac/hannaWebshop/graph/resolver"
	"github.com/AntonioTrupac/hannaWebshop/service/users"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AntonioTrupac/hannaWebshop/graph/generated"
	logDB "github.com/AntonioTrupac/hannaWebshop/loggers"
)

const defaultPort = "8080"

var database *gorm.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	initDB()
	usersService := users.New(database)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(usersService)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB() {
	var err error

	dbSql, err := sql.Open("mysql", "e3as53baoujr:pscale_pw_-FP67XNcWA9wEnYtVEAYKSKGzOlbzej-ETam_vFurtU@tcp(kqmwg2ezxey8.eu-west-3.psdb.cloud)/hannawebshop?tls=true&charset=utf8&parseTime=true&loc=Local")
	database, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbSql,
	}), &gorm.Config{
		Logger: logDB.LogGORM(),
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNECT TO DB")
	}

	err = database.AutoMigrate(&generated.User{}, &generated.Address{})

	if err != nil {
		fmt.Println(err)
		panic("MODELS NOT ADDED")
	}
}
