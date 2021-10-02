package chrome

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

type Options struct {
	Url string
	Width int
	Height int
	Quality int
	Format string
}

func ScreenshotTasks(options Options, imageBuf *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(
			int64(options.Width),
			int64(options.Height),
			1.0,
			false),
		chromedp.Navigate(options.Url),
		chromedp.ActionFunc(func(ctx context.Context) (err error) {
			*imageBuf, err = page.CaptureScreenshot().
				WithFormat(page.CaptureScreenshotFormat(options.Format)).
				WithQuality(int64(options.Quality)).
				Do(ctx)
			return err
		}),
	}
}
