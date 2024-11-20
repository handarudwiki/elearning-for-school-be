package response

import "github.com/handarudwiki/models"

type EventResponse struct {
	ID       uint          `json:"id"`
	Title    string        `json:"title"`
	Location string        `json:"location"`
	Body     string        `json:"body"`
	UserID   uint          `json:"user_id"`
	User     *UserResponse `json:"user"`
	Date     string        `json:"date"`
	Time     string        `json:"time"`
}

func ToEventResponse(event *models.Event) *EventResponse {
	var response *UserResponse

	if event.User != nil {
		userResponse := ToUserResponse(event.User)
		response = &userResponse
	}

	return &EventResponse{
		ID:       event.ID,
		Title:    event.Title,
		Location: event.Location,
		Body:     event.Body,
		UserID:   event.UserID,
		User:     response,
		Date:     event.Date.Format("2006-01-02"),
		Time:     event.Time,
	}
}

func ToEventResponseSlice(events []*models.Event) []*EventResponse {
	var eventResponse []*EventResponse

	for _, e := range events {
		event := ToEventResponse(e)

		eventResponse = append(eventResponse, event)
	}

	return eventResponse
}
