package theaudienceclub

import (
	"theatrenow/pkg/tools"
)

func (t *TheAudienceClubWebsite) Login() error {

	err := t.browser.Open(t.loginUrl)
	if err != nil {
		return err
	}
	tools.Logger().Info("landed on login webpage")

	fm, err := t.browser.Form("form[id='dialogForm']")
	if err != nil {
		return err
	}

	err = fm.Input("userName", "simonaquino@gmail.com")
	if err != nil {
		return err
	}

	err = fm.Input("password", "k0ljas1m0n")
	if err != nil {
		return err
	}

	err = fm.Submit()
	if err != nil {
		return err
	}

	tools.Logger().Info("Login successful")

	return nil
}
