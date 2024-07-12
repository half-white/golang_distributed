package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "A",
			LastName:  "B",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
			},
		},
		{
			ID:        2,
			FirstName: "X",
			LastName:  "Y",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 82,
				},
				{
					Title: "Final Exam",
					Type:  GradeExam,
					Score: 90,
				},
			},
		},
	}
}
