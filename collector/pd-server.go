/*
 * Created: 2020-01-03 14:18:23
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
	"encoding/json"
	"fmt"

	"github.com/Win-Man/tidb-healthy-check/pkg/requests"
)

type HttpPDInfo struct {
	Name       string   `json:"name"`
	MemberId   uint64   `json:"member_id"`
	PeerUrls   []string `json:"peer_urls"`
	ClientUrls []string `json:"client_urls"`
}

type HttpPDClusterInfo struct {
	Members    []HttpPDInfo `json:"members"`
	Leader     HttpPDInfo   `json:"leader"`
	EtcdLeader HttpPDInfo   `json:"etcd_leader"`
}

type PDClusterInfo struct {
	ServerNum  int
	NodeList   []HttpPDInfo
	LeaderNode HttpPDInfo
}

func GetPDClusterInfo(pdUrl string) (pdClusterInfo *PDClusterInfo, statusCode int) {

	res, statusCode := requests.HttpGet(fmt.Sprintf("http://%v/pd/api/v1/members", pdUrl))
	if 200 != statusCode {
		pdClusterInfo = nil
		return
	}
	var httpPDClusterInfo HttpPDClusterInfo
	err := json.Unmarshal(res, &httpPDClusterInfo)
	if err != nil {
		fmt.Println(err)
	}
	pdClusterInfo = new(PDClusterInfo)
	pdClusterInfo.NodeList = httpPDClusterInfo.Members
	pdClusterInfo.LeaderNode = httpPDClusterInfo.Leader
	pdClusterInfo.ServerNum = len(httpPDClusterInfo.Members)
	//fmt.Printf("%+v\n", pdClusterInfo)
	return
}
