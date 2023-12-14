package timeclient

import (
	"syscall/js"

	"github.com/cdvelop/timetools"
)

func (t *timeCLient) DateToDay() string {

	t.setDate()

	t.only_date, _ = timetools.DateToDayHour(t.real_date, t.fake_date)

	return t.only_date
}

func (t *timeCLient) DateToDayHour(seconds ...bool) (date, hour string) {

	t.setDate()

	return timetools.DateToDayHour(t.real_date, t.fake_date, seconds...)
}

func (t *timeCLient) setDate() {

	// Obtener la fecha actual en JavaScript
	t.jsDate = js.Global().Get("Date").New()

	// hora en formato 24hras HH:MM:SS ej:15:04:58
	t.hour = t.jsDate.Call("toLocaleTimeString", "en", map[string]interface{}{"hour12": false}).String()
	// js.Global().Get("console").Call("log", "hora formateada", t.hour)

	t.real_date = t.jsDate.Call("toISOString").String()[0:10] + " " + t.hour
	// js.Global().Get("console").Call("log", "fecha formateada", t.real_date)

}
