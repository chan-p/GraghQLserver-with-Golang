package main

import (
	"log"
	_ "net/http"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/go-yaml/yaml"
	"github.com/yuuu/gqlgen-echo-sample/graph"
	"github.com/yuuu/gqlgen-echo-sample/graph/generated"                          
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

	db,err := gorm.Open(
		config.Env.Dbms, 
		fmt.Sprintf(
			"host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
			config.Env.Host, config.Env.User, config.Env.Password, config.Env.Port, config.Env.Dbname,
		),
	)
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

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err := e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}