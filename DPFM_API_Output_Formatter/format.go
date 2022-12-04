package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm/DPFM_API_Input_Reader"
)

func ConvertToHeader(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.Header

	header := Header{
		Payer:                     *data.Payer,
		PayerPaymentDate:          *data.PayerPaymentDate,
		PayerPaymentRequisitionID: *data.PayerPaymentRequisitionID,
		//TransactionCode:           *data.,
		PaymentApplicantCode: *data.PaymentApplicantCode,
		PaymentApplicantName: *data.PayerFinInstAccountName,
		PaymentDate:          *data.PayerPaymentDate,
		PayingBankNo:         *data.PayerFinInstCode,
		PayingBankName:       *data.PayerFinInstName,
		PayingBranchNo:       *data.PayerFinInstBranchCode,
		PayingBranchName:     *data.PayerFinInstBranchName,
		PayingDepositType:    *data.PayerFinInstControlKey,
		PayingAccountNo:      *data.PayerFinInstAccount,
		//TotalNumber:           *data.TotalNumber,
		//TotalAmount: *int(data.PaytReqnAmtInTransCrcy), #floatをintに変換する
	}

	return &header, nil
}

/*
func ConvertToItem(sdc *dpfm_api_input_reader.SDC) (*Item, error) {
	data := sdc.Item

	item := Item{
		Payer:                       *data.Payer,
		PayerPaymentDate:            *data.PayerPaymentDate,
		PayerPaymentRequisitionID:   *data.PayerPaymentRequisitionID,
		PayerPaymentRequisitionItem: *data.PayerPaymentRequisitionItem,
		Payee:                       *data.Payee,
		receivingBankNo:             *data.PayeeFinInstCode,
		receivingBankName:           *data.PayeeFinInstName,
		receivingBranchNo:           *data.PayeeFinInstBranchCode,
		receivingBranchName:         *data.PayeeFinInstBranchName,
		receivingDepositType:        *data.PayeeFinInstControlKey,
		receivingAccountNo:          *data.PayeeFinInstAccount,
		receivingBeneficiaryName:    *data.PayeeFinInstAccountName,
		paymentAmount:               *data.PaytReqnItemAmtInTransCrcy,
	}

	return &item, nil
}
*/
