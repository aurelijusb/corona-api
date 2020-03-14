package app

import (
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {
	DescribeTable("should parse filename to time",
		func(fileName, date string) {
			By(fileName)
			Expect(FileNameToDate(fileName)).To(Equal(date))
			Expect(DateToFileName(date)).To(Equal(fileName))
		},
		entry("koronavirusas-2020_03_05_14_45_02_00_00.html", "2020-03-05T14:45:02Z"),
		entry("koronavirusas-2020_03_14_11_45_02_00_00.html", "2020-03-14T11:45:02Z"),
	)
})

func entry(input, expected string) TableEntry {
	return Entry(input, input, expected)
}
