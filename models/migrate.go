package models

import "CRUDProject/cfg"

func Migrate() {
	cfg.DB.AutoMigrate(&Student{}, &Class{}, &Subject{}, &Teacher{})
}
