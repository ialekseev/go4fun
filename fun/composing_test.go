package fun

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComposingCompose2(t *testing.T) {
	//given
	f := func(a int) bool { return a != 0 }
	g := func(b bool) string { return fmt.Sprint(b) }

	//when
	h := Compose2(f, g)

	//then
	assert.Equal(t, g(f(1)), h(1))
}

func TestComposingCompose3(t *testing.T) {
	//given
	f := func(a int) string { return fmt.Sprint(a) }
	g := func(b string) bool { return b != "" }
	h := func(c bool) string { return fmt.Sprint(c) }

	//when
	j := Compose3(f, g, h)

	//then
	assert.Equal(t, h(g(f(1))), j(1))
}
