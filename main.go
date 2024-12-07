//go:build js && wasm
// +build js,wasm

package main

import (
	"strconv"
	"syscall/js"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n;
	}
	return fibonacci(n - 1) + fibonacci(n - 2)
}

func calculate(this js.Value, i []js.Value) interface{} {
	value := js.Global().Get("document").Call("getElementById", "fibonacci_input").Get("value").String()
	num, _ := strconv.Atoi(value);
	
	js.Global().Get("document").Call("getElementById", "fibonacci_result").Set("innerText", fibonacci(num));
	return js.ValueOf(0);
}

func registerCallback() {
	js.Global().Set("calculate", js.FuncOf(calculate));
}

func main() {
	c1 := make(chan struct{}, 0);
	registerCallback();
	<- c1;
}