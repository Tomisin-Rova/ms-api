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
