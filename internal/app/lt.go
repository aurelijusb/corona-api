package app

import (
	"regexp"
	"fmt"
)

// FileNameToDate - extracts timestamp from file name
func FileNameToDate(fileName string) string {
	re := regexp.MustCompile(`.+(\d{4})_(\d{2})_(\d{2})_(\d{2})_(\d{2})_(\d{2})_00_00\.html`)
	match := re.FindStringSubmatch(fileName)
	return fmt.Sprintf("%s-%s-%sT%s:%s:%sZ", match[1], match[2], match[3], match[4], match[5], match[6]);
}

// DateToFileName - reverse of NameToDate
func DateToFileName(date string) string {
	re := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2}):(\d{2})Z`)
	match := re.FindStringSubmatch(date)
	return fmt.Sprintf("koronavirusas-%s_%s_%s_%s_%s_%s_00_00.html", match[1], match[2], match[3], match[4], match[5], match[6]);
}