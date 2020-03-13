package theaudienceclub

import (
	"fmt"
	"strings"
)

type show struct {
	title string
	description string
	genre string
	venue string
	date string
}


func (t *TheAudienceClubWebsite) getShowInfo(attribute string) []string {
	e := t.browser.Find("." + attribute)
	eSlice := make([]string, e.Length())
	b := e.Find("a[target='_self']")

	s := strings.Split(b.Text(), "\n")

	for _, e := range s {
		fmt.Println(strings.TrimSpace(e))
	}
	//for i := 0; i < e.Nodes; i++ {
//	for _, n := range e. {
//		fmt.Println(n.Data)
		//eSlice = append(eSlice, strings.Trim(e.Text(), " \n\t"))
//	}

	return eSlice
}

func (t *TheAudienceClubWebsite) DownloadShows() error {

	err := t.browser.Click("a[href='/member/theatreLadderAllGenre.php']")
	if err != nil {
		return err
	}
	//fmt.Println(t.browser.Body())
	_ = t.getShowInfo("showtitle")
	//genres := t.getShowInfo("genre")
	//descriptions := t.getShowInfo("showdescription")
	//venues := t.getShowInfo("venue")
	//_ = t.getShowInfo("dateTime")

	//fmt.Println(titles)
	//fmt.Println(genres)
	//fmt.Println(descriptions)
	//fmt.Println(venues)
	//fmt.Println(dates)


	return nil
}