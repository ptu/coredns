package debug

import clog "github.com/ptu/coredns/plugin/pkg/log"

func init() { clog.Discard() }
