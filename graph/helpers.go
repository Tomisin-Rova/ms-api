package graph

import (
	customerTypes "ms.api/protos/pb/types"
	"ms.api/types"
)

var (
	// Response messages
	authFailedMessage = "User authentication failed"
)

type Helper interface {
	GetQuestionaryStatusIndex(val types.QuestionaryStatuses) int32
	GetQuestionaryTypesIndex(val types.QuestionaryTypes) int32
	GetDeveicePreferenceTypesIndex(val types.DevicePreferencesTypes) int32
	GetCustomer_CustomerStatusIndex(val customerTypes.Customer_CustomerStatuses) int32
	GetCustomerStatusIndex(val types.CustomerStatuses) int32
	DeviceTokenInputFromModel(tokenType types.DeviceTokenTypes) customerTypes.DeviceToken_DeviceTokenTypes
	PreferenceInputFromModel(input types.DevicePreferencesTypes) customerTypes.DevicePreferences_DevicePreferencesTypes
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

func (h *helpersfactory) GetCustomer_CustomerStatusIndex(val customerTypes.Customer_CustomerStatuses) int32 {
	switch val {
	case customerTypes.Customer_SIGNEDUP:
		return 0
	case customerTypes.Customer_REGISTERED:
		return 1
	case customerTypes.Customer_VERIFIED:
		return 2
	case customerTypes.Customer_ONBOARDED:
		return 3
	case customerTypes.Customer_REJECTED:
		return 4
	case customerTypes.Customer_EXITED:
		return 5
	default:
		// should never happen
		return -1
	}
}

func (h *helpersfactory) GetCustomerStatusIndex(val types.CustomerStatuses) int32 {
	switch val {
	case types.CustomerStatusesSignedup:
		return 0
	case types.CustomerStatusesRegistered:
		return 1
	case types.CustomerStatusesVerified:
		return 2
	case types.CustomerStatusesOnboarded:
		return 3
	case types.CustomerStatusesRejected:
		return 4
	case types.CustomerStatusesExited:
		return 5
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
