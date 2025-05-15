# mtest

Small unit testing helper library.

A vastly simplified version of [Testify](https://github.com/stretchr/testify). Custom assertions are nice, but most of the time they are overkill.

This library only implements:

- `True(t *testing.TB, ok bool)`
- `False(t *testing.TB, ok bool)`
- `Equal(t *testing.TB, expected, actual any)`
- `Error(t *testing.TB, err error)`
- `NoError(t *testing.TB, err error)`
- `Panics(t *testing.TB, fn func())`

Install:

```
go get github.com/mikattack/mtest
```

Example:

```
package yours

import (
  "testing"

  "github.com/mikattack/mtest/assert"
)

func TestSomething(t *testing.T) {
  assert.Equal(t, 123, 123) // Pass
  assert.Equal(t, 123, 321) // Fail

  assert.Error(t, errors.New("intentional")) // Pass
  assert.NoError(t, "not an error") // Pass
}
```