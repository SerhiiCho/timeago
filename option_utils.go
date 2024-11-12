package timeago

import "github.com/SerhiiCho/timeago/v3/option"

func enableOption(opt option.Option) {
	options = append(options, opt)
}

func enableOptions(opts []option.Option) {
	for _, opt := range opts {
		enableOption(opt)
	}
}

func optionIsEnabled(opt option.Option) bool {
	for _, option := range options {
		if option == opt {
			return true
		}
	}

	return false
}
