package mock

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

func GetTest() echo.HandlerFunc {
	return getEchoHandler("GetTest")
}

func PostTest() echo.HandlerFunc {
	return getEchoHandler("PostTest")
}

// レスポンス作成
func getEchoHandler(filename string) echo.HandlerFunc {
	return func(c echo.Context) error {
		// read json file
		bytes, err := readFile("./mock/json/" + filename + ".json")
		if err != nil {
			fmt.Println("error:", err)
			return c.String(http.StatusOK, "file read error")
		}

		// json parse
		var response interface{}
		errUnmarshal := json.Unmarshal(bytes, &response)
		if errUnmarshal != nil {
			fmt.Println("error:", errUnmarshal)
			return c.String(http.StatusOK, "json parse error")
		} else {
			return c.JSON(http.StatusOK, response)
		}
	}
}

// ファイルを読み込む
func readFile(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	return bytes, err
}