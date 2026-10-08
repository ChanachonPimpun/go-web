package main

import (
	"gorm.io/gorm"
)

// ประกาศ struct ที่มีโครงสร้างเดียวกับตารางที่เราสร้างใน database
type Product struct {
	gorm.Model // จะสร้าง Id, create_at, update_at, delete_at แบบ auto
	Name       string
	Price      float32
	Quantity   int
}
