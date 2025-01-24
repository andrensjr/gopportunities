package config

import (
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func InitDB() error {
	return nil
	// return errors.New("fake error")
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
