package messagetype

type MessageType string

const (
	Alert  MessageType = "alert"
	Update MessageType = "update"
	Cancel MessageType = "cancel"
)
