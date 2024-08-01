package models

import (
	"CRUDProject/cfg"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	Name      string
	Students  []Student `gorm:"many2many:student_classes;"`
	SubjectID uint      `gorm:"not null;unique"` // Ensure Class has a unique SubjectID
	Subject   Subject   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	TeacherID uint
	Teacher   Teacher
}

func CreateClass(class *Class) error {
	if err := cfg.DB.Create(class).Error; err != nil {
		return err
	}
	return nil
}

func GetClassByID(id uint) (*Class, error) {
	var class Class
	if err := cfg.DB.Preload("Students").Preload("Subject").Preload("Teacher").First(&class, id).Error; err != nil {
		return nil, err
	}
	return &class, nil
}

func GetAllClasses() ([]Class, error) {
	var classes []Class
	if err := cfg.DB.Preload("Students").Preload("Subject").Preload("Teacher").Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func UpdateClass(id uint, updatedClass *Class) error {
	existingClass, err := GetClassByID(id)
	if err != nil {
		return err
	}

	// Update information
	existingClass.Name = updatedClass.Name
	existingClass.SubjectID = updatedClass.SubjectID
	existingClass.TeacherID = updatedClass.TeacherID
	// Optionally, update the students as well
	existingClass.Students = updatedClass.Students

	if err := cfg.DB.Save(existingClass).Error; err != nil {
		return err
	}

	return nil
}

func DeleteClass(id uint) error {
	if err := cfg.DB.Delete(&Class{}, id).Error; err != nil {
		return err
	}
	return nil
}

func RegisterClass(classID uint, studentID uint) error {
	var class Class
	if err := cfg.DB.First(&class, classID).Error; err != nil {
		return err
	}

	var student Student
	if err := cfg.DB.First(&student, studentID).Error; err != nil {
		return err
	}

	// Add the class to the student's classes
	if err := cfg.DB.Model(&student).Association("Classes").Append(&class); err != nil {
		return err
	}

	// Add the student to the class's students
	if err := cfg.DB.Model(&class).Association("Students").Append(&student); err != nil {
		return err
	}

	return nil
}
