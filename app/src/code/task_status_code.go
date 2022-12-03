package code

type taskStatusCode struct {
	name string
}

func (code *taskStatusCode) Name() string {
	return code.name
}

var (
	NO_PROCESSING = taskStatusCode{name: "NO_PROCESSING"}
	PROCESSING    = taskStatusCode{name: "PROCESSING"}
	DONE          = taskStatusCode{name: "DONE"}
)

func TaskStatusCodeValues() []taskStatusCode {
	values := make([]taskStatusCode, 0)
	values = append(values, NO_PROCESSING)
	values = append(values, PROCESSING)
	values = append(values, DONE)
	return values
}

func NameToTaskStatusCode(name string) *taskStatusCode {
	values := TaskStatusCodeValues()
	for _, v := range values {
		if v.name == name {
			return &v
		}
	}
	return nil
}
