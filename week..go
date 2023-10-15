package timeclient

import (
	"syscall/js"

	"github.com/cdvelop/timetools"
)

// date ej: 2006-01-02,  0 para Domingo, 1 para Lunes, etc.
func (TimeCLient) WeekDayNumber(date_in string) (int, error) {

	format_date, err := timetools.ChangeDateSeparator(date_in, "/")
	if err != nil {
		return 0, err
	}

	jsDate := js.Global().Get("Date").New(format_date)

	return jsDate.Call("getDay").Int(), nil
}
