package main

import (
	"fmt"
	"imgwasm/imgtxt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("createImages", createImages()) // эти функции будут видны в Global области JavaScript
	js.Global().Set("getImageFromId", getImageFromId())
	<-make(chan bool) //  останавливаемся навсегда
}

// Js вызовет
func createImages() js.Func {
	// возвращаем фунцию, которая сможет принять параметры JavaScript
	// this - это реальный this js
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		str := args[0].String()

		arr := createImg(str)
		// ff := []byte{'d', 's', 't'}
		// x := jsutil.TemporaryUint8Array(len(ff), ff) // jsutil - смотреть
		// _ = x
		mm := make(map[string]interface{})
		mm["count"] = len(arr)
		mm["img"] = arr[0]
		return mm
	})
	return jsonFunc
}

// Основной запуск, получение всех картинок
func createImg(str string) []string {
	imgtxt.CreateImages(str)
	return imgtxt.Images64 // массив
}

// Js вызовет
func getImageFromId() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		numInt := args[0].Int()
		str := getImageNum(numInt)
		mm := make(map[string]interface{})
		mm["img"] = str
		return mm
	})
	return jsonFunc
}

// отдаем картинку по Id
func getImageNum(num int) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error: ", err)
		}
	}()
	if len(imgtxt.Images64) > 0 {
		//	imgtxt.ImageToDom(imgtxt.Images64[num])
		return imgtxt.Images64[num]
	}
	return ""
}
