package cronplan

import (
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

func (v *Expression) IterFrom(from time.Time) *Iterator {
	iter := &Iterator{
		expr: v,
		from: from,
	}
	return iter
}
