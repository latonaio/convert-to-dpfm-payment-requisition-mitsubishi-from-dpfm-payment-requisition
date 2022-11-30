package dpfm_api_input_reader

import (
	"convert-to-dpfm-payment-requisition-mitsubishi-from-dpfm/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToBpExistenceConf() {

}

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header

	return &requests.Header{
		Payer:                          data.Payer,
		PayerPaymentDate:               data.PayerPaymentDate,
		PayerPaymentRequisitionID:      data.PayerPaymentRequisitionID,
		PaymentReqnStatus:              data.PaymentReqnStatus,
		AcceptanceNoByFinInst:          data.AcceptanceNoByFinInst,
		PaytReqnAmtInTransCrcy:         data.PaytReqnAmtInTransCrcy,
		Currency:                       data.Currency,
		PaymentMethod:                  data.PaymentMethod,
		PayerPostingDate:               data.PayerPostingDate,
		PayerHouseBank:                 data.PayerHouseBank,
		PayerHouseBankAccount:          data.PayerHouseBankAccount,
		PayerFinInstCountry:            data.PayerFinInstCountry,
		PayerFinInstCode:               data.PayerFinInstCode,
		PayerFinInstBranchCode:         data.PayerFinInstBranchCode,
		PayerFinInstFullCode:           data.PayerFinInstFullCode,
		PayerFinInstSWIFTCode:          data.PayerFinInstSWIFTCode,
		PayerInternalFinInstCustomerID: data.PayerInternalFinInstCustomerID,
		PayerInternalFinInstAccountID:  data.PayerInternalFinInstAccountID,
		PayerFinInstControlKey:         data.PayerFinInstControlKey,
		PayerFinInstAccount:            data.PayerFinInstAccount,
		PayerFinInstAccountName:        data.PayerFinInstAccountName,
		PayerFinInstName:               data.PayerFinInstName,
		PayerFinInstBranchName:         data.PayerFinInstBranchName,
		PaymentApplicantCode:           data.PaymentApplicantCode,
		CreationDateTime:               data.CreationDateTime,
		ChangedOnDateTime:              data.ChangedOnDateTime,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	data := sdc.Item

	return &requests.Item{
		Payer:                          data.Payer,
		PayerPaymentDate:               data.PayerPaymentDate,
		PayerPaymentRequisitionID:      data.PayerPaymentRequisitionID,
		PayerPaymentRequisitionItem:    data.PayerPaymentRequisitionItem,
		Payee:                          data.Payee,
		PayeeFinInstCountry:            data.PayeeFinInstCountry,
		PayeeFinInstCode:               data.PayeeFinInstCode,
		PayeeFinInstBranchCode:         data.PayeeFinInstBranchCode,
		PayeeFinInstFullCode:           data.PayeeFinInstFullCode,
		PayeeFinInstSWIFTCode:          data.PayeeFinInstSWIFTCode,
		PaytReqnItemAmtInTransCrcy:     data.PaytReqnItemAmtInTransCrcy,
		PayeeInternalFinInstCustomerID: data.PayeeInternalFinInstCustomerID,
		PayeeInternalFinInstAccountID:  data.PayeeInternalFinInstAccountID,
		PayeeFinInstControlKey:         data.PayeeFinInstControlKey,
		PayeeFinInstAccount:            data.PayeeFinInstAccount,
		PayeeFinInstAccountName:        data.PayeeFinInstAccountName,
		PayeeFinInstName:               data.PayeeFinInstName,
		PayeeFinInstBranchName:         data.PayeeFinInstBranchName,
		PayerAccountingDocument:        data.PayerAccountingDocument,
		PayerAccountingDocumentItem:    data.PayerAccountingDocumentItem,
		CreationDateTime:               data.CreationDateTime,
		ChangedOnDateTime:              data.ChangedOnDateTime,
	}
}
