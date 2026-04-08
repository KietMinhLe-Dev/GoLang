package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Khai báo cấu trúc các biến giống với .env để load dữ liệu từ .env
type Config struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	PORT        string
}

// Tạo hàm để load config để load dữ liệu từ .env
func LoadConfig() *Config {

	// godotenv.Load() có nghĩa là load file .env
	err := godotenv.Load()

	// Kiểm tra lỗi nếu không load được file .env thì sẽ in ra lỗi
	//nil là không có lỗi
	if err != nil {
		log.Println("Error loading .env file, using default environments", err)
	}

	// &config là con trỏ trỏ đến config	
	return &Config{
		DB_HOST:     getEnv("DB_HOST", "127.0.0.1"),
		DB_PORT:     getEnv("DB_PORT", "3306"),
		DB_USER:     getEnv("DB_USER", "root"),
		DB_PASSWORD: getEnv("DB_PASSWORD", ""),
		DB_NAME:     getEnv("DB_NAME", "go_db"),
		PORT:        getEnv("PORT", "8080"),
	}
}

// Tạo hàm để lấy giá trị từ .env
// key là tên biến trong .env
// fallback là giá trị mặc định nếu không tìm thấy biến trong .env
func getEnv(key, fallback string) string {
	// os.LookupEnv(key) có nghĩa là lấy giá trị từ .env
	// ok là true nếu tìm thấy biến trong .env
	if value, ok := os.LookupEnv(key); ok {
		return value // Trả về giá trị nếu tìm thấy
	}
	return fallback // Trả về giá trị mặc định nếu không tìm thấy
}
