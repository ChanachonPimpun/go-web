package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = ""
	dbname   = "stock_db"
)

func main() {
	//สร้างตัวแปรเก็บ database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local")

	//สร้าง log เอาไว้ดู
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)

	// สร้าง database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	//ตรวจสอบว่า connect database ได้ไหม
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println(db)
	fmt.Println("Database connection success...")

	// ทำ Migration ตัว model กับ database
	db.AutoMigrate(&Product{})
	fmt.Println("Migration success")

	//ทดสอบเรียกใช้งานบันทึกข้อมูลลงตาราง
	newProduct := &Product{
		Name:     "Iphone 1",
		Price:    10000,
		Quantity: 10,
	}
	createProductDB(db, newProduct)

	//ทดสอบเรียกใช้งานดึงข้อมูลทั้งหมดจากตาราง
	// getProoductAllDB(db)

	//ทดสอบเรียกใช้งานดึงข้อมูลโดยมีเงื่อนไข ID
	// getProductByIDDB(db, 1)

	//ทดสอบเรียกใช้งานดึงข้อมูลโดยมีเงื่อนไข name
	// getProductByNameDB(db, "Iphone")

	//ทดสอบเรียกใช้งานอัพเดทข้อมูลโดยมีเงื่อนไข ID
	// updateProduct := &Product{
	// 	Name:        "Bluemi",
	// 	Price:       25000,
	// 	Quantity:    10,
	// }
	// updateProductByIDDB(db, 1, updateProduct)

	//ทดสอบเรียกใช้งานอัพเดทข้อมูลโดยมีเงื่อนไข name
	updateProduct := &Product{
		Name:     "SUMSAMG",
		Price:    10000,
		Quantity: 4,
	}
	updateProductByNameDB(db, "SUMSAMG", updateProduct)

	//ทดสอบเรียกใช้งานลบข้อมูลโดยมีเงื่อนไข ID
	deleteProductByIDDB(db, 1)
}

func createProductDB(db *gorm.DB, product *Product) {
	// คำสั่งบันทึกข้อมูลลง table
	result := db.Create(product)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Create product success")
}

// ฟังก์ชั่นเรียกดข้อมูลทั้งหมดจาก table
func getProoductAllDB(db *gorm.DB) {
	//สร้างตัวแปรเก็บข้อมูลที่ดึงมาจาก table
	var products []Product

	//คำสั่งดึงข้อมูลทั้งหมดจาก table
	result := db.Find(&products)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product all success")
	fmt.Println(products)
	//return products
}

// ฟังก์ชั่นดึงข้อมูลโดยมีเงื่อนไข ณ ทีนี้ใช้ id จาก table
func getProductByIDDB(db *gorm.DB, id int) {
	//สร้างตัวแปรเก็บข้อมูลที่ดึงมาจาก table
	var product Product

	//คำสั่งดึงข้อมูลโดยมีเงื่อนไข
	result := db.First(&product, id)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by id success")
	fmt.Println(product)
	//return product
}

// ฟงัก์ชั่นดึงข้อมูลโดยมีเงื่อนไข ณ ทีนี้ใช้ name จาก table
func getProductByNameDB(db *gorm.DB, name string) {
	//สร้างตัวแปรเก็บข้อมูลที่ดึงมาจาก table
	var product Product

	//คำสั่งดึงข้อมูลโดยมีเงื่อนไข name
	result := db.Where("name LIKE ?", "%"+name+"%").Find(&product)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Get product by name success")
	fmt.Println(product)
	//return product
}

// ฟังก์ชั่นอัพเดทข้อมูลโดยมีเงื่อนไข ณ ทีนี้ใช้ id จาก table
func updateProductByIDDB(db *gorm.DB, id int, product *Product) {
	//คำสั่งอัพเดตข้อมูล
	result := db.Model(&Product{}).Where("id = ?", id).Updates(product)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product by id success")
}

// ฟงัก์ชั่นแก้ไขข้อมูลโดยมีเงื่อนไข ณ ทีนี้ใช้ name จาก table
func updateProductByNameDB(db *gorm.DB, name string, product *Product) {
	//คำสั่งอัพเดตข้อมูล
	result := db.Model(&Product{}).Where("name Like ?", "%"+name+"%").Updates(product)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Update product by name success")
}

// ฟังก์ชั่นลบข้อมูลโดยมีเงื่อนไข ณ ทีนี้ใช้ id จาก table
func deleteProductByIDDB(db *gorm.DB, id int) {
	//คำสั่งลบข้อมูล (soft delete)
	// result := db.Delete(&Product{	}, id)
	//คำสั่งลบข้อมูล (hard delete)
	result := db.Unscoped().Delete(&Product{}, id)

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete product by id success")
}

// ฟงัก์ชั่นลบข้อมูลโดยมีเงื่อนไข ณ ทีนี้ใช้ name จาก table
func deleteProductByNameDB(db *gorm.DB, name string) {
	//คำสั่งลบข้อมูล (soft delete)
	// result := db.Where("name LIKE ?", "%"+name+"%").Delete(&Product{})
	//คำสั่งลบข้อมูล (hard delete)
	result := db.Unscoped().Where("name LIKE ?", "%"+name+"%").Delete(&Product{})

	//ตรวจสอบผล
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Delete product by name success")
}
