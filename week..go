package timeclient

import "syscall/js"

// date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
func (TimeCLient) WeekDayNumber(date_in string) (int, error) {
	jsDate := js.Global().Get("Date").New(date_in)
	return jsDate.Call("getDay").Int(), nil
}
