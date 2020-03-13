package theaudienceclub

import (
	"github.com/headzoo/surf/errors"
	"theatrenow/pkg/tools"
)

func (t *TheAudienceClubWebsite) SkipAnnouncements() error {
	tools.Logger().Info("skipping announcement page...")

	fm, err := t.browser.Form("form[value='Click Here']")
	// the page might not appear
	if err == nil {
		err2 := fm.Click("Click Here")
		if err2 != nil {
			return err2
		}
		tools.Logger().Info("announcement page skipped")
	} else {
		if _, ok := err.(errors.ElementNotFound); !ok {
			return err
		}
		tools.Logger().Info("announcement page not found. Continuing...")
	}

	return nil
}
