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
