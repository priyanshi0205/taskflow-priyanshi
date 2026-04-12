package task

var validStatuses = map[string]struct{}{
	"todo":        {},
	"in_progress": {},
	"done":        {},
}

var validPriorities = map[string]struct{}{
	"low":    {},
	"medium": {},
	"high":   {},
}
