package main

import (
	"github.com/rdav27/go-biz-admin/database"
	"github.com/rdav27/go-biz-admin/routes"
)

func main() {
	database.Connect()

	r := routes.SetupRouter()
	_ = r.Run(":8080")
}
