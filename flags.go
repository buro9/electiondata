package electiondata

import (
	"flag"
	"sync"
)

var registerFlags sync.Once

var (
	Raw2010 *string
	Raw2015 *string
	Raw2016 *string
)

func RegisterFlags() {
	registerFlags.Do(
		func() {
			Raw2010 = flag.String("2010", "", "path to the 2010 election results data")
			Raw2015 = flag.String("2015", "", "path to the 2015 election results data")
			Raw2016 = flag.String("2016", "", "path to the 2016 referendum results data")
		},
	)
}
