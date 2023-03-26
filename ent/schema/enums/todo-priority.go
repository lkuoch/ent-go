package enums

type TodoPriority string

const (
	High   TodoPriority = "High"
	Medium TodoPriority = "Medium"
	Low    TodoPriority = "Low"
	None   TodoPriority = "None"
)

func (TodoPriority) Values() []string {
	return []string{string(High), string(Medium), string(Low), string(None)}
}
