package fun

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var f1 = func(a int) string {
	time.Sleep(time.Millisecond * time.Duration(a))
	return fmt.Sprint(a)
}
var memo1 = Memo1(f1)

var f2 = func(a, b int) string {
	time.Sleep(time.Millisecond * time.Duration(a))
	return fmt.Sprint(a)
}
var memo2 = Memo2(f2)

var f3 = func(a, b, c int) string {
	time.Sleep(time.Millisecond * time.Duration(a))
	return fmt.Sprint(a)
}
var memo3 = Memo3(f3)

func TestMemo1(t *testing.T) {
	//when
	r := memo1(1)
	r = memo1(1)

	//then
	assert.Equal(t, "1", r)
}

func TestMemo2(t *testing.T) {
	//when
	r := memo2(1, 1)
	r = memo2(1, 1)

	//then
	assert.Equal(t, "1", r)
}

func TestMemo3(t *testing.T) {
	//when
	r := memo3(1, 1, 1)
	r = memo3(1, 1, 1)

	//then
	assert.Equal(t, "1", r)
}

func BenchmarkMemo1(b *testing.B) {
	b.Run(fmt.Sprintf("original function (f1)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f1(5)
		}
	})

	b.Run(fmt.Sprintf("memo function (m1)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			memo1(5)
		}
	})
}

func BenchmarkMemo2(b *testing.B) {
	b.Run(fmt.Sprintf("original function (f2)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f2(5, 5)
		}
	})

	b.Run(fmt.Sprintf("memo function (m2)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			memo2(5, 5)
		}
	})
}

func BenchmarkMemo3(b *testing.B) {
	b.Run(fmt.Sprintf("original function (f3)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			f3(5, 5, 5)
		}
	})

	b.Run(fmt.Sprintf("memo function (m3)"), func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			memo3(5, 5, 5)
		}
	})
}
