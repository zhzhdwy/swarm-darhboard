package service

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
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

type Addresses struct {
	Overlay      string   `json:"overlay"`
	Underlay     []string `json:"underlay"`
	Ethereum     string   `json:"ethereum"`
	PublicKey    string   `json:"publicKey"`
	PssPublicKey string   `json:"pssPublicKey"`
}

func (bs *BeeService) GetPeers() map[string]int {
	p := make(map[string]int)
	for _, node := range bs.Nodes {
		split := strings.Split(node, ":")
		url := fmt.Sprintf("http://%v/peers", node)
		body, _, err := public.HttpGet(url)
		if err != nil {
			p[split[0]] = -1
		}
		peers := Peers{}
		err = json.Unmarshal([]byte(body), &peers)
		if err != nil {
			p[split[0]] = -1
		}
		p[split[0]] = len(peers.Peers)
	}
	return p
}

func (bs *BeeService) GetEAddress() map[string]string {
	a := make(map[string]string)
	for _, node := range bs.Nodes {
		split := strings.Split(node, ":")
		url := fmt.Sprintf("http://%v/addresses", node)
		body, _, err := public.HttpGet(url)
		if err != nil {
			a[split[0]] = "nil"
		}
		address := Addresses{}
		err = json.Unmarshal([]byte(body), &address)
		if err != nil {
			a[split[0]] = "nil"
		}
		a[split[0]] = address.Ethereum
	}
	return a
}

func (bs *BeeService) GetVersion() map[string]string {
	v := make(map[string]string)
	regVersion := regexp.MustCompile(`bee_info{version="(?P<version>\d+\.\d+\.\d+)\-\w+"}`)
	for _, node := range bs.Nodes {
		split := strings.Split(node, ":")
		url := fmt.Sprintf("http://%v/metrics", node)
		body, _, err := public.HttpGet(url)
		if err != nil {
			v[split[0]] = "0"
		}
		match := regVersion.FindStringSubmatch(string(body))
		v[split[0]] = match[1]
	}
	return v
}

func (bs *BeeService) GetPort() map[string]string {
	port := make(map[string]string)
	for _, node := range bs.Nodes {
		split := strings.Split(node, ":")
		if len(split) < 2 {
			log.Fatal("主机信息有误")
		}
		port[split[0]] = split[1]
	}
	return port
}

func (bs *BeeService) GetAlive() (map[string]int, int, int) {
	s := make(map[string]int)
	alive := 0
	dead := 0
	for _, node := range bs.Nodes {
		split := strings.Split(node, ":")
		url := fmt.Sprintf("http://%v/metrics", node)
		_, code, err := public.HttpGet(url)
		if err != nil || code != 200 {
			s[split[0]] = 0
			dead = dead + 1
		}
		s[split[0]] = 1
		alive = alive + 1
	}
	return s, alive, dead
}
