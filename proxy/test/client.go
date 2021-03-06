package main

import (
	"flag"
	"log"
	"os"
	"time"

	mysql_proxy "github.com/huoshan017/mysql-go/proxy/client"
	"github.com/huoshan017/mysql-go/proxy/test/game_db"
)

func main() {
	if len(os.Args) < 1 {
		log.Printf("args not enough, must specify a config file for db define\n")
		return
	}

	host_arg := flag.String("h", "", "config host server")
	flag.Parse()

	var host string
	if nil != host_arg {
		host = *host_arg
		log.Printf("config file path %v\n", host)
	} else {
		log.Printf("not specified config file arg\n")
		return
	}

	var proxy_addr string = host
	var db_proxy mysql_proxy.DB
	err := db_proxy.Connect(proxy_addr)
	if err != nil {
		log.Printf("db proxy connect err %v\n", err.Error())
		return
	}

	db_proxy.RunBackground()

	tb_mgr := game_db.NewTablesProxyManager(&db_proxy, 1, "game_db")
	player_table_proxy := tb_mgr.GetT_PlayerTableProxy()

	//field_name := "id"

	var tp game_db.T_Player
	for id := uint32(1); id <= 100000; id++ {
		tp.Set_id(id)
		tp.Set_role_id(uint64(10000 + id))
		player_table_proxy.Insert(&tp)
	}

	/*go func() {
		var err error
		var p *game_db.T_Player
		var ps []*game_db.T_Player
		var id int = 1
		for ; id < 10; id++ {
			p, err = player_table_proxy.Select(field_name, id)
			if err != nil {
				log.Printf("select id %v err %v\n", id, err.Error())
				continue
			}

			log.Printf("selected player: %v\n", p)
		}

		for i := 0; i < 100000; i++ {
			ps, err = player_table_proxy.SelectRecords("level", 1)
			if err != nil {
				log.Printf("selected player records err %v\n", err.Error())
			} else {
				log.Printf("selected players: %v\n", ps)
			}

			ps, err = player_table_proxy.SelectAllRecords()
			if err != nil {
				log.Printf("selected all player records err %v\n", err.Error())
			} else {
				log.Printf("selected all players: %v\n", ps)
			}

			ps, err = player_table_proxy.SelectRecordsCondition("vip_level", 1, nil)
			if err != nil {
				log.Printf("selected records condition err %v\n", err.Error())
			} else {
				log.Printf("selected records condition players: %v\n", ps)
			}

			time.Sleep(time.Millisecond * 1)
		}
	}()*/

	for {
		time.Sleep(time.Second)
	}
}
