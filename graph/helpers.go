package graph

import (
	"ms.api/protos/pb/auth"
	"ms.api/protos/pb/onboarding"
	protoTypes "ms.api/protos/pb/types"
	"ms.api/types"
)

var (
	// Response messages
	authFailedMessage = "User authentication failed"
)

type Helper interface {
	MapQuestionaryStatus(val types.QuestionaryStatuses) protoTypes.Questionary_QuestionaryStatuses
	MapQuestionaryType(val types.QuestionaryTypes) protoTypes.Questionary_QuestionaryTypes
	MapProtoQuesionaryStatus(val protoTypes.Questionary_QuestionaryStatuses) types.QuestionaryStatuses
	MapProtoQuestionaryType(val protoTypes.Questionary_QuestionaryTypes) types.QuestionaryTypes
	GetDeveicePreferenceTypesIndex(val types.DevicePreferencesTypes) int32
	GetProtoCustomerStatuses(val types.CustomerStatuses) protoTypes.Customer_CustomerStatuses
	DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) protoTypes.DeviceToken_DeviceTokenTypes
	PreferenceInputFromModel(input types.DevicePreferencesTypes) protoTypes.DevicePreferences_DevicePreferencesTypes
	StaffLoginTypeFromModel(input types.AuthType) auth.StaffLoginRequest_AuthType
	MakeFeesFromProto(protoFees []*protoTypes.Fee) []*types.Fee
	MapFeeTypes(val protoTypes.Fee_FeeTypes) types.FeeTypes
	MapProtoFeeTypes(val types.FeeTypes) protoTypes.Fee_FeeTypes
	MapTransactionTypeStatus(val protoTypes.TransactionType_TransactionTypeStatuses) types.TransactionTypeStatuses
	GetProtoTransactionTypesStatuses(val types.TransactionTypeStatuses) protoTypes.TransactionType_TransactionTypeStatuses
	GetProtoDeviceTokenType(val types.DeviceTokenTypes) protoTypes.DeviceToken_DeviceTokenTypes
	GetProtoDevicePreferencesType(val types.DevicePreferencesTypes) protoTypes.DevicePreferences_DevicePreferencesTypes
	MapCustomerTitle(val types.CustomerTitle) protoTypes.Customer_CustomerTitle
	MapProtoCustomerTitle(val protoTypes.Customer_CustomerTitle) types.CustomerTitle
	MakeExchangeRateFromProto(protoExchangeRate *protoTypes.ExchangeRate) *types.ExchangeRate
	MapProtoCredentialTypes(val protoTypes.IdentityCredentials_IdentityCredentialsTypes) types.IdentityCredentialsTypes
	MapCredentialTypes(val types.IdentityCredentialsTypes) protoTypes.IdentityCredentials_IdentityCredentialsTypes
	MapStaffAuditLogType(val protoTypes.StaffAuditLog_StaffAuditLogTypes) types.StaffAuditLogType
	MapProtoStaffAuditLogType(val types.StaffAuditLogType) protoTypes.StaffAuditLog_StaffAuditLogTypes
	MapScheduledTransactionRepeatType(val protoTypes.ScheduledTransaction_ScheduledTransactionRepeatType) types.ScheduledTransactionRepeatType
	MapProtoScheduledTransactionRepeatType(val types.ScheduledTransactionRepeatType) protoTypes.ScheduledTransaction_ScheduledTransactionRepeatType
	MapScheduledTransactionStatus(val protoTypes.ScheduledTransaction_ScheduledTransactionStatuses) types.ScheduledTransactionStatus
	MapProtoScheduledTransactionStatus(val types.ScheduledTransactionStatus) protoTypes.ScheduledTransaction_ScheduledTransactionStatuses
	MapProtoCustomerPreferenceType(val types.CustomerPreferencesTypes) protoTypes.CustomerPreferences_CustomerPreferencesTypes
	MapProtoFAQTypes(val types.FilterType) onboarding.GetFAQRequest_FAQFilter
}

type helpersfactory struct{}

func (h *helpersfactory) MapCustomerTitle(val types.CustomerTitle) protoTypes.Customer_CustomerTitle {
	switch val {
	case types.CustomerTitleMr:
		return protoTypes.Customer_MR
	case types.CustomerTitleMrs:
		return protoTypes.Customer_MRS
	case types.CustomerTitleMiss:
		return protoTypes.Customer_MISS
	case types.CustomerTitleMs:
		return protoTypes.Customer_MS
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) MapProtoCustomerTitle(val protoTypes.Customer_CustomerTitle) types.CustomerTitle {
	switch val {
	case protoTypes.Customer_MR:
		return types.CustomerTitleMr
	case protoTypes.Customer_MRS:
		return types.CustomerTitleMrs
	case protoTypes.Customer_MISS:
		return types.CustomerTitleMiss
	case protoTypes.Customer_MS:
		return types.CustomerTitleMs
	default:
		// should never happen
		return ""
	}
}

func (h *helpersfactory) MapQuestionaryStatus(val types.QuestionaryStatuses) protoTypes.Questionary_QuestionaryStatuses {
	switch val {
	case types.QuestionaryStatusesActive:
		return protoTypes.Questionary_ACTIVE
	case types.QuestionaryStatusesInactive:
		return protoTypes.Questionary_INACTIVE
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) MapQuestionaryType(val types.QuestionaryTypes) protoTypes.Questionary_QuestionaryTypes {
	switch val {
	case types.QuestionaryTypesReasons:
		return protoTypes.Questionary_REASONS
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) MapProtoQuesionaryStatus(val protoTypes.Questionary_QuestionaryStatuses) types.QuestionaryStatuses {
	switch val {
	case protoTypes.Questionary_ACTIVE:
		return types.QuestionaryStatusesActive
	case protoTypes.Questionary_INACTIVE:
		return types.QuestionaryStatusesInactive
	}

	return ""
}

func (h *helpersfactory) MapProtoQuestionaryType(val protoTypes.Questionary_QuestionaryTypes) types.QuestionaryTypes {
	switch val {
	case protoTypes.Questionary_REASONS:
		return types.QuestionaryTypesReasons
	}

	return ""
}

func (h *helpersfactory) GetDeveicePreferenceTypesIndex(val types.DevicePreferencesTypes) int32 {
	switch val {
	case types.DevicePreferencesTypesPush:
		return 0
	case types.DevicePreferencesTypesBiometrics:
		return 1
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) MapProtoCustomerStatuses(val protoTypes.Customer_CustomerStatuses) types.CustomerStatuses {
	switch val {
	case protoTypes.Customer_SIGNEDUP:
		return types.CustomerStatusesSignedup
	case protoTypes.Customer_REGISTERED:
		return types.CustomerStatusesRegistered
	case protoTypes.Customer_VERIFIED:
		return types.CustomerStatusesVerified
	case protoTypes.Customer_ONBOARDED:
		return types.CustomerStatusesOnboarded
	case protoTypes.Customer_REJECTED:
		return types.CustomerStatusesRegistered
	case protoTypes.Customer_EXITED:
		return types.CustomerStatusesExited
	case protoTypes.Customer_NGN_ONBOARDED:
		return types.CustomerStatusesNgnOnboarded
	default:
		// should never happen
		return ""
	}
}

func (h *helpersfactory) GetProtoCustomerStatuses(val types.CustomerStatuses) protoTypes.Customer_CustomerStatuses {
	switch val {
	case types.CustomerStatusesSignedup:
		return protoTypes.Customer_SIGNEDUP
	case types.CustomerStatusesRegistered:
		return protoTypes.Customer_REGISTERED
	case types.CustomerStatusesVerified:
		return protoTypes.Customer_VERIFIED
	case types.CustomerStatusesOnboarded:
		return protoTypes.Customer_ONBOARDED
	case types.CustomerStatusesRejected:
		return protoTypes.Customer_REJECTED
	case types.CustomerStatusesExited:
		return protoTypes.Customer_EXITED
	case types.CustomerStatusesNgnOnboarded:
		return protoTypes.Customer_NGN_ONBOARDED
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) GetProtoProductStatuses(val types.ProductStatuses) protoTypes.Product_ProductStatuses {
	switch val {
	case types.ProductStatusesActive:
		return protoTypes.Product_ACTIVE
	case types.ProductStatusesInactive:
		return protoTypes.Product_INACTIVE
	default:
		return -1
	}
}

func (h *helpersfactory) GetProtoProductTypes(val types.ProductTypes) protoTypes.Product_ProductTypes {
	switch val {
	case types.ProductTypesCurrentAccount:
		return protoTypes.Product_CURRENT_ACCOUNT
	case types.ProductTypesFixedDeposit:
		return protoTypes.Product_FIXED_DEPOSIT
	default:
		return -1
	}
}

func (h *helpersfactory) GetProtoTransactionStatuses(val types.TransactionStatuses) protoTypes.Transaction_TransactionStatuses {
	switch val {
	case types.TransactionStatusesApproved:
		return protoTypes.Transaction_APPROVED
	case types.TransactionStatusesPending:
		return protoTypes.Transaction_PENDING
	case types.TransactionStatusesRejected:
		return protoTypes.Transaction_REJECTED
	default:
		return -1
	}
}

func (h *helpersfactory) GetProtoTransactionTypesStatuses(val types.TransactionTypeStatuses) protoTypes.TransactionType_TransactionTypeStatuses {
	switch val {
	case types.TransactionTypeStatusesActive:
		return protoTypes.TransactionType_ACTIVE
	case types.TransactionTypeStatusesInactive:
		return protoTypes.TransactionType_INACTIVE
	default:
		return -1
	}
}

func (h *helpersfactory) GetProtoDeviceTokenType(val types.DeviceTokenTypes) protoTypes.DeviceToken_DeviceTokenTypes {
	switch val {
	case types.DeviceTokenTypesFirebase:
		return protoTypes.DeviceToken_FIREBASE
	default:
		return -1
	}
}

func (h *helpersfactory) GetProtoDevicePreferencesType(val types.DevicePreferencesTypes) protoTypes.DevicePreferences_DevicePreferencesTypes {
	switch val {
	case types.DevicePreferencesTypesBiometrics:
		return protoTypes.DevicePreferences_BIOMETRICS
	case types.DevicePreferencesTypesPush:
		return protoTypes.DevicePreferences_PUSH
	default:
		return -1
	}
}

func (h *helpersfactory) DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) protoTypes.DeviceToken_DeviceTokenTypes {
	switch tokenType {
	default:
		return protoTypes.DeviceToken_FIREBASE
	}
}

func (h *helpersfactory) PreferenceInputFromModel(input types.DevicePreferencesTypes) protoTypes.DevicePreferences_DevicePreferencesTypes {
	switch input {
	case types.DevicePreferencesTypesPush:
		return protoTypes.DevicePreferences_PUSH
	case types.DevicePreferencesTypesBiometrics:
		return protoTypes.DevicePreferences_BIOMETRICS
	default:
		return protoTypes.DevicePreferences_PUSH
	}
}

func (h *helpersfactory) MapProtoStaffStatuses(val protoTypes.Staff_StaffStatuses) types.StaffStatuses {
	switch val {
	case protoTypes.Staff_ACTIVE:
		return types.StaffStatusesActive
	case protoTypes.Staff_INACTIVE:
		return types.StaffStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) StaffLoginTypeFromModel(input types.AuthType) auth.StaffLoginRequest_AuthType {
	switch input {
	case types.AuthTypeGoogle:
		return auth.StaffLoginRequest_GOOGLE
	default:
		return -1
	}
}

func (h *helpersfactory) MapCDDStatusesFromModel(val types.CDDStatuses) protoTypes.CDD_CDDStatuses {
	switch val {
	case types.CDDStatusesPending:
		return protoTypes.CDD_PENDING
	case types.CDDStatusesManualReview:
		return protoTypes.CDD_MANUAL_REVIEW
	case types.CDDStatusesApproved:
		return protoTypes.CDD_APPROVED
	case types.CDDStatusesDeclined:
		return protoTypes.CDD_DECLINED
	default:
		return -1
	}
}

func (h *helpersfactory) MapProtoCDDStatuses(val protoTypes.CDD_CDDStatuses) types.CDDStatuses {
	switch val {
	case protoTypes.CDD_PENDING:
		return types.CDDStatusesPending
	case protoTypes.CDD_MANUAL_REVIEW:
		return types.CDDStatusesManualReview
	case protoTypes.CDD_APPROVED:
		return types.CDDStatusesApproved
	case protoTypes.CDD_DECLINED:
		return types.CDDStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) MapProtoProductStatuses(val protoTypes.Product_ProductStatuses) types.ProductStatuses {
	switch val {
	case protoTypes.Product_ACTIVE:
		return types.ProductStatusesActive
	case protoTypes.Product_INACTIVE:
		return types.ProductStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) MapProtoBeneficiaryStatuses(val protoTypes.Beneficiary_BeneficiaryStatuses) types.BeneficiaryStatuses {
	switch val {
	case protoTypes.Beneficiary_ACTIVE:
		return types.BeneficiaryStatusesActive
	case protoTypes.Beneficiary_INACTIVE:
		return types.BeneficiaryStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) MapBeneficiaryStatuses(val types.BeneficiaryStatuses) protoTypes.Beneficiary_BeneficiaryStatuses {
	switch val {
	case types.BeneficiaryStatusesActive:
		return protoTypes.Beneficiary_ACTIVE
	case types.BeneficiaryStatusesInactive:
		return protoTypes.Beneficiary_INACTIVE
	default:
		return -1
	}
}

func (h *helpersfactory) MapProtoProductTypes(val protoTypes.Product_ProductTypes) types.ProductTypes {
	switch val {
	case protoTypes.Product_CURRENT_ACCOUNT:
		return types.ProductTypesCurrentAccount
	case protoTypes.Product_FIXED_DEPOSIT:
		return types.ProductTypesFixedDeposit
	default:
		return ""
	}
}

func (h *helpersfactory) MapTransactionStatuses(val protoTypes.Transaction_TransactionStatuses) types.TransactionStatuses {
	switch val {
	case protoTypes.Transaction_APPROVED:
		return types.TransactionStatusesApproved
	case protoTypes.Transaction_PENDING:
		return types.TransactionStatusesPending
	case protoTypes.Transaction_REJECTED:
		return types.TransactionStatusesRejected
	default:
		return ""
	}
}

func (h *helpersfactory) MapBeneficiaryAccountStatuses(val protoTypes.BeneficiaryAccount_BeneficiaryAccountStatuses) types.BeneficiaryAccountStatuses {
	switch val {
	case protoTypes.BeneficiaryAccount_ACTIVE:
		return types.BeneficiaryAccountStatusesActive
	case protoTypes.BeneficiaryAccount_INACTIVE:
		return types.BeneficiaryAccountStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) MapLinkedTransactionStatuses(val protoTypes.LinkedTransaction_LinkedTransactionStatuses) types.LinkedTransactionStatuses {
	switch val {
	case protoTypes.LinkedTransaction_APPROVED:
		return types.LinkedTransactionStatusesApproved
	case protoTypes.LinkedTransaction_PENDING:
		return types.LinkedTransactionStatusesPending
	case protoTypes.LinkedTransaction_REJECTED:
		return types.LinkedTransactionStatusesRejected
	default:
		return ""
	}
}

func (h *helpersfactory) MapTransactionTypeStatus(val protoTypes.TransactionType_TransactionTypeStatuses) types.TransactionTypeStatuses {
	switch val {
	case protoTypes.TransactionType_ACTIVE:
		return types.TransactionTypeStatusesActive
	case protoTypes.TransactionType_INACTIVE:
		return types.TransactionTypeStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) makeAddressFromProto(adddresses []*protoTypes.Address) []*types.Address {
	if adddresses == nil {
		return make([]*types.Address, 0)
	}

	addresses_ := make([]*types.Address, len(adddresses))
	for i, address := range adddresses {
		addresses_[i] = &types.Address{
			Primary: address.Primary,
			Country: &types.Country{
				ID:         address.Country.Id,
				CodeAlpha2: address.Country.CodeAlpha2,
				CodeAlpha3: address.Country.CodeAlpha3,
				Name:       address.Country.Name,
			},
			State:    &address.State,
			City:     &address.City,
			Street:   address.Street,
			Postcode: address.Postcode,
			Cordinates: func() *types.Cordinates {
				if address.Coordinates == nil {
					return &types.Cordinates{}
				}
				return &types.Cordinates{
					Latitude:  float64(address.Coordinates.Latitude),
					Longitude: float64(address.Coordinates.Longitude),
				}
			}(),
		}
	}
	return addresses_
}

func (h *helpersfactory) makePhonesFromProto(phones []*protoTypes.Phone) []*types.Phone {
	if phones == nil {
		return make([]*types.Phone, 0)
	}

	phones_ := make([]*types.Phone, len(phones))
	for i, phone := range phones {
		phones_[i] = &types.Phone{
			Primary:  phone.Primary,
			Number:   phone.Number,
			Verified: phone.Verified,
		}
	}
	return phones_
}

func (h *helpersfactory) makeCustomerFromProto(customer *protoTypes.Customer) *types.Customer {
	result := &types.Customer{}

	if customer != nil {
		email := &types.Email{}
		if customer.Email != nil {
			email.Address = customer.Email.Address
			email.Verified = customer.Email.Verified
		}

		result = &types.Customer{
			ID:        customer.Id,
			Title:     h.MapProtoCustomerTitle(customer.Title),
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Dob:       customer.Dob,
			Bvn:       &customer.Bvn,
			Addresses: h.makeAddressFromProto(customer.Addresses),
			Phones:    h.makePhonesFromProto(customer.Phones),
			Email:     email,
			Status:    h.MapProtoCustomerStatuses(customer.Status),
			StatusTs:  customer.StatusTs.AsTime().Unix(),
			Ts:        customer.Ts.AsTime().Unix(),
		}
	}
	return result
}

func (h *helpersfactory) MakeCurrencyFromProto(currency *protoTypes.Currency) *types.Currency {
	if currency == nil {
		return &types.Currency{}
	}

	return &types.Currency{
		ID:     currency.Id,
		Name:   currency.Name,
		Code:   currency.Code,
		Symbol: currency.Symbol,
	}
}

func (h *helpersfactory) makeStaffFromProto(staff *protoTypes.Staff) *types.Staff {
	return &types.Staff{
		ID:        staff.Id,
		Name:      staff.Name,
		LastName:  staff.LastName,
		Dob:       &staff.Dob,
		Addresses: h.makeAddressFromProto(staff.Addresses),
		Phones:    h.makePhonesFromProto(staff.Phones),
		Email:     staff.Email,
		Status:    h.MapProtoStaffStatuses(staff.Status),
		StatusTs:  staff.StatusTs.AsTime().Unix(),
		Ts:        staff.Ts.AsTime().Unix(),
	}
}

func (h *helpersfactory) mapOrganizationStatuses(val protoTypes.Organization_OrganizationStatuses) types.OrganizationStatuses {
	switch val {
	case protoTypes.Organization_ACTIVE:
		return types.OrganizationStatusesActive
	case protoTypes.Organization_INACTIVE:
		return types.OrganizationStatusesInactive

	default:
		return ""
	}
}

func (h *helpersfactory) makeOrganisationFromProto(organization *protoTypes.Organization) *types.Organization {
	return &types.Organization{
		ID:       organization.Id,
		Name:     organization.Name,
		Status:   h.mapOrganizationStatuses(organization.Status),
		StatusTs: organization.StatusTs.AsTime().Unix(),
		Ts:       organization.Ts.AsTime().Unix(),
	}
}

func (h *helpersfactory) mapPOAStatuses(val protoTypes.POA_POAStatuses) types.POAStatuses {
	switch val {
	case protoTypes.POA_APPROVED:
		return types.POAStatusesApproved
	case protoTypes.POA_MANUAL_REVIEW:
		return types.POAStatusesManualReview
	case protoTypes.POA_PENDING:
		return types.POAStatusesPending
	case protoTypes.POA_DECLINED:
		return types.POAStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapPOAActionStatuses(val protoTypes.POAAction_POAStatuses) types.POAStatuses {
	switch val {
	case protoTypes.POAAction_APPROVED:
		return types.POAStatusesApproved
	case protoTypes.POAAction_MANUAL_REVIEW:
		return types.POAStatusesManualReview
	case protoTypes.POAAction_PENDING:
		return types.POAStatusesPending
	case protoTypes.POAAction_DECLINED:
		return types.POAStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapPOAActionTypes(val protoTypes.POAAction_POAActionTypes) types.POAActionTypes {
	switch val {
	case protoTypes.POAAction_CHANGE_STATUS:
		return types.POAActionTypesChangeStatus
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCStatuses(val protoTypes.KYC_KYCStatuses) types.KYCStatuses {
	switch val {
	case protoTypes.KYC_APPROVED:
		return types.KYCStatusesApproved
	case protoTypes.KYC_MANUAL_REVIEW:
		return types.KYCStatusesManualReview
	case protoTypes.KYC_PENDING:
		return types.KYCStatusesPending
	case protoTypes.KYC_DECLINED:
		return types.KYCStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCActionTypes(val protoTypes.KYCAction_KYCActionTypes) types.KYCActionTypes {
	switch val {
	case protoTypes.KYCAction_CHANGE_STATUS:
		return types.KYCActionTypesChangeStatus
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCActionStatuses(val protoTypes.KYCAction_KYCStatuses) types.KYCStatuses {
	switch val {
	case protoTypes.KYCAction_APPROVED:
		return types.KYCStatusesApproved
	case protoTypes.KYCAction_MANUAL_REVIEW:
		return types.KYCStatusesManualReview
	case protoTypes.KYCAction_PENDING:
		return types.KYCStatusesPending
	case protoTypes.KYCAction_DECLINED:
		return types.KYCStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCReportStatuses(val protoTypes.Report_ReportStatuses) types.ReportStatuses {
	switch val {
	case protoTypes.Report_APPROVED:
		return types.ReportStatusesApproved
	case protoTypes.Report_MANUAL_REVIEW:
		return types.ReportStatusesManualReview
	case protoTypes.Report_PENDING:
		return types.ReportStatusesPending
	case protoTypes.Report_DECLINED:
		return types.ReportStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCReportTypes(val protoTypes.Report_ReportTypes) types.KYCTypes {
	switch val {
	case protoTypes.Report_DOCUMENT:
		return types.KYCTypesDocument
	case protoTypes.Report_FACIAL_VIDEO:
		return types.KYCTypesFacialVideo
	default:
		return ""
	}
}

func (h *helpersfactory) mapAMLStatuses(val protoTypes.AML_AMLStatuses) types.AMLStatuses {
	switch val {
	case protoTypes.AML_APPROVED:
		return types.AMLStatusesApproved
	case protoTypes.AML_MANUAL_REVIEW:
		return types.AMLStatusesManualReview
	case protoTypes.AML_PENDING:
		return types.AMLStatusesPending
	case protoTypes.AML_DECLINED:
		return types.AMLStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapAMLActionStatuses(val protoTypes.AMLAction_AMLStatuses) types.AMLStatuses {
	switch val {
	case protoTypes.AMLAction_APPROVED:
		return types.AMLStatusesApproved
	case protoTypes.AMLAction_MANUAL_REVIEW:
		return types.AMLStatusesManualReview
	case protoTypes.AMLAction_PENDING:
		return types.AMLStatusesPending
	case protoTypes.AMLAction_DECLINED:
		return types.AMLStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapAMLActionTypes(val protoTypes.AMLAction_AMLActionTypes) types.AMLActionTypes {
	switch val {
	case protoTypes.AMLAction_CHANGE_STATUS:
		return types.AMLActionTypesChangeStatus
	default:
		return ""
	}
}

func (h *helpersfactory) makeAMLsFromProto(amls []*protoTypes.AML) []*types.Aml {
	amls_ := make([]*types.Aml, len(amls))
	for i, aml := range amls {
		amlActions := make([]*types.AMLAction, len(aml.Actions))
		for j, action := range aml.Actions {
			amlActions[j] = &types.AMLAction{
				Type:         h.mapAMLActionTypes(action.Type),
				Reporter:     h.makeStaffFromProto(action.Reporter),
				TargetStatus: h.mapAMLActionStatuses(action.TargetStatus),
				Message:      action.Message,
				Ts:           action.Ts.AsTime().Unix(),
			}
		}

		amls_[i] = &types.Aml{
			Organization: h.makeOrganisationFromProto(aml.Organization),
			Identifier:   aml.Identifier,
			File:         &aml.File,
			Result:       aml.Result,
			PublicURL:    &aml.PublicUrl,
			Actions:      amlActions,
			Status:       h.mapAMLStatuses(aml.Status),
			StatusTs:     aml.StatusTs.AsTime().Unix(),
			Ts:           aml.Ts.AsTime().Unix(),
		}
	}
	return amls_
}

func (h *helpersfactory) makeKYCsFromProto(kycs []*protoTypes.KYC) []*types.Kyc {
	kycs_ := make([]*types.Kyc, len(kycs))
	for i, kyc := range kycs {
		reports := make([]*types.Reports, len(kyc.Reports))
		for j, report := range kyc.Reports {
			reports[j] = &types.Reports{
				Identifier: report.Identifier,
				Type:       h.mapKYCReportTypes(report.Type),
				File:       &report.File,
				Result:     &report.Result,
				SubResult:  &report.SubResult,
				PublicURL:  &report.PublicUrl,
				Review: &types.Review{
					Resubmit: report.Review.Resubmit,
					Message:  &report.Review.Message,
					Ts:       report.Ts.AsTime().Unix(),
				},
				Status:   h.mapKYCReportStatuses(report.Status),
				StatusTs: report.StatusTs.AsTime().Unix(),
				Ts:       report.Ts.AsTime().Unix(),
			}
		}

		actions := make([]*types.KYCAction, len(kyc.Actions))
		for k, action := range kyc.Actions {
			actions[k] = &types.KYCAction{
				Type:         h.mapKYCActionTypes(action.Type),
				Reporter:     h.makeStaffFromProto(action.Reporter),
				TargetStatus: h.mapKYCActionStatuses(action.TargetStatus),
				Message:      action.Message,
				Ts:           action.Ts.AsTime().Unix(),
			}
		}

		kycs_[i] = &types.Kyc{
			Organization: h.makeOrganisationFromProto(kyc.Organization),
			Identifier:   kyc.Identifier,
			PublicURL:    &kyc.PublicUrl,
			Reports:      reports,
			Actions:      actions,
			Status:       h.mapKYCStatuses(kyc.Status),
			StatusTs:     kyc.StatusTs.AsTime().Unix(),
			Ts:           kyc.Ts.AsTime().Unix(),
		}
	}

	return kycs_
}

func (h *helpersfactory) makePOAsFromProto(poas []*protoTypes.POA) []*types.Poa {
	poas_ := make([]*types.Poa, len(poas))
	for i, poa := range poas {
		poaActions := make([]*types.POAAction, len(poa.Actions))
		for j, poaAction := range poa.Actions {
			poaActions[j] = &types.POAAction{
				Type:         h.mapPOAActionTypes(poaAction.Type),
				Reporter:     h.makeStaffFromProto(poaAction.Reporter),
				TargetStatus: h.mapPOAActionStatuses(poaAction.TargetStatus),
				Message:      poaAction.Message,
				Ts:           poaAction.Ts.AsTime().Unix(),
			}
		}
		poas_[i] = &types.Poa{
			Organization: h.makeOrganisationFromProto(poa.Organization),
			Identifier:   poa.Identifier,
			File:         &poa.File,
			Result:       &poa.Result,
			Review: &types.Review{
				Resubmit: poa.Review.Resubmit,
				Message:  &poa.Review.Message,
				Ts:       poa.Review.Ts.AsTime().Unix(),
			},
			Actions:  poaActions,
			Status:   h.mapPOAStatuses(poa.Status),
			StatusTs: poa.StatusTs.AsTime().Unix(),
			Ts:       poa.Ts.AsTime().Unix(),
		}
	}

	return poas_
}

func (h *helpersfactory) MakeBeneficiaryFromProto(beneficiary *protoTypes.Beneficiary) *types.Beneficiary {
	result := types.Beneficiary{}
	if beneficiary != nil {

		beneficiaryAccounts := make([]*types.BeneficiaryAccount, len(beneficiary.Accounts))
		for index, beneficiaryAccount := range beneficiary.Accounts {
			beneficiaryAccounts[index] = h.MakeBeneficiaryAccountFromProto(beneficiaryAccount)
		}

		result = types.Beneficiary{
			ID:                beneficiary.Id,
			Customer:          h.makeCustomerFromProto(beneficiary.Customer),
			Name:              beneficiary.Name,
			Accounts:          beneficiaryAccounts,
			TransactionsCount: int64(beneficiary.TransactionsCount),
			Status:            h.MapProtoBeneficiaryStatuses(beneficiary.Status),
			StatusTs:          beneficiary.StatusTs.AsTime().Unix(),
			Ts:                beneficiary.Ts.AsTime().Unix(),
		}
	}
	return &result
}

func (h *helpersfactory) MakeBeneficiaryAccountFromProto(beneficiaryAccount *protoTypes.BeneficiaryAccount) *types.BeneficiaryAccount {
	result := types.BeneficiaryAccount{}

	if beneficiaryAccount != nil {

		currency := &types.Currency{}
		if beneficiaryAccount.Currency != nil {
			currency = &types.Currency{
				ID:     beneficiaryAccount.Currency.Id,
				Name:   beneficiaryAccount.Currency.Name,
				Code:   beneficiaryAccount.Currency.Code,
				Symbol: beneficiaryAccount.Currency.Symbol,
			}
		}

		result = types.BeneficiaryAccount{
			ID:            beneficiaryAccount.Id,
			Beneficiary:   h.MakeBeneficiaryFromProto(beneficiaryAccount.Beneficiary),
			Name:          &beneficiaryAccount.Name,
			Account:       h.MakeAccountFromProto(beneficiaryAccount.Account),
			Currency:      currency,
			AccountNumber: beneficiaryAccount.AccountNumber,
			Code:          beneficiaryAccount.Code,
			Status:        h.MapBeneficiaryAccountStatuses(beneficiaryAccount.Status),
			StatusTs:      beneficiaryAccount.StatusTs.AsTime().Unix(),
			Ts:            beneficiaryAccount.Ts.AsTime().Unix(),
		}
	}

	return &result
}

func (h *helpersfactory) GetProtoBeneficiaryStatuses(val types.BeneficiaryStatuses) protoTypes.Beneficiary_BeneficiaryStatuses {
	switch val {
	case types.BeneficiaryStatusesActive:
		return protoTypes.Beneficiary_ACTIVE
	case types.BeneficiaryStatusesInactive:
		return protoTypes.Beneficiary_INACTIVE
	default:
		return -1
	}
}

func (h *helpersfactory) MakeProductFromProto(product *protoTypes.Product) *types.Product {
	if product == nil {
		return &types.Product{
			Currency: &types.Currency{},
		}
	}

	termLength := int64(product.TermLength)
	interestRate := float64(product.InterestRate)
	minimumOpeningBalance := float64(product.MinimumOpeningBalance)

	mambu := &types.ProductMambu{}
	if product.Mambu != nil {
		mambu.EncodedKey = &product.Mambu.EncodedKey
	}
	return &types.Product{
		ID:                    product.Id,
		Type:                  h.MapProtoProductTypes(product.Type),
		Currency:              h.MakeCurrencyFromProto(product.Currency),
		TermLength:            &termLength,
		TermUnit:              &product.TermUnit,
		InterestRate:          &interestRate,
		MinimumOpeningBalance: &minimumOpeningBalance,
		Mambu:                 mambu,
		Status:                h.MapProtoProductStatuses(product.Status),
		StatusTs:              product.StatusTs.AsTime().Unix(),
		Ts:                    product.Ts.AsTime().Unix(),
	}

}

func (h *helpersfactory) MapLinkedTransactionTypes(val protoTypes.LinkedTransaction_LinkedTransactionTypes) types.LinkedTransactionTypes {
	switch val {
	case protoTypes.LinkedTransaction_APPLY_FEE:
		return types.LinkedTransactionTypesApplyFee
	case protoTypes.LinkedTransaction_DEPOSIT:
		return types.LinkedTransactionTypesDeposit
	case protoTypes.LinkedTransaction_WITHDRAWAL:
		return types.LinkedTransactionTypesDeposit
	default:
		return ""
	}
}

func (h *helpersfactory) MapFeeTypes(val protoTypes.Fee_FeeTypes) types.FeeTypes {
	switch val {
	case protoTypes.Fee_FIXED:
		return types.FeeTypesFixed
	case protoTypes.Fee_VARIABLE:
		return types.FeeTypesVariable
	default:
		return ""
	}
}

func (h *helpersfactory) MapProtoFeeTypes(val types.FeeTypes) protoTypes.Fee_FeeTypes {
	switch val {
	case types.FeeTypesFixed:
		return protoTypes.Fee_FIXED
	case types.FeeTypesVariable:
		return protoTypes.Fee_VARIABLE
	default:
		return -100
	}
}

func (h *helpersfactory) MapFeeStatuses(val protoTypes.Fee_FeeStatuses) types.FeeStatuses {
	switch val {
	case protoTypes.Fee_ACTIVE:
		return types.FeeStatusesActive
	case protoTypes.Fee_INACTIVE:
		return types.FeeStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) MapStaffAuditLogType(val protoTypes.StaffAuditLog_StaffAuditLogTypes) types.StaffAuditLogType {
	switch val {
	case protoTypes.StaffAuditLog_FEES:
		return types.StaffAuditLogTypeFees
	case protoTypes.StaffAuditLog_FX_RATE:
		return types.StaffAuditLogTypeFxRate
	case protoTypes.StaffAuditLog_CUSTOMER_DETAILS_UPDATE:
		return types.StaffAuditLogTypeCustomerDetailsUpdate
	default:
		return ""
	}
}

func (h *helpersfactory) MapProtoStaffAuditLogType(val types.StaffAuditLogType) protoTypes.StaffAuditLog_StaffAuditLogTypes {
	switch val {
	case types.StaffAuditLogTypeFees:
		return protoTypes.StaffAuditLog_FEES
	case types.StaffAuditLogTypeFxRate:
		return protoTypes.StaffAuditLog_FX_RATE
	case types.StaffAuditLogTypeCustomerDetailsUpdate:
		return protoTypes.StaffAuditLog_CUSTOMER_DETAILS_UPDATE
	default:
		return -1
	}
}

func (h *helpersfactory) MakeAccountFromProto(account *protoTypes.Account) *types.Account {
	if account == nil {
		return &types.Account{
			ID:       "",
			Customer: &types.Customer{},
			Product: &types.Product{
				Currency: &types.Currency{},
			},
		}
	}

	balances := &types.AccountBalances{}
	if account.Balances != nil {
		balances.TotalBalance = float64(account.Balances.TotalBalance)
	}

	fcmb := &types.AccountFcmb{}
	if account.Fcmb != nil {
		fcmb.CifID = &account.Fcmb.CifId
		fcmb.NgnAccountNumber = &account.Fcmb.NgnAccountNumber
	}

	vault := &types.AccountVault{}
	if account.Vault != nil {
		principalAmount := float64(account.Vault.PrincipalAmount)
		interestAmount := float64(account.Vault.InterestAccumulated)

		vault.InterestAccumulated = &interestAmount
		vault.PrincipalAmount = &principalAmount
	}

	return &types.Account{
		ID:            account.Id,
		Customer:      h.makeCustomerFromProto(account.Customer),
		Product:       h.MakeProductFromProto(account.Product),
		Name:          account.Name,
		Iban:          &account.Iban,
		AccountNumber: &account.AccountNumber,
		Code:          &account.Code,
		MaturityDate:  &account.MaturityDate,
		Balances:      balances,
		Mambu:         h.makeMambuAccountFromProto(account.Mambu),
		Fcmb:          fcmb,
		Vault:         vault,
		Status:        h.MapProtoAccountStatuses(account.Status),
		StatusTs:      account.StatusTs.AsTime().Unix(),
		Ts:            account.Ts.AsTime().Unix(),
	}
}

func (h *helpersfactory) makeMambuAccountFromProto(mambu *protoTypes.AccountMambu) *types.AccountMambu {
	if mambu == nil {
		return &types.AccountMambu{}
	}

	return &types.AccountMambu{
		BranchKey:  &mambu.BranchKey,
		EncodedKey: &mambu.EncodedKey,
	}
}

func (h *helpersfactory) MapProtoAccountStatuses(val protoTypes.Account_AccountStatuses) types.AccountStatuses {
	switch val {
	case protoTypes.Account_ACTIVE:
		return types.AccountStatusesActive
	case protoTypes.Account_INACTIVE:
		return types.AccountStatusesInactive
	default:
		return ""
	}
}

func (h *helpersfactory) MapAccountStatuses(val types.AccountStatuses) protoTypes.Account_AccountStatuses {
	switch val {
	case types.AccountStatusesActive:
		return protoTypes.Account_ACTIVE
	case types.AccountStatusesInactive:
		return protoTypes.Account_INACTIVE
	default:
		return -1
	}
}

func (h *helpersfactory) MakeTransactionFromProto(transaction *protoTypes.Transaction) *types.Transaction {
	result := &types.Transaction{}

	if transaction != nil {

		fees := make([]*types.TransactionFee, len(transaction.Fees))
		for index, fee := range transaction.Fees {
			fees[index] = &types.TransactionFee{
				ID:     fee.Id,
				Amount: float64(fee.Amount),
			}
		}

		linkedTransactions := make([]*types.LinkedTransaction, len(transaction.LinkedTransactions))
		for index, linkedTransaction := range transaction.LinkedTransactions {
			linkedTransactions[index] = &types.LinkedTransaction{
				ID:   linkedTransaction.Id,
				Type: h.MapLinkedTransactionTypes(linkedTransaction.Type),
				Currency: &types.Currency{
					ID:     linkedTransaction.Currency.Id,
					Symbol: linkedTransaction.Currency.Symbol,
					Code:   linkedTransaction.Currency.Code,
					Name:   linkedTransaction.Currency.Name,
				},
				Amount: float64(linkedTransaction.Amount),
				Source: &types.LinkedTransactionSource{
					Customer:    h.makeCustomerFromProto(linkedTransaction.Source.Customer),
					Account:     h.MakeAccountFromProto(linkedTransaction.Source.Account),
					AccountData: linkedTransaction.Source.AccountData,
				},
				Target: &types.LinkedTransactionTarget{
					Account:            h.MakeAccountFromProto(linkedTransaction.Target.Account),
					BeneficiaryAccount: h.MakeBeneficiaryAccountFromProto(linkedTransaction.Target.BeneficiaryAccount),
					AccountData:        linkedTransaction.Target.AccountData,
				},
				Mambu: &types.LinkedTransactionMambu{
					TransactionEncodedKey: linkedTransaction.Mambu.TransactionEncodedKey,
				},
				Fcmb: &types.LinkedTransactionFcmb{
					TransactionIdentifier: linkedTransaction.Fcmb.TransactionIdentifier,
				},
				Status:   h.MapLinkedTransactionStatuses(linkedTransaction.Status),
				StatusTs: linkedTransaction.StatusTs.AsTime().Unix(),
				Ts:       linkedTransaction.Ts.AsTime().Unix(),
			}
		}

		var exchangeRate *types.ExchangeRate
		if transaction.ExchangeRate.Id != "" {
			exchangeRate = &types.ExchangeRate{
				ID: transaction.ExchangeRate.Id,
				BaseCurrency: &types.Currency{
					ID:     transaction.ExchangeRate.BaseCurrency.Id,
					Symbol: transaction.ExchangeRate.BaseCurrency.Symbol,
					Code:   transaction.ExchangeRate.BaseCurrency.Code,
					Name:   transaction.ExchangeRate.BaseCurrency.Name,
				},
				TargetCurrency: &types.Currency{
					ID:     transaction.ExchangeRate.TargetCurrency.Id,
					Symbol: transaction.ExchangeRate.TargetCurrency.Symbol,
					Code:   transaction.ExchangeRate.TargetCurrency.Code,
					Name:   transaction.ExchangeRate.TargetCurrency.Name,
				},
				BuyPrice:  float64(transaction.ExchangeRate.BuyPrice),
				SalePrice: float64(transaction.ExchangeRate.SalePrice),
				Ts:        transaction.ExchangeRate.Ts.AsTime().Unix(),
			}
		}

		result = &types.Transaction{
			ID: transaction.Id,
			TransactionType: &types.TransactionType{
				ID:       transaction.TransactionType.Id,
				Name:     transaction.TransactionType.Name,
				Status:   h.MapTransactionTypeStatus(transaction.TransactionType.Status),
				StatusTs: transaction.StatusTs.AsTime().Unix(),
				Ts:       transaction.Ts.AsTime().Unix(),
			},
			Reference:    transaction.Reference,
			Fees:         fees,
			ExchangeRate: exchangeRate,
			Source: &types.TransactionSource{
				Customer:                h.makeCustomerFromProto(transaction.Source.Customer),
				Account:                 h.MakeAccountFromProto(transaction.Source.Account),
				Amount:                  float64(transaction.Source.Amount),
				BalanceAfterTransaction: float64(transaction.Source.BalanceAfterTransaction),
			},
			Target: &types.TransactionTarget{
				Customer:                h.makeCustomerFromProto(transaction.Target.Customer),
				Beneficiary:             h.MakeBeneficiaryFromProto(transaction.Target.Beneficiary),
				Account:                 h.MakeAccountFromProto(transaction.Target.Account),
				BeneficiaryAccount:      h.MakeBeneficiaryAccountFromProto(transaction.Target.BeneficiaryAccount),
				Amount:                  float64(transaction.Target.Amount),
				BalanceAfterTransaction: float64(transaction.Target.BalanceAfterTransaction),
			},
			IdempotencyKey:     transaction.IdempotencyKey,
			LinkedTransactions: linkedTransactions,
			Status:             h.MapTransactionStatuses(transaction.Status),
			StatusTs:           transaction.StatusTs.AsTime().Unix(),
			Ts:                 transaction.Ts.AsTime().Unix(),
		}
	}

	return result
}

func (h *helpersfactory) MakeFeesFromProto(protoFees []*protoTypes.Fee) []*types.Fee {
	fees := make([]*types.Fee, len(protoFees))

	for index, fee := range protoFees {

		feeBoundaries := make([]*types.FeeBoundaries, len(fee.Boundaries))
		for i, boundary := range fee.Boundaries {
			lower := float64(boundary.Lower)
			upper := float64(boundary.Upper)
			amount := float64(boundary.Amount)
			percentage := float64(boundary.Percentage)

			feeBoundaries[i] = &types.FeeBoundaries{
				Lower:      &lower,
				Upper:      &upper,
				Amount:     &amount,
				Percentage: &percentage,
			}
		}

		fees[index] = &types.Fee{
			ID: fee.Id,
			TransactionType: &types.TransactionType{
				ID:       fee.TransactionType.Id,
				Name:     fee.TransactionType.Name,
				Status:   h.MapTransactionTypeStatus(fee.TransactionType.Status),
				StatusTs: fee.TransactionType.StatusTs.AsTime().Unix(),
				Ts:       fee.TransactionType.Ts.AsTime().Unix(),
			},
			Type:       h.MapFeeTypes(fee.Type),
			Boundaries: feeBoundaries,
			Status:     h.MapFeeStatuses(fee.Status),
			StatusTs:   fee.StatusTs.AsTime().Unix(),
			Ts:         fee.Ts.AsTime().Unix(),
		}
	}

	return fees
}

func (h *helpersfactory) makeFeeFromProto(protoFee *protoTypes.Fee) *types.Fee {
	result := &types.Fee{}
	if protoFee != nil {
		feeBoundaries := make([]*types.FeeBoundaries, len(protoFee.Boundaries))
		for i, boundary := range protoFee.Boundaries {
			lower := float64(boundary.Lower)
			upper := float64(boundary.Upper)
			amount := float64(boundary.Amount)
			percentage := float64(boundary.Percentage)

			feeBoundaries[i] = &types.FeeBoundaries{
				Lower:      &lower,
				Upper:      &upper,
				Amount:     &amount,
				Percentage: &percentage,
			}
		}

		result = &types.Fee{
			ID: protoFee.Id,
			TransactionType: &types.TransactionType{
				ID:       protoFee.TransactionType.Id,
				Name:     protoFee.TransactionType.Name,
				Status:   h.MapTransactionTypeStatus(protoFee.TransactionType.Status),
				StatusTs: protoFee.TransactionType.StatusTs.AsTime().Unix(),
				Ts:       protoFee.TransactionType.Ts.AsTime().Unix(),
			},
			Type:       h.MapFeeTypes(protoFee.Type),
			Boundaries: feeBoundaries,
			Status:     h.MapFeeStatuses(protoFee.Status),
			StatusTs:   protoFee.StatusTs.AsTime().Unix(),
			Ts:         protoFee.Ts.AsTime().Unix(),
		}
	}

	return result
}

func (h *helpersfactory) MakeExchangeRateFromProto(exchangeRate *protoTypes.ExchangeRate) *types.ExchangeRate {
	return &types.ExchangeRate{
		ID: exchangeRate.Id,
		BaseCurrency: &types.Currency{
			ID:     exchangeRate.BaseCurrency.Id,
			Symbol: exchangeRate.BaseCurrency.Symbol,
			Code:   exchangeRate.BaseCurrency.Code,
			Name:   exchangeRate.BaseCurrency.Name,
		},
		TargetCurrency: &types.Currency{
			ID:     exchangeRate.TargetCurrency.Id,
			Symbol: exchangeRate.TargetCurrency.Symbol,
			Code:   exchangeRate.TargetCurrency.Code,
			Name:   exchangeRate.TargetCurrency.Name,
		},
		BuyPrice:  float64(exchangeRate.BuyPrice),
		SalePrice: float64(exchangeRate.SalePrice),
		Ts:        exchangeRate.Ts.AsTime().Unix(),
	}
}

func (h *helpersfactory) MapProtoCredentialTypes(val protoTypes.IdentityCredentials_IdentityCredentialsTypes) types.IdentityCredentialsTypes {
	switch val {
	case protoTypes.IdentityCredentials_LOGIN:
		return types.IdentityCredentialsTypesLogin
	case protoTypes.IdentityCredentials_PIN:
		return types.IdentityCredentialsTypesPin
	default:
		return ""
	}
}

func (h *helpersfactory) MapCredentialTypes(val types.IdentityCredentialsTypes) protoTypes.IdentityCredentials_IdentityCredentialsTypes {
	switch val {
	case types.IdentityCredentialsTypesLogin:
		return protoTypes.IdentityCredentials_LOGIN
	case types.IdentityCredentialsTypesPin:
		return protoTypes.IdentityCredentials_PIN
	default:
		return -1
	}
}

func (h *helpersfactory) MakeStaffAuditLogFromProto(staffAuditLog *protoTypes.StaffAuditLog) *types.StaffAuditLog {
	result := &types.StaffAuditLog{}
	if staffAuditLog != nil {
		var oldStaffAuditLogValue, newStaffAuditLogValue types.StaffAuditLogValue
		switch staffAuditLog.Type {
		case protoTypes.StaffAuditLog_FEES:
			// oldValue
			if staffAuditLog.OldValue != nil && staffAuditLog.OldValue.Data != nil {
				oldFeeProto := staffAuditLog.OldValue.Data.(*protoTypes.StaffAuditLogValue_Fee).Fee
				if oldFeeProto != nil {
					oldStaffAuditLogValue = h.makeFeeFromProto(oldFeeProto)
				}
			}

			// newValue
			if staffAuditLog.NewValue != nil && staffAuditLog.NewValue.Data != nil {
				newFeeProto := staffAuditLog.NewValue.Data.(*protoTypes.StaffAuditLogValue_Fee).Fee
				if newFeeProto != nil {
					newStaffAuditLogValue = h.makeFeeFromProto(newFeeProto)
				}
			}

		case protoTypes.StaffAuditLog_FX_RATE:
			// oldValue
			if staffAuditLog.OldValue != nil && staffAuditLog.OldValue.Data != nil {
				oldFxRateProto := staffAuditLog.OldValue.Data.(*protoTypes.StaffAuditLogValue_ExchangeRate).ExchangeRate
				if oldFxRateProto != nil {
					oldStaffAuditLogValue = h.MakeExchangeRateFromProto(oldFxRateProto)
				}
			}

			// newValue
			if staffAuditLog.NewValue != nil && staffAuditLog.NewValue.Data != nil {
				newFxRateProto := staffAuditLog.NewValue.Data.(*protoTypes.StaffAuditLogValue_ExchangeRate).ExchangeRate
				if newFxRateProto != nil {
					newStaffAuditLogValue = h.MakeExchangeRateFromProto(newFxRateProto)
				}
			}

		case protoTypes.StaffAuditLog_CUSTOMER_DETAILS_UPDATE:
			// oldValue
			if staffAuditLog.OldValue != nil && staffAuditLog.OldValue.Data != nil {
				oldCustomerProto := staffAuditLog.OldValue.Data.(*protoTypes.StaffAuditLogValue_Customer).Customer
				if oldCustomerProto != nil {
					oldStaffAuditLogValue = h.makeCustomerFromProto(oldCustomerProto)
				}
			}

			// newValue
			if staffAuditLog.NewValue != nil && staffAuditLog.NewValue.Data != nil {
				newCustomerProto := staffAuditLog.NewValue.Data.(*protoTypes.StaffAuditLogValue_Customer).Customer
				if newCustomerProto != nil {
					newStaffAuditLogValue = h.makeCustomerFromProto(newCustomerProto)
				}
			}

		default:
			return result
		}

		result = &types.StaffAuditLog{
			ID:       staffAuditLog.Id,
			Staff:    h.makeStaffFromProto(staffAuditLog.Staff),
			OldValue: oldStaffAuditLogValue,
			NewValue: newStaffAuditLogValue,
			Type:     h.MapStaffAuditLogType(staffAuditLog.Type),
			Ts:       staffAuditLog.Ts.AsTime().Unix(),
		}
	}
	return result
}

func (h *helpersfactory) makeStatesFromProto(states []*protoTypes.State) []*types.State {
	if states == nil {
		return make([]*types.State, 0)
	}

	newStates := make([]*types.State, len(states))
	for i, state := range states {
		newStates[i] = &types.State{
			IsoCode: state.IsoCode,
			Name:    state.Name,
		}
	}
	return newStates
}

func (h *helpersfactory) MapScheduledTransactionRepeatType(val protoTypes.ScheduledTransaction_ScheduledTransactionRepeatType) types.ScheduledTransactionRepeatType {
	switch val {
	case protoTypes.ScheduledTransaction_ONE_TIME:
		return types.ScheduledTransactionRepeatTypeOneTime
	case protoTypes.ScheduledTransaction_WEEKLY:
		return types.ScheduledTransactionRepeatTypeWeekly
	case protoTypes.ScheduledTransaction_MONTHLY:
		return types.ScheduledTransactionRepeatTypeMonthly
	case protoTypes.ScheduledTransaction_ANNUALLY:
		return types.ScheduledTransactionRepeatTypeAnnually
	default:
		return ""
	}
}

func (h *helpersfactory) MapProtoScheduledTransactionRepeatType(val types.ScheduledTransactionRepeatType) protoTypes.ScheduledTransaction_ScheduledTransactionRepeatType {
	switch val {
	case types.ScheduledTransactionRepeatTypeOneTime:
		return protoTypes.ScheduledTransaction_ONE_TIME
	case types.ScheduledTransactionRepeatTypeWeekly:
		return protoTypes.ScheduledTransaction_WEEKLY
	case types.ScheduledTransactionRepeatTypeMonthly:
		return protoTypes.ScheduledTransaction_MONTHLY
	case types.ScheduledTransactionRepeatTypeAnnually:
		return protoTypes.ScheduledTransaction_ANNUALLY
	default:
		return -100
	}
}

func (h *helpersfactory) MapScheduledTransactionStatus(val protoTypes.ScheduledTransaction_ScheduledTransactionStatuses) types.ScheduledTransactionStatus {
	switch val {
	case protoTypes.ScheduledTransaction_ACTIVE:
		return types.ScheduledTransactionStatusActive
	case protoTypes.ScheduledTransaction_INACTIVE:
		return types.ScheduledTransactionStatusInactive
	default:
		return ""
	}
}

func (h *helpersfactory) MapProtoScheduledTransactionStatus(val types.ScheduledTransactionStatus) protoTypes.ScheduledTransaction_ScheduledTransactionStatuses {
	switch val {
	case types.ScheduledTransactionStatusActive:
		return protoTypes.ScheduledTransaction_ACTIVE
	case types.ScheduledTransactionStatusInactive:
		return protoTypes.ScheduledTransaction_INACTIVE
	default:
		return -100
	}
}

func (h *helpersfactory) MapProtoCustomerPreferenceType(val types.CustomerPreferencesTypes) protoTypes.CustomerPreferences_CustomerPreferencesTypes {
	switch val {
	case types.CustomerPreferencesTypesMarketing:
		return protoTypes.CustomerPreferences_MARKETING
	default:
		return protoTypes.CustomerPreferences_MARKETING
	}
}

func (h *helpersfactory) MapProtoFAQTypes(val types.FilterType) onboarding.GetFAQRequest_FAQFilter {
	switch val {
	case types.FilterTypeSearch:
		return onboarding.GetFAQRequest_SEARCH
	case types.FilterTypeAll:
		return onboarding.GetFAQRequest_ALL
	case types.FilterTypeSpecific:
		return onboarding.GetFAQRequest_SPECIFIC
	default:
		return onboarding.GetFAQRequest_FEATURED
	}
}

func (h *helpersfactory) makeFAQFromProto(faq *protoTypes.FAQ) *types.Faq {
	var result *types.Faq
	tags, topic := h.makeTagsAndTopicFromProto(faq.Tags, faq.Topic)
	result = &types.Faq{
		ID:         faq.Id,
		Question:   faq.Question,
		Answer:     faq.Answer,
		IsFeatured: faq.IsFeatured,
		Tags:       tags,
		Ts:         faq.Ts.AsTime().Unix(),
		UpdateTs:   faq.UpdateTs.AsTime().Unix(),
		Topic:      topic,
	}
	return result
}

func (h *helpersfactory) makeTagsAndTopicFromProto(tags []string, topics protoTypes.FAQ_FAQTopic) ([]*string, types.FAQTopic) {
	var faqTag []*string
	if tags != nil {
		faqTag = make([]*string, len(tags))
		for i, tag := range tags {
			faqTag[i] = &tag
		}
	}
	switch topics {
	case protoTypes.FAQ_ACCOUNT_OPENING:
		return faqTag, types.FAQTopicAccountOpening
	case protoTypes.FAQ_FUNDING:
		return faqTag, types.FAQTopicFunding
	case protoTypes.FAQ_PAYMENTS:
		return faqTag, types.FAQTopicPayments
	case protoTypes.FAQ_STATEMENT:
		return faqTag, types.FAQTopicStatement
	case protoTypes.FAQ_SECURITY:
		return faqTag, types.FAQTopicSecurity
	default:
		return faqTag, types.FAQTopicAboutRova
	}
}
func (r *helpersfactory) paginationDetails(keywords *string, first *int64, after *string, last *int64, before *string) *onboarding.GetFAQRequest {
	var request onboarding.GetFAQRequest
	//var helper helpersfactory
	if keywords != nil {
		request.Keywords = *keywords
	}
	if first != nil {
		request.First = int32(*first)
	}
	if after != nil {
		request.After = *after
	}
	if last != nil {
		request.Last = int32(*last)
	}
	if before != nil {
		request.Before = *before
	}
	return &request
}
