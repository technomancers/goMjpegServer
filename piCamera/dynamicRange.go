package piCamera

//DRCType is to set the dynamic range compression
type DRCType int

const (
	//DRCOff turns of DRC
	DRCOff DRCType = iota
	//DRCLow compress the range slightly
	DRCLow
	//DRCMedium compress the range more
	DRCMedium
	//DRCHigh compress the range even more
	DRCHigh
)
