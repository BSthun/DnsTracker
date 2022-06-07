package messages

type InboundMessage struct {
	Event   InboundEvent   `json:"event" validate:"required"`
	Payload map[string]any `json:"payload"`
}

type OutboundMessage struct {
	Event   OutboundEvent `json:"event"`
	Payload any           `json:"payload"`
}

type InboundEvent string
type OutboundEvent string
