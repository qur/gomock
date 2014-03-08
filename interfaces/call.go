package interfaces

type Call interface {
	After(preReq Call) Call
	AnyTimes() Call
	Do(f interface{}) Call
	Return(rets ...interface{}) Call
	SetArg(n int, value interface{}) Call
	String() string
	Times(n int) Call
}
