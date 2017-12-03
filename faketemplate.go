package fakeexample

type FuncName int

const (
	Func1 FuncName = iota
	Func2
	FUNC_COUNT
)

var (
	CallCount     [FUNC_COUNT]int
	FuncIO        funcIO
	ApiCallRecord []FuncName
)

type FakeFunc1Call struct {
	Receives struct {
		Arg1 int
	}
	Returns struct {
		Return1 int
		Error   error
	}
}

type FakeFunc2Call struct {
	Receives struct {
		Arg1 string
	}
	Returns struct {
		Return1 string
		Error   error
	}
}

type funcIO struct {
	FakeFunc1Calls []FakeFunc1Call
	FakeFunc2Calls []FakeFunc2Call
}

func Init() {
	CallCount = [FUNC_COUNT]int{}
	FuncIO = funcIO{}
	ApiCallRecord = []FuncName{}
}

func SetFake(funcToFake FuncName) {
	switch funcToFake {
	case Func1:
		mypkg.Func1 = FakeFunc1
	case Func2:
		mypkg.Func2 = FakeFunc2
	}
}

func SetReturns(fn FuncName, index int, returns ...interface{}) {

	if CallCount[fn] > 0 {
		panic("cannot set returns after function has been called")
	}

	switch fn {
	case Func1:
		cp := FuncIO.FakeFunc1Calls
		FuncIO.FakeFunc1Calls = make([]FakeFunc1Call, index+1, index+1)
		copy(FuncIO.FakeFunc1Calls, cp)
		FuncIO.FakeFunc1Calls[index].Returns.Return1 = returns[0].(int)
		FuncIO.FakeFunc1Calls[index].Returns.Error = returns[1].(error)
	case Func2:
		cp := FuncIO.FakeFunc2Calls
		FuncIO.FakeFunc2Calls = make([]FakeFunc2Call, index+1, index+1)
		copy(FuncIO.FakeFunc2Calls, cp)
		FuncIO.FakeFunc2Calls[index].Returns.Return1 = returns[0].(string)
		FuncIO.FakeFunc2Calls[index].Returns.Error = returns[1].(error)
	}
}

func FakeFunc1(arg1 int) (int, error) {
	defer increment(Func1)

	if len(FuncIO.FakeFunc1Calls) == CallCount[Func1] {
		FuncIO.FakeFunc1Calls = append(FuncIO.FakeFunc1Calls, FakeFunc1Call{})
	}

	FuncIO.FakeFunc1Calls[CallCount[Func1]].Receives.Arg1 = arg1

	return FuncIO.FakeFunc1Calls[CallCount[Func1]].Returns.Return1, FuncIO.FakeFunc1Calls[CallCount[Func1]].Returns.Error
}

func FakeFunc2(arg1 string) (string, error) {
	defer increment(Func2)

	if len(FuncIO.FakeFunc1Calls) == CallCount[Func2] {
		FuncIO.FakeFunc2Calls = append(FuncIO.FakeFunc2Calls, FakeFunc2Call{})
	}

	FuncIO.FakeFunc2Calls[CallCount[Func2]].Receives.Arg1 = arg1

	return FuncIO.FakeFunc2Calls[CallCount[Func2]].Returns.Return1, FuncIO.FakeFunc2Calls[CallCount[Func2]].Returns.Error
}

func increment(fn FuncName) {
	CallCount[fn]++
	ApiCallRecord = append(ApiCallRecord, fn)
}
