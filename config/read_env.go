package config

//
//import (
//	"os"
//)
//
//type DataBaseConfig struct {
//	Name     string
//	Username string
//	Passwd   string
//	Host     string
//	SSLmode  string
//}
//
//type Config struct {
//	GitHub DataBaseConfig
//}
//
//// New returns a new Config struct
//func New() *Config {
//	return &Config{
//		GitHub: DataBaseConfig{
//			Name: getEnv(),
//			Username string
//			Passwd   string
//			Host     string
//			SSLmode  string
//
//			Username: getEnv("GITHUB_USERNAME", ""),
//			APIKey:   getEnv("GITHUB_API_KEY", ""),
//		},
//	}
//}
//
//// Simple helper function to read an environment or return a default value
//func getEnv(key string, defaultVal string) string {
//	if value, exists := os.LookupEnv(key); exists {
//		return value
//	}
//
//	return defaultVal
//}
