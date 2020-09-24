package main

import (
	"GoWebServiceWithCRUD/pkg/config"
	"GoWebServiceWithCRUD/pkg/routes"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)




func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	config.InitialMigration()
	routes.UserEndPointsRequestHandler()
	routes.UserEndPointsRequestHandler()
}










