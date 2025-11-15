package gamemap

type HexName int

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
	hexNameCount
)

var hexNameList = map[HexName]string{
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

func GetHexNameList() map[HexName]string {
	return hexNameList
}

type HexagonCoordinates struct {
	x int32
	y int32
	z int32
}

func NewHexagonCoordinatesList() []HexagonCoordinates {
	coordinates := [hexNameCount]HexagonCoordinates{}
	coordinates[NevishLine].x = 0
	coordinates[Stonecradle].x = 0
	coordinates[LinnOfMercy].x = 0
	coordinates[Deadlands].x = 0
	coordinates[DrownedVale].x = 0
	coordinates[AllodsBight].x = 0
	coordinates[ReaversPass].x = 0

	coordinates[CallumsCape].x = 1
	coordinates[TheMoors].x = 1
	coordinates[CallahansPassage].x = 1
	coordinates[MarbanHollow].x = 1
	coordinates[Clahstra].x = 1
	coordinates[EndlessShore].x = 1
	coordinates[TheFingers].x = 1

	coordinates[SpeakingWoods].x = 2
	coordinates[ReachingTrail].x = 2
	coordinates[ViperPit].x = 2
	coordinates[WeatheredExpanse].x = 2
	coordinates[StlicanShelf].x = 2
	coordinates[TempestIsland].x = 2

	coordinates[BasinSionnach].x = 3
	coordinates[HowlCounty].x = 3
	coordinates[ClansheadValley].x = 3
	coordinates[MorgensCrossing].x = 3
	coordinates[Godcrofts].x = 3

	coordinates[Oarbreaker].x = -1
	coordinates[FarranacCoast].x = -1
	coordinates[KingsCage].x = -1
	coordinates[LochMor].x = -1
	coordinates[UmbralWildwood].x = -1
	coordinates[ShackledChasm].x = -1
	coordinates[Terminus].x = -1

	coordinates[FishermansRow].x = -2
	coordinates[Westgate].x = -2
	coordinates[Sableport].x = -2
	coordinates[Heartlands].x = -2
	coordinates[GreatMarch].x = -2
	coordinates[Acrithia].x = -2

	coordinates[StemaLanding].x = -3
	coordinates[Origin].x = -3
	coordinates[AshFields].x = -3
	coordinates[RedRiver].x = -3
	coordinates[Kalokai].x = -3

	coordinates[Origin].z = 0
	coordinates[Sableport].z = 0
	coordinates[LochMor].z = 0
	coordinates[Deadlands].z = 0
	coordinates[MarbanHollow].z = 0
	coordinates[WeatheredExpanse].z = 0
	coordinates[MorgensCrossing].z = 0

	coordinates[StemaLanding].z = 1
	coordinates[Westgate].z = 1
	coordinates[KingsCage].z = 1
	coordinates[LinnOfMercy].z = 1
	coordinates[CallahansPassage].z = 1
	coordinates[ViperPit].z = 1
	coordinates[ClansheadValley].z = 1

	coordinates[FishermansRow].z = 2
	coordinates[FarranacCoast].z = 2
	coordinates[Stonecradle].z = 2
	coordinates[TheMoors].z = 2
	coordinates[ReachingTrail].z = 2
	coordinates[HowlCounty].z = 2

	coordinates[Oarbreaker].z = 3
	coordinates[NevishLine].z = 3
	coordinates[CallumsCape].z = 3
	coordinates[SpeakingWoods].z = 3
	coordinates[BasinSionnach].z = 3

	coordinates[AshFields].z = -1
	coordinates[Heartlands].z = -1
	coordinates[UmbralWildwood].z = -1
	coordinates[DrownedVale].z = -1
	coordinates[Clahstra].z = -1
	coordinates[StlicanShelf].z = -1
	coordinates[Godcrofts].z = -1

	coordinates[RedRiver].z = -2
	coordinates[GreatMarch].z = -2
	coordinates[ShackledChasm].z = -2
	coordinates[AllodsBight].z = -2
	coordinates[EndlessShore].z = -2
	coordinates[TempestIsland].z = -2

	coordinates[Kalokai].z = -3
	coordinates[Acrithia].z = -3
	coordinates[Terminus].z = -3
	coordinates[ReaversPass].z = -3
	coordinates[TheFingers].z = -3

	coordinates[BasinSionnach].y = 0
	coordinates[ReachingTrail].y = 0
	coordinates[CallahansPassage].y = 0
	coordinates[Deadlands].y = 0
	coordinates[UmbralWildwood].y = 0
	coordinates[GreatMarch].y = 0
	coordinates[Kalokai].y = 0

	coordinates[HowlCounty].y = 1
	coordinates[ViperPit].y = 1
	coordinates[MarbanHollow].y = 1
	coordinates[DrownedVale].y = 1
	coordinates[ShackledChasm].y = 1
	coordinates[Acrithia].y = 1

	coordinates[ClansheadValley].y = 2
	coordinates[WeatheredExpanse].y = 2
	coordinates[Clahstra].y = 2
	coordinates[AllodsBight].y = 2
	coordinates[Terminus].y = 2

	coordinates[MorgensCrossing].y = 3
	coordinates[StlicanShelf].y = 3
	coordinates[EndlessShore].y = 3
	coordinates[ReaversPass].y = 3

	coordinates[Godcrofts].y = 4
	coordinates[TempestIsland].y = 4
	coordinates[TheFingers].y = 4

	coordinates[SpeakingWoods].y = -1
	coordinates[TheMoors].y = -1
	coordinates[LinnOfMercy].y = -1
	coordinates[LochMor].y = -1
	coordinates[Heartlands].y = -1
	coordinates[RedRiver].y = -1

	coordinates[CallumsCape].y = -2
	coordinates[Stonecradle].y = -2
	coordinates[KingsCage].y = -2
	coordinates[Sableport].y = -2
	coordinates[AshFields].y = -2

	coordinates[NevishLine].y = -3
	coordinates[FarranacCoast].y = -3
	coordinates[Westgate].y = -3
	coordinates[Origin].y = -3

	coordinates[Oarbreaker].y = -4
	coordinates[FishermansRow].y = -4
	coordinates[StemaLanding].y = -4

	return coordinates[:]
}
