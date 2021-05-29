package ccmasker

import "regexp"

// filterGroup was necessary in order to eliminate false positives with the luhn check
// Some PANs are of variable length, which causes issues with the combination of
// matching and checking the luhn number.
//
// mask is a string denoting the card type and is used to mask the PAN number in the
// message.
// variable is the variable length compiled Regexp. This is matched first to reduce
// the total number iterations for each message.
// fixed is an array of fixed length compiled Regexp. Message is iterated upon with
// this filter until a luhn check matches.
type filterGroup struct {
	mask     string
	variable *regexp.Regexp
	fixed    []*regexp.Regexp
}

// compileFilters returns a slice of filterGroup values
// Details on PANs taken from https://en.wikipedia.org/wiki/Payment_card_number
// Order of fixed length filters ** MUST ** be from longer to shorter
func compileFilters() []filterGroup {
	s := " +=_-"
	filters := []filterGroup{
		{
			mask:     "XXXX-VISA-XXXX",
			variable: regexp.MustCompile("4[0-9]{3}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?([0-9]{4}|[0-9])"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("4[0-9]{3}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
				regexp.MustCompile("4[0-9]{3}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]"),
			},
		},
		{
			mask:     "XXXX-Master5xxx-XXXX",
			variable: regexp.MustCompile("5[1-5][0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("5[1-5][0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
			},
		},
		{
			mask:     "XXXX-Maestro-XXXX",
			variable: regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2}"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]"),
				regexp.MustCompile("(5018|5020|5038|5893|6304|6759|6761|6762|6763)[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
			},
		},
		{
			mask:     "XXXX-MaestroUK-XXXX",
			variable: regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{0,4}[" + s + "]?[0-9]{0,3}"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2}"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2}"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}[" + s + "]?[0-9]"),
				regexp.MustCompile("(6767[" + s + "]?70[0-9]{2}|6767[" + s + "]?74[0-9]{2})[" + s + "]?[0-9]{4}"),
			},
		},
		{
			mask:     "XXXX-Master2xxx-XXXX",
			variable: regexp.MustCompile("2[2-7][0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("2[2-7][0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
			},
		},
		{
			mask:     "XXXX-AmEx-XXXX",
			variable: regexp.MustCompile("(34|37)[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("(34|37)[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
			},
		},
		{
			mask:     "XXXX-DinersInt-XXXX",
			variable: regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2,4}[" + s + "]?[0-9]{0,3}"),
			fixed: []*regexp.Regexp{
				regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
				regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2}"),
				regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]"),
				regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}"),
				regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{3}"),
				regexp.MustCompile("36[0-9]{2}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{4}[" + s + "]?[0-9]{2}"),
			},
		},
	}

	return filters
}
