package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

// Remove tracking/SEO parameters from the URL:
var urlParamBlacklist = []string{
	"wt_mc", "wtmc", "WT.mc_id", "wt_zmc",

	"ocid", "xid",

	"at_medium", "at_campaign", "at_custom1", "at_custom2",
	"at_custom3", "at_custom4",

	"utm_source", "utm_medium", "utm_campaign", "utm_term",
	"utm_content", "utm_name", "utm_referrer", "utm_brand",
	"utm_social-type", "utm_kxconfid",

	"guce_referrer", "guce_referrer_sig", "guccounter",

	"ga_source", "ga_medium", "ga_term", "ga_content",
	"ga_campaign", "ga_place",

	"pk_campaign", "pk_keyword", "pk_medium", "pk_source",

	"fb_action_ids", "fb_action_types", "fb_source", "fb_ref",
	"fbclid", "fbc",

	"hmb_campaign", "hmb_medium", "hmb_source",

	"newsticker", "CMP", "feature", "camp", "cid", "source",
	"ns_campaign", "ns_mchannel", "ito", "xg_source", "__tn__",
	"__twitter_impression", "share", "ncid", "rnd", "feed_id",
	"_unique_id", "GEPC", "pt", "xtor", "wtrid", "medium", "sara_ecid",
	"from", "inApp", "ssm", "campaign", "mbid", "s_campaign", "rss_id",
	"cmpid", "s_cid", "mv2", "scid", "sc2id", "sdid", "s_iid", "ssm",
	"spi_ref", "referrerlane",

	"share_bandit_exp", "share_bandit_var",

	"igshid", "idpartenaire",

	"aff_code", "affID",

	"recruited_by_id", "recruiter",
}

// Shortened URLs need to be undone first:
var shortenerList = []string{
	"bit.ly", "buff.ly", "dlvr.it",
	"goo.gl", "youtu.be", "tinyurl.com",
	"ow.ly", "amzn.to", "ift.tt", "zpr.io",
	"apple.co", "mol.im", "redd.it",
	"shar.es", "is.gd", "dld.bz",
	"trib.al", "fb.me", "tumblr.co",
	"cutt.ly", "app.link", "twib.in",
	"kko.to", "rsci.app.link", "upflow.co",
	"snip.ly", "lnk.to", "1jux.net", "gscoff.co",

	// News providers shortening themselves:
	"dw.com", "rt.com", "buzz.de", "crackm.ag",
	"glm.io", "opr.news", "sz.de", "rp-online.de",
	"mktw.net", "naver.me", "abcn.ws", "wapo.st",
	"vm.tiktok.com", "www.rsi.ch", "en.fut.ec",
	"kurz.zdf.de", "on.rtl.de", "a.msn.com",
	"va.newsrepublic.net", "www.cityam.com",
	"www.sosvox.org",
}

func contains(arr []string, str string) bool {
	// Array-containing check: Returns true if found.
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func ExpandUrl(url string) (string, error) {
	// URL expander with x509 checks disabled.
	expandedUrl := url

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			expandedUrl = req.URL.String()
			return nil
		},
		Transport: tr,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	return expandedUrl, nil
}

func processUrlItem(urlToCheck string) string {
	// Processes an URL item, returns the cleaned, unshortened
	// "expanded" URL.
	u, err := url.Parse(urlToCheck)
	if err != nil {
		log.Fatal(err)
	}

	// Some URL shorteners are not known to us (yet?).
	// Chances are that URLs with a path that ends in
	// "/SoMeSTRiNG123" are shortened. Catch them as well.
	re, _ := regexp.Compile("^/[A-Za-z0-9_-]{5,}$")
	potentialUrlShortener := re.MatchString(u.Path)

	if potentialUrlShortener || contains(shortenerList, u.Hostname()) {
		expandedUrl, err := ExpandUrl(urlToCheck)
		if err != nil {
			// Cannot reach the URL:
			return urlToCheck
		}

		// Overwrite the original URL by the expanded one:
		urlToCheck = expandedUrl

		// Parse again, just in case:
		u, err = url.Parse(urlToCheck)
		if err != nil {
			// Error in the updated domain:
			return urlToCheck
		}
	}

	// Remove tracking parameters:
	q := u.Query()
	for _, param := range urlParamBlacklist {
		q.Del(param)
		u.RawQuery = q.Encode()
	}

	return u.String()
}
