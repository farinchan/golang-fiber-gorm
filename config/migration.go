package config

import (
	"golang-fiber-gorm/models"
	"log"
)

func RunMigration() {
    err := DB.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Gagal melakukan migrasi:", err)
    }
    log.Println("Migrasi database sukses!")
}
