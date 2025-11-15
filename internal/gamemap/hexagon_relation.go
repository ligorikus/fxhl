package gamemap

type HexagonRelation struct {
	hex         HexName
	top         *HexagonRelation
	topRight    *HexagonRelation
	topLeft     *HexagonRelation
	bottom      *HexagonRelation
	bottomRight *HexagonRelation
	bottomLeft  *HexagonRelation
}

func (hr *HexagonRelation) GetHexName() HexName {
	return hr.hex
}

func (hr *HexagonRelation) GetTop() *HexagonRelation {
	return hr.top
}

func (hr *HexagonRelation) GetTopRight() *HexagonRelation {
	return hr.topRight
}

func (hr *HexagonRelation) GetTopLeft() *HexagonRelation {
	return hr.topLeft
}

func (hr *HexagonRelation) GetBottom() *HexagonRelation {
	return hr.bottom
}

func (hr *HexagonRelation) GetBottomRight() *HexagonRelation {
	return hr.bottomRight
}

func (hr *HexagonRelation) GetBottomLeft() *HexagonRelation {
	return hr.bottomLeft
}

func NewHexagonRelationList() []HexagonRelation {
	relations := [hexNameCount]HexagonRelation{}
	for i := range hexNameList {
		relations[i] = HexagonRelation{
			hex: i,
		}
	}

	relations[BasinSionnach].bottomLeft = &relations[SpeakingWoods]
	relations[BasinSionnach].bottom = &relations[ReachingTrail]
	relations[BasinSionnach].bottomRight = &relations[HowlCounty]

	relations[SpeakingWoods].bottomLeft = &relations[CallumsCape]
	relations[SpeakingWoods].bottom = &relations[TheMoors]
	relations[SpeakingWoods].bottomRight = &relations[ReachingTrail]
	relations[SpeakingWoods].topRight = &relations[BasinSionnach]

	relations[HowlCounty].bottomLeft = &relations[ReachingTrail]
	relations[HowlCounty].bottom = &relations[ViperPit]
	relations[HowlCounty].bottomRight = &relations[ClansheadValley]
	relations[HowlCounty].topLeft = &relations[BasinSionnach]

	relations[CallumsCape].bottomLeft = &relations[NevishLine]
	relations[CallumsCape].bottom = &relations[Stonecradle]
	relations[CallumsCape].bottomRight = &relations[TheMoors]
	relations[CallumsCape].topRight = &relations[SpeakingWoods]

	relations[ReachingTrail].bottomLeft = &relations[TheMoors]
	relations[ReachingTrail].bottom = &relations[CallahansPassage]
	relations[ReachingTrail].bottomRight = &relations[ViperPit]
	relations[ReachingTrail].topRight = &relations[HowlCounty]
	relations[ReachingTrail].top = &relations[BasinSionnach]
	relations[ReachingTrail].topLeft = &relations[SpeakingWoods]

	relations[ClansheadValley].bottomLeft = &relations[ViperPit]
	relations[ClansheadValley].bottom = &relations[WeatheredExpanse]
	relations[ClansheadValley].bottomRight = &relations[MorgensCrossing]
	relations[ClansheadValley].topLeft = &relations[HowlCounty]

	relations[NevishLine].bottomLeft = &relations[Oarbreaker]
	relations[NevishLine].bottom = &relations[FarranacCoast]
	relations[NevishLine].bottomRight = &relations[Stonecradle]
	relations[NevishLine].topRight = &relations[CallumsCape]

	relations[TheMoors].bottomLeft = &relations[Stonecradle]
	relations[TheMoors].bottom = &relations[LinnOfMercy]
	relations[TheMoors].bottomRight = &relations[CallahansPassage]
	relations[TheMoors].topRight = &relations[ReachingTrail]
	relations[TheMoors].top = &relations[SpeakingWoods]
	relations[TheMoors].topLeft = &relations[CallumsCape]

	relations[ViperPit].bottomLeft = &relations[CallahansPassage]
	relations[ViperPit].bottom = &relations[MarbanHollow]
	relations[ViperPit].bottomRight = &relations[WeatheredExpanse]
	relations[ViperPit].topRight = &relations[ClansheadValley]
	relations[ViperPit].top = &relations[HowlCounty]
	relations[ViperPit].topLeft = &relations[ReachingTrail]

	relations[MorgensCrossing].bottomLeft = &relations[WeatheredExpanse]
	relations[MorgensCrossing].bottom = &relations[StlicanShelf]
	relations[MorgensCrossing].bottomRight = &relations[Godcrofts]
	relations[MorgensCrossing].topLeft = &relations[ClansheadValley]

	relations[Oarbreaker].bottom = &relations[FishermansRow]
	relations[Oarbreaker].bottomRight = &relations[FarranacCoast]
	relations[Oarbreaker].topRight = &relations[NevishLine]

	relations[Stonecradle].bottomLeft = &relations[FarranacCoast]
	relations[Stonecradle].bottom = &relations[KingsCage]
	relations[Stonecradle].bottomRight = &relations[LinnOfMercy]
	relations[Stonecradle].topRight = &relations[TheMoors]
	relations[Stonecradle].top = &relations[CallumsCape]
	relations[Stonecradle].topLeft = &relations[NevishLine]

	relations[CallahansPassage].bottomLeft = &relations[LinnOfMercy]
	relations[CallahansPassage].bottom = &relations[Deadlands]
	relations[CallahansPassage].bottomRight = &relations[MarbanHollow]
	relations[CallahansPassage].topRight = &relations[ViperPit]
	relations[CallahansPassage].top = &relations[ReachingTrail]
	relations[CallahansPassage].topLeft = &relations[TheMoors]

	relations[WeatheredExpanse].bottomLeft = &relations[MarbanHollow]
	relations[WeatheredExpanse].bottom = &relations[Clahstra]
	relations[WeatheredExpanse].bottomRight = &relations[StlicanShelf]
	relations[WeatheredExpanse].topRight = &relations[MorgensCrossing]
	relations[WeatheredExpanse].top = &relations[ClansheadValley]
	relations[WeatheredExpanse].topLeft = &relations[ViperPit]

	relations[Godcrofts].bottomLeft = &relations[StlicanShelf]
	relations[Godcrofts].bottom = &relations[TempestIsland]
	relations[Godcrofts].topLeft = &relations[MorgensCrossing]

	relations[FarranacCoast].bottomLeft = &relations[FishermansRow]
	relations[FarranacCoast].bottom = &relations[Westgate]
	relations[FarranacCoast].bottomRight = &relations[KingsCage]
	relations[FarranacCoast].topRight = &relations[Stonecradle]
	relations[FarranacCoast].top = &relations[NevishLine]
	relations[FarranacCoast].topLeft = &relations[Oarbreaker]

	relations[LinnOfMercy].bottomLeft = &relations[KingsCage]
	relations[LinnOfMercy].bottom = &relations[LochMor]
	relations[LinnOfMercy].bottomRight = &relations[Deadlands]
	relations[LinnOfMercy].topRight = &relations[CallahansPassage]
	relations[LinnOfMercy].top = &relations[TheMoors]
	relations[LinnOfMercy].topLeft = &relations[Stonecradle]

	relations[MarbanHollow].bottomLeft = &relations[Deadlands]
	relations[MarbanHollow].bottom = &relations[DrownedVale]
	relations[MarbanHollow].bottomRight = &relations[Clahstra]
	relations[MarbanHollow].topRight = &relations[WeatheredExpanse]
	relations[MarbanHollow].top = &relations[ViperPit]
	relations[MarbanHollow].topLeft = &relations[CallahansPassage]

	relations[StlicanShelf].bottomLeft = &relations[Clahstra]
	relations[StlicanShelf].bottom = &relations[EndlessShore]
	relations[StlicanShelf].bottomRight = &relations[TempestIsland]
	relations[StlicanShelf].topRight = &relations[Godcrofts]
	relations[StlicanShelf].top = &relations[MorgensCrossing]
	relations[StlicanShelf].topLeft = &relations[WeatheredExpanse]

	relations[FishermansRow].bottom = &relations[StemaLanding]
	relations[FishermansRow].bottomRight = &relations[Westgate]
	relations[FishermansRow].topRight = &relations[FarranacCoast]
	relations[FishermansRow].top = &relations[Oarbreaker]

	relations[KingsCage].bottomLeft = &relations[Westgate]
	relations[KingsCage].bottom = &relations[Sableport]
	relations[KingsCage].bottomRight = &relations[LochMor]
	relations[KingsCage].topRight = &relations[LinnOfMercy]
	relations[KingsCage].top = &relations[Stonecradle]
	relations[KingsCage].topLeft = &relations[FarranacCoast]

	relations[Deadlands].bottomLeft = &relations[LochMor]
	relations[Deadlands].bottom = &relations[UmbralWildwood]
	relations[Deadlands].bottomRight = &relations[DrownedVale]
	relations[Deadlands].topRight = &relations[MarbanHollow]
	relations[Deadlands].top = &relations[CallahansPassage]
	relations[Deadlands].topLeft = &relations[LinnOfMercy]

	relations[Clahstra].bottomLeft = &relations[DrownedVale]
	relations[Clahstra].bottom = &relations[AllodsBight]
	relations[Clahstra].bottomRight = &relations[EndlessShore]
	relations[Clahstra].topRight = &relations[StlicanShelf]
	relations[Clahstra].top = &relations[WeatheredExpanse]
	relations[Clahstra].topLeft = &relations[MarbanHollow]

	relations[TempestIsland].bottomLeft = &relations[EndlessShore]
	relations[TempestIsland].bottom = &relations[TheFingers]
	relations[TempestIsland].top = &relations[Godcrofts]
	relations[TempestIsland].topLeft = &relations[StlicanShelf]

	relations[Westgate].bottomLeft = &relations[StemaLanding]
	relations[Westgate].bottom = &relations[Origin]
	relations[Westgate].bottomRight = &relations[Sableport]
	relations[Westgate].topRight = &relations[KingsCage]
	relations[Westgate].top = &relations[FarranacCoast]
	relations[Westgate].topLeft = &relations[FishermansRow]

	relations[LochMor].bottomLeft = &relations[Sableport]
	relations[LochMor].bottom = &relations[Heartlands]
	relations[LochMor].bottomRight = &relations[UmbralWildwood]
	relations[LochMor].topRight = &relations[Deadlands]
	relations[LochMor].top = &relations[LinnOfMercy]
	relations[LochMor].topLeft = &relations[KingsCage]

	relations[DrownedVale].bottomLeft = &relations[UmbralWildwood]
	relations[DrownedVale].bottom = &relations[ShackledChasm]
	relations[DrownedVale].bottomRight = &relations[AllodsBight]
	relations[DrownedVale].topRight = &relations[Clahstra]
	relations[DrownedVale].top = &relations[MarbanHollow]
	relations[DrownedVale].topLeft = &relations[Deadlands]

	relations[EndlessShore].bottomLeft = &relations[AllodsBight]
	relations[EndlessShore].bottom = &relations[ReaversPass]
	relations[EndlessShore].bottomRight = &relations[TheFingers]
	relations[EndlessShore].topRight = &relations[TempestIsland]
	relations[EndlessShore].top = &relations[StlicanShelf]
	relations[EndlessShore].topLeft = &relations[Clahstra]

	relations[StemaLanding].bottomRight = &relations[Origin]
	relations[StemaLanding].topRight = &relations[Westgate]
	relations[StemaLanding].top = &relations[FishermansRow]

	relations[Sableport].bottomLeft = &relations[Origin]
	relations[Sableport].bottom = &relations[AshFields]
	relations[Sableport].bottomRight = &relations[Heartlands]
	relations[Sableport].topRight = &relations[LochMor]
	relations[Sableport].top = &relations[KingsCage]
	relations[Sableport].topLeft = &relations[Westgate]

	relations[UmbralWildwood].bottomLeft = &relations[Heartlands]
	relations[UmbralWildwood].bottom = &relations[GreatMarch]
	relations[UmbralWildwood].bottomRight = &relations[ShackledChasm]
	relations[UmbralWildwood].topRight = &relations[DrownedVale]
	relations[UmbralWildwood].top = &relations[Deadlands]
	relations[UmbralWildwood].topLeft = &relations[LochMor]

	relations[AllodsBight].bottomLeft = &relations[ShackledChasm]
	relations[AllodsBight].bottom = &relations[Terminus]
	relations[AllodsBight].bottomRight = &relations[ReaversPass]
	relations[AllodsBight].topRight = &relations[EndlessShore]
	relations[AllodsBight].top = &relations[Clahstra]
	relations[AllodsBight].topLeft = &relations[DrownedVale]

	relations[TheFingers].bottomLeft = &relations[ReaversPass]
	relations[TheFingers].top = &relations[TempestIsland]
	relations[TheFingers].topLeft = &relations[EndlessShore]

	relations[Origin].bottomRight = &relations[AshFields]
	relations[Origin].topRight = &relations[Sableport]
	relations[Origin].top = &relations[Westgate]
	relations[Origin].topLeft = &relations[StemaLanding]

	relations[Heartlands].bottomLeft = &relations[AshFields]
	relations[Heartlands].bottom = &relations[RedRiver]
	relations[Heartlands].bottomRight = &relations[GreatMarch]
	relations[Heartlands].topRight = &relations[UmbralWildwood]
	relations[Heartlands].top = &relations[LochMor]
	relations[Heartlands].topLeft = &relations[Sableport]

	relations[ShackledChasm].bottomLeft = &relations[GreatMarch]
	relations[ShackledChasm].bottom = &relations[Acrithia]
	relations[ShackledChasm].bottomRight = &relations[Terminus]
	relations[ShackledChasm].topRight = &relations[AllodsBight]
	relations[ShackledChasm].top = &relations[DrownedVale]
	relations[ShackledChasm].topLeft = &relations[UmbralWildwood]

	relations[ReaversPass].bottomLeft = &relations[Terminus]
	relations[ReaversPass].topRight = &relations[TheFingers]
	relations[ReaversPass].top = &relations[EndlessShore]
	relations[ReaversPass].topLeft = &relations[AllodsBight]

	relations[AshFields].bottomRight = &relations[RedRiver]
	relations[AshFields].topRight = &relations[Heartlands]
	relations[AshFields].top = &relations[Sableport]
	relations[AshFields].topLeft = &relations[Origin]

	relations[GreatMarch].bottomLeft = &relations[RedRiver]
	relations[GreatMarch].bottom = &relations[Kalokai]
	relations[GreatMarch].bottomRight = &relations[Acrithia]
	relations[GreatMarch].topRight = &relations[ShackledChasm]
	relations[GreatMarch].top = &relations[UmbralWildwood]
	relations[GreatMarch].topLeft = &relations[Heartlands]

	relations[Terminus].bottomLeft = &relations[Acrithia]
	relations[Terminus].topRight = &relations[ReaversPass]
	relations[Terminus].top = &relations[AllodsBight]
	relations[Terminus].topLeft = &relations[ShackledChasm]

	relations[RedRiver].bottomRight = &relations[Kalokai]
	relations[RedRiver].topRight = &relations[GreatMarch]
	relations[RedRiver].top = &relations[Heartlands]
	relations[RedRiver].topLeft = &relations[AshFields]

	relations[Acrithia].bottomLeft = &relations[Kalokai]
	relations[Acrithia].topRight = &relations[Terminus]
	relations[Acrithia].top = &relations[ShackledChasm]
	relations[Acrithia].topLeft = &relations[GreatMarch]

	relations[Kalokai].topRight = &relations[Acrithia]
	relations[Kalokai].top = &relations[GreatMarch]
	relations[Kalokai].topLeft = &relations[RedRiver]

	return relations[:]
}
