package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	var date string
	var birhday string
	app := cli.NewApp()
	app.Name = "tgh"
	app.Description = "persian calender"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "convert , c",
			Usage:       "convert miladi to shmasi date",
			Destination: &date,
		},
		cli.StringFlag{
			Name:        "brtihday, b",
			Usage:       "age calculator",
			Destination: &birhday,
		},
	}
	app.Action = func(c *cli.Context) error {
		convert := c.String("convert")
		birh := c.String("brtihday")
		if convert != "" && birh != "" {
			fmt.Println("please choose one option!")
		} else if convert != "" {
			if date != "" {
				t, err := parse(date)
				if err != nil {
					fmt.Println(err)
				} else {
					y, m, d := calc(int(t.Year()), int(t.Month()), int(t.Day()), is_leap(int(t.Year())))

					fmt.Println(d, persian_month[m], y)
				}
			}

		} else if birhday != "" {
			b, err := parse(birhday)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(time.Now().Year()-b.Year(), "Year", int(time.Now().Month()-b.Month()), "Month", time.Now().Day()-b.Day(), "Day", time.Now().Hour()-b.Hour(), "Hour", time.Now().Minute()-b.Minute(), "Minute", time.Now().Second()-b.Second(), "Second")
			}
		} else {
			cprint(print_str_time())
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
