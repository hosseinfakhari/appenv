package appenv

import (
	"testing"
)

func TestAppEnv_setSystemEnvs(t *testing.T) {
	a := NewAppEnv(".appenv", false)
	t.Log(a)
}