package helper

import "eniqilo_store/types"

func ValidateRegisterRequest(request *types.RegisterRequest) map[string]string {
	errList := make(map[string]string)

	ValidatePhoneNumber(request.PhoneNumber, errList)
	ValidateName(request.Name, errList)
	validatePassword(request.Password, errList)

	return errList
}

func ValidateLoginRequest(request *types.LoginRequest) map[string]string {
	errList := make(map[string]string)

	ValidatePhoneNumber(request.PhoneNumber, errList)
	validatePassword(request.Password, errList)

	return errList
}

func validatePassword(pass string, errList map[string]string) {
	if pass == "" {
		errList["password"] = "Password can't be null."
		return
	}
	if len(pass) < 5 || len(pass) > 15 {
		errList["password"] = "Password must be between 5 and 15 characters in length."
		return
	}
}