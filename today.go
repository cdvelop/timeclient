package timeclient

import (
	"syscall/js"
)

func (t timeCLient) DateToDay() string {
	return t.date(false)
}

func (t timeCLient) DateToDayHour() string {
	return t.date(true)
}

func (t timeCLient) date(with_hour bool) string {

	var hour string

	if t.current_date != "" {
		return t.current_date
	}

	// Obtener la fecha actual en JavaScript
	jsDate := js.Global().Get("Date").New()

	if with_hour {
		hour = " " + currentHour(&jsDate)
		// js.Global().Get("console").Call("log", "hora formateada", hour)

	}

	// fecha formateada
	date := jsDate.Call("toISOString").String()[0:10]
	// js.Global().Get("console").Call("log", "fecha formateada", date)

	return date + hour
}

// hora en formato 24hras HH:MM:SS ej:15:04:58
func currentHour(date *js.Value) string {
	return date.Call("toLocaleTimeString", "en", map[string]interface{}{"hour12": false}).String()
}
