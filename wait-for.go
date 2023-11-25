package timeclient

import "syscall/js"

// WaitFor espera el número especificado de milisegundos y luego ejecuta la función de retorno de llamada.
func (timeCLient) WaitFor(milliseconds int, callback func()) {
	js.Global().Call("setTimeout", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
		callback() // Llamar a la función de retorno de llamada después de esperar
		return nil
	}), milliseconds)
}
