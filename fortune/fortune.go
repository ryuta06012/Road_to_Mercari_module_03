/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   fortune.go                                         :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hryuuta <hryuuta@student.42tokyo.jp>       +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2022/01/14 12:49:13 by hryuuta           #+#    #+#             */
/*   Updated: 2022/01/14 13:28:30 by hryuuta          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package fortune

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Ping struct {
	Fortune string `json:"fortune"`
	Study   string `json:"study"`
}

func (p Ping) Init() {
	p.Fortune = ""
	p.Study = ""
}

var fortuneList = []string{"Dai-kichi", "Kichi", "Chuu-kichi", "Sho-kichi", "Sue-kichi", "Kyo", "Dai-kyo"}

func selectContentRandom(fortune string) Ping {
	var p Ping
	switch fortune {
	case "Dai-kichi":
		p = Ping{Fortune: fortune, Study: "Everything will be fine."}
	case "Kichi":
		p = Ping{Fortune: fortune, Study: "Basically, things are going well."}
	case "Chuu-kichi":
		p = Ping{Fortune: fortune, Study: "Not so bad."}
	case "Sho-kichi":
		p = Ping{Fortune: fortune, Study: "Let's stay alert and Study."}
	case "Sue-kichi":
		p = Ping{Fortune: fortune, Study: "If we don't, we'll fail."}
	case "Kyo":
		p = Ping{Fortune: fortune, Study: "It's hopeless."}
	case "Dai-kyo":
		p = Ping{Fortune: fortune, Study: "Give it up"}
	}
	return p
}

func (p Ping) selectFortuneRandom() Ping {
	rand.Seed(time.Now().UnixNano())
	content := fortuneList[rand.Int()%len(fortuneList)]

	return selectContentRandom(content)
}

func Oracle() Ping {
	var p Ping
	p.Init()
	t := time.Now()
	if t.Month() == 1 && (t.Day() >= 1 && t.Day() <= 3) {
		return selectContentRandom("Dai-kichi")
	}
	res := p.selectFortuneRandom()
	return res
}

func PingHandler(w http.ResponseWriter, r *http.Request) {

	ping := Oracle()
	output, err := json.MarshalIndent(ping, "", "  ")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Println(string(output))
	w.Write([]byte(ping.Fortune))
}
