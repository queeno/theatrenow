package app

import (
	"theatrenow/pkg/theaudienceclub"
)

func Run() error {

	// Create a browser
	website := theaudienceclub.NewTheAudienceClubWebsite("https://www.theaudienceclub.com/login.php")

	err := website.Login()
	if err != nil {
		return err
	}

	err = website.SkipAnnouncements()
	if err != nil {
		return err
	}

	err = website.DownloadShows()
	if err != nil {
		return err
	}

	return nil
}
