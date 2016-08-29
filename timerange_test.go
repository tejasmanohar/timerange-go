package timerange

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	start := time.Date(2016, 8, 28, 9, 0, 0, 0, time.UTC)
	end := time.Date(2016, 8, 28, 11, 0, 0, 0, time.UTC)

	i := 0
	iter := New(start, end, time.Hour)
	for iter.Next() {
		current := iter.Current()
		switch {
		case i == 0:
			assert.Equal(t, true, current.Equal(start))
		case i == 1:
			assert.Equal(t, true, current.Equal(start.Add(time.Hour)))
		case i == 2:
			assert.Equal(t, true, current.Equal(start.Add(2*time.Hour)))
		}
		i++
	}
	assert.Equal(t, 3, i)
}

func TestUnevenRange(t *testing.T) {
	start := time.Date(2016, 8, 28, 9, 0, 0, 0, time.UTC)
	end := time.Date(2016, 8, 28, 10, 30, 0, 0, time.UTC)

	i := 0
	iter := New(start, end, time.Hour)
	for iter.Next() {
		current := iter.Current()
		switch {
		case i == 0:
			assert.Equal(t, true, current.Equal(start))
		case i == 1:
			assert.Equal(t, true, current.Equal(start.Add(time.Hour)))
		}
		i++
	}
	assert.Equal(t, 2, i)
}

func TestEmptyRange(t *testing.T) {
	start := time.Date(2016, 8, 28, 9, 0, 0, 0, time.UTC)
	end := time.Date(2016, 8, 28, 8, 0, 0, 0, time.UTC)

	iter := New(start, end, time.Nanosecond)
	assert.Equal(t, false, iter.Next())
}
