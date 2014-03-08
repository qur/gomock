package interfaces

type MockController interface {
	Call(receiver interface{}, method string, args ...interface{}) []interface{}
	RecordCall(receiver interface{}, method string, args ...interface{}) Call
}
