package helper

import (
	"eniqilo_store/types"
	"net/url"
)

func ValidateProductRequest(request *types.ProductRequest) map[string]string {
	errList := make(map[string]string)

	validateString(request.Name, "name", 1, 30, errList)
	validateString(request.SKU, "sku", 1, 30, errList)
	validateCategory(request.Category, errList)
	validateUrl(request.ImageUrl, errList)
	validateString(request.Notes, "notes", 1, 200, errList)
	validateInt(&request.Price, "price", 1, -1, errList)
	validateInt(&request.Stock, "stock", 0, 100000, errList)
	validateString(request.Location, "location", 1, 200, errList)
	if !request.IsAvailable {
		errList["isAvailable"] = "isAvailable can't be null."
	}

	return errList
}

func validateString(str string, reqType string, min int, max int, errList map[string]string){
	if str == "" {
		errList[reqType] = reqType+" can't be null."
		return
	}
	if len(str) < min || len(str) > max {
		errList[reqType] = "Bad request"
		return
	}
}

func validateInt(num *int, reqType string, min int, max int, errList map[string]string) {
	if num == nil {
		errList[reqType] = reqType+" can't be null."
		return
	}

	val := *num
	if min != -1 && val < min {
		errList[reqType] = "Bad request"
		return
	}
	if max != -1 && val > max {
		errList[reqType] = "Bad request"
		return
	}
}

func validateCategory(category string, errList map[string]string) {
	if category == "" {
		errList["category"] = "Category can't be null."
		return
	}
	validCategory := map[string]bool {
		"Clothing": true,
		"Accessories": true,
		"Footwear": true,
		"Beverages": true,
	}

	if !validCategory[category] || category == "" {
		errList["category"] = "Invalid Category"
		return
	}
}

func validateUrl(u string, errList map[string]string) {
	if u == "" {
		errList["url"] = "Image Url can't be null."
		return
	}
	_, err := url.Parse(u)
	if err != nil {
		errList["url"] = "Invalid image Url."
	}
}