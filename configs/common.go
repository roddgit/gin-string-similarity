package configs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func StartProcess() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func EndProcess(start int64) string {
	end := time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
	diff := (end - start) / 1000
	str_diff := strconv.FormatInt(diff, 10)

	return str_diff
}

func DateTimeNow() string {
	tn := time.Now()

	return tn.Format("2006-01-02 15:04:05")
}

func ValidationString(reqString string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	leadtrailstring := re_leadclose_whtsp.ReplaceAllString(reqString, "")
	final_trim := re_inside_whtsp.ReplaceAllString(leadtrailstring, " ")
	result := strings.ToLower(final_trim)

	return result
}

func TrimFloat(numFloat float64) float64 {
	var s string = strconv.FormatFloat(numFloat, 'f', 4, 64)
	subStr := SubstrStr(s, 1, 2)
	if subStr == ".0" {
		return numFloat
	} else {
		nStr := SubstrStr(s, 0, 4)
		result, _ := strconv.ParseFloat(nStr, 64)
		return result
	}

}

func SubstrStr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func LogsIdGenerator(uniq string) string {
	t := time.Now()
	ts := t.Format("060102150405.000000")
	result := uniq + strings.Replace(ts, ".", "", -1)

	return result
}

func CreateKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func ArrayToString(m ...interface{}) string {
	arr_json, _ := json.Marshal(m)

	return string(arr_json)
}

func StringToInt(s string) int {
	intVar, _ := strconv.Atoi(s)

	return intVar
}

func HandleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
