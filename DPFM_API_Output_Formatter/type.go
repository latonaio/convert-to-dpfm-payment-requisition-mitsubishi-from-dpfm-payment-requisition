package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey    string      `json:"connection_key"`
	Result           bool        `json:"result"`
	RedisKey         string      `json:"redis_key"`
	Filepath         string      `json:"filepath"`
	APIStatusCode    int         `json:"api_status_code"`
	RuntimeSessionID string      `json:"runtime_session_id"`
	BusinessPartner  int         `json:"business_partner"`
	ServiceLabel     string      `json:"service_label"`
	Message          Message     `json:"Message"`
	APISchema        string      `json:"api_schema"`
	Accepter         []string    `json:"accepter"`
	OrderID          interface{} `json:"order_id"`
	Deleted          bool        `json:"deleted"`
}

type Message struct {
	Header Header `json:"Header"`
}

type Header struct {
	Payer                     int     `json:"Payer"`
	PayerPaymentDate          string  `json:"PayerPaymentDate"`
	PayerPaymentRequisitionID int     `json:"PayerPaymentRequisitionID"`
	TransactionCode           *string `json:"transactionCode"`
	PaymentApplicantCode      string  `json:"paymentApplicantCode"`
	PaymentApplicantName      string  `json:"paymentApplicantName"`
	PaymentDate               string  `json:"paymentDate"`
	PayingBankNo              string  `json:"payingBankNo"`
	PayingBankName            string  `json:"payingBankName"`
	PayingBranchNo            string  `json:"payingBranchNo"`
	PayingBranchName          string  `json:"payingBranchName"`
	PayingDepositType         string  `json:"payingDepositType"`
	PayingAccountNo           string  `json:"payingAccountNo"`
	TotalNumber               *int    `json:"totalNumber"`
	TotalAmount               *int    `json:"totalAmount"`
}

type Item struct {
	Payer                       int    `json:"Payer"`
	PayerPaymentDate            string `json:"PayerPaymentDate"`
	PayerPaymentRequisitionID   int    `json:"PayerPaymentRequisitionI"`
	PayerPaymentRequisitionItem int    `json:"PayerPaymentRequisitionItem"`
	Payee                       int    `json:"Payee"`
	receivingBankNo             string `json:"receivingBankNo"`
	receivingBankName           string `json:"receivingBankName"`
	receivingBranchNo           string `json:"receivingBranchNo"`
	receivingBranchName         string `json:"receivingBranchName"`
	clearingHouseNo             string `json:"clearingHouseNo"`
	receivingDepositType        string `json:"receivingDepositType"`
	receivingAccountNo          string `json:"receivingAccountNo"`
	receivingBeneficiaryName    string `json:"receivingBeneficiaryName"`
	paymentAmount               int    `json:"paymentAmount"`
	newAmendCode                string `json:"newAmendCode"`
	customerCode1               string `json:"customerCode1"`
	customerCode2               string `json:"customerCode2"`
	paymentApplicantNo          string `json:"paymentApplicantNo"`
	ediInformation              string `json:"ediInformation"`
	paymentType                 string `json:"paymentType"`
}
