package mock

import (
	"time"
)

// Start of method calls and parameter capture
var stats = make(map[string]*funcCalls, 0)

type funcCalls struct {
	Count  int
	Params [][]interface{}
}

type CallInfo struct {
	Ok     bool
	Name   string
	Params []interface{}
}

type Params []interface{}

func NumCalls(name string) int {
	call := forCall(name)
	return call.Count
}

func CallParams(name string) []Params {
	call := forCall(name)
	if call.Count > 0 {
		calls := make([]Params, 0)
		for i := 0; i < call.Count; i++ {
			calls = append(calls, call.Params[i])
		}
		return calls
	}
	return []Params{}
}

func capture(key string, params []interface{}) {
	val, ok := stats[key]
	if !ok {
		val = &funcCalls{}
		val.Count = 0
		val.Params = make([][]interface{}, 0)
		stats[key] = val
	}
	val.Count++
	val.Params = append(val.Params, params)

}

func forCall(key string) funcCalls {
	if val, ok := stats[key]; ok {
		return *val
	}
	return funcCalls{}
}

// End of method calls and parameter capture

// Begin of mock for L2 and its methods
type MockL2 struct {
	s    string
	time time.Duration
	Il3  Il3
}

type LM21 func(i1 int, f2 float32) string

var MockL2_LM21 LM21

func (v MockL2) LM21(i1 int, f2 float32) string {
	capture("MockL2_LM21", []interface{}{i1, f2})
	return MockL2_LM21(i1, f2)
}

// End of mock for L2 and its methods

// Begin of mock for L1 and its methods
type MockL1 struct {
	s     string
	S1    string
	time  time.Duration
	Time2 time.Duration
	L2    L2
	Il2   Il2
	PL2   *L2
	DEPS_ interface{}
}

type LM1 func(i1 int, f2 float32) (string, *int)

var MockL1_LM1 LM1

func (v MockL1) LM1(i1 int, f2 float32) (string, *int) {
	capture("MockL1_LM1", []interface{}{i1, f2})
	return MockL1_LM1(i1, f2)
}

type LM2 func(t1 time.Duration, f2 float32) (string, time.Duration)

var MockL1_LM2 LM2

func (p *MockL1) LM2(t1 time.Duration, f2 float32) (string, time.Duration) {
	capture("MockL1_LM2", []interface{}{t1, f2})
	return MockL1_LM2(t1, f2)
}

type LM3 func(pf1 *float32) (string, time.Duration)

var MockL1_LM3 LM3

func (p *MockL1) LM3(pf1 *float32) (string, time.Duration) {
	capture("MockL1_LM3", []interface{}{pf1})
	return MockL1_LM3(pf1)
}

// End of mock for L1 and its methods
