/*
 * Created: 2020-01-03 11:04:07
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

// StoreStatus contains status about a store.
type StoreStatus struct {
	Capacity        string  `json:"capacity"`
	Available       string  `json:"available"`
	LeaderCount     int     `json:"leader_count"`
	LeaderWeight    float64 `json:"leader_weight"`
	LeaderScore     float64 `json:"leader_score"`
	LeaderSize      int64   `json:"leader_size"`
	RegionCount     int     `json:"region_count"`
	RegionWeight    float64 `json:"region_weight"`
	RegionScore     float64 `json:"region_score"`
	RegionSize      int64   `json:"region_size"`
	StartTS         string  `json:"start_ts"`
	LastHeartbeatTS string  `json:"last_heartbeat_ts"`
	Uptime          string  `json:"uptime"`
}

type StoreBaseInfo struct {
	Id      int    `json:"id"`
	Ip      string `json:"address"`
	Version string `json:"version"`
	State   string `json:"state_name"`
}

type StoreInfo struct {
	Store  StoreBaseInfo `json:"store"`
	Status StoreStatus   `json:"status"`
}

type StoreClusterInfo struct {
	ServerNum  int         `json:"count"`
	ServerList []StoreInfo `json:"stores"`
}

type TiKVNode struct {
	StoreId     int
	Address     string
	Version     string
	State       string
	LeaderCount int
	RegionCount int
	Uptime      string
}

type TiKVCluster struct {
	ServerNum  int
	ServerList []TiKVNode
}

func GetTiKVServerInfo(pdUrl string) (tikvCluster *TiKVCluster, statusCode int) {
	res, statusCode := requests.HttpGet(fmt.Sprintf("http://%v/pd/api/v1/stores", pdUrl))
	if 200 != statusCode {
		tikvCluster = nil
		return
	}
	var storeClusterInfo StoreClusterInfo
	err := json.Unmarshal(res, &storeClusterInfo)
	if err != nil {
		fmt.Println(err)
	}
	tikvCluster = new(TiKVCluster)
	var serverList []TiKVNode
	for _, v := range storeClusterInfo.ServerList {
		serverList = append(serverList, TiKVNode{v.Store.Id,
			v.Store.Ip, v.Store.Version, v.Store.State,
			v.Status.LeaderCount, v.Status.RegionCount, v.Status.Uptime,
		})
	}
	tikvCluster.ServerNum = storeClusterInfo.ServerNum
	tikvCluster.ServerList = serverList
	//fmt.Printf("%+v\n", tikvCluster)
	return
}
