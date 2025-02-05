package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User 定义模型
type User struct {
	Id      int    `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(255)"`
	Age     int
	Created time.Time `gorm:"autoCreateTime"`
	Updated time.Time `gorm:"autoUpdateTime"`
}

var DB *gorm.DB

func init() {
	var err error
	// 配置连接字符串
	dsn := "user=postgres password=123456 dbname=postgres host=localhost port=5432 sslmode=disable"
	// 连接到 PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	// 自动迁移，确保数据库表与模型同步
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}
}

// Add 插入数据
func Add() {
	user := User{Name: "张三", Age: 25}
	result := DB.Create(&user)
	if result.Error != nil {
		fmt.Println("Error inserting user:", result.Error)
	} else {
		fmt.Println("User added successfully:", user)
	}
}

// Delete 删除数据
func Delete() {
	id := 7
	var user User
	result := DB.Delete(&user, id)
	if result.Error != nil {
		fmt.Println("Error deleting user:", result.Error)
	} else {
		fmt.Println("User deleted successfully")
	}
}

// Update 更新数据
func Update() {
	user := User{Id: 6, Name: "李四", Age: 24}
	result := DB.First(&user, user.Id)
	if result.Error != nil {
		fmt.Println("Error finding user:", result.Error)
		return
	}

	// 更新字段
	user.Name = "张三"
	user.Age = 24
	result = DB.Save(&user)
	if result.Error != nil {
		fmt.Println("Error updating user:", result.Error)
	} else {
		fmt.Println("User updated successfully:", user)
	}
}

// Find 查找数据
func Find() User {
	user := User{Name: "张三"}
	result := DB.Where("name = ?", user.Name).First(&user)
	if result.Error != nil {
		fmt.Println("Error finding user:", result.Error)
		var null User
		return null
	}

	fmt.Println("Found user:", user)
	return user
}
