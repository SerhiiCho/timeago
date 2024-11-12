package ctx

import (
	"github.com/SerhiiCho/timeago/v3/config"
	"github.com/SerhiiCho/timeago/v3/option"
)

type Ctx struct {
	conf    *config.Config
	options []option.Option
}

func New(conf *config.Config, options []option.Option) *Ctx {
	return &Ctx{
		conf:    conf,
		options: options,
	}
}

func (c Ctx) Config() *config.Config {
	return c.conf
}

func (c Ctx) Options() []option.Option {
	return c.options
}
