/*
 * Created: 2020-01-03 10:15:30
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

type HttpTiDBNode struct {
	Version       string `json:"version"`
	GitHash       string `json:"git_hash"`
	DdlId         string `json:"ddl_id"`
	Ip            string `json:"ip"`
	ListeningPort int    `json:"listening_port"`
	StatusPort    int    `json:"status_port"`
	Lease         string `json:"lease"`
}

type HttpTiDBCluster struct {
	ServersNum        int64                   `json:"servers_num"`
	OwnerId           string                  `json:"owner_id"`
	VersionConsistent bool                    `json:"is_all_server_version_consistent"`
	AllServersInfo    map[string]HttpTiDBNode `json:"all_servers_info"`
}

type TiDBNode struct {
	Version    string
	Ip         string
	Port       int
	StatusPort int
	Lease      string
}

type TiDBCluster struct {
	ServersNum        int64
	VersionConsistent bool
	ServerList        []TiDBNode
}

func GetTiDBServerInfo(tidbStatusUrl string) (tidbCluster *TiDBCluster, statusCode int) {
	res, statusCode := requests.HttpGet(fmt.Sprintf("http://%v/info/all", tidbStatusUrl))
	if 200 != statusCode {
		tidbCluster = nil
		return
	}
	var httpTiDBCluster HttpTiDBCluster
	err := json.Unmarshal(res, &httpTiDBCluster)
	if err != nil {
		fmt.Println(err)
	}
	var serverList []TiDBNode
	for _, v := range httpTiDBCluster.AllServersInfo {
		serverList = append(serverList, TiDBNode{v.Version, v.Ip, v.ListeningPort, v.StatusPort, v.Lease})
	}
	tidbCluster = new(TiDBCluster)
	tidbCluster.ServersNum = httpTiDBCluster.ServersNum
	tidbCluster.VersionConsistent = httpTiDBCluster.VersionConsistent
	tidbCluster.ServerList = serverList
	//fmt.Printf("%+v\n", tidbCluster)
	return
}
