package withdrawal_system_status_test

import (
	"github.com/pdepip/go-binance/binance"
	"testing"
)

func TestWithdrawalSystemStatus(t *testing.T) {
	client := binance.New("", "")

	if status, err := client.GetWithdrawalSystemStatus(); err != nil {
		t.Fatal(err)
	} else {
		switch {
		case status.Status == binance.SystemStatusNormal:
			t.Logf("Status normal - full response: %+v", status)
		case status.Status == binance.SystemStatusMaintenance:
			t.Logf("Status maintenance - full response: %+v", status)
		default:
			t.Errorf("Unexpected status: %+v", status)
		}
	}
}
