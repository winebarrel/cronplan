package cronplan

import (
	"iter"
	"time"
)

type Iterator struct {
	expr *Expression
	from time.Time
}

func (iter *Iterator) HasNext() bool {
	next := iter.expr.Next(iter.from)
	return !next.IsZero()
}

func (iter *Iterator) Next() time.Time {
	next := iter.expr.Next(iter.from)
	if !next.IsZero() {
		iter.from = next.Add(1 * time.Minute)
	}
	return next
}

func (iter *Iterator) Seq() iter.Seq[time.Time] {
	return func(yield func(time.Time) bool) {
		for {
			next := iter.Next()
			if next.IsZero() {
				break
			}
			if !yield(next) {
				break
			}
		}
	}
}

func (v *Expression) IterFrom(from time.Time) *Iterator {
	iter := &Iterator{
		expr: v,
		from: from,
	}
	return iter
}
