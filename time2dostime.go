package time2dostime

import (
	"time"
)

type BasicDate struct {
	Year  uint16
	Month time.Month
	Day   uint8
}

func (d BasicDate) MonthByte() uint8 { return uint8(d.Month) }

func (d BasicDate) DosYear() uint8 {
	var y uint16 = d.Year - 1980
	return uint8(y & 0xff)
}

func (d BasicDate) DosDate() uint16 {
	var y uint16 = uint16(d.DosYear()) << 9
	var m uint16 = uint16(d.MonthByte()) << 5
	return y | m | uint16(d.Day)
}

type BasicTime struct {
	Hour   uint8
	Minute uint8
	Second uint8
}

func (t BasicTime) DosTime() uint16 {
	var h uint16 = uint16(t.Hour) << 11
	var m uint16 = uint16(t.Minute) << 5
	var s uint16 = uint16(t.Second) >> 1
	return h | m | s
}

type BasicDateTime struct {
	BasicDate
	BasicTime
}

func (b BasicDateTime) DosTime() uint32 {
	var d uint32 = uint32(b.BasicDate.DosDate()) << 16
	var t uint32 = uint32(b.BasicTime.DosTime())
	return d | t
}

type Time struct{ time.Time }

func (t Time) ToBasicDateTime() BasicDateTime {
	var y uint16 = uint16(t.Time.Year() & 0xffff)
	var oy uint16 = y - 1980
	var by uint8 = uint8(oy & 0xff)

	var bd BasicDate = BasicDate{
		Year:  uint16(by) + 1980,
		Month: t.Time.Month(),
		Day:   uint8(t.Time.Day()),
	}

	var bt BasicTime = BasicTime{
		Hour:   uint8(t.Time.Hour()),
		Minute: uint8(t.Time.Minute()),
		Second: uint8(t.Time.Second()),
	}

	return BasicDateTime{
		bd,
		bt,
	}
}

func TimeToBasic(t time.Time) BasicDateTime { return Time{t}.ToBasicDateTime() }
