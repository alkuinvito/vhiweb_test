package main

import (
	"vhiweb_test/app/users"
	"vhiweb_test/app/vendors"
	"vhiweb_test/lib/adapters"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	db := adapters.NewDB()
	db.AutoMigrate(&vendors.VendorModel{})
	db.AutoMigrate(&users.UserModel{})
}
