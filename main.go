package main

import (
	"log"
	_ "net/http"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-yaml/yaml"                          
    "io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type env struct {  
    Env profile	
}

type profile struct {
	Dbms	string
	User	string
	Password	string
	Dbname	string
	Host	string
	Port	string
}

func connectDatabase(env string) *gorm.DB {
	buf, err := ioutil.ReadFile("./db/config_"+env+".yaml")
	if err != nil {
		fmt.Println("Error")                                    
		fmt.Println(err)
		panic(err.Error())                             
	}

	config, err := ReadOnStruct(buf)                  
    if err != nil {                                
		fmt.Println(err)
		panic(err.Error())                                                           
	}  

	DBMS     := config.Env.Dbms
	USER     := config.Env.User
	PASS     := config.Env.Password
	HOST 	 := config.Env.Host
	PORT	 := config.Env.Port
	DBNAME   := config.Env.Dbname
  
	CONNECT := "host="+HOST+" port="+PORT+" user="+USER+" dbname="+DBNAME+" password="+PASS+" sslmode=disable"
	db,err := gorm.Open(DBMS, CONNECT)
  
	if err != nil {
	  panic(err.Error())
	}
	return db
}

func ReadOnStruct(fileBuffer []byte) (env, error) {
	var	data env
               
	err := yaml.Unmarshal(fileBuffer, &data)    
    if err != nil {   
		fmt.Println("Error")                                
        fmt.Println(err)                                                             
	}                                            
    return data, nil                                  
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := connectDatabase("dev")
	fmt.Println(db)

	err := e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}