package response

import "github.com/handarudwiki/models"

type SubjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToSubjectResponse(subject *models.Subject) SubjectResponse {
	return SubjectResponse{
		ID:          subject.ID,
		Name:        subject.Name,
		Description: subject.Description,
	}
}

func TosSubjectResponseSlice(subjects []*models.Subject) []SubjectResponse {
	var subjectResponses []SubjectResponse
	for _, subject := range subjects {
		subjectResponses = append(subjectResponses, ToSubjectResponse(subject))
	}
	return subjectResponses
}
