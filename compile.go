package main

import "regexp"

// FilterGroup was necessary in order to eliminate false positives with the luhn check
// Some PANs are of variable length, which causes issues with the combination of
// matching and checking the luhn number.
//
// Mask is a string denoting the card type and is used to mask the PAN number in the
// message.
// Variable is the variable length compiled Regexp. This is matched first to reduce
// the total number iterations for each message.
// Fixed is an array of fixed length compiled Regexp. Message is iterated upon with
// this filter until a luhn check matches.
type FilterGroup struct {
	Mask     string
	Variable *regexp.Regexp
	Fixed    []*regexp.Regexp
}

// CompileFilters returns a slice of FilterGroup values
// Details on PANs taken from https://en.wikipedia.org/wiki/Payment_card_number
// Since we are forced to iterate over all filters regardless of matching, order does
// not matter.
func CompileFilters() []FilterGroup {
	s := "[ +=_-]"
	filters := []FilterGroup{
		{
			Mask:     "XXXX-VISA-XXXX",
			Variable: regexp.MustCompile("4[0-9]{3}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?([0-9]{4}|[0-9])"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("4[0-9]{3}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]"),
				regexp.MustCompile("4[0-9]{3}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			},
		},
		{
			Mask:     "XXXX-Master5xxx-XXXX",
			Variable: regexp.MustCompile("5[1-5][0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("5[1-5][0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			},
		},
		{
			Mask:     "XXXX-Maestro-XXXX",
			Variable: regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{0,4}" + s + "?[0-9]{0,3}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + ""),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
			},
		},
		{
			Mask:     "XXXX-MaestroUK-XXXX",
			Variable: regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{0,4}" + s + "?[0-9]{0,3}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + ""),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]"),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{2}"),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]"),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2}"),
				regexp.MustCompile("(6767" + s + "?70[0-9]{2}|6767" + s + "?74[0-9]{2})" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
			},
		},
		{
			Mask:     "XXXX-Master2xxx-XXXX",
			Variable: regexp.MustCompile("2[2-7][0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("2[2-7][0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			},
		},
		{
			Mask:     "XXXX-AmEx-XXXX",
			Variable: regexp.MustCompile("(34|37)[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("(34|37)[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
			},
		},
		{
			Mask:     "XXXX-DinersUSC-XXXX",
			Variable: regexp.MustCompile("54[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("54[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
			},
		},
		{
			Mask:     "XXXX-DinersInt-XXXX",
			Variable: regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2,4}" + s + "?[0-9]{0,3}"),
			Fixed: []*regexp.Regexp{
				regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2}"),
				regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{3}"),
				regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{4}"),
				regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2,4}" + s + "?[0-9]"),
				regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2,4}" + s + "?[0-9]{2}"),
				regexp.MustCompile("36[0-9]{2}" + s + "?[0-9]{4}" + s + "?[0-9]{4}" + s + "?[0-9]{2,4}" + s + "?[0-9]{3}"),
			},
		},
	}

	return filters
}
