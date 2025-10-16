package ticker

import (
	"time"
)

type Ticker interface {
	Tick() time.Duration
}
