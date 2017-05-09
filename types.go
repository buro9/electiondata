package electiondata

type Results struct {
	Elections []Election `json:"elections"`
}

type Election struct {
	Year           int            `json:"year"`
	Constituencies []Constituency `json:"constituencies"`
}

type Constituency struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Results []Result `json:"results"`
}

type Result struct {
	Candidate string `json:"candidate"`
	Party     string `json:"party"`
	Votes     int    `json:"votes"`
	Incumbent bool   `json:"incumbent"`
	Elected   bool   `json:"elected"`
}

func (e *Election) Totals() (
	constituencies int,
	candidates int,
	votes int,
) {
	constituencies = len(e.Constituencies)

	for _, c := range e.Constituencies {
		for _, r := range c.Results {
			candidates += 1
			votes += r.Votes
		}
	}

	return
}
