package app

import (
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {
	DescribeTable("should parse filename to time",
		func(fileName, date string) {
			Expect(FileNameToDate(fileName)).To(Equal(date))
			Expect(DateToFileName(date)).To(Equal(fileName))
		},
		entry("koronavirusas-2020_03_05_14_45_02_00_00.html", "2020-03-05T14:45:02Z"),
		entry("koronavirusas-2020_03_14_11_45_02_00_00.html", "2020-03-14T11:45:02Z"),
	)
	DescribeTable("should extract ConfirmedCases",
		func(raw string, value string) {
			structured := ExtractData(raw, "koronavirusas-2000_01_01_01_01_01_00_00.html", false)
			expected, err := strconv.Atoi(value)
			Expect(err).To(Not(HaveOccurred()))
			Expect(structured.ConfirmedCases).To(Equal(expected))
		},
		entry("<li>Lietuvoje patvirtintų ligos atvejų: <strong>1 (patvirtinta 2020-02-28 4.00 val.)</strong></li>", "1"),
		entry("<li>Lietuvoje patvirtintų ligos atvejų: 7 (1 patvirtintas 2020-02-28 4.00 val., 2 patvirtinti 2020-03-10 22 val., 3 patvirtinti 2020-03-13 13 val., 1 patvirtintas 2020-03-14 01.30 val.)</li>", "7"),
		entry("<li><strong>Lietuvoje patvirtintų ligos atvejų:</strong> <strong>9 </strong>(1 patvirtintas 2020-02-28 4.00 val., 2 patvirtinti 2020-03-10 22 val., 3 patvirtinti 2020-03-13 13 val., 1 patvirtintas 2020-03-14 01.30 val., 1 patvirtintas 2020-03-14</li>", "9"),
	)
	DescribeTable("should extract LastDayChecked",
		func(raw string, value string) {
			structured := ExtractData(raw, "koronavirusas-2000_01_01_01_01_01_00_00.html", false)
			expected, err := strconv.Atoi(value)
			Expect(err).To(Not(HaveOccurred()))
			Expect(structured.LastDayChecked).To(Equal(expected))
		},
		entry("<li>Per vakar dieną ištirta: 76 (iš jų 44 iki vidurnakčio, 32 po vidurnakčio)</li>", "76"),
		entry("<li>Per vakar dieną ištirta: 15 </li>", "15"),
	)
	DescribeTable("should extract TotalTests",
		func(raw string, value string) {
			structured := ExtractData(raw, "koronavirusas-2000_01_01_01_01_01_00_00.html", false)
			expected, err := strconv.Atoi(value)
			Expect(err).To(Not(HaveOccurred()))
			Expect(structured.TotalTests).To(Equal(expected))
		},
		entry("<li>Iki šiol iš viso atlikta tyrimų dėl įtariamo koronaviruso: 231</li>", "231"),
		entry("<li>Iki šiol iš viso ištirta ėminių dėl įtariamo koronaviruso: 442</li>", "442"),
	)
	DescribeTable("should extract KnownCandidates",
		func(raw string, value string) {
			structured := ExtractData(raw, "koronavirusas-2000_01_01_01_01_01_00_00.html", false)
			expected, err := strconv.Atoi(value)
			Expect(err).To(Not(HaveOccurred()))
			Expect(structured.KnownCandidates).To(Equal(expected))
		},
		entry("<li>NVSC turimi duomenys apie keliavusius į teritorijas, kur vyksta COVID-19 plitimas visuomenėje (stebima jų sveikatos būklė): 4981 asmuo (2020-03-04 duomenimis)<br />", "4981"),
		entry("<li>NVSC turimi duomenys apie keliavusius į teritorijas, kur vyksta COVID-19 plitimas visuomenėje (stebima jų sveikatos būklė): 5175 asmenys (2020-03-06 duomenys)</li>", "5175"),
		entry("<li>NVSC turimi duomenys apie keliavusius į teritorijas, kur vyksta COVID-19 plitimas visuomenėje (stebima jų sveikatos būklė): 9176 (2020-03-13 duomenimis)</li>", "9176"),
	)
	DescribeTable("should extract LocalTime",
		func(raw string, expected string) {
			structured := ExtractData(raw, "koronavirusas-2000_01_01_01_01_01_00_00.html", false)
			Expect(structured.LocalTime).To(Equal(expected))
		},
		entry("<p><p><u><strong>Kovo 4 d. 21.50 val. duomenimis:</strong></u></p>", "2020-03-04T21:50:00+02:00"),
		entry("<p><p><u><strong>Kovo 8 d. 21.30 val. duomenimis:</strong></u></p>", "2020-03-08T21:30:00+02:00"),
		entry("<p><p><u><strong>Kovo 13 d. 8.00 val. duomenimis:</strong></u></p>", "2020-03-13T08:00:00+02:00"),
		entry("<strong>Kovo 14 d. 9.00 val. duomenimis: </strong>", "2020-03-14T09:00:00+02:00"),
		entry("<strong>Gruodžio 31 d. 23.59 val. duomenimis: </strong>", "2020-12-31T23:59:00+02:00"),
	)
})

func entry(input, expected string) TableEntry {
	return Entry(input, input, expected)
}
