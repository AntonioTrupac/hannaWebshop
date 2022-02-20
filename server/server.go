package main

import (
	"database/sql"
	"fmt"
	"github.com/AntonioTrupac/hannaWebshop/graph/resolver"
	"github.com/AntonioTrupac/hannaWebshop/model"
	productService "github.com/AntonioTrupac/hannaWebshop/service/products"
	userService "github.com/AntonioTrupac/hannaWebshop/service/users"
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

var database *gorm.DB

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("Cannot read port from environment!")
	}

	initDB()
	productsService := productService.NewProducts(database)
	usersService := userService.NewUsers(database)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver.NewResolver(usersService, productsService)}))

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
		Logger:                                   logDB.LogGORM(),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic("FAILED TO CONNECT TO DB")
	}

	err = migrate(database)
	if err != nil {
		return
	}

	if err != nil {
		fmt.Println(err)
		panic("MODELS NOT ADDED")
	}
}

func migrate(db *gorm.DB) error {
	err := db.Debug().AutoMigrate(&model.User{}, &model.Address{}, &model.Product{}, &model.Image{}, &model.Category{})
	if err != nil {
		return err
	}

	return nil
}
