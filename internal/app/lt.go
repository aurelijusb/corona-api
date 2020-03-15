package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var parsingCache map[string]CoronaReport = map[string]CoronaReport{}

type (
	// CoronaReport - structued representation of official website data
	CoronaReport struct {
		CrawlDate       string    `json:"crawlDate"`
		ConfirmedCases  int       `json:"confirmedCases"`
		LastDayChecked  int       `json:"lastDayChecked"`
		TotalTests      int       `json:"totalTests"`
		KnownCandidates int       `json:"knownCandidates"`
		Time            time.Time `json:"time"`
		LocalTime       string    `json:"localTime"`
		LinkToRaw       string    `json:"linkToRaw"`
	}
)

const currentYear = "2020"

// FileNameToDate - extracts timestamp from file name
func FileNameToDate(fileName string) string {
	re := regexp.MustCompile(`.+(\d{4})_(\d{2})_(\d{2})_(\d{2})_(\d{2})_(\d{2})_00_00\.html`)
	match := re.FindStringSubmatch(fileName)
	return fmt.Sprintf("%s-%s-%sT%s:%s:%sZ", match[1], match[2], match[3], match[4], match[5], match[6])
}

// DateToFileName - reverse of NameToDate
func DateToFileName(date string) string {
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})Z`)
	match := re.FindStringSubmatch(date)
	return fmt.Sprintf("koronavirusas-%s_%s_%s_%s_%s_%s_00_00.html", match[1], match[2], match[3], match[4], match[5], match[6])
}

// ExtractData - parse HTML into structured data
func ExtractData(raw string, fileName string) CoronaReport {
	if cached, exits := parsingCache[fileName]; exits {
		return cached
	}

	crawlDate := FileNameToDate(fileName)
	result := CoronaReport{
		CrawlDate: crawlDate,
		LinkToRaw: fmt.Sprintf("/api/v1/raw/%s", fileName),
	}
	for _, line := range strings.Split(raw, "\n") {
		if value, ok := extractNumber(line, `.*<li>.*Lietuvoje patvirtintų ligos atvejų:\D+(\d+)[^(]+\([^)]+.+`, 1); ok {
			result.ConfirmedCases = value
		}
		if value, ok := extractNumber(line, `.*<li>Per vakar dieną ištirta: (\d+).*</li>.*`, 1); ok {
			result.LastDayChecked = value
		}
		if value, ok := extractNumber(line, `.*<li>Iki šiol iš viso .+ dėl įtariamo koronaviruso: (\d+).*</li>.*`, 1); ok {
			result.TotalTests = value
		}
		if value, ok := extractNumber(line, `.*<li>NVSC turimi duomenys apie keliavusius.*: (\d+).*<.*`, 1); ok {
			result.KnownCandidates = value
		}
		if match, ok := extractAll(line, `.*<strong>(Sausio|Vasario|Kovo|Balandžio|Gegužės|Birželio|Liepos|Rugpjūčio|Rugsėjo|Spalio|Lapkričio|Gruodžio) (\d+) d\. (\d+)\.(\d+) val\. duomenimis:.*`); ok {
			localTime := fmt.Sprintf("%s-%02d-%02sT%02s:%02s:00+02:00", currentYear, month(match[1]), match[2], match[3], match[4]) // Assuming Vilnius time zone and year 2020
			result.LocalTime = localTime
			parsedTime, err := time.Parse(time.RFC3339, localTime)
			if err != nil {
				fmt.Printf("Error: Could not translate date: %s: %s\n", match[1], err.Error())
				continue
			}
			result.Time = parsedTime.UTC()
		}
	}

	parsingCache[fileName] = result

	return result
}

func extractNumber(line string, pattern string, groupNr int) (int, bool) {
	re := regexp.MustCompile(pattern)
	if re.Match([]byte(line)) {
		match := re.FindStringSubmatch(line)
		numericValue, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("Error: Could not convert confirmed cases to integer: %#v from %s", match[1], strings.Trim(line, "\r\n\t "))
			return -1, false
		}
		return numericValue, true
	}
	return -1, false
}

func extractAll(line string, pattern string) ([]string, bool) {
	re := regexp.MustCompile(pattern)
	if re.Match([]byte(line)) {
		match := re.FindStringSubmatch(line)
		return match, true
	}
	return []string{}, false
}

func month(inLithuania string) int {
	months := []string{
		"Sausio",
		"Vasario",
		"Kovo",
		"Balandžio",
		"Gegužės",
		"Birželio",
		"Liepos",
		"Rugpjūčio",
		"Rugsėjo",
		"Spalio",
		"Lapkričio",
		"Gruodžio",
	}
	for nr, name := range months {
		if strings.ToLower(inLithuania) == strings.ToLower(name) {
			return nr + 1
		}
	}
	return 0
}
