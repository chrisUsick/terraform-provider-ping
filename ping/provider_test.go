package ping

import (
	"testing"
)

func TestProvider(t *testing.T) {
	if err := ProviderFactory(nil).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
