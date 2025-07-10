package data

import "strconv"

func Bool2Ptr(bool2 bool) *bool {
	return &bool2
}

func Bool2Str2Ptr(bool2 bool) *string {
	return String2Ptr(strconv.FormatBool(bool2))
}

func String2Ptr(string2 string) *string {
	return &string2
}
