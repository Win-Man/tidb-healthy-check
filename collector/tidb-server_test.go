/*
 * Created: 2020-01-03 10:20:00
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

func TestGetTiDBServerInfo(t *testing.T) {
	var tests = []struct {
		url  string
		code int
	}{
		{"10.10.0.3:10080", 200},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("TestGetTiDBServerInfo:%v", tt.url)
		t.Run(testname, func(t *testing.T) {
			_, statusCode := GetTiDBServerInfo(tt.url)
			if statusCode != tt.code {
				t.Errorf("got %d", statusCode)
			}
		})
	}
}
