package main

import (
	"elasticsearch/api"
	"elasticsearch/seed"
	"flag"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	err     error
	service string
)

func init() {
	dsn := "host=localhost port=5432 user=admin password=password dbname=workers sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	flag.StringVar(&service, "service", "", "run service: migration, seeding")
}

func main() {
	flag.Parse()

	switch service {
	case "seed":
		seed.Run(db)
	case "api":
		api.Run()
	}
}
