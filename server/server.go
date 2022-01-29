package main

import (
	"database/sql"
	"fmt"
	"github.com/AntonioTrupac/hannaWebshop/graph/resolver"
	"github.com/AntonioTrupac/hannaWebshop/service/users"
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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
	dsn := os.Getenv("DSN")

	if dsn == "" {
		log.Fatalf("DSN string is empty")
	}

	dbSql, err := sql.Open("mysql", dsn)
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
