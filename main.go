package main

import (
	"fmt"
	"ginlearn/db"
	"ginlearn/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "helllo world")
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
func main() {
	//http.HandleFunc("/hello", sayHello)
	//err := http.ListenAndServe(":9090", nil)
	//
	//if err != nil {
	//	fmt.Printf("http server start faild, err: %v\n", err)
	//	return
	//}

	//router := gin.Default()
	//router.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "Hello World")
	//})
	//router.Run(":8000")

	//env := os.Environ()
	//var cmdEnv []string
	//fmt.Println(env)
	//for _, e := range env {
	//	i := strings.Index(e, "=")
	//	if i > 0 && (e[:i] == "ENV_NAME") {
	//		fmt.Println(e)
	//	} else {
	//		cmdEnv = append(cmdEnv, e)
	//	}
	//}
	//fmt.Println(cmdEnv)

	db.InitDB()
	router := routers.InitRouter()

	// use ginSwagger middleware to
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()

}
