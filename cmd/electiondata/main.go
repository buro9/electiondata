package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/buro9/electiondata"
)

func main() {
	electiondata.RegisterFlags()
	csvFile := flag.Bool("csv", false, `output in .csv`)
	flag.Parse()

	if (electiondata.Raw2010 == nil || *electiondata.Raw2010 == "") &&
		(electiondata.Raw2015 == nil || *electiondata.Raw2015 == "") &&
		(electiondata.Raw2016 == nil || *electiondata.Raw2016 == "") {
		fmt.Print("One of 2010 or 2015 or 2016 must be supplied:\n\n")
		flag.Usage()
		os.Exit(1)
	}

	var results electiondata.Results

	if *electiondata.Raw2010 != "" {
		election, err := electiondata.Parse2010(*electiondata.Raw2010)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		cons, cands, votes := election.Totals()
		fmt.Printf(
			"2010: %d votes across %d candidates in %d constituencies\n",
			votes,
			cands,
			cons,
		)
		results.Elections = append(results.Elections, election)
	}

	if *electiondata.Raw2015 != "" {
		election, err := electiondata.Parse2015(*electiondata.Raw2015)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		cons, cands, votes := election.Totals()
		fmt.Printf(
			"2015: %d votes across %d candidates in %d constituencies\n",
			votes,
			cands,
			cons,
		)
		results.Elections = append(results.Elections, election)
	}

	if *electiondata.Raw2016 != "" {
		referendum, err := electiondata.Parse2016(*electiondata.Raw2016)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		cons, answers, votes := referendum.Totals()
		fmt.Printf(
			"2016: %d votes across %d answers in %d constituencies\n",
			votes,
			answers,
			cons,
		)
		results.Referendums = append(results.Referendums, referendum)
	}

	var (
		b        []byte
		filename string
		err      error
	)
	if csvFile != nil && *csvFile {
		// CSV
		records, err := results.ToCSV()
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		filename = `out/elections.csv`
		file, err := os.Create(filename)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer file.Close()

		w := csv.NewWriter(file)
		for _, record := range records {
			err := w.Write(record)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		}
		w.Flush()

	} else {
		//JSON
		b, err = json.MarshalIndent(results, "", "\t")
		//b, err := json.Marshal(results)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		filename = `out/elections.js`
		file, err := os.Create(filename)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		defer file.Close()

		_, err = file.Write(b)
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}

		// we also want to output a file for every constituency result, named
		// after the ONS_ID, i.e. W07000049.json for Aberavon
		cr := make(map[string]electiondata.ByConstituency)
		for _, e := range results.Elections {
			for _, c := range e.Constituencies {
				if _, ok := cr[c.ID]; !ok {
					cr[c.ID] = electiondata.ByConstituency{
						ID:   c.ID,
						Name: c.Name,
					}
				}
				r := cr[c.ID]
				switch e.Year {
				case 2010:
					r.Y2010 = c.Results
				case 2015:
					r.Y2015 = c.Results
				}
				cr[c.ID] = r
			}
		}

		for k, v := range cr {
			b, err = json.MarshalIndent(v, "", "\t")
			//b, err := json.Marshal(results)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}

			filename = fmt.Sprintf(`out/%s.json`, k)
			file, err := os.Create(filename)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
			defer file.Close()

			_, err = file.Write(b)
			if err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		}
	}
}
