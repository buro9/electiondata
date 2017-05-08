package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	CON  = "CON"
	DUP  = "DUP"
	GRN  = "GRN"
	IND  = "IND"
	LAB  = "LAB"
	LIB  = "LIB"
	PC   = "PC"
	SDLP = "SDLP"
	SF   = "SF"
	SNP  = "SNP"
	SPK  = "SPK"
)

type Results struct {
	Constituencies []Constituency
}

type Constituency struct {
	Name    string
	Results []Result
}

type Result struct {
	Candidate string
	Party     string
	Votes     int
	Incumbent bool
	Elected   bool
}

var (
	raw2010 = flag.String("2010", "", "path to the 2010 constituency results data")
	raw2015 = flag.String("2015", "", "path to the 2015 constituency results data")
)

func main() {
	flag.Parse()

	if (raw2010 == nil || *raw2010 == "") && (raw2015 == nil || *raw2015 == "") {
		fmt.Print("One of 2010 or 2015 must be supplied:\n\n")
		flag.Usage()
		os.Exit(1)
	}

	var results Results

	if *raw2010 != "" {
		err := results.parse2010()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	if *raw2015 != "" {
		err := results.parse2015()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
	}

	fmt.Printf("results for %d constituencies loaded\n", len(results.Constituencies))
}

func (results *Results) parse2010() error {
	b, err := readFile(*raw2010)
	if err != nil {
		return err
	}

	r := csv.NewReader(bytes.NewReader(b))
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	const (
		ORDER = iota
		NO
		SEAT
		CANDIDATE
		PARTY
		VOTE
		PERCENT
		INCUMBENT
		ELECTED
		WINNER
		MAJORITY
		TURNOUT
		CHANGE
		LAB_CON_SWING
		CON_LIB_SWING
		LAB_LIB_SWING
		HOLD_WIN
	)

	var cur Constituency
	for ii, record := range records {
		if ii == 0 {
			// header row
			continue
		}

		if ii == 1 {
			cur.Name = record[SEAT]
		} else if cur.Name != record[SEAT] {
			// add the last one, and work on the new
			results.Constituencies = append(results.Constituencies, cur)
			cur = Constituency{}
			cur.Name = record[SEAT]
		}

		var res Result
		res.Candidate = record[CANDIDATE]
		res.Elected = (record[ELECTED] == "*")
		res.Incumbent = (record[INCUMBENT] == "*")
		switch record[PARTY] {
		case "C":
			res.Party = CON
		case "DUP":
			res.Party = DUP
		case "Green":
			res.Party = GRN
		case "Ind":
			res.Party = IND
		case "Lab":
			res.Party = LAB
		case "LD":
			res.Party = LIB
		case "PC":
			res.Party = PC
		case "SDLP":
			res.Party = SDLP
		case "SF":
			res.Party = SF
		case "SNP":
			res.Party = SNP
		case "Speaker":
			res.Party = SPK
		default:
			res.Party = record[PARTY]
		}
		votes, err := strconv.Atoi(strings.Replace(record[VOTE], ",", "", -1))
		if err != nil {
			return err
		}
		res.Votes = votes
		cur.Results = append(cur.Results, res)
	}
	results.Constituencies = append(results.Constituencies, cur)

	for _, c := range results.Constituencies {
		for _, r := range c.Results {
			if r.Elected {
				// winning parties sanitisation test
				switch r.Party {
				case "Alliance":
				case CON:
				case DUP:
				case GRN:
				case IND:
				case LAB:
				case "Lab Co-op":
				case LIB:
				case PC:
				case SDLP:
				case SF:
				case SNP:
				case SPK:
				default:
					return fmt.Errorf("%s not recognised for %s\n", r.Party, c.Name)
				}
			}
		}
	}

	return nil
}

func (results *Results) parse2015() error {
	return nil
}

func readFile(path string) ([]byte, error) {
	fmt.Printf("reading file at %s\n", path)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}
