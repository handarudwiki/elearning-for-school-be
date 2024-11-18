package response

import "github.com/handarudwiki/models"

type ClassroomResponse struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	TeacherID uint         `json:"teacher_id"`
	Teacher   *models.User `json:"teacher"`
	Grade     string       `json:"grade"`
	Group     string       `json:"group"`
	Settings  string       `json:"settings"`
}

func ToClassroomResponse(classroom *models.Classroom) ClassroomResponse {
	return ClassroomResponse{
		ID:        classroom.ID,
		Name:      classroom.Name,
		TeacherID: classroom.TeacherID,
		Grade:     classroom.Grade,
		Group:     classroom.Group,
		Settings:  classroom.Settings,
		Teacher:   classroom.Teacher,
	}
}

func ToClassroomResponseSlice(classrooms []*models.Classroom) []ClassroomResponse {
	var classroomResponses []ClassroomResponse
	for _, classroom := range classrooms {
		classroomResponses = append(classroomResponses, ToClassroomResponse(classroom))
	}
	return classroomResponses
}
