package status

type Status string

const (
	Actual   Status = "actual"
	Exercise Status = "exercise"
	System   Status = "system"
	Test     Status = "test"
	Draft    Status = "draft"
)
