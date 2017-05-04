package piCamera

//MeteringType is used for setting the Metering Mode.
type MeteringType int

const (
	//MeterNone tell this package to use whatever the default is
	MeterNone MeteringType = iota
	//MeterAverage average the whole frame for metering
	MeterAverage
	//MeterSpot use spot metering
	MeterSpot
	//MeterBacklit will assume a backlit image
	MeterBacklit
	//MeterMatrix use matrix metering
	MeterMatrix
)
