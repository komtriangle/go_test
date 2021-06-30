package main

type Response struct {
	Success bool   `json: Success`
	ErrCode string `json: ErrCode`
	Value   int    `json: Value`
}

type TestModel struct {
	A        int
	B        int
	Response Response
}
