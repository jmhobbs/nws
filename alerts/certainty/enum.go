package certainty

type Certainty string

const (
	Observed Certainty = "observed"
	Likely   Certainty = "likely"
	Possible Certainty = "possible"
	Unlikely Certainty = "unlikely"
	Unknown  Certainty = "unknown"
)
