/*
 * Created: 2020-01-03 15:10:23
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
 
 func TestGetPDServerInfo(t *testing.T) {
	 var tests = []struct {
		 url  string
		 code int
	 }{
		 {"10.10.0.1:2379", 200},
	 }
 
	 for _, tt := range tests {
		 testname := fmt.Sprintf("TestGetPDServerInfo:%v", tt.url)
		 t.Run(testname, func(t *testing.T) {
			 _, statusCode := GetPDClusterInfo(tt.url)
			 if statusCode != tt.code {
				 t.Errorf("got %d", statusCode)
			 }
		 })
	 }
 }