package electiondata

type Results struct {
	Elections   []Election   `json:"elections"`
	Referendums []Referendum `json:"referendums"`
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

type Referendum struct {
	Year                int
	Question            string
	ConstituencyAnswers []ConstituencyAnswer
}

type ConstituencyAnswer struct {
	ID      string
	Name    string
	Answers []Answer
}

type Answer struct {
	Answer string
	Votes  int
}

func (e *Election) Totals() (
	constituencies int,
	candidates int,
	votes int,
) {
	constituencies = len(e.Constituencies)

	for _, c := range e.Constituencies {
		for _, r := range c.Results {
			candidates++
			votes += r.Votes
		}
	}

	return
}

func (ref *Referendum) Totals() (
	constituencies int,
	answers int,
	votes int,
) {
	constituencies = len(ref.ConstituencyAnswers)

	for _, c := range ref.ConstituencyAnswers {
		if answers < len(c.Answers) {
			answers = len(c.Answers)
		}
		for _, r := range c.Answers {
			votes += r.Votes
		}
	}

	return
}
