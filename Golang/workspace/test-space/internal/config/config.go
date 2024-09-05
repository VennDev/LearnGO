package config

type Config struct {
	Port int
	// Thêm các cấu hình khác nếu cần
}

func Load() (*Config, error) {
	// Đọc cấu hình từ file hoặc biến môi trường
	return &Config{
		Port: 8080,
	}, nil
}
