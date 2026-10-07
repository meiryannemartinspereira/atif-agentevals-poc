package atif

type Trajectory struct {
	Steps []Step
}

type Step struct {
	Source      string
	Message     *Message
	Reasoning   string
	ToolCalls   []ToolCall
	Observation *Observation
	Metrics     map[string]float64
	Metadata    map[string]string
}

type Message struct {
	Content string
}

type ToolCall struct {
	ToolCallID   string
	FunctionName string
	Arguments    string
}

type Observation struct {
	Results []ObservationResult
}

type ObservationResult struct {
	SourceCallID string
	Content      string
}
