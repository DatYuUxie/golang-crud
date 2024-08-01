package models

import (
	"CRUDProject/cfg"
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name        string
	Email       string `gorm:"unique"`
	PhoneNumber string
	Classes     []Class
}

func CreateTeacher(teacher *Teacher) error {
	if err := cfg.DB.Create(teacher).Error; err != nil {
		return err
	}
	return nil
}

func GetTeacherByID(id uint) (*Teacher, error) {
	var teacher Teacher
	if err := cfg.DB.Preload("Classes").First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func GetAllTeachers() ([]Teacher, error) {
	var teachers []Teacher
	if err := cfg.DB.Preload("Classes").Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func UpdateTeacher(id uint, teacher *Teacher) error {
	existingTeacher, err := GetTeacherByID(id)
	if err != nil {
		return err
	}

	existingTeacher.Name = teacher.Name
	existingTeacher.Email = teacher.Email
	existingTeacher.PhoneNumber = teacher.PhoneNumber

	if err := cfg.DB.Save(existingTeacher).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTeacher(id uint) error {
	if err := cfg.DB.Delete(&Teacher{}, id).Error; err != nil {
		return err
	}
	return nil
}
