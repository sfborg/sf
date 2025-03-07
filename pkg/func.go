package sf

import (
	"github.com/gnames/gnparser"
	"github.com/gnames/gnparser/ent/nomcode"
)

// ParserPool initializes and returns a map of GNparser pools.
// It creates two channels, one for botanical parsers and one for general
// parsers, each with a buffer size specified by jobsNum. The parsers are
// configured with options based on the provided details flag.
//
// Parameters:
//   - jobsNum: The number of parser instances to create for each
//     channel.
//   - details: A boolean flag to include detailed parsing options.
//
// Returns: A map with two channels:
//   - "botanical": A channel for GNparser instances configured
//     for botanical parsing.
//   - "any": A channel for GNparser instances configured for general parsing.
func ParserPool(jobsNum int, details bool) map[string]chan gnparser.GNparser {
	botChan := make(chan gnparser.GNparser, jobsNum)
	anyChan := make(chan gnparser.GNparser, jobsNum)
	opts := []gnparser.Option{gnparser.OptWithDetails(details)}
	botOpts := append(opts, gnparser.OptCode(nomcode.Cultivar))
	for range jobsNum {
		botChan <- gnparser.New(gnparser.NewConfig(botOpts...))
		anyChan <- gnparser.New(gnparser.NewConfig(opts...))
	}
	return map[string]chan gnparser.GNparser{
		"botanical": botChan,
		"any":       anyChan,
	}
}
