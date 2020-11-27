package scrapper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/structs"
)

type extractedJob struct {
	ID       string
	Title    string
	Location string
	Salary   string
	Summary  string
}

// Scrape scraps job-posting
func Scrape(term, fileFmt string) {
	// csv, json
	var jobs []extractedJob
	c := make(chan []extractedJob)
	baseURL := "https://kr.indeed.com/jobs?q=" + "term" + "&limit=50"

	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		go getPage(baseURL, i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs, fileFmt)
	fmt.Println("Done, extracted", len(jobs))
}

func getApplyURL(id string) string {
	baseApplyURL := "https://kr.indeed.com/viewjob?jk="
	return baseApplyURL + id
}

func getPage(baseURL string, page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageUnit := 50
	pageURL := baseURL + "&start=" + strconv.Itoa(page*pageUnit)
	fmt.Println("Requesting ", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find("sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	c <- extractedJob{
		ID:       id,
		Title:    title,
		Location: location,
		Salary:   salary,
		Summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) (pageCount int) {
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	// response의 body는 io 처리된다. (io.Reader type)
	// TODO: python에서의 HTTP.get처리도 with을 사용해야 할까?
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pageCount = s.Find("a").Length()
	})

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}

}

// write to csv
func writeJobs(jobs []extractedJob, fmt string) {
	fileName := "jobs" + "." + fmt
	if fmt == "json" {
		jsonString, _ := json.MarshalIndent(jobs, "", "  ")
		err := ioutil.WriteFile(fileName, jsonString, 0644)
		checkErr(err)
	} else if fmt == "csv" {
		file, err := os.Create(fileName)
		checkErr(err)

		w := csv.NewWriter(file)
		headers := structs.Names(&extractedJob{})
		wErr := w.Write(headers)
		checkErr(wErr)

		for _, job := range jobs {
			jobSlice := []string{
				getApplyURL(job.ID),
				job.Title,
				job.Location,
				job.Salary,
				job.Summary,
			}
			jwErr := w.Write(jobSlice)
			checkErr(jwErr)
		}
		defer func() {
			w.Flush()
			file.Close()
		}()
	}

}
