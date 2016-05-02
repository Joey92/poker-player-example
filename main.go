package main

import (
    poker "github.com/lean-poker/poker-player-go"
    "github.com/lean-poker/poker-player-go/leanpoker"
)


type Player struct{}

func (p *Player) BetRequest(state *leanpoker.Game) int {
    return 0;
}

func (p *Player) Showdown(state *leanpoker.Game) {
    
}

func (p *Player) Version() string {
    return "Default folding Go player"
}

func main() {
    p := &Player{}
    
    poker.Start(p)
}