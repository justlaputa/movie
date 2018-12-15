package pt

import (
	"bytes"
	"encoding/json"
)

type DigitalFormat uint8
type DigitalResolution uint8
type DigitalFileSize uint64

const (
	UnknownDigitalFormat DigitalFormat = iota
	Blueray
	HDTV
	WebDL
	UHDTV
	Blueray3D
)

const (
	//UnknownResolution unknow resolution
	UnknownResolution DigitalResolution = iota
	//FHD 1080p video
	FHD
	//HD 720p video
	HD
	//UHD4K 4K video
	UHD4K
)

var digitalFormatToString = map[DigitalFormat]string{
	Blueray:              "blueray",
	HDTV:                 "hdtv",
	WebDL:                "webdl",
	UHDTV:                "uhdtv",
	Blueray3D:            "3D",
	UnknownDigitalFormat: "unknown",
}

var stringToDigitalFormat = map[string]DigitalFormat{
	"blueray": Blueray,
	"hdtv":    HDTV,
	"webdl":   WebDL,
	"uhdtv":   UHDTV,
	"3D":      Blueray3D,
	"unknown": UnknownDigitalFormat,
}

func (f DigitalFormat) toString() string {
	return digitalFormatToString[f]
}

// MarshalJSON marshals the enum as a quoted json string
func (f *DigitalFormat) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(digitalFormatToString[*f])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (f *DigitalFormat) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*f = stringToDigitalFormat[j]
	return nil
}

var digitalResolutionToString = map[DigitalResolution]string{
	FHD:               "1080p",
	HD:                "720p",
	UHD4K:             "4K",
	UnknownResolution: "unknown",
}

var stringToDigitalResolution = map[string]DigitalResolution{
	"1080p":   FHD,
	"720p":    HD,
	"4K":      UHD4K,
	"unknown": UnknownResolution,
}

func (f DigitalResolution) toString() string {
	return digitalResolutionToString[f]
}

// MarshalJSON marshals the enum as a quoted json string
func (f *DigitalResolution) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(digitalResolutionToString[*f])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (f *DigitalResolution) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*f = stringToDigitalResolution[j]
	return nil
}
