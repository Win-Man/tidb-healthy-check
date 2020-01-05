/*
 * Created: 2019-12-31 15:32:19
 * Author : Win-Man
 * Email : gang.shen0423@gmail.com
 * -----
 * Last Modified:
 * Modified By:
 * -----
 * Description:
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Win-Man/tidb-healthy-check/collector"
	"github.com/bndr/gotabulate"
)

var (
	h               bool
	config_path     string
	output_type     string
	pd_url          string
	tidb_status_url string
)

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	printTiDBClusterTable()
	printPDClusterTable()
	printTiKVClusterTable()
}

func init() {
	flag.StringVar(&config_path, "c", "", "tidb-ansible inventory.ini path")
	flag.StringVar(&output_type, "o", "print", "dest type: print|txt")
	flag.StringVar(&pd_url, "p", "", "pd cluster info")
	flag.StringVar(&tidb_status_url, "t", "", "tidb cluster node")
	flag.BoolVar(&h, "h", false, "help info")
	flag.Usage = usage
}

func printTiDBClusterTable() {
	tidbCluster, statusCode := collector.GetTiDBServerInfo(tidb_status_url)
	if 200 != statusCode {
		fmt.Printf("get tidb cluster status error\nplease try curl http://%v/info/all",tidb_status_url)
	}
	var tidbTableRows [][]interface{}
	for _, node := range tidbCluster.ServerList {
		tmpNode := []interface{}{fmt.Sprintf("%v:%v", node.Ip, node.Port), node.StatusPort, node.Lease}
		tidbTableRows = append(tidbTableRows, tmpNode)
	}

	tidbTable := gotabulate.Create(tidbTableRows)
	tidbTable.SetHeaders([]string{"IP","tidb_status_port","lease"})
	tidbTable.SetEmptyString("None")
	tidbTable.SetAlign("right")
	fmt.Println(tidbTable.Render("grid"))
}

func printPDClusterTable(){
	pdCluster,statusCode := collector.GetPDClusterInfo(pd_url)
	if 200 != statusCode {
		fmt.Printf("get pd cluster status error\nplease try curl")
	}
	var pdTableRows [][]interface{}
	for _,node := range pdCluster.NodeList{
		tmpNode := []interface{}{node.Name,node.PeerUrls}
		pdTableRows = append(pdTableRows,tmpNode)
	}

	pdTable := gotabulate.Create(pdTableRows)
	pdTable.SetHeaders([]string{"name","peer_url"})
	pdTable.SetEmptyString("None")
	pdTable.SetAlign("right")
	fmt.Println(pdTable.Render("grid"))
}

func printTiKVClusterTable(){
	tikvCluster,statusCode := collector.GetTiKVServerInfo(pd_url)
	if 200 != statusCode{
		fmt.Printf("get tikv cluster status error\nplease try curl")
	}
	var tikvTableRows [][]interface{}
	for _,node := range tikvCluster.ServerList{
		tmpNode := []interface{}{node.Address,node.StoreId,node.Version,node.State,node.RegionCount,node.LeaderCount}
		tikvTableRows = append(tikvTableRows,tmpNode)
	}
	tikvTable := gotabulate.Create(tikvTableRows)
	tikvTable.SetHeaders([]string{"IP","Store_id","Version","State","Region_Count","Leader_Count"})
	tikvTable.SetEmptyString("None")
	tikvTable.SetAlign("right")
	fmt.Println(tikvTable.Render("grid"))
}
// type TiKVNode struct {
// 	StoreId     int
// 	Address     string
// 	Version     string
// 	State       string
// 	LeaderCount int
// 	RegionCount int
// 	Uptime      string
// }


func usage() {
	fmt.Fprintf(os.Stderr, `tidb-healthy-check version: v0.1.0
Usage: 

Options:
`)

	flag.PrintDefaults()
}
