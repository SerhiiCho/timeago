package timeago

type opt string

const (
	// OptUpcoming option removes the suffix "ago" when the date is
	// in the future. This option is enabled by default, there
	// is no need to pass it. It's available to keep backward
	// compatibility with the previous versions.
	OptUpcoming opt = "upcoming"

	// Upcoming option displays "OptOnline" if date interval withing
	// 60 seconds. For example instead of "13 seconds ago"
	// prints "OptOnline".
	OptOnline opt = "online"

	// OptJustNow option displays "Just now" if date interval withing
	// 60 seconds. For example instead of "32 seconds ago" prints
	// "Just now".
	OptJustNow opt = "justNow"

	// OptNoSuffix option removes suffix from datetime result and get
	// for example "5 minutes" instead of "5 minutes ago".
	OptNoSuffix opt = "noSuffix"
)

func enableOption(o opt) {
	options = append(options, o)
}

func enableOptions(opts []opt) {
	for _, opt := range opts {
		enableOption(opt)
	}
}

func optionIsEnabled(o opt) bool {
	for _, option := range options {
		if option == o {
			return true
		}
	}

	return false
}
