package rfc1436

import "strings"

type Datatype string

const (
	Text            Datatype = "0"
	Directory       Datatype = "1"
	CSOPhoneBook    Datatype = "2"
	Error           Datatype = "3"
	BinHex          Datatype = "4"
	DOSBinary       Datatype = "5" // doesn't get postamble
	UUEncoded       Datatype = "6"
	IndexSearch     Datatype = "7"
	TextTelnet      Datatype = "8"
	Binary          Datatype = "9" // doesn't get postamble
	RedundantServer Datatype = "+"
	TextTN3270      Datatype = "T"
	GIF             Datatype = "g"
	Image           Datatype = "I"
	EndSentinel     string   = "\r\n.\r\n"
)

func HasPostamble(d Datatype) bool {
	return !(d == DOSBinary || d == Binary)
}

func DotEscape(d Datatype, s string) string {
	if HasPostamble(d) {
		s = strings.Replace(s, EndSentinel, "\r\n..\r\n", -1)
	}
	return s
}

func Postamble(d Datatype, s string) string {
	if HasPostamble(d) {
		s += EndSentinel
	}
	return s
}
