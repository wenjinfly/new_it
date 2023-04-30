// package main
package testexample

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Student struct {
	StudentIds uint       `gorm:"primarykey"`
	CreatedAt  *time.Time `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time   comment:            version:2023-03-12 22:39
	UpdatedAt  *time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time   comment:            version:2023-03-12 22:39
	DeletedAt  *time.Time `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time   comment:            version:2023-03-12 22:39

	Code  string
	Name  string
	Books []Book `gorm:"many2many:student_books;joinForeignKey:StudentIds;joinReferences:BookIds"`
}

type Book struct {
	BookIds   uint       `gorm:"primarykey"`
	CreatedAt *time.Time `gorm:"column:created_at" json:"CreatedAt"` //type:*time.Time   comment:            version:2023-03-12 22:39
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"UpdatedAt"` //type:*time.Time   comment:            version:2023-03-12 22:39
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"DeletedAt"` //type:*time.Time   comment:            version:2023-03-12 22:39

	Name     string
	Students []Student `gorm:"many2many:student_books;joinForeignKey:BookIds;joinReferences:StudentIds"`
}

// User 拥有并属于多种 language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}

// 检索 User 列表并预加载 Language
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("Languages").Find(&users).Error
	return users, err
}

// 检索 Language 列表并预加载 User
func GetAllLanguages(db *gorm.DB) ([]Language, error) {
	var languages []Language
	err := db.Model(&Language{}).Preload("Users").Find(&languages).Error
	return languages, err
}

func InitDB() *gorm.DB {
	//创建一个数据库的连接
	mysqlConfig := mysql.Config{
		DSN:                       "root:PXDN93VRKUm8TeE7@tcp(192.168.0.169:33069)/test_db?charset=utf8&parseTime=true", // DSN data source name
		DefaultStringSize:         191,                                                                                  // string 类型字段的默认长度
		SkipInitializeWithVersion: false,                                                                                // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)}); err != nil {
		panic("can not open db error: " + err.Error())
		//return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(1)
		sqlDB.SetMaxOpenConns(2)
		return db
	}
}

func test_gorm() {
	db := InitDB()
	db.AutoMigrate(&Student{}, &Book{})

	student1 := Student{Code: "000001", Name: "张三"}
	student2 := Student{Code: "000002", Name: "李四"}
	student3 := Student{Code: "000003", Name: "王五"}
	db.Save(&student1)
	db.Save(&student2)
	db.Save(&student3)

	var studentlist []Student
	db.Table("students").Where("id = 1").Or("id = 2").Find(&studentlist)

	book1 := Book{Name: "笑傲江湖", Students: studentlist}
	book2 := Book{
		Name: "神雕侠侣", Students: []Student{
			student3,
		},
	}

	db.Save(&book1)
	db.Save(&book2)

	var student Student
	db.Table("students").Where("id = 1").First(&student)
	book := []Book{}
	db.Preload("Students").Find(&book)
	fmt.Println(book)

	//db.Model(&student).Related(&book, "Books")
	//fmt.Println(book)
	db.Model(&student).Association("Books").Find(&book)
	fmt.Println(book)

	var bookQ Book
	db.Table("books").Where("id = 1").First(&bookQ)
	db.Model(&bookQ).Association("Students").Find(&studentlist)
	fmt.Println(studentlist)

	db.Model(&bookQ).Association("Students").Append(&student3)
	db.Model(&bookQ).Association("Students").Append(&Student{Code: "000005", Name: "西门吹雪"})

	db.Model(&bookQ).Association("Students").Delete(student3)
	db.Model(&bookQ).Association("Students").Clear()
}
