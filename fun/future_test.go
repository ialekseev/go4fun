package fun

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewFuture(t *testing.T) {
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	time.Sleep(time.Millisecond * 50)
	assert.Equal(t, Some("abc"), *f.value)
}

func TestIsCompleted(t *testing.T) {
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	assert.False(t, f.IsCompleted())
	time.Sleep(time.Millisecond * 50)
	assert.True(t, f.IsCompleted())
}

func TestOnComplete(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	var result string
	//when
	f.OnComplete(func(s string) { result = s + "def" })
	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "abcdef", result)
}

func TestMap(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})
	//when
	r := f.Map(func(s string) string { return s + "def" })
	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "abcdef", r.Result())
}

func TestFlatMap(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	//when
	r := f.FlatMap(func(s string) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return s + "def"
		})
	})

	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "abcdef", r.Result())
}
