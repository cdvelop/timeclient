package timeclient

import (
	"syscall/js"
)

// layout ej: 2006-01-02
func (TimeCLient) ToDay(layout string) string {

	// Obtener la fecha actual en JavaScript
	jsDate := js.Global().Get("Date").New()

	// Obtener la hora en formato HH:MM:SS
	// formattedTime := jsDate.Call("toLocaleTimeString", "en", map[string]interface{}{"hour12": false}).String()

	date := jsDate.Call("toISOString").String()[0:10]
	// Imprimir la hora formateada

	// js.Global().Get("console").Call("log", "hora formateada", date)

	return date
}
