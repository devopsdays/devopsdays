package helpers

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

// ValidateField performs field validation for common types
func ValidateField(input, field string) bool {
	switch field {
	case "city":
		if strings.Count(input, "") > 100 {
			return false
		}
		return true
	case "year":
		if strings.Count(input, "") != 5 {
			return false
		} else if s, err := strconv.ParseInt(input, 10, 32); err == nil {
			if s < 2016 || s > 2030 {
				return false
			}
			return true
		}
	case "twitter":
		if strings.ContainsAny(input, " ") {
			return false
		}
		return true
	case "website":
		if govalidator.IsRequestURL(input) {
			return true
		}
		return false

	case "facebook":
		if govalidator.IsRequestURL(input) {
			return true
		}
		return false

	case "linkedin":
		if govalidator.IsRequestURL(input) {
			return true
		}
		return false

	case "github":
		if (govalidator.IsRequestURL(input)) || (strings.ContainsAny(input, " ")) {
			return false
		}
		return true
	case "gitlab":
		if (govalidator.IsRequestURL(input)) || (strings.ContainsAny(input, " ")) {
			return false
		}
		return true
	case "filepath":
		ret, _ := govalidator.IsFilePath(input)
		return ret
	case "date":
		ret := govalidator.IsTime(input, "2006-01-02")
		return ret
	case "googleanalytics":
		ret, _ := regexp.MatchString(`UA-[0-9]{8}-[0-9]`, input)
		return ret
	case "coordinates":
		ret, _ := regexp.MatchString(`-?[0-9]+\.[0-9]+,\W?-?[0-9]+.[0-9]+`, input)
		return ret
	}
	return true // TODO: Make this return an error if no field was matched
}
