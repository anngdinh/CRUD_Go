package main

import (
	"fmt"
	// "html"
	// "log"
	// "net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	name  string
	email string
}

// CURD
func test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "ping !",
	})
}
func GetAllUser(ctx *gin.Context) {
	db := sqlConnect()
	defer db.Close()
	
	var users []User
	all, _ := db.Query("SELECT * FROM Users")
	for all.Next() {
		var temp User
		all.Scan(&temp.name, &temp.email)
		users = append(users, temp)
	}
	fmt.Println(users)
}


func main() {
	db := sqlConnect()
	db.Query("CREATE TABLE Users (name VARCHAR(255), email VARCHAR(255))")

	db.Query("INSERT INTO Users VALUES ( 'minh', 'idol' )")
	db.Query("INSERT INTO Users VALUES ( 'an', 'ga bap' )")
	db.Close()

	router := gin.Default()

	router.GET("/test", test)
	router.GET("/", GetAllUser)

	router.Run(":8081")

}


func sqlConnect() *sql.DB {
	fmt.Println("Connect MySQL !!")
		DBMS := "mysql"
	USER := "go_test"
	PASS := "password"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "go_database"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	db, err := sql.Open(DBMS, CONNECT)

	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("!!!!!!!!!!!!!! ERROR !!!!!!!!!!!!!!!!!!!!!")
		panic(err.Error())
	}

	return db
	// defer db.Close()
}