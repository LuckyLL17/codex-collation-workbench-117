package lexicon

import "fmt"

func (l Location) String() string {
	return fmt.Sprintf("卷%d/叶%d/行%d/字%d", l.Volume, l.Leaf, l.Line, l.Character)
}
func (l Location) Valid() bool { return l.Volume > 0 && l.Leaf > 0 && l.Line > 0 && l.Character > 0 }
func Distance(a, b Location) int {
	return abs(a.Volume-b.Volume)*100000 + abs(a.Leaf-b.Leaf)*1000 + abs(a.Line-b.Line)*100 + abs(a.Character-b.Character)
}
func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
