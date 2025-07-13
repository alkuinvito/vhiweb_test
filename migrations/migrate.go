package main

import (
	"vhiweb_test/app/users"
	"vhiweb_test/lib/adapters"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	db := adapters.NewDB()
	db.AutoMigrate(&users.UserModel{})
}
