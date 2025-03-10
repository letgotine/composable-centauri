package types

import (
	"cosmossdk.io/math"
)

// CheckExceedsQuota checks if new in/out flow is going to reach the max in/out or not
func (q *Quota) CheckExceedsQuota(direction PacketDirection, amount math.Int, totalValue math.Int) bool {
	// If there's no channel value (this should be almost impossible), it means there is no
	// supply of the asset, so we shoudn't prevent inflows/outflows
	if totalValue.IsZero() {
		return false
	}
	var threshold math.Int
	if direction == PACKET_RECV {
		threshold = totalValue.Mul(q.MaxPercentRecv).Quo(math.NewInt(100))
	} else {
		threshold = totalValue.Mul(q.MaxPercentSend).Quo(math.NewInt(100))
	}

	return amount.GT(threshold)
}
