package implementations

type MockLogger struct {
	InfoCalled   bool
	InfoMessage  string
	InfoData     any
	ErrorCalled  bool
	ErrorMessage string
	ErrorData    any
}

func (m *MockLogger) Info(message string, data interface{}) {
	m.InfoCalled = true
	m.InfoMessage = message
	m.InfoData = data
}

func (m *MockLogger) Error(message string, data interface{}) {
	m.ErrorCalled = true
	m.ErrorMessage = message
	m.ErrorData = data
}
