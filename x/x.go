package x

import (
	"fmt"
	"os"
	"strings"

	"github.com/webview/webview"
)

// Client webview client
var Client webview.WebView

// CreateWebview create webview
func CreateWebview(url string) {
	debug := true
	Client = webview.New(debug)
	defer Client.Destroy()
	var w, h = 440, 600
	Client.SetSize(w, h, webview.HintNone)
	Client.Navigate(url)
	var x = strings.Join(os.Args, ", ")
	var b = fmt.Sprintf(`console.log("%v")`, x)
	Client.Init(b)
	Client.SetTitle("小夕")
	Client.Run()
}
