package theaudienceclub

import (
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
)

type TheAudienceClubWebsite struct {
	browser *browser.Browser
	loginUrl string
}

func NewTheAudienceClubWebsite(loginUrl string) *TheAudienceClubWebsite {
	return &TheAudienceClubWebsite{
		browser: surf.NewBrowser(),
		loginUrl: loginUrl,
	}
}
