package dialekt

// Represents the actual wildcard portion of a pattern expression.
type PatternWildcard struct {
	PatternChildInterface
}

func NewPatternWildcard() {
	return &PatternWildcard{}
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (wild *PatternWildcard) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitPatternWildcard(wild)
}
