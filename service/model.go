package service

import (
	"student/dbaccess"

	"gorm.io/gorm"
)

type StudentService struct {
	Repository StudentRepository
}

type StudentViewModel struct {
	ID       int                `json:"id"`
	Name     string             `json:"name"`
	Subjects []SubjectViewModel `json:"subjects"`
}

type SubjectViewModel struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Schedule string  `json:"schedule"`
	Score    float64 `json:"score"`
}

type StudentRepository interface {
	GetAllStudents() ([]dbaccess.Student, error)
	GetStudentByID(id int) (dbaccess.Student, error)
	CreateStudent(student dbaccess.Student) error
	UpdateStudent(id int, student dbaccess.Student) error
	DeleteStudent(id int) error
}

type Student struct {
	gorm.Model
	Name     string    `json:"name"`
	Subjects []Subject `gorm:"foreignKey:StudentID"`
}

type Subject struct {
	gorm.Model
	Name      string  `json:"name"`
	Schedule  string  `json:"schedule"`
	Score     float64 `json:"score"`
	StudentID uint   
}
