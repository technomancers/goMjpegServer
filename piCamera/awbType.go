package piCamera

//AWBType is for setting the Automatic White Balance setting.
type AWBType int

const (
	//AwbAuto is for automatic white balance
	AwbAuto AWBType = iota
	//AwbOff is for turning automatic white balance off
	AwbOff
	//AwbSun is for sunny mode
	AwbSun
	//AwbCloud is for cloudy mode
	AwbCloud
	//AwbShade is for shade mode
	AwbShade
	//AwbTungsten tungsten lighting mode
	AwbTungsten
	//AwbFluorescent fluorescent lighting mode
	AwbFluorescent
	//AwbIncandescent incandescent lighting mode
	AwbIncandescent
	//AwbFlash flash mode
	AwbFlash
	//AwbHorizon horizon mode
	AwbHorizon
)

//AWBGains sets the blue and red gains to be applied when AWBOff is set.
type AWBGains struct {
	b float32
	r float32
}

//NewAWBGains creates gains to set to Red and Blue.
//Values are multiplied to the values so 1.5 is 150% and .5 is 50%.
func NewAWBGains(blue, red float32) *AWBGains {
	return &AWBGains{
		b: blue,
		r: red,
	}
}
