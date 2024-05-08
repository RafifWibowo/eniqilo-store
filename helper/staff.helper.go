package helper

import "eniqilo_store/types"

func ValidateRegisterRequest(request *types.RegisterRequest) map[string]string {
	errList := make(map[string]string)

	validatePhoneNumber(request.PhoneNumber, errList)
	validateName(request.Name, errList)
	validatePassword(request.Password, errList)

	return errList
}

func ValidateLoginRequest(request *types.LoginRequest) map[string]string {
	errList := make(map[string]string)

	validatePhoneNumber(request.PhoneNumber, errList)
	validatePassword(request.Password, errList)

	return errList
}

func validatePhoneNumber(phone string, errList map[string]string) {
	if phone == "" {
		errList["phoneNumber"] = "Phone number can't be null."
		return
	}
	if phone[0:1] != "+" {
		errList["phoneNumber"] = "Not in phone number format."
		return
	}
	if len(phone) < 10 || len(phone) > 16 {
		errList["phoneNumber"] = "Phone number must be between 10 and 16 characters in length."
		return
	}
}

func validateName(name string, errList map[string]string) {
	if name == "" {
		errList["name"] = "Name can't be null."
		return
	}
	if len(name) < 5 || len(name) > 50 {
		errList["name"] = "Name must be between 5 and 50 characters in length."
		return
	}
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