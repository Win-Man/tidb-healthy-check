/*
 * Created: 2020-01-02 15:34:11
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */

package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (body []byte,code  int) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("error:%v", err)
	}
	defer response.Body.Close()
	body, _ = ioutil.ReadAll(response.Body)
	code = response.StatusCode
	return
}

func httpPost(url string) (body []byte) {
	return
}
