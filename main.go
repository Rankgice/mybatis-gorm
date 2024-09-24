package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"mybatis-gorm/utils"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func initDB() *gorm.DB {
	dsn := "root:123584679y@tcp(47.100.102.21:3306)/gorm_sql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 自动迁移模式
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil
	}

	return db
}

func main() {
	db := initDB()
	// 测试数据
	sql := `select * from users where email like #{email} and name like "王%" order by created_at limit #{pageSize}`
	values := map[string]interface{}{
		"email":    "%qq.com",
		"pageSize": 2,
	}
	var users []User
	utils.TranSql(db, sql, values).Scan(&users)
	// 打印返回值
	for _, user := range users {
		log.Printf("User: %+v", user)
	}
}
