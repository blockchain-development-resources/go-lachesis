// +build debug

package poset

import (
	"github.com/Fantom-foundation/go-lachesis/src/common"
	"github.com/Fantom-foundation/go-lachesis/src/peers"
	"github.com/Fantom-foundation/go-lachesis/src/state"
)

// Store provides an interface for persistent and non-persistent stores
// to store key lachesis consensus information on a node.
type Store interface {
	TopologicalEvents() ([]Event, error)
	CacheSize() int
	Participants() (*peers.Peers, error)
	RootsBySelfParent() (map[string]Root, error)
	GetEventBlock(string) (Event, error)
	SetEvent(Event) error
	ParticipantEvents(string, int64) ([]string, error)
	ParticipantEvent(string, int64) (string, error)
	LastEventFrom(string) (string, bool, error)
	LastConsensusEventFrom(string) (string, bool, error)
	KnownEvents() map[int64]int64
	ConsensusEvents() []string
	ConsensusEventsCount() int64
	AddConsensusEvent(Event) error
	GetRound(int64) (RoundCreated, error)
	SetRound(int64, RoundCreated) error
	LastRound() int64
	RoundClothos(int64) []string
	RoundEvents(int64) int
	GetRoot(string) (Root, error)
	GetBlock(int64) (Block, error)
	SetBlock(Block) error
	LastBlockIndex() int64
	GetFrame(int64) (Frame, error)
	SetFrame(Frame) error
	Reset(map[string]Root) error
	Close() error
	NeedBootstrap() bool // Was the store loaded from existing db
	StorePath() string
	// StateDB returns state database
	StateDB() state.Database
	StateRoot() common.Hash
}
