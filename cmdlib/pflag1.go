package main

import (
	"fmt"
	"github.com/gosuri/uitable"
	flag "github.com/spf13/pflag"
)

var ip *int = flag.Int("flagname", 1234, "help message for flagname")

func main() {
	//flag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	//flag.Parse()
	////fmt.Printf("ip: %v", *ip)
	//
	//flagset := flag.NewFlagSet("ok", flag.ContinueOnError)
	//flagset.Set("ipp", "123")
	//getInt, err := flagset.GetInt("ipp")
	//if err != nil {
	//	fmt.Printf("err: %v", err)
	//} else {
	//	fmt.Printf("ipp: %v", getInt)
	//}

	table := uitable.New()
	table.MaxColWidth = 50

	table.AddRow("NAME", "BIRTHDAY", "BIO")
	hackers := []struct {
		Name     string
		Birthday string
		Bio      string
	}{
		{
			Name:     "Luca",
			Birthday: "2021-12-12",
			Bio:      "Oops!",
		},
		{
			Name:     "Luca1",
			Birthday: "2021-12-02",
			Bio:      "Oops!1",
		},
		{
			Name:     "Luca2",
			Birthday: "2021-12-13",
			Bio:      "Oops!2",
		},
	}
	for _, hacker := range hackers {
		table.AddRow(hacker.Name, hacker.Birthday, hacker.Bio)
	}
	fmt.Println(table)
}
