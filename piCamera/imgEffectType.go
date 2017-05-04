package piCamera

//ImgEffectType is used for setting the image effect to use.
type ImgEffectType int

const (
	//ImfxNone no effect
	ImfxNone ImgEffectType = iota
	//ImfxNegative invert the image colours
	ImfxNegative
	//ImfxSolarise solarise the image
	ImfxSolarise
	//ImfxPosterise posterise the image
	ImfxPosterise
	//ImfxWhiteboard whiteboard effect
	ImfxWhiteboard
	//ImfxBlackboard blackboard effect
	ImfxBlackboard
	//ImfxSketch sketch effect
	ImfxSketch
	//ImfxDenoise denoise the image
	ImfxDenoise
	//ImfxEmboss the image
	ImfxEmboss
	//ImfxOilpaint oil paint effect
	ImfxOilpaint
	//ImfxHatch hatch sektch effect
	ImfxHatch
	//ImfxGPen graphite sketch effect
	ImfxGPen
	//ImfxPastel pastel effect
	ImfxPastel
	//ImfxWatercolour watercolour effect
	ImfxWatercolour
	//ImfxFilm film grain effect
	ImfxFilm
	//ImfxBlur blur the image
	ImfxBlur
	//ImfxSaturation colour saturate the image
	ImfxSaturation
)
