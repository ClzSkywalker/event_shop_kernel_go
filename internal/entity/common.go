package entity

type CommonResponseId struct {
	OnlyCode string `json:"oc,omitempty"`
}

type CommonRequestId struct {
	OnlyCode string `json:"oc,omitempty" binding:"required" validate:"required"`
}
