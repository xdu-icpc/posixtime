// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs clock_c2go_linux.go

package posixtime

const (
	CLOCK_MONOTONIC_RAW ClockID = 0x4

	CLOCK_REALTIME_COARSE ClockID = 0x5

	CLOCK_MONOTONIC_COARSE ClockID = 0x6

	CLOCK_BOOTTIME ClockID = 0x7

	CLOCK_REALTIME_ALARM ClockID = 0x8

	CLOCK_BOOTTIME_ALARM ClockID = 0x9

	CLOCK_TAI ClockID = 0xb
)

var platformClocks = [...]ClockID{CLOCK_MONOTONIC_RAW,
	CLOCK_REALTIME_COARSE, CLOCK_BOOTTIME, CLOCK_REALTIME_ALARM,
	CLOCK_BOOTTIME_ALARM, CLOCK_TAI}
