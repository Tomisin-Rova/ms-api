package graph

import (
	"ms.api/protos/pb/auth"
	customerTypes "ms.api/protos/pb/types"
	"ms.api/types"
)

var (
	// Response messages
	authFailedMessage = "User authentication failed"
)

type Helper interface {
	MapQuestionaryStatus(val types.QuestionaryStatuses) customerTypes.Questionary_QuestionaryStatuses
	MapQuestionaryType(val types.QuestionaryTypes) customerTypes.Questionary_QuestionaryTypes
	MapProtoQuesionaryStatus(val customerTypes.Questionary_QuestionaryStatuses) types.QuestionaryStatuses
	MapProtoQuestionaryType(val customerTypes.Questionary_QuestionaryTypes) types.QuestionaryTypes
	GetDeveicePreferenceTypesIndex(val types.DevicePreferencesTypes) int32
	GetProtoCustomerStatuses(val types.CustomerStatuses) customerTypes.Customer_CustomerStatuses
	DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) customerTypes.DeviceToken_DeviceTokenTypes
	PreferenceInputFromModel(input types.DevicePreferencesTypes) customerTypes.DevicePreferences_DevicePreferencesTypes
	StaffLoginTypeFromModel(input types.AuthType) auth.StaffLoginRequest_AuthType
}

type helpersfactory struct{}

func (h *helpersfactory) MapQuestionaryStatus(val types.QuestionaryStatuses) customerTypes.Questionary_QuestionaryStatuses {
	switch val {
	case types.QuestionaryStatusesActive:
		return customerTypes.Questionary_ACTIVE
	case types.QuestionaryStatusesInactive:
		return customerTypes.Questionary_INACTIVE
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) MapQuestionaryType(val types.QuestionaryTypes) customerTypes.Questionary_QuestionaryTypes {
	switch val {
	case types.QuestionaryTypesReasons:
		return customerTypes.Questionary_REASONS
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) MapProtoQuesionaryStatus(val customerTypes.Questionary_QuestionaryStatuses) types.QuestionaryStatuses {
	switch val {
	case customerTypes.Questionary_ACTIVE:
		return types.QuestionaryStatusesActive
	case customerTypes.Questionary_INACTIVE:
		return types.QuestionaryStatusesInactive
	}

	return ""
}

func (h *helpersfactory) MapProtoQuestionaryType(val customerTypes.Questionary_QuestionaryTypes) types.QuestionaryTypes {
	switch val {
	case customerTypes.Questionary_REASONS:
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

func (h *helpersfactory) MapProtoCustomerStatuses(val customerTypes.Customer_CustomerStatuses) types.CustomerStatuses {
	switch val {
	case customerTypes.Customer_SIGNEDUP:
		return types.CustomerStatusesSignedup
	case customerTypes.Customer_REGISTERED:
		return types.CustomerStatusesRegistered
	case customerTypes.Customer_VERIFIED:
		return types.CustomerStatusesVerified
	case customerTypes.Customer_ONBOARDED:
		return types.CustomerStatusesOnboarded
	case customerTypes.Customer_REJECTED:
		return types.CustomerStatusesRegistered
	case customerTypes.Customer_EXITED:
		return types.CustomerStatusesExited
	default:
		// should never happen
		return ""
	}
}

func (h *helpersfactory) GetProtoCustomerStatuses(val types.CustomerStatuses) customerTypes.Customer_CustomerStatuses {
	switch val {
	case types.CustomerStatusesSignedup:
		return customerTypes.Customer_SIGNEDUP
	case types.CustomerStatusesRegistered:
		return customerTypes.Customer_REGISTERED
	case types.CustomerStatusesVerified:
		return customerTypes.Customer_VERIFIED
	case types.CustomerStatusesOnboarded:
		return customerTypes.Customer_ONBOARDED
	case types.CustomerStatusesRejected:
		return customerTypes.Customer_REJECTED
	case types.CustomerStatusesExited:
		return customerTypes.Customer_EXITED
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) customerTypes.DeviceToken_DeviceTokenTypes {
	switch tokenType {
	default:
		return customerTypes.DeviceToken_FIREBASE
	}
}

func (h *helpersfactory) PreferenceInputFromModel(input types.DevicePreferencesTypes) customerTypes.DevicePreferences_DevicePreferencesTypes {
	switch input {
	case types.DevicePreferencesTypesPush:
		return customerTypes.DevicePreferences_PUSH
	case types.DevicePreferencesTypesBiometrics:
		return customerTypes.DevicePreferences_BIOMETRICS
	default:
		return customerTypes.DevicePreferences_PUSH
	}
}

func (h *helpersfactory) MapProtoStaffStatuses(val customerTypes.Staff_StaffStatuses) types.StaffStatuses {
	switch val {
	case customerTypes.Staff_ACTIVE:
		return types.StaffStatusesActive
	case customerTypes.Staff_INACTIVE:
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

func (h *helpersfactory) MapCDDStatusesFromModel(val types.CDDStatuses) customerTypes.CDD_CDDStatuses {
	switch val {
	case types.CDDStatusesPending:
		return customerTypes.CDD_PENDING
	case types.CDDStatusesManualReview:
		return customerTypes.CDD_MANUAL_REVIEW
	case types.CDDStatusesApproved:
		return customerTypes.CDD_APPROVED
	case types.CDDStatusesDeclined:
		return customerTypes.CDD_DECLINED
	default:
		return -1
	}
}

func (h *helpersfactory) MapProtoCDDStatuses(val customerTypes.CDD_CDDStatuses) types.CDDStatuses {
	switch val {
	case customerTypes.CDD_PENDING:
		return types.CDDStatusesPending
	case customerTypes.CDD_MANUAL_REVIEW:
		return types.CDDStatusesManualReview
	case customerTypes.CDD_APPROVED:
		return types.CDDStatusesApproved
	case customerTypes.CDD_DECLINED:
		return types.CDDStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) makeAddressFromProto(adddresses []*customerTypes.Address) []*types.Address {
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
		}
	}
	return addresses_
}

func (h *helpersfactory) makePhonesFromProto(phones []*customerTypes.Phone) []*types.Phone {
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

func (h *helpersfactory) makeCustomerFromProto(customer *customerTypes.Customer) *types.Customer {
	return &types.Customer{
		ID:        customer.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Dob:       customer.Dob,
		Bvn:       &customer.Bvn,
		Addresses: h.makeAddressFromProto(customer.Addresses),
		Phones:    h.makePhonesFromProto(customer.Phones),
		Email: &types.Email{
			Address:  customer.Email.Address,
			Verified: customer.Email.Verified,
		},
		Status:   h.MapProtoCustomerStatuses(customer.Status),
		StatusTs: customer.StatusTs.AsTime().Unix(),
		Ts:       customer.Ts.AsTime().Unix(),
	}
}

func (h *helpersfactory) makeStaffFromProto(staff *customerTypes.Staff) *types.Staff {
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

func (h *helpersfactory) mapOrganizationStatuses(val customerTypes.Organization_OrganizationStatuses) types.OrganizationStatuses {
	switch val {
	case customerTypes.Organization_ACTIVE:
		return types.OrganizationStatusesActive
	case customerTypes.Organization_INACTIVE:
		return types.OrganizationStatusesInactive

	default:
		return ""
	}
}

func (h *helpersfactory) makeOrganisationFromProto(organization *customerTypes.Organization) *types.Organization {
	return &types.Organization{
		ID:       organization.Id,
		Name:     organization.Name,
		Status:   h.mapOrganizationStatuses(organization.Status),
		StatusTs: organization.StatusTs.AsTime().Unix(),
		Ts:       organization.Ts.AsTime().Unix(),
	}
}

func (h *helpersfactory) mapPOAStatuses(val customerTypes.POA_POAStatuses) types.POAStatuses {
	switch val {
	case customerTypes.POA_APPROVED:
		return types.POAStatusesApproved
	case customerTypes.POA_MANUAL_REVIEW:
		return types.POAStatusesManualReview
	case customerTypes.POA_PENDING:
		return types.POAStatusesPending
	case customerTypes.POA_DECLINED:
		return types.POAStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapPOAActionStatuses(val customerTypes.POAAction_POAStatuses) types.POAStatuses {
	switch val {
	case customerTypes.POAAction_APPROVED:
		return types.POAStatusesApproved
	case customerTypes.POAAction_MANUAL_REVIEW:
		return types.POAStatusesManualReview
	case customerTypes.POAAction_PENDING:
		return types.POAStatusesPending
	case customerTypes.POAAction_DECLINED:
		return types.POAStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapPOAActionTypes(val customerTypes.POAAction_POAActionTypes) types.POAActionTypes {
	switch val {
	case customerTypes.POAAction_CHANGE_STATUS:
		return types.POAActionTypesChangeStatus
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCStatuses(val customerTypes.KYC_KYCStatuses) types.KYCStatuses {
	switch val {
	case customerTypes.KYC_APPROVED:
		return types.KYCStatusesApproved
	case customerTypes.KYC_MANUAL_REVIEW:
		return types.KYCStatusesManualReview
	case customerTypes.KYC_PENDING:
		return types.KYCStatusesPending
	case customerTypes.KYC_DECLINED:
		return types.KYCStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCActionTypes(val customerTypes.KYCAction_KYCActionTypes) types.KYCActionTypes {
	switch val {
	case customerTypes.KYCAction_CHANGE_STATUS:
		return types.KYCActionTypesChangeStatus
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCActionStatuses(val customerTypes.KYCAction_KYCStatuses) types.KYCStatuses {
	switch val {
	case customerTypes.KYCAction_APPROVED:
		return types.KYCStatusesApproved
	case customerTypes.KYCAction_MANUAL_REVIEW:
		return types.KYCStatusesManualReview
	case customerTypes.KYCAction_PENDING:
		return types.KYCStatusesPending
	case customerTypes.KYCAction_DECLINED:
		return types.KYCStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCReportStatuses(val customerTypes.Report_ReportStatuses) types.ReportStatuses {
	switch val {
	case customerTypes.Report_APPROVED:
		return types.ReportStatusesApproved
	case customerTypes.Report_MANUAL_REVIEW:
		return types.ReportStatusesManualReview
	case customerTypes.Report_PENDING:
		return types.ReportStatusesPending
	case customerTypes.Report_DECLINED:
		return types.ReportStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapKYCReportTypes(val customerTypes.Report_ReportTypes) types.KYCTypes {
	switch val {
	case customerTypes.Report_DOCUMENT:
		return types.KYCTypesDocument
	case customerTypes.Report_FACIAL_VIDEO:
		return types.KYCTypesFacialVideo
	default:
		return ""
	}
}

func (h *helpersfactory) mapAMLStatuses(val customerTypes.AML_AMLStatuses) types.AMLStatuses {
	switch val {
	case customerTypes.AML_APPROVED:
		return types.AMLStatusesApproved
	case customerTypes.AML_MANUAL_REVIEW:
		return types.AMLStatusesManualReview
	case customerTypes.AML_PENDING:
		return types.AMLStatusesPending
	case customerTypes.AML_DECLINED:
		return types.AMLStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapAMLActionStatuses(val customerTypes.AMLAction_AMLStatuses) types.AMLStatuses {
	switch val {
	case customerTypes.AMLAction_APPROVED:
		return types.AMLStatusesApproved
	case customerTypes.AMLAction_MANUAL_REVIEW:
		return types.AMLStatusesManualReview
	case customerTypes.AMLAction_PENDING:
		return types.AMLStatusesPending
	case customerTypes.AMLAction_DECLINED:
		return types.AMLStatusesDeclined
	default:
		return ""
	}
}

func (h *helpersfactory) mapAMLActionTypes(val customerTypes.AMLAction_AMLActionTypes) types.AMLActionTypes {
	switch val {
	case customerTypes.AMLAction_CHANGE_STATUS:
		return types.AMLActionTypesChangeStatus
	default:
		return ""
	}
}

func (h *helpersfactory) makeAMLsFromProto(amls []*customerTypes.AML) []*types.Aml {
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

func (h *helpersfactory) makeKYCsFromProto(kycs []*customerTypes.KYC) []*types.Kyc {
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

func (h *helpersfactory) makePOAsFromProto(poas []*customerTypes.POA) []*types.Poa {
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
