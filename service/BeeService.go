package service

import (
	"encoding/json"
	"fmt"
	"swarm/public"
)

type BeeService struct {
	Nodes []string
}

type Peers struct {
	Peers []struct {
		Address  string `json:"address"`
		FullNode bool   `json:"fullNode"`
	} `json:"peers"`
}

func (bs *BeeService) GetPeers() (p map[string]int) {
	for _, node := range bs.Nodes {
		url := fmt.Sprintf("http://%v:1635/peers", node)
		body, err := public.HttpGet(url)
		if err != nil {
			p[node] = -1
			return
		}
		peers := Peers{}
		err = json.Unmarshal([]byte(body), &peers)
		if err != nil {
			p[node] = -1
			return
		}
		fmt.Println(node)
		fmt.Println(len(peers.Peers))
		p[node] = len(peers.Peers)
	}
	fmt.Println(p)
	return
}

func (bs *BeeService) GetStatus(baseUrl string) int {
	url := "http://" + baseUrl + ":1635/"
	body, err := public.HttpGet(url)
	if err != nil {
		return 0
	}
	peers := Peers{}
	json.Unmarshal([]byte(body), &peers)
	if err != nil {
		return 0
	}
	number := len(peers.Peers)
	return number
}
