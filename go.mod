module github.com/cdvelop/timeclient

go 1.20

require (
	github.com/cdvelop/model v0.0.107
	github.com/cdvelop/timetools v0.0.34
)

require github.com/cdvelop/strings v0.0.9 // indirect

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/timetools => ../timetools
