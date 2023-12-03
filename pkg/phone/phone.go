package phone

import (
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	digitsOnlyRegexp         = regexp.MustCompile(`\D`)
	leadZeroRegexp           = regexp.MustCompile(`^0+`)
	rusLocalePrefixRegexp    = regexp.MustCompile(`^8+`)
	rusLocaleMobPrefixRegexp = regexp.MustCompile(`^89`)
)

// Parse mobile number by country
func Parse(number string, country string) string {
	return parseInternal(number, country, false)
}

// ParseUint mobile number by country and return uint64
func ParseUint(number string, country string) (uint64, error) {
	r := regexp.MustCompile(`[^0-9]`)
	number = r.ReplaceAllString(number, "")
	number = parseInternal(number, country, false)
	i, err := strconv.ParseUint(number, 10, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// ParseWithLandLine is Parse mobile and landline number by country
func ParseWithLandLine(number string, country string) string {
	return parseInternal(number, country, true)
}

func parseInternal(number string, country string, landLineInclude bool) string {
	number = strings.Replace(number, " ", "", -1)
	country = strings.Replace(country, " ", "", -1)
	plusSign := false
	if strings.HasPrefix(number, "+") {
		plusSign = true
	}

	// remove any non-digit character, included the +
	number = digitsOnlyRegexp.ReplaceAllString(number, "")

	iso3166 := getISO3166ByCountry(country)

	if indexOfString(iso3166.Alpha3, []string{"GAB", "CIV", "COG"}) == -1 {
		number = leadZeroRegexp.ReplaceAllString(number, "")
	}

	if iso3166.Alpha3 == "RUS" && len(number) == 11 && rusLocaleMobPrefixRegexp.MatchString(number) {
		number = rusLocalePrefixRegexp.ReplaceAllString(number, "")
	}
	if plusSign {
		iso3166 = GetISO3166ByNumber(number, landLineInclude)
	} else {
		if indexOfInt(len(number), iso3166.PhoneNumberLengths) != -1 {
			number = iso3166.CountryCode + number
		}
	}
	if validatePhoneISO3166(number, iso3166, landLineInclude) {
		return number
	}
	return ""
}

func getISO3166ByCountry(country string) ISO3166 {
	iso3166 := ISO3166{}
	uppperCaseCountry := strings.ToUpper(country)
	switch len(country) {
	case 0:
		iso3166 = GetISO3166()[0]
	case 2:
		for _, i := range GetISO3166() {
			if i.Alpha2 == uppperCaseCountry {
				iso3166 = i
				break
			}
		}
	case 3:
		for _, i := range GetISO3166() {
			if i.Alpha3 == uppperCaseCountry {
				iso3166 = i
				break
			}
		}
	default:
		for _, i := range GetISO3166() {
			if strings.ToUpper(i.CountryName) == uppperCaseCountry {
				iso3166 = i
				break
			}
		}
	}
	return iso3166
}

// GetISO3166ByNumber ...
func GetISO3166ByNumber(number string, withLandLine bool) ISO3166 {
	iso3166 := ISO3166{}
	for _, i := range GetISO3166() {
		r := getRegexpByCountryCode(i.CountryCode)
		for _, l := range i.PhoneNumberLengths {
			if r.MatchString(number) && len(number) == len(i.CountryCode)+l {
				// Check match with mobile codes
				for _, w := range i.MobileBeginWith {
					rm := getRegexpByCountryCode(i.CountryCode + w)
					if rm.MatchString(number) {
						// Match by mobile codes
						return i
					}
				}

				// Match by country code only for landline numbers only
				if withLandLine {
					iso3166 = i
					break
				}
			}
		}
	}
	return iso3166
}

func validatePhoneISO3166(number string, iso3166 ISO3166, withLandLine bool) bool {
	if len(iso3166.PhoneNumberLengths) == 0 {
		return false
	}

	if withLandLine {
		r := getRegexpByCountryCode(iso3166.CountryCode)
		for _, l := range iso3166.PhoneNumberLengths {
			if r.MatchString(number) && len(number) == len(iso3166.CountryCode)+l {
				return true
			}
		}
		return false
	}

	r := getRegexpByCountryCode(iso3166.CountryCode)
	number = r.ReplaceAllString(number, "")
	for _, l := range iso3166.PhoneNumberLengths {
		if l == len(number) {
			for _, w := range iso3166.MobileBeginWith {
				rm := getRegexpByCountryCode(w)
				if rm.MatchString(number) {
					return true
				}
			}
		}
	}
	return false
}

func indexOfString(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func indexOfInt(word int, data []int) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

var rMap = map[string]*regexp.Regexp{}
var rLock = sync.RWMutex{}

func getRegexpByCountryCode(countryCode string) *regexp.Regexp {
	rLock.Lock()
	defer rLock.Unlock()
	regex, exists := rMap[countryCode]
	if exists {
		return regex
	} else {
		rMap[countryCode] = regexp.MustCompile(`^` + countryCode)
	}
	return rMap[countryCode]
}
