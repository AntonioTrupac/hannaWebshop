package executeDb

import (
	"database/sql"
	"fmt"
	logDb "github.com/AntonioTrupac/hannaWebshop/dbLogger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Age       int
}

func Execute() (*gorm.DB, *sql.DB) {
	db, err := sql.Open("mysql", "84oq5zrl3er9:pscale_pw_1aThg2_Feg_LwWBFcfq-ENYUDsxep_-uGHsXXQAev3Y@tcp(kqmwg2ezxey8.eu-west-3.psdb.cloud)/hannawebshop?tls=true")
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logDb.DbLog(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	sqlDB, _ := gormDB.DB()

	user := User{FirstName: "Luka", LastName: "Lukic", Age: 20}
	result := gormDB.Create(&user)

	var users []User
	gormDB.Find(&users)
	for _, u := range users {
		fmt.Println(u.FirstName, u.LastName)
	}

	// returns error
	err = result.Error

	if err != nil {
		panic(err.Error()) //this should be a proper error handler
	}

	defer func(sqlDB *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(sqlDB)

	return gormDB, sqlDB
}
