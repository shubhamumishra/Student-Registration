package service

import "student/dbaccess"

func (s *StudentService) GetAllStudents() ([]StudentViewModel, error) {
	students, err := s.Repository.GetAllStudents()
	if err != nil {
		return nil, err
	}

	viewModels := make([]StudentViewModel, len(students))
	for i, student := range students {
		viewModels[i] = ToViewModel(&student)
	}

	return viewModels, nil
}

func (s *StudentService) GetStudentByID(id int) (StudentViewModel, error) {
	student, err := s.Repository.GetStudentByID(id)
	if err != nil {
		return StudentViewModel{}, err
	}

	return ToViewModel(&student), nil
}

func (s *StudentService) CreateStudent(student StudentViewModel) error {
	subjects := make([]dbaccess.Subject, len(student.Subjects))
	for i, subject := range student.Subjects {
		subjects[i] = dbaccess.Subject{
			Name:     subject.Name,
			Schedule: subject.Schedule,
			Score:    subject.Score,
		}
	}

	return s.Repository.CreateStudent(dbaccess.Student{
		Name:     student.Name,
		Subjects: subjects,
	})
}

func (s *StudentService) UpdateStudent(id int, student StudentViewModel) error {
	subjects := make([]dbaccess.Subject, len(student.Subjects))
	for i, subject := range student.Subjects {
		subjects[i] = dbaccess.Subject{
			ID:       subject.ID,
			Name:     subject.Name,
			Schedule: subject.Schedule,
			Score:    subject.Score,
		}
	}

	return s.Repository.UpdateStudent(id, dbaccess.Student{
		Name:     student.Name,
		Subjects: subjects,
	})
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.Repository.DeleteStudent(id)
}

func ToViewModel(s *dbaccess.Student) StudentViewModel {
	subjectViewModels := make([]SubjectViewModel, len(s.Subjects))
	for i, subject := range s.Subjects {
		subjectViewModels[i] = SubjectViewModel{
			ID:       subject.ID,
			Name:     subject.Name,
			Schedule: subject.Schedule,
			Score:    subject.Score,
		}
	}

	return StudentViewModel{
		ID:       s.ID,
		Name:     s.Name,
		Subjects: subjectViewModels,
	}
}
