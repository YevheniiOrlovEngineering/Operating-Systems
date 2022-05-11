package process

const (
	PNumMax     int = 10
	PArrTimeMax int = 5
	PBurTimeMax int = 10
)

type Process struct {
	Id string
	// Arrival time
	At int
	// Burst time
	Bt int
	// Start time
	St int
	// Finish time
	Ft int
	// Waiting time
	Wt int
	// Turn Around time
	Tat int
}
