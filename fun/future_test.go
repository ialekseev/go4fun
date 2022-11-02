package fun

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFuture_ApplyFuture1(t *testing.T) {
	// given
	f := FutureValue(func() int {
		time.Sleep(time.Millisecond * 10)
		return 123
	})

	// when
	r := ApplyFuture1(f, func(a int) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return fmt.Sprint(a) + "456"
		})
	})

	time.Sleep(time.Millisecond * 50)
	// then
	assert.Equal(t, "123456", r.Result())
}

func TestFuture_ApplyFuture2(t *testing.T) {
	// given
	f1 := FutureValue(func() int {
		time.Sleep(time.Millisecond * 10)
		return 123
	})

	f2 := FutureValue(func() bool {
		time.Sleep(time.Millisecond * 10)
		return true
	})

	// when
	r := ApplyFuture2(f1, f2, func(a int, b bool) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return fmt.Sprint(a) + " " + fmt.Sprint(b)
		})
	})

	time.Sleep(time.Millisecond * 60)
	// then
	assert.Equal(t, "123 true", r.Result())
}

func TestFuture_ApplyFuture3(t *testing.T) {
	// given
	f1 := FutureValue(func() int {
		time.Sleep(time.Millisecond * 10)
		return 123
	})

	f2 := FutureValue(func() bool {
		time.Sleep(time.Millisecond * 10)
		return true
	})

	f3 := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	// when
	r := ApplyFuture3(f1, f2, f3, func(a int, b bool, c string) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return fmt.Sprint(a) + " " + fmt.Sprint(b) + " " + fmt.Sprint(c)
		})
	})

	time.Sleep(time.Millisecond * 60)
	// then
	assert.Equal(t, "123 true abc", r.Result())
}

func TestFuture_FlatMap(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	//when
	r := f.FlatMap(func(a string) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return a + "def"
		})
	})

	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "abcdef", r.Result())
}

func TestFuture_FlatMapFuture(t *testing.T) {
	//given
	f := FutureValue(func() int {
		time.Sleep(time.Millisecond * 10)
		return 123
	})

	//when
	r := FlatMapFuture(f, func(a int) Future[string] {
		return FutureValue(func() string {
			time.Sleep(time.Millisecond * 10)
			return fmt.Sprint(a) + "456"
		})
	})

	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "123456", r.Result())
}

func TestFuture_FutureValue(t *testing.T) {
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	time.Sleep(time.Millisecond * 50)
	assert.Equal(t, Some("abc"), *f.value)
}

func TestFuture_IsCompleted(t *testing.T) {
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})
	assert.False(t, f.IsCompleted())
	time.Sleep(time.Millisecond * 50)
	assert.True(t, f.IsCompleted())
}

func TestFuture_Map(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})
	//when
	r := f.Map(func(a string) string { return a + "def" })
	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "abcdef", r.Result())
}

func TestFuture_MapFuture(t *testing.T) {
	//given
	f := FutureValue(func() int {
		time.Sleep(time.Millisecond * 10)
		return 123
	})
	//when
	r := MapFuture(f, func(a int) string { return fmt.Sprint(a) + "456" })
	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "123456", r.Result())
}

func TestFuture_OnComplete(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 10)
		return "abc"
	})

	var result string
	//when
	f.OnComplete(func(a string) { result = a + "def" })
	time.Sleep(time.Millisecond * 50)
	//then
	assert.Equal(t, "abcdef", result)
}

func TestFuture_Result(t *testing.T) {
	//given
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 50)
		return "abc"
	})
	//when
	r := f.Result()
	//then
	assert.Equal(t, "abc", r)
}

func TestFuture_Value(t *testing.T) {
	f := FutureValue(func() string {
		time.Sleep(time.Millisecond * 20)
		return "abc"
	})
	assert.Equal(t, None[string](), f.Value())
	time.Sleep(time.Millisecond * 30)
	assert.Equal(t, Some("abc"), f.Value())
}
