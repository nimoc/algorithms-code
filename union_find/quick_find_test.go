package union_find

import (
	"log"
	"testing"
)

func TestNewQuickFind(t *testing.T) {
	a := NewQuickFind(9)
	a.Union(5,6)
	a.Union(6,8)
	log.Print(a.String())
}
