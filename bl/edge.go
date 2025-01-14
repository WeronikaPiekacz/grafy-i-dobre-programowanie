package bl

type Edge struct {
	RightPoint string
	LeftPoint  string
}

func NewEdge(RightPoint string, LeftPoint string) *Edge {
	return &Edge{RightPoint: RightPoint, LeftPoint: LeftPoint}
}

func (edge *Edge) equals(other Edge) bool {
	return (edge.LeftPoint == other.LeftPoint && edge.RightPoint == other.RightPoint) ||
		(edge.LeftPoint == other.RightPoint && edge.RightPoint == other.LeftPoint)
}
