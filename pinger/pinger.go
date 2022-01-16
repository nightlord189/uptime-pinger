package pinger

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"time"
)

type Pinger struct {
}

type CheckUrlResult struct {
	URL    string
	Result string
	Err    error
	Time   time.Duration
}

func (p *Pinger) CheckUrl(url string) CheckUrlResult {
	start := time.Now()
	result := CheckUrlResult{
		URL: url,
	}
	browser := rod.New()
	err := browser.Connect()
	if err != nil {
		result.Result = "browser_error"
		result.Err = err
		return result
	}
	defer browser.MustClose()
	browser = browser.Timeout(5 * time.Second)
	target := proto.TargetCreateTarget{
		URL: url,
	}
	page, err := browser.Page(target)
	if err != nil {
		result.Result = "page_error"
		result.Err = err
		return result
	}
	err = page.WaitLoad()
	if err != nil {
		result.Result = "wait_load_error"
		result.Err = err
		return result
	}
	result.Result = "success"
	result.Time = time.Since(start)
	return result
}
