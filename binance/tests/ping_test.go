package ping_test

import (
	"github.com/pdepip/go-binance/binance"
	"testing"
)

func TestPing(t *testing.T) {
	client := binance.New("", "")

	if ping, err := client.Ping(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%+v\n", ping)
	}

}
