package ccmasker

import (
	"bytes"
	"errors"
	"io"
	"log"
	"regexp"

	"github.com/mailru/easyjson/jwriter"
	"github.com/theplant/luhn"
)

//easyjson:json
type Message struct {
	Msg []byte
}

// ProcessMessage filters the message through regexp filters and returns appropriate response for rsyslog
// The iterations appear wasteful, but there are edge cases which make iterating for
// all possible PAN lengths necessary.
func ProcessMessage(out io.Writer, message []byte, filters []filterGroup, numFilter *regexp.Regexp) error {
	validated := false

	for _, group := range filters {
		// If variable length pattern matches move on
		if group.variable.Match(message) {
			for _, fixedPattern := range group.fixed {
				// If fixed length pattern matches move on
				if fixedPattern.Match(message) {
					matchStrings := fixedPattern.FindAll(message, -1)
					for _, match := range matchStrings {
						// Prepare string for Luhn check
						cleanMatch := numFilter.ReplaceAll(match, []byte{})
						cleanInt, err := parseIntFromBytes(cleanMatch)
						if err != nil {
							return err
						}
						// Check with Luhn
						if luhn.Valid(cleanInt) {
							validated = true
							message = fixedPattern.ReplaceAllLiteral(message, group.mask)
						}
					}
				}
			}
		}
	}

	// If PAN data isn't found, return empty JSON
	if validated == false {
		_, err := out.Write([]byte{'{', '}', '\n'})
		if err != nil {
			log.Fatalf("could not write to stdout: %s", err)
		}
		return nil
	}

	// If PAN data is found, wrap to JSON and return
	jw := jwriter.Writer{}
	message = bytes.TrimSuffix(message, []byte{'\n'})
	jsonMessage := Message{Msg: message}
	jsonMessage.MarshalEasyJSON(&jw)
	_, err := jw.DumpTo(out)
	if err != nil {
		return err
	}
	_, err = out.Write([]byte{'\n'})
	if err != nil {
		return err
	}

	return nil
}

// parseIntFromBytes is the equivalent of strconv.Atoi but for byteslices
func parseIntFromBytes(b []byte) (int, error) {
	if len(b) == 0 {
		return 0, errors.New("empty byte slice")
	}

	var num int
	var sign = 1

	start := 0
	if b[0] == '-' {
		sign = -1
		start = 1
	}

	for i := start; i < len(b); i++ {
		if b[i] < '0' || b[i] > '9' {
			return 0, errors.New("invalid byte slice")
		}
		num = num*10 + int(b[i]-'0')
	}

	return sign * num, nil
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (in Message) MarshalEasyJSON(out *jwriter.Writer) {
	out.RawByte('{')

	const prefix string = ",\"Msg\":"
	out.RawString(prefix[1:])

	// Manually handle byte slice as a JSON string
	out.RawByte('"')
	for _, b := range in.Msg {
		switch b {
		case '"':
			out.RawString(`\"`)
		case '\\':
			out.RawString(`\\`)
		case '\b':
			out.RawString(`\b`)
		case '\f':
			out.RawString(`\f`)
		case '\n':
			out.RawString(`\n`)
		case '\r':
			out.RawString(`\r`)
		case '\t':
			out.RawString(`\t`)
		default:
			if b < 0x20 {
				out.RawString(`\u00`)
				out.RawByte("0123456789abcdef"[b>>4])
				out.RawByte("0123456789abcdef"[b&0xF])
			} else {
				out.RawByte(b)
			}
		}
	}
	out.RawByte('"')

	out.RawByte('}')
}
