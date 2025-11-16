package gamemap

type HexName int32

const (
	Acrithia HexName = iota
	AllodsBight
	AshFields
	BasinSionnach
	CallahansPassage
	CallumsCape
	Clahstra
	ClansheadValley
	Deadlands
	DrownedVale
	EndlessShore
	FarranacCoast
	FishermansRow
	Godcrofts
	GreatMarch
	Heartlands
	HowlCounty
	Kalokai
	KingsCage
	LinnOfMercy
	LochMor
	MarbanHollow
	TheMoors
	MorgensCrossing
	NevishLine
	Oarbreaker
	Origin
	ReachingTrail
	ReaversPass
	RedRiver
	Sableport
	ShackledChasm
	SpeakingWoods
	StemaLanding
	StlicanShelf
	Stonecradle
	TempestIsland
	Terminus
	TheFingers
	UmbralWildwood
	ViperPit
	WeatheredExpanse
	Westgate
	HexNameCount
)

var HexNameList = map[HexName]string{
	Acrithia:         "Acrithia",
	AllodsBight:      "AllodsBight",
	AshFields:        "AshFields",
	BasinSionnach:    "BasinSionnach",
	CallahansPassage: "CallahansPassage",
	CallumsCape:      "CallumsCape",
	Clahstra:         "Clahstra",
	ClansheadValley:  "ClansheadValley",
	Deadlands:        "Deadlands",
	DrownedVale:      "DrownedVale",
	EndlessShore:     "EndlessShore",
	FarranacCoast:    "FarranacCoast",
	FishermansRow:    "FishermansRow",
	Godcrofts:        "Godcrofts",
	GreatMarch:       "GreatMarch",
	Heartlands:       "Heartlands",
	HowlCounty:       "HowlCounty",
	Kalokai:          "Kalokai",
	KingsCage:        "KingsCage",
	LinnOfMercy:      "LinnOfMercy",
	LochMor:          "LochMor",
	MarbanHollow:     "MarbanHollow",
	TheMoors:         "TheMoors",
	MorgensCrossing:  "MorgensCrossing",
	NevishLine:       "NevishLine",
	Oarbreaker:       "Oarbreaker",
	Origin:           "Origin",
	ReachingTrail:    "ReachingTrail",
	ReaversPass:      "ReaversPass",
	RedRiver:         "RedRiver",
	Sableport:        "Sableport",
	ShackledChasm:    "ShackledChasm",
	SpeakingWoods:    "SpeakingWoods",
	StemaLanding:     "StemaLanding",
	StlicanShelf:     "StlicanShelf",
	Stonecradle:      "Stonecradle",
	TempestIsland:    "TempestIsland",
	Terminus:         "Terminus",
	TheFingers:       "TheFingers",
	UmbralWildwood:   "UmbralWildwood",
	ViperPit:         "ViperPit",
	WeatheredExpanse: "WeatheredExpanse",
	Westgate:         "Westgate",
}

type HexagonCoordinates struct {
	x int32
	y int32
}

func (hex *HexagonCoordinates) GetX() int32 {
	return hex.x
}

func (hex *HexagonCoordinates) GetY() int32 {
	return hex.y
}

func NewHexagonCoordinatesList() []HexagonCoordinates {
	coordinates := [HexNameCount]HexagonCoordinates{}
	coordinates[Acrithia] = HexagonCoordinates{
		x: 1,
		y: -3,
	}
	coordinates[AllodsBight] = HexagonCoordinates{
		x: 2,
		y: -2,
	}
	coordinates[AshFields] = HexagonCoordinates{
		x: -2,
		y: -1,
	}
	coordinates[BasinSionnach] = HexagonCoordinates{
		x: 0,
		y: 3,
	}
	coordinates[CallahansPassage] = HexagonCoordinates{
		x: 0,
		y: 1,
	}
	coordinates[CallumsCape] = HexagonCoordinates{
		x: -2,
		y: 3,
	}
	coordinates[Clahstra] = HexagonCoordinates{
		x: 2,
		y: -1,
	}
	coordinates[ClansheadValley] = HexagonCoordinates{
		x: 2,
		y: 1,
	}
	coordinates[Deadlands] = HexagonCoordinates{
		x: 0,
		y: 0,
	}
	coordinates[DrownedVale] = HexagonCoordinates{
		x: 1,
		y: -1,
	}
	coordinates[EndlessShore] = HexagonCoordinates{
		x: 3,
		y: -2,
	}
	coordinates[FarranacCoast] = HexagonCoordinates{
		x: -3,
		y: 2,
	}
	coordinates[FishermansRow] = HexagonCoordinates{
		x: -4,
		y: 2,
	}
	coordinates[Godcrofts] = HexagonCoordinates{
		x: 4,
		y: -1,
	}
	coordinates[GreatMarch] = HexagonCoordinates{
		x: 0,
		y: -2,
	}
	coordinates[Heartlands] = HexagonCoordinates{
		x: -1,
		y: -1,
	}
	coordinates[HowlCounty] = HexagonCoordinates{
		x: 1,
		y: 2,
	}
	coordinates[Kalokai] = HexagonCoordinates{
		x: 0,
		y: -3,
	}
	coordinates[KingsCage] = HexagonCoordinates{
		x: -2,
		y: 1,
	}
	coordinates[LinnOfMercy] = HexagonCoordinates{
		x: -1,
		y: 1,
	}
	coordinates[LochMor] = HexagonCoordinates{
		x: -1,
		y: 0,
	}
	coordinates[MarbanHollow] = HexagonCoordinates{
		x: 1,
		y: 0,
	}
	coordinates[TheMoors] = HexagonCoordinates{
		x: -1,
		y: 2,
	}
	coordinates[MorgensCrossing] = HexagonCoordinates{
		x: 3,
		y: 0,
	}
	coordinates[NevishLine] = HexagonCoordinates{
		x: -3,
		y: 3,
	}
	coordinates[Oarbreaker] = HexagonCoordinates{
		x: -4,
		y: 3,
	}
	coordinates[Origin] = HexagonCoordinates{
		x: -3,
		y: 0,
	}
	coordinates[ReachingTrail] = HexagonCoordinates{
		x: 0,
		y: 2,
	}
	coordinates[ReaversPass] = HexagonCoordinates{
		x: 3,
		y: -3,
	}
	coordinates[RedRiver] = HexagonCoordinates{
		x: -1,
		y: -2,
	}
	coordinates[Sableport] = HexagonCoordinates{
		x: -2,
		y: 0,
	}
	coordinates[ShackledChasm] = HexagonCoordinates{
		x: 1,
		y: -2,
	}
	coordinates[SpeakingWoods] = HexagonCoordinates{
		x: -1,
		y: 3,
	}
	coordinates[StemaLanding] = HexagonCoordinates{
		x: -4,
		y: 1,
	}
	coordinates[StlicanShelf] = HexagonCoordinates{
		x: 3,
		y: -1,
	}
	coordinates[Stonecradle] = HexagonCoordinates{
		x: -2,
		y: 2,
	}
	coordinates[TempestIsland] = HexagonCoordinates{
		x: 4,
		y: -2,
	}
	coordinates[Terminus] = HexagonCoordinates{
		x: 2,
		y: -3,
	}
	coordinates[TheFingers] = HexagonCoordinates{
		x: 4,
		y: -3,
	}
	coordinates[UmbralWildwood] = HexagonCoordinates{
		x: 0,
		y: -1,
	}
	coordinates[ViperPit] = HexagonCoordinates{
		x: 1,
		y: 1,
	}
	coordinates[WeatheredExpanse] = HexagonCoordinates{
		x: 2,
		y: 0,
	}
	coordinates[Westgate] = HexagonCoordinates{
		x: -3,
		y: 1,
	}

	return coordinates[:]
}
