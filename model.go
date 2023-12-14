package timeclient

import "syscall/js"

type timeCLient struct {
	jsDate    js.Value
	real_date string

	fake_date string
	only_date string
	hour      string
}
