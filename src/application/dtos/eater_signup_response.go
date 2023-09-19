package dtos

type EaterSignupResponse struct {
	EaterID string `json:"eater_id"`
}

func NewEaterSignupResponse(eaterID string) *EaterSignupResponse {
	return &EaterSignupResponse{
		EaterID: eaterID,
	}
}
