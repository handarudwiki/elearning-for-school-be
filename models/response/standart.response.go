package response

import "github.com/handarudwiki/models"

type StandartResponse struct {
	ID         uint   `json:"id"`
	TeacherID  uint   `json:"teacher_id"`
	StandertID uint   `json:"standart_id"`
	SubjectID  uint   `json:"subject_id"`
	Type       string `json:"type"`
	Body       string `json:"body"`
	Code       string `json:"code"`
}

func ToStandartResponse(standart *models.Standart) *StandartResponse {
	return &StandartResponse{
		ID:         standart.ID,
		TeacherID:  standart.TeacherID,
		StandertID: standart.StandartID,
		SubjectID:  standart.SubjectID,
		Type:       standart.Type,
		Body:       standart.Body,
		Code:       standart.Code,
	}
}

func ToStandartResponseSlice(standarts []*models.Standart) []*StandartResponse {
	var standartResponse []*StandartResponse

	for _, s := range standarts {
		standart := ToStandartResponse(s)

		standartResponse = append(standartResponse, standart)
	}

	return standartResponse
}
