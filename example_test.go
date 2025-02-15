package tsid

import (
	"testing"
)

func TestSeqID(t *testing.T) {
	if c, e := SeqID(10, 10); e == nil {
		var n int64
		for i := 0; i < 100; i++ {
			d := c()
			if d < 1 {
				t.Error("an error ID(zero) was generated")
			}
			if d < n {
				t.Error("the ID generated by SeqID are not incremental")
			}
			n = d
		}
	} else {
		t.Error(e)
	}
}

func BenchmarkSeqID(b *testing.B) {
	c, e := SeqID(10, 10)
	if e != nil {
		b.Fatal(e)
		return
	}
	var n int64
	for i := 0; i < b.N; i++ {
		d := c()
		if d < 1 {
			b.Error("an error ID(zero) was generated")
		}
		if d < n {
			b.Errorf("the ID generated by SeqID are not incremental, old: %d, new: %d", n, d)
		}
		n = d
	}
}
