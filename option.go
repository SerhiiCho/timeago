package timeago

type Option string

const (
	// Upcoming option removes the suffix "ago" when the date is
	// in the future. This option is enabled by default, there
	// is no need to pass it. It's available to keep backward
	// compatibility with the previous versions.
	Upcoming Option = "upcoming"

	// Upcoming option displays "Online" if date interval withing
	// 60 seconds. For example instead of "13 seconds ago"
	// prints "Online".
	Online Option = "online"

	// JustNow option displays "Just now" if date interval withing
	// 60 seconds. For example instead of "32 seconds ago" prints
	// "Just now".
	JustNow Option = "justNow"

	// NoSuffix option removes suffix from datetime result and get
	// for example "5 minutes" instead of "5 minutes ago".
	NoSuffix Option = "noSuffix"
)

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
