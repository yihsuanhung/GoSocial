package db

import (
	"gorm.io/gorm"
)

var SqlInstance *gorm.DB

func Init() {
	// Singleton
	if SqlInstance != nil {
		return
	}
	cfg := DefaultConfig()
	SqlInstance = cfg.Build()
}
