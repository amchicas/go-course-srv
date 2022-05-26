package domain

type Status int

const (
	Inactive Status = iota
	Active
)

func (s Status) String() (status string) {
	switch s {
	case Inactive:
		status = "Inactive"
	case Active:
		status = "Active"

	}
	return
}
