package diff

import (
	"github.com/gnames/gnlib/ent/verifier"
)

type Record struct {
	DataSet          string                  `json:"dataSet"`
	Index            int                     `json:"index"`
	EditDistance     int                     `json:"editDistance,omitempty"`
	ID               string                  `json:"id,omitempty"`
	Name             string                  `json:"name"`
	ParsingQuality   int                     `json:"parsingQuality"`
	Cardinality      int                     `json:"cardinality,omitempty"`
	CanonicalSimple  string                  `json:"canonicalSimple,omitempty"`
	CanonicalFull    string                  `json:"canonicalFull,omitempty"`
	CanonicalStemmed string                  `json:"canonicalStemmed,omitempty"`
	Authors          []string                `json:"authors,omitempty"`
	Year             int                     `json:"year,omitempty"`
	Family           string                  `json:"family,omitempty"`
	MatchType        verifier.MatchTypeValue `json:"matchType,omitempty"`
	Score            float64                 `json:"score,omitempty"`
	ScoreDetails     *verifier.ScoreDetails  `json:"scoreDetails,omitempty"`
}
