package mocks

import (
	"encoding/json"

	"github.com/OrderMyGear/taxjar-go"
)

// ListCustomers - mock response
var ListCustomers = new(taxjar.ListCustomersResponse)
var _ = json.Unmarshal([]byte(ListCustomersJSON), &ListCustomers)

// ListCustomersJSON - mock ListCustomers JSON
var ListCustomersJSON = `{
	"customers": [
		"123",
		"456"
	]
}`
