package response

import "github.com/handarudwiki/models"

type ClassroomSubjectResponse struct {
	ID          uint               `json:"id"`
	ClassroomID uint               `json:"classroom_id"`
	Classroom   *ClassroomResponse `json:"classroom,omitempty"`
	SubjectID   uint               `json:"subject_id"`
	Subject     *SubjectResponse   `json:"subject,omitempty"`
	TeacherID   uint               `json:"teacher_id"`
	Teacher     *UserResponse      `json:"teacher,omitempty"`
}

func ToClassroomSubjectResponse(classroomSubject *models.ClassroomSubject) ClassroomSubjectResponse {
	var classroom *ClassroomResponse
	var subject *SubjectResponse
	var teacher *UserResponse

	if classroomSubject.Classroom != nil {
		classroomResp := ToClassroomResponse(classroomSubject.Classroom)
		classroom = &classroomResp
	}

	if classroomSubject.Subject != nil {
		subjectResp := ToSubjectResponse(classroomSubject.Subject)
		subject = &subjectResp
	}

	if classroomSubject.Teacher != nil {
		userResp := ToUserResponse(classroomSubject.Teacher)
		teacher = &userResp
	}

	return ClassroomSubjectResponse{
		ID:          classroomSubject.ID,
		ClassroomID: classroomSubject.ClassroomID,
		Classroom:   classroom,
		SubjectID:   classroomSubject.SubjectID,
		Subject:     subject,
		TeacherID:   classroomSubject.TeacherID,
		Teacher:     teacher,
	}
}

func ToclassroomSubjectResponseSlice(classroomSubjects []*models.ClassroomSubject) []ClassroomSubjectResponse {
	var classroomSubjectsSlice []ClassroomSubjectResponse
	for _, classroomSubject := range classroomSubjects {
		classroomSubjectsSlice = append(classroomSubjectsSlice, ToClassroomSubjectResponse(classroomSubject))
	}

	return classroomSubjectsSlice
}
