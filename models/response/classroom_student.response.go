package response

import "github.com/handarudwiki/models"

type ClassroomStudentResponse struct {
	ID          uint               `json:"id"`
	StudentID   uint               `json:"student_id"`
	ClassroomID uint               `json:"classroom_id"`
	Student     *UserResponse      `json:"student"`
	Classroom   *ClassroomResponse `json:"classroom"`
}

func ToClassroomStudentResponse(classroomStudent *models.ClassroomStudent) ClassroomStudentResponse {

	var clasroom ClassroomResponse
	var student UserResponse

	if classroomStudent.Classroom != nil {
		clasroom = ToClassroomResponse(classroomStudent.Classroom)
	}

	if classroomStudent.Student != nil {
		student = ToUserResponse(classroomStudent.Student)
	}

	return ClassroomStudentResponse{
		ID:          classroomStudent.ID,
		StudentID:   classroomStudent.StudentID,
		ClassroomID: classroomStudent.ClassroomID,
		Student:     &student,
		Classroom:   &clasroom,
	}
}

func ToClassroomStudentResponseSlice(classroomStudents []*models.ClassroomStudent) []ClassroomStudentResponse {
	var classroomStudentResponses []ClassroomStudentResponse
	for _, classroomStudent := range classroomStudents {
		classroomStudentResponses = append(classroomStudentResponses, ToClassroomStudentResponse(classroomStudent))
	}
	return classroomStudentResponses
}
