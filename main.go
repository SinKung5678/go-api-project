// main.go
package main

import (
	"encoding/json" // ใช้ในการเข้ารหัสและถอดรหัสข้อมูล JSON
	"fmt"           // ใช้สำหรับการพิมพ์ข้อความลงคอนโซล
	"log"           // ใช้สำหรับการบันทึก log
	"net/http"      // ใช้สำหรับการสร้าง HTTP server
	"strconv"       // ใช้สำหรับการแปลงค่าประเภท string เป็น int
)

// กำหนดโครงสร้างของข้อมูลผู้ใช้
type User struct {
	FirstName string `json:"first_name"` // ชื่อ
	LastName  string `json:"last_name"`  // นามสกุล
	Email     string `json:"email"`      // อีเมล
	Age       int    `json:"age"`        // อายุ
	Address   string `json:"address"`    // ที่อยู่
	Gender    string `json:"gender"`     // เพศ
	Phone     string `json:"phone"`      // เบอร์โทร
}

// ข้อมูลผู้ใช้ตัวอย่าง
var users = []User{
	{"John", "Doe", "john@example.com", 25, "123 Main St", "Male", "555-1234"},
	{"Jane", "Smith", "jane@example.com", 30, "456 Elm St", "Female", "555-5678"},
	// เพิ่มข้อมูลอื่นๆ ตามต้องการ
}

// ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้โดยใช้ Pagination
func getUsers(w http.ResponseWriter, r *http.Request) {
	// รับพารามิเตอร์ page และ limit จาก URL
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	// แปลงค่าพารามิเตอร์เป็นจำนวนเต็ม
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	// คำนวณตำแหน่งเริ่มต้นและสิ้นสุดของข้อมูลที่จะดึง
	start := (page - 1) * limit
	end := start + limit

	// ตรวจสอบขอบเขตของข้อมูล
	if start > len(users) {
		start = len(users)
	}
	if end > len(users) {
		end = len(users)
	}

	// ดึงข้อมูลผู้ใช้ตามขอบเขตที่คำนวณ
	paginatedUsers := users[start:end]

	// กำหนด header ให้เป็น JSON
	w.Header().Set("Content-Type", "application/json")
	// เข้ารหัสข้อมูลผู้ใช้เป็น JSON และส่งกลับ
	json.NewEncoder(w).Encode(paginatedUsers)
}

// ฟังก์ชันหลักสำหรับการเริ่มต้นเซิร์ฟเวอร์
func main() {
	// กำหนดเส้นทางของ API
	http.HandleFunc("/users", getUsers)
	// พิมพ์ข้อความบอกสถานะเซิร์ฟเวอร์
	fmt.Println("Server started at :8080")
	// เริ่มต้นเซิร์ฟเวอร์และฟังที่พอร์ต 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
