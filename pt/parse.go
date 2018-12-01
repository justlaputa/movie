package pt

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const (
	HDCSite   = "HDChina"
	PutaoSite = "Putao"
)

//DigitalResolutionString string name of each resolution type
var DigitalResolutionString = []string{
	"1080p",
	"720p",
	"4K",
	"Unknown",
}

func (i DigitalResolution) String() string {
	if i >= DigitalResolution(len(DigitalResolutionString)-1) {
		return "DigitalResolution(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return DigitalResolutionString[i]
}

var DigitalFormatMap = map[DigitalFormat][]string{
	Blueray:   []string{"bluray", "blu-ray", "blueray", "bd"},
	HDTV:      []string{"hdtv"},
	WebDL:     []string{"webdl", "web-dl", "webrip", "web"},
	UHDTV:     []string{"uhdtv"},
	Blueray3D: []string{"3d", "sbs"},
}

var DigitalResolutionMap = map[DigitalResolution][]string{
	FHD:   []string{"1080", "1080p", "1080i"},
	HD:    []string{"720", "720p"},
	UHD4K: []string{"4k", "2160p"},
}

type MovieInfo struct {
	Title      string
	Year       int
	Group      string
	Source     DigitalFormat
	Resolution DigitalResolution
	Size       DigitalFileSize
	ID         string
	Site       string
}

//ParseHDCTitle parse a HDC movie title string into structured movie information
// example:
// movie := ParseHDCTitle("Man.in.Black.1997.UHDTV.4K.HEVC-HDCTV[7.33 GB]")
// movie will have following value:
// {
// 	Title:      "Man in Black",
// 	Year:       1997,
// 	Group:      "HDCTV",
// 	Source:     UHDTV,
// 	Resolution: UHD4K,
// 	Size:       7330000000, //in Bytes
// }
func ParseHDCTitle(title string) MovieInfo {
	if title == "" {
		return MovieInfo{}
	}

	var size DigitalFileSize
	title, sizeString := removeEndBracket(title)
	if sizeString != "" {
		size = parseSize(sizeString)
	}

	fields := split(title)

	year, yearIndex := findYear(fields)
	source, sourceIndex := findSource(fields)
	resolution, resIndex := findResolution(fields)
	group := findGroup(fields)

	minIndex := minPositive(yearIndex, sourceIndex, resIndex)

	movieTitle := strings.Join(fields[:minIndex], " ")

	return MovieInfo{
		movieTitle, year, group, source, resolution, size, "", HDCSite,
	}
}

//ParsePutaoTitle parse putao movie item title
func ParsePutaoTitle(title string) MovieInfo {
	info := MovieInfo{Site: PutaoSite}
	if title == "" {
		return info
	}

	title = removeBeginBracket(title)

	title, sizeString := removeEndBracket(title)
	if sizeString != "" {
		info.Size = parseSize(sizeString)
	}

	fields := split(title)

	year, yearIndex := findYear(fields)
	info.Year = year

	source, sourceIndex := findSource(fields)
	info.Source = source

	resolution, resIndex := findResolution(fields)
	info.Resolution = resolution

	info.Group = findGroup(fields)

	minIndex := minPositive(yearIndex, sourceIndex, resIndex)

	info.Title = strings.Join(fields[:minIndex], " ")

	return info
}

func removeBeginBracket(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 0 && s[0] == '[' {
		if i := strings.Index(s, "]"); i > 0 {
			return s[i+1:]
		}
	}
	return s
}

func removeEndBracket(s string) (string, string) {
	s = strings.TrimSpace(s)
	l := len(s)
	if s[l-1:] == "]" {
		if i := strings.LastIndex(s, "["); i > 0 {
			return s[0:i], s[i+1 : l-1]
		}
	}
	return s, ""
}

func parseSize(s string) DigitalFileSize {
	s = strings.TrimSpace(s)
	fields := strings.Split(s, " ")
	if len(fields) < 2 {
		return 0
	}

	var size DigitalFileSize

	ssize, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		log.Printf("failed to parse digital size from %s: %v", fields[0], err)
		return 0
	}

	switch strings.ToLower(fields[1]) {
	case "gb":
		size = DigitalFileSize(ssize * 1e9)
	case "mb":
		size = DigitalFileSize(ssize * 1e6)
	}

	return size
}

func split(title string) []string {
	return strings.FieldsFunc(title, func(r rune) bool {
		return (r == '.' || r == ' ')
	})
}

func findYear(fields []string) (int, int) {
	for i := len(fields) - 1; i >= 0; i-- {
		f := fields[i]
		if year, err := tryParseYear(f); err == nil {
			return year, i
		}
	}
	return -1, -1
}

func tryParseYear(yyyy string) (int, error) {
	if len(yyyy) != 4 {
		return -1, fmt.Errorf("%s is not 4 digit year", yyyy)
	}

	if yyyy[0] != '1' && yyyy[0] != '2' {
		return -1, fmt.Errorf("%s is not 1xxx or 2xxx, not supported year range", yyyy)
	}

	return strconv.Atoi(yyyy)
}

func findSource(fields []string) (DigitalFormat, int) {
	for i, field := range fields {
		for format, names := range DigitalFormatMap {
			if contains(names, strings.ToLower(field)) {
				return format, i
			}
		}
	}
	return UnknownDigitalFormat, -1
}

func findResolution(fields []string) (DigitalResolution, int) {
	for i, field := range fields {
		for format, names := range DigitalResolutionMap {
			if contains(names, strings.ToLower(field)) {
				return format, i
			}
		}
	}
	return UnknownResolution, -1
}

func findGroup(fields []string) string {
	if len(fields) <= 0 {
		return ""
	}

	last := fields[len(fields)-1]
	if last == "iso" {
		last = fields[len(fields)-2]
	}

	if i := strings.LastIndex(last, "-"); i >= 0 {
		group := last[i+1:]
		if ii := strings.LastIndex(group, "@"); ii >= 0 {
			group = group[ii+1:]
		}
		return group
	}

	return ""
}

func contains(array []string, s string) bool {
	for _, e := range array {
		if e == s {
			return true
		}
	}
	return false
}

func minPositive(ints ...int) int {
	m := math.MaxInt32
	for _, i := range ints {
		if i >= 0 && i < m {
			m = i
		}
	}
	return m
}
