package models

import (
	"CRUDProject/cfg"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name        string
	Email       string `gorm:"unique"`
	PhoneNumber string
	Classes     []Class `gorm:"many2many:student_classes;"`
}

func CreateStudent(student *Student) error {
	if err := cfg.DB.Create(student).Error; err != nil {
		return err
	}
	return nil
}

func GetStudentByID(id uint) (*Student, error) {
	var student Student
	if err := cfg.DB.Preload("Classes").First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func GetAllStudents() ([]Student, error) {
	var students []Student
	if err := cfg.DB.Preload("Classes").Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func UpdateStudent(id uint, student *Student) error {
	existingStudent, err := GetStudentByID(id)
	if err != nil {
		return err
	}

	// update information
	existingStudent.Name = student.Name
	existingStudent.Email = student.Email
	existingStudent.PhoneNumber = student.PhoneNumber

	if err := cfg.DB.Save(existingStudent).Error; err != nil {
		return err
	}

	return nil
}

func DeleteStudent(id uint) error {
	if err := cfg.DB.Delete(&Student{}, id).Error; err != nil {
		return err
	}
	return nil
}
