package timeago

const (
	Upcoming Option = "upcoming"
	Online   Option = "online"
	JustNow  Option = "justNow"
	NoSuffix Option = "noSuffix"
)

type Option string

func enableOption(opt Option) {
	options = append(options, opt)
}

func enableOptions(opts []Option) {
	for _, opt := range opts {
		enableOption(opt)
	}
}

func optionIsEnabled(opt Option) bool {
	for _, option := range options {
		if option == opt {
			return true
		}
	}

	return false
}
