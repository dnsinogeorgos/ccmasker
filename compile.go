package main

import "regexp"

type FilterGroup struct {
	Mask     string
	Variable *regexp.Regexp
	Fixed    []*regexp.Regexp
}

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
