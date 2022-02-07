package graph

import (
	"ms.api/protos/pb/auth"
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
}

type helpersfactory struct{}

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
	email := &types.Email{}
	if customer.Email != nil {
		email.Address = customer.Email.Address
		email.Verified = customer.Email.Verified
	}

	return &types.Customer{
		ID:        customer.Id,
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

func (h *helpersfactory) MakeProductFromProto(product *protoTypes.Product) *types.Product {
	if product == nil {
		return &types.Product{}
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
		InterestRate:          &interestRate,
		MinimumOpeningBalance: &minimumOpeningBalance,
		Mambu:                 mambu,
		Status:                h.MapProtoProductStatuses(product.Status),
		StatusTs:              product.StatusTs.AsTime().Unix(),
		Ts:                    product.Ts.AsTime().Unix(),
	}

}

func (h *helpersfactory) MakeAccountFromProto(account *protoTypes.Account) *types.Account {
	if account == nil {
		return &types.Account{}
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
