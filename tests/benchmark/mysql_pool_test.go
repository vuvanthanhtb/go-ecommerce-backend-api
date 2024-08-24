package main

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "Thanh"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	dsn := "root:Thanh2024@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Handle error if you want
			// fmt.Println("Error dropping table:", err)
		}
	}

	// Create table if the table exists
	db.AutoMigrate(&User{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Thiết lập các tham số kết nối
	sqlDb.SetMaxOpenConns(1) // Số lượng kết nối tối đa
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns10(b *testing.B) {
	dsn := "root:Thanh2024@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Handle error if you want
			// fmt.Println("Error dropping table:", err)
		}
	}

	// Create table if the table exists
	db.AutoMigrate(&User{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}

	// Thiết lập các tham số kết nối
	sqlDb.SetMaxOpenConns(10) // Số lượng kết nối tối đa
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
