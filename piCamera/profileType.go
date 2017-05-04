package piCamera

//ProfileType sets the H264 profile to be used for the encoding.
type ProfileType int

const (
	//ProfileNone tells this package to use whatever the default is.
	ProfileNone ProfileType = iota
	//ProfileBaseline is for the baseline profile
	ProfileBaseline
	//ProfileMain is for the main profile
	ProfileMain
	//ProfileHigh is for a high profile
	ProfileHigh
)
