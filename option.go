package timeago

type opt string

const (
	// OptUpcoming option removes the suffix "ago" when the date is
	// in the future. This option is enabled by default, there
	// is no need to pass it. It's available to keep backward
	// compatibility with the previous versions.
	OptUpcoming opt = "upcoming"

	// Upcoming option displays "OptOnline" if date interval withing
	// the OnlineThreshold provided in config. By default, it's
	// 60 seconds. For example instead of "13 seconds ago" it will
	// print "Online".
	OptOnline opt = "online"

	// OptJustNow option displays "Just now" if date interval withing
	// the JustNowThreshold provided in config. By default, it's
	// 60 seconds. For example instead of "32 seconds ago" it will
	// print "Just now".
	OptJustNow opt = "justNow"

	// OptNoSuffix option removes suffix from datetime result and get
	// for example "5 minutes" instead of "5 minutes ago".
	OptNoSuffix opt = "noSuffix"
)

func enableOption(opt opt) {
	options = append(options, opt)
}

func enableOptions(opts []opt) {
	for i := range opts {
		enableOption(opts[i])
	}
}

func optionIsEnabled(opt opt) bool {
	for i := range options {
		if options[i] == opt {
			return true
		}
	}

	return false
}
