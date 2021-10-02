package routes

import (
	"context"
	"fmt"
	"github.com/Sh4yy/ImageRender/chrome"
	"github.com/Sh4yy/ImageRender/utils"
	"github.com/chromedp/chromedp"
	"github.com/gorilla/mux"
	"github.com/joncalhoun/qson"
	"log"
	"net/http"
)

// createRenderURL generates target url for chrome to render the image
func createRenderURL(template, query string, port int) string {
	return fmt.Sprintf("http://localhost:%d/render/%s?%s", port, template, query)
}

// createOptions creates default options for the chrome renderer
func createOptions(r *http.Request) chrome.Options {
	template := mux.Vars(r)["template"]
	url := createRenderURL(template, r.URL.RawQuery, utils.Config.Server.Port)
	data := struct {
		Width   int    `json:"width,omitempty"`
		Height  int    `json:"weight,omitempty"`
		Quality int    `json:"quality,omitempty"`
		Format  string `json:"format,omitempty"`
	}{
		Width:   utils.Config.Render.Width,
		Height:  utils.Config.Render.Height,
		Quality: utils.Config.Render.Quality,
		Format:  utils.Config.Render.Format,
	}
	err := qson.Unmarshal(&data, r.URL.RawQuery)
	if err != nil {
		fmt.Println(err)
	}
	return chrome.Options{
		Url:     url,
		Width:   data.Width,
		Height:  data.Height,
		Quality: data.Quality,
		Format:  data.Format,
	}
}

func Capture(w http.ResponseWriter, r *http.Request) {

	// target options for chromedp
	options := createOptions(r)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var imageBuf []byte
	if err := chromedp.Run(ctx, chrome.ScreenshotTasks(options, &imageBuf)); err != nil {
		log.Fatal(err)
	}

	// content type format header
	switch options.Format {
	case "jpeg":
		w.Header().Set("Content-Type", "image/jpeg")
		break
	case "png":
		w.Header().Set("Content-Type", "image/png")
		break
	default:
		http.Error(w, "invalid content format", 400)
		return
	}

	_, err := w.Write(imageBuf)
	if err != nil {
		fmt.Println(err)
	}
}
