package mocks

import (
	"encoding/json"

	"github.com/OrderMyGear/taxjar-go-2"
)

// ListRefunds - mock response
var ListRefunds = new(taxjar.ListRefundsResponse)
var _ = json.Unmarshal([]byte(ListRefundsJSON), &ListRefunds)

// ListRefundsJSON - mock ListRefunds JSON
var ListRefundsJSON = `{
	"refunds": [
		"123-refund",
		"246-refund",
		"359-refund"
	]
}`
