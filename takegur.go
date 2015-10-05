package main

/*
Made with <3 from picturesarenice and emi8ly.
*/

import (
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/codegangsta/cli"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

type ImgurJson struct {
	Data    `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type Data []struct {
	ID             string      `json:"id"`
	Title          string      `json:"title"`
	Description    interface{} `json:"description"`
	Datetime       int         `json:"datetime"`
	Type           string      `json:"type"`
	Animated       bool        `json:"animated"`
	Width          int         `json:"width"`
	Height         int         `json:"height"`
	Size           int         `json:"size"`
	Views          int         `json:"views"`
	Bandwidth      int64       `json:"bandwidth"`
	Vote           interface{} `json:"vote"`
	Favorite       bool        `json:"favorite"`
	Nsfw           bool        `json:"nsfw"`
	Section        string      `json:"section"`
	AccountURL     string      `json:"account_url"`
	AccountID      int         `json:"account_id"`
	CommentPreview interface{} `json:"comment_preview"`
	Topic          string      `json:"topic"`
	TopicID        int         `json:"topic_id"`
	Link           string      `json:"link"`
	CommentCount   int         `json:"comment_count"`
	Ups            int         `json:"ups"`
	Downs          int         `json:"downs"`
	Points         int         `json:"points"`
	Score          int         `json:"score"`
	IsAlbum        bool        `json:"is_album"`
}

func Reader(regex *regexp.Regexp) {
	var jsonResults ImgurJson
	var wg sync.WaitGroup

	fmt.Println("compiling searcher.")
	time.Sleep(1 * time.Second)

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.imgur.com/3/gallery.json", nil)
	// my API key.
	req.Header.Add()

	fmt.Println("downloading data from imgur.")
	resp, err := client.Do(req)
	fmt.Println("server status: ", resp.StatusCode)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()


	err = json.NewDecoder(resp.Body).Decode(&jsonResults)

	for key := range jsonResults.Data {
		//				fmt.Println(jsonResults.Data[key].Title)
		if regex.MatchString(jsonResults.Data[key].Title) == true {
			fmt.Println(jsonResults.Data[key].Title)

			/*
				start parsing through the album images and increment the waitgroup counter.
			*/
			if jsonResults.Data[key].IsAlbum == true {
				albumUri := fmt.Sprintf("https://api.imgur.com/3/album/%s/images", jsonResults.Data[key].ID)
				fmt.Println(albumUri)
				req, err := http.NewRequest("GET", albumUri, nil)
				req.Header.Add()
				if err != nil {
					fmt.Println(err)
				}
				resp, err := client.Do(req)
				fmt.Println("image status: ", resp.StatusCode)
				defer resp.Body.Close()
				//handles the new album json.
				var albumResults ImgurJson
				err = json.NewDecoder(resp.Body).Decode(&albumResults)

				for index := range albumResults.Data {
					wg.Add(1)
					//				fmt.Println(albumResults.Data[index].Link)
					go imageDownloader(albumResults.Data[index].Link, albumResults.Data[index].ID, &wg)
				}
				wg.Wait()
			}

			/*
			this is for indidividual images.
			 */
			if jsonResults.Data[key].IsAlbum == false {
				imageUri := fmt.Sprintf("https://api.imgur.com/3/image/%s/", jsonResults.Data[key].ID)
				fmt.Println(imageUri)
				req, err := http.NewRequest("GET", imageUri, nil)
				req.Header.Add()
				if err != nil {
					fmt.Println(err)
				}
				resp, err := client.Do(req)
				fmt.Println("image status: ", resp.StatusCode)
				defer resp.Body.Close()
				wg.Add(1)
				go imageDownloader(jsonResults.Data[key].Link, jsonResults.Data[key].ID, &wg)
				wg.Wait()
			}

			fmt.Println("done.")
		}
	}
	fmt.Println("done!")
}

func imageDownloader(uri string, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	tokens := strings.Split(uri, "/")
	fileName := tokens[len(tokens)-1]

	outFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", uri, nil)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	header := resp.ContentLength
	bar := pb.New(int(header)).SetUnits(pb.U_BYTES)
	bar.SetRefreshRate(time.Millisecond)
	//	bar.Start()
	rd := bar.NewProxyReader(resp.Body)
	// and copy from reader
	io.Copy(outFile, rd)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	/*
		let's download some imgur!
	*/

	wp, err := regexp.Compile(`\b[wW]all[pP]aper\b|[wW]all[pP]apers\b`)
	bt, err := regexp.Compile(`\b[bB]lack\b\s[tT]wi[t]{1,2}er`)
	mobile, err := regexp.Compile(`\b[Mm]obile\s\[wW]all[pP]apers\b`)
	db, err := regexp.Compile(`\b[Dd]ickbu[t]{1,2}\b|\b[Dd]ick\b\s\bbu[t]{1,2}\b`)
	sc, err := regexp.Compile(`\b[sS]tay\b\s[cC]la[s]{1,2}y\b`)
	da, err := regexp.Compile(`\b[Dd]arwin\b\s[aA]wards\b`)
	fails, err := regexp.Compile(`\b[fF]ails\b|[fF]ails`)
	reactions, err := regexp.Compile(`\bmrw|MRW\b`)

	if err != nil {
		fmt.Println(err)
	}

	app := cli.NewApp()

	app.Name = "takegur"
	app.Authors = []cli.Author{
		{Name: "picturesarenice"},
		{Name: "emi8ly"},
	}
	app.Version = "1.0"

	app.Commands = []cli.Command{
		{
			Name:  "black-twitter",
			Usage: "when u tell her 2 stop n she keeps suckin...",
			Action: func(c *cli.Context) {
				Reader(bt)
			},
		},
		{
			Name:  "wallpapers",
			Usage: "ALL THE DESKTOP BACKGROUNDS ARE BELONG TO YOU.",
			Action: func(c *cli.Context) {
				Reader(wp)
			},
		},
		{
			Name:  "mobile",
			Usage: "for when you need to keep looking at your phone to avoid meetings.",
			Action: func(c *cli.Context) {
				Reader(mobile)
			},
		},
		{
			Name:  "dickbutt",
			Usage: "please don't ever use this.",
			Action: func(c *cli.Context) {
				Reader(db)
			},
		},
		{
			Name:  "stay-classy",
			Usage: "you're a sick bastard.",
			Action: func(c *cli.Context) {
				Reader(sc)
			},
		},
		{
			Name:  "darwin-awards",
			Usage: "let's watch some stupid people!",
			Action: func(c *cli.Context) {
				Reader(da)
			},
		},
		{
			Name:  "fails",
			Usage: "ouch.",
			Action: func(c *cli.Context) {
				Reader(fails)
			},
		},
		{
			Name:  "mrw",
			Usage: "your reaction when...",
			Action: func(c *cli.Context) {
				Reader(reactions)
			},
		},
	}

	app.Usage = "used to download the awesomeness of imgur."

	app.Run(os.Args)
}
