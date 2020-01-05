/*
 * Created: 2020-01-03 11:40:43
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */
package collector

import (
	"fmt"
	"testing"
)

func TestGetTiKVServerInfo(t *testing.T) {
	var tests = []struct {
		url  string
		code int
	}{
		{"10.10.0.3:2379", 200},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("TestGetTiKVServerInfo:%v", tt.url)
		t.Run(testname, func(t *testing.T) {
			_, statusCode := GetTiKVServerInfo(tt.url)
			if statusCode != tt.code {
				t.Errorf("got %d", statusCode)
			}
		})
	}
}
