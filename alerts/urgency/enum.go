package urgency

type Urgency string

const (
	Immediate Urgency = "immediate"
	Expected  Urgency = "expected"
	Future    Urgency = "future"
	Past      Urgency = "past"
	Unknown   Urgency = "unknown"
)
