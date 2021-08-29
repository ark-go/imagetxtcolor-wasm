//g o:build wasm && js

package imgtxt

import (
	"fmt"

	"github.com/ark-go/imgtxtcolor/pkg/imgtxtcolor"
)

var Images64 []string

func CreateImages(str string) {
	images, err := imgtxtcolor.CreateImageTextLog(str, nil, imgtxtcolor.LogConsole) //CreateImageText(str, nil)
	if err != nil {
		fmt.Println("[createImages]", err)
	}
	fmt.Println("[createImages готово]")
	Images64 = getBase64(images)

	// for _, txt := range images64 {
	// 	imageToDom(txt)
	// }

}

// func ImageToDom(txt64 string) {
// 	jsDoc := js.Global().Get("document")
// 	if !jsDoc.Truthy() {
// 		fmt.Println("Не получить document")
// 	}
// 	zoneImage := jsDoc.Call("getElementById", "zoneImage")
// 	if !zoneImage.Truthy() {
// 		fmt.Println("не получить zoneImage")
// 	}
// 	zoneImage.Set("innerHTML", "")
// 	img := jsDoc.Call("createElement", "img")

// 	zoneImage.Call("appendChild", img)

// 	img.Set("src", "data:image/png;base64,"+txt64)

// }
