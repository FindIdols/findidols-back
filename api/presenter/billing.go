package presenter

//Customer data
type Customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Billing struct {
	BillingType string `json:"billingType"`
	BankSlipUrl string `json:"bankSlipUrl"`
}
