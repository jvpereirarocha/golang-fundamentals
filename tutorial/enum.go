package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string {
	StateIdle: "idle",
	StateConnected: "connected",
	StateError: "error",
	StateRetrying: "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println("NS: ", ns)
	
	ns2 := transition(ns)
	fmt.Println("NS2: ", ns2)

	ns3 := transition(ns2)
	fmt.Println("NS3: ", ns3)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}