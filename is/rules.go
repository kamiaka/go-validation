package is

import (
	"net"
	"regexp"
	"strings"
	"unicode"

	"github.com/asaskevich/govalidator"
	validation "github.com/kamiaka/go-validation"
)

// String rules
var (
	// Email validates if a string is a valid email address.
	Email = validation.NewStringRule(govalidator.IsEmail, "%[1]v must be a valid email address")
	// URL validates if a string is a valid URL.
	URL = validation.NewStringRule(govalidator.IsURL, "%[1]v must be a valid URL")
	// RequestURL validates if a string is a valid request URL
	RequestURL = validation.NewStringRule(govalidator.IsRequestURL, "%[1]v must be a valid request URL")
	// RequestURI validates if a string is a valid request URI
	RequestURI = validation.NewStringRule(govalidator.IsRequestURI, "%[1]v must be a valid request URI")
	// Alpha validates if a string contains English letters only. (a-zA-Z)
	Alpha = validation.NewStringRule(govalidator.IsAlpha, "%[1]v must contains English letters only")
	// Digit validates if a string contains digits only. (0-9)
	Digit = validation.NewStringRule(isDigit, "%[1]v must contains digits only")
	// Alphanumeric validates if a string contains English letters and digits only. (a-zA-Z0-9)
	Alphanumeric = validation.NewStringRule(govalidator.IsAlphanumeric, "%[1]v must contains English letters and digits only")
	// UTFLetter validates if a string contains unicode letters only
	UTFLetter = validation.NewStringRule(govalidator.IsUTFLetter, "%[1]v must contain unicode letter characters only")
	// UTFDigit validates if a string contains unicode decimal digits only
	UTFDigit = validation.NewStringRule(govalidator.IsUTFDigit, "%[1]v must contain unicode decimal digits only")
	// UTFLetterNumeric validates if a string contains unicode letters and numbers only
	UTFLetterNumeric = validation.NewStringRule(govalidator.IsUTFLetterNumeric, "%[1]v must contain unicode letters and numbers only")
	// UTFNumeric validates if a string contains unicode number characters (category N) only
	UTFNumeric = validation.NewStringRule(isUTFNumeric, "%[1]v must contain unicode number characters only")
	// LowerCase validates if a string contains lower case unicode letters only
	LowerCase = validation.NewStringRule(govalidator.IsLowerCase, "%[1]v must be in lower case")
	// UpperCase validates if a string contains upper case unicode letters only
	UpperCase = validation.NewStringRule(govalidator.IsUpperCase, "%[1]v must be in upper case")
	// Hexadecimal validates if a string is a valid hexadecimal number
	Hexadecimal = validation.NewStringRule(govalidator.IsHexadecimal, "%[1]v must be a valid hexadecimal number")
	// HexColor validates if a string is a valid hexadecimal color code
	HexColor = validation.NewStringRule(govalidator.IsHexcolor, "%[1]v must be a valid hexadecimal color code")
	// RGBColor validates if a string is a valid RGB color in the form of rgb(R, G, B)
	RGBColor = validation.NewStringRule(govalidator.IsRGBcolor, "%[1]v must be a valid RGB color code")
	// Int validates if a string is a valid integer number
	Int = validation.NewStringRule(govalidator.IsInt, "%[1]v must be an integer number")
	// Float validates if a string is a floating point number
	Float = validation.NewStringRule(govalidator.IsFloat, "%[1]v must be a floating point number")
	// UUIDv3 validates if a string is a valid version 3 UUID
	UUIDv3 = validation.NewStringRule(govalidator.IsUUIDv3, "%[1]v must be a valid UUID v3")
	// UUIDv4 validates if a string is a valid version 4 UUID
	UUIDv4 = validation.NewStringRule(govalidator.IsUUIDv4, "%[1]v must be a valid UUID v4")
	// UUIDv5 validates if a string is a valid version 5 UUID
	UUIDv5 = validation.NewStringRule(govalidator.IsUUIDv5, "%[1]v must be a valid UUID v5")
	// UUID validates if a string is a valid UUID
	UUID = validation.NewStringRule(govalidator.IsUUID, "%[1]v must be a valid UUID")
	// CreditCard validates if a string is a valid credit card number
	CreditCard = validation.NewStringRule(govalidator.IsCreditCard, "%[1]v must be a valid credit card number")
	// ISBN10 validates if a string is an ISBN version 10
	ISBN10 = validation.NewStringRule(govalidator.IsISBN10, "%[1]v must be a valid ISBN-10")
	// ISBN13 validates if a string is an ISBN version 13
	ISBN13 = validation.NewStringRule(govalidator.IsISBN13, "%[1]v must be a valid ISBN-13")
	// ISBN validates if a string is an ISBN (either version 10 or 13)
	ISBN = validation.NewStringRule(isISBN, "%[1]v must be a valid ISBN")
	// JSON validates if a string is in valid JSON format
	JSON = validation.NewStringRule(govalidator.IsJSON, "%[1]v must be in valid JSON format")
	// ASCII validates if a string contains ASCII characters only
	ASCII = validation.NewStringRule(govalidator.IsASCII, "%[1]v must contain ASCII characters only")
	// PrintableASCII validates if a string contains printable ASCII characters only
	PrintableASCII = validation.NewStringRule(govalidator.IsPrintableASCII, "%[1]v must contain printable ASCII characters only")
	// Multibyte validates if a string contains multibyte characters
	Multibyte = validation.NewStringRule(govalidator.IsMultibyte, "%[1]v must contain multibyte characters")
	// FullWidth validates if a string contains full-width characters
	FullWidth = validation.NewStringRule(govalidator.IsFullWidth, "%[1]v must contain full-width characters")
	// HalfWidth validates if a string contains half-width characters
	HalfWidth = validation.NewStringRule(govalidator.IsHalfWidth, "%[1]v must contain half-width characters")
	// VariableWidth validates if a string contains both full-width and half-width characters
	VariableWidth = validation.NewStringRule(govalidator.IsVariableWidth, "%[1]v must contain both full-width and half-width characters")
	// Base64 validates if a string is encoded in Base64
	Base64 = validation.NewStringRule(govalidator.IsBase64, "%[1]v must be encoded in Base64")
	// DataURI validates if a string is a valid base64-encoded data URI
	DataURI = validation.NewStringRule(govalidator.IsDataURI, "%[1]v must be a Base64-encoded data URI")
	// CountryCode2 validates if a string is a valid ISO3166 Alpha 2 country code
	CountryCode2 = validation.NewStringRule(govalidator.IsISO3166Alpha2, "%[1]v must be a valid two-letter country code")
	// CountryCode3 validates if a string is a valid ISO3166 Alpha 3 country code
	CountryCode3 = validation.NewStringRule(govalidator.IsISO3166Alpha3, "%[1]v must be a valid three-letter country code")
	// DialString validates if a string is a valid dial string that can be passed to Dial()
	DialString = validation.NewStringRule(govalidator.IsDialString, "%[1]v must be a valid dial string")
	// MAC validates if a string is a MAC address
	MAC = validation.NewStringRule(govalidator.IsMAC, "%[1]v must be a valid MAC address")
	// IP validates if a string is a valid IP address (either version 4 or 6)
	IP = validation.NewStringRule(govalidator.IsIP, "%[1]v must be a valid IP address")
	// IPv4 validates if a string is a valid version 4 IP address
	IPv4 = validation.NewStringRule(govalidator.IsIPv4, "%[1]v must be a valid IPv4 address")
	// IPv6 validates if a string is a valid version 6 IP address
	IPv6 = validation.NewStringRule(govalidator.IsIPv6, "%[1]v must be a valid IPv6 address")
	// CIDR validates if as string is a valid CIDR
	CIDR = validation.NewStringRule(isCIDR, "%[1]v must be a valid CIDR")
	// IPv4CIDR validates if as string is a valid version 4 IP address's CIDR
	IPv4CIDR = validation.NewStringRule(isIPv4CIDR, "%[1]v must be a valid IPv4 CIDR")
	// IPv6CIDR validates if as string is a valid version 6 IP address's CIDR
	IPv6CIDR = validation.NewStringRule(isIPv6CIDR, "%[1]v must be a valid IPv6 CIDR")
	// DNSName validates if a string is valid DNS name
	DNSName = validation.NewStringRule(govalidator.IsDNSName, "%[1]v must be a valid DNS name")
	// Host validates if a string is a valid IP (both v4 and v6) or a valid DNS name
	Host = validation.NewStringRule(govalidator.IsHost, "%[1]v must be a valid IP address or DNS name")
	// Port validates if a string is a valid port number
	Port = validation.NewStringRule(govalidator.IsPort, "%[1]v must be a valid port number")
	// MongoID validates if a string is a valid Mongo ID
	MongoID = validation.NewStringRule(govalidator.IsMongoID, "%[1]v must be a valid hex-encoded MongoDB ObjectId")
	// Latitude validates if a string is a valid latitude
	Latitude = validation.NewStringRule(govalidator.IsLatitude, "%[1]v must be a valid latitude")
	// Longitude validates if a string is a valid longitude
	Longitude = validation.NewStringRule(govalidator.IsLongitude, "%[1]v must be a valid longitude")
	// SSN validates if a string is a social security number (SSN)
	SSN = validation.NewStringRule(govalidator.IsSSN, "%[1]v must be a valid social security number")
	// Semver validates if a string is a valid semantic version
	Semver = validation.NewStringRule(govalidator.IsSemver, "%[1]v must be a valid semantic version")
)

// patterns
var (
	isDigitRegEx = regexp.MustCompile("^[0-9]$")
)

func isDigit(s string) bool {
	return isDigitRegEx.MatchString(s)
}

func isISBN(value string) bool {
	return govalidator.IsISBN(value, 10) || govalidator.IsISBN(value, 13)
}

func isUTFNumeric(value string) bool {
	for _, c := range value {
		if unicode.IsNumber(c) == false {
			return false
		}
	}
	return true
}

func isCIDR(s string) bool {
	_, _, err := net.ParseCIDR(s)
	return err == nil
}

func isIPv4CIDR(s string) bool {
	if !isCIDR(s) {
		return false
	}
	return strings.ContainsRune(s, '.')
}

func isIPv6CIDR(s string) bool {
	if !isCIDR(s) {
		return false
	}
	return strings.ContainsRune(s, ':')
}
