package libfrp

import "testing"

func TestVersion(t *testing.T) {
	t.Log(Version())
}

func TestIsFrpsRunning(t *testing.T) {
	t.Log(IsFrpcRunning())
}
