package models

import (
	"CRUDProject/cfg"
	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Name string
}

func CreateSubject(subject *Subject) error {
	if err := cfg.DB.Create(subject).Error; err != nil {
		return err
	}
	return nil
}

func GetSubjectByID(id uint) (*Subject, error) {
	var subject Subject
	if err := cfg.DB.First(&subject, id).Error; err != nil {
		return nil, err
	}
	return &subject, nil
}

func GetAllSubjects() ([]Subject, error) {
	var subjects []Subject
	if err := cfg.DB.Find(&subjects).Error; err != nil {
		return nil, err
	}
	return subjects, nil
}

func UpdateSubjectByID(id uint, subject *Subject) error {
	existSubject, err := GetSubjectByID(id)
	if err != nil {
		return err
	}

	existSubject.Name = subject.Name
	if err := cfg.DB.Save(existSubject).Error; err != nil {
		return err
	}
	return nil
}

func DeleteSubject(id uint) error {
	if err := cfg.DB.Delete(&Subject{}, id).Error; err != nil {
		return err
	}
	return nil
}
