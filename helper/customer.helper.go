package helper

import "eniqilo_store/types"

func ValidateCustomerRegisterRequest(request *types.CustomerRegisterRequest) map[string]string {
	errList := make(map[string]string)

	ValidatePhoneNumber(request.PhoneNumber, errList)
	ValidateName(request.Name, errList)

	return errList
}