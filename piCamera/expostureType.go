package piCamera

//ExposureType is for setting the exposure mode.
type ExposureType int

const (
	//ExpAuto use automatic exposure mode
	ExpAuto ExposureType = iota
	//ExpNight select setting for night shooting
	ExpNight
	//ExpBacklight select setting for backlit subject
	ExpBacklight
	//ExpSpotlight select setting for spotlit subject
	ExpSpotlight
	//ExpSports select setting for sports
	ExpSports
	//ExpSnow select setting optimised for snowy scenery
	ExpSnow
	//ExpBeach select setting optimised for beach
	ExpBeach
	//ExpVerylong select setting for long exposures
	ExpVerylong
	//ExpFixedfps constrain fps to a fixed value
	ExpFixedfps
	//ExpAntishake turns on antishake mode
	ExpAntishake
	//ExpFireworks select setting optimised for fireworks
	ExpFireworks
)
