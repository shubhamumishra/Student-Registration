package dbaccess

import (
	"gorm.io/gorm"
)

type Student struct {
	ID       int       `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Subjects []Subject `gorm:"foreignKey:StudentID"`
}

type Subject struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	Name      string  `json:"name"`
	Schedule  string  `json:"schedule"`
	Score     float64 `json:"score"`
	StudentID int
}

type SQLStudentRepository struct {
	DB *gorm.DB
}

func (r *SQLStudentRepository) Migrate() error {

	return r.DB.AutoMigrate(&Student{}, &Subject{})
}

func (r *SQLStudentRepository) GetAllStudents() ([]Student, error) {
	var students []Student
	err := r.DB.Find(&students).Error
	return students, err
}

func (r *SQLStudentRepository) GetStudentByID(id int) (Student, error) {
	var student Student
	err := r.DB.Preload("Subjects").First(&student, id).Error
	return student, err
}

func (r *SQLStudentRepository) CreateStudent(student Student) error {
	return r.DB.Create(&student).Error
}

func (r *SQLStudentRepository) UpdateStudent(id int, student Student) error {

	err := r.DB.Where("student_id = ?", id).Delete(&Subject{}).Error
	if err != nil {
		return err
	}

	student.ID = id
	return r.DB.Omit("Subjects").Save(&student).Error
}

func (r *SQLStudentRepository) DeleteStudent(id int) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Where("student_id = ?", id).Delete(&Subject{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&Student{}, id).Error; err != nil {
			return err
		}
		return nil
	})
}
