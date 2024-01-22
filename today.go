package timeclient

import (
	"syscall/js"

	"github.com/cdvelop/timetools"
)

func (t *timeCLient) DateToDay(dateFormatStructPointer any) string {
	t.setDate()
	t.only_date, _ = timetools.DateToDayHour(t.real_date, t.fake_date, dateFormatStructPointer)
	return t.only_date
}

func (t *timeCLient) DateToDayHour(dateFormatStructPointer any) (date, hour string) {
	t.setDate()
	return timetools.DateToDayHour(t.real_date, t.fake_date, dateFormatStructPointer)
}

func (t *timeCLient) setDate() {
	t.real_date = js.Global().Call("currentDate").String()
	// js.Global().Get("console").Call("log", "fecha formateada", t.real_date)
}
