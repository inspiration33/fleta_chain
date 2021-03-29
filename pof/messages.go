package pof

import (
	"github.com/fletaio/fleta/common"
	"github.com/fletaio/fleta/common/hash"
	"github.com/fletaio/fleta/core/types"
)

var (
	BlockGenMessageType = types.DefineHashedType("pof.BlockGenMessage")
)

// RoundVote is a message for a round vote
type RoundVote struct {
	ChainID              uint8
	LastHash             hash.Hash256
	TargetHeight         uint32
	TimeoutCount         uint32
	Formulator           common.Address
	FormulatorPublicHash common.PublicHash
	Timestamp            uint64
	IsReply              bool
}

// RoundVoteMessage is a message for a round vote
type RoundVoteMessage struct {
	RoundVote *RoundVote
}

// RoundVoteAck is a message for a round vote ack
type RoundVoteAck struct {
	ChainID              uint8
	LastHash             hash.Hash256
	TargetHeight         uint32
	TimeoutCount         uint32
	Formulator           common.Address
	FormulatorPublicHash common.PublicHash
	PublicHash           common.PublicHash
	Timestamp            uint64
	IsReply              bool
}

// RoundVoteAckMessage is a message for a round vote
type RoundVoteAckMessage struct {
	RoundVoteAck *RoundVoteAck
}

// NextRoundVote is a message for a next round vote
type NextRoundVote struct {
	ChainID              uint8
	TimeoutCount         uint32
	Formulator           common.Address
	FormulatorPublicHash common.PublicHash
	Timestamp            uint64
	IsReply              bool
}

// NextRoundVoteMessage is a message for a next round vote
type NextRoundVoteMessage struct {
	RoundVote *RoundVote
}

// NextRoundVoteAck is a message for a next round vote ack
type NextRoundVoteAck struct {
	ChainID              uint8
	TimeoutCount         uint32
	Formulator           common.Address
	FormulatorPublicHash common.PublicHash
	PublicHash           common.PublicHash
	Timestamp            uint64
	IsReply              bool
}

// NextRoundVoteAckMessage is a message for a next round vote
type NextRoundVoteAckMessage struct {
	RoundVoteAck *RoundVoteAck
}

// BlockReqMessage is a message for a block request
type BlockReqMessage struct {
	PrevHash             hash.Hash256
	TargetHeight         uint32
	TimeoutCount         uint32
	Formulator           common.Address
	FormulatorPublicHash common.PublicHash
}

// BlockGenMessage is a message for a block generation
type BlockGenMessage struct {
	Block              *types.Block
	GeneratorSignature common.Signature
	IsReply            bool
}

// BlockVote is message for a block vote
type BlockVote struct {
	TargetHeight       uint32
	Header             *types.Header
	GeneratorSignature common.Signature
	ObserverSignature  common.Signature
	IsReply            bool
}

// BlockVoteMessage is a message for a round vote
type BlockVoteMessage struct {
	BlockVote *BlockVote
}

// BlockObSignMessage is a message for a block observer signatures
type BlockObSignMessage struct {
	TargetHeight       uint32
	BlockSign          *types.BlockSign
	ObserverSignatures []common.Signature
}

// BlockGenRequest is a message to request block gen
type BlockGenRequest struct {
	ChainID              uint8
	LastHash             hash.Hash256
	TargetHeight         uint32
	TimeoutCount         uint32
	Formulator           common.Address
	FormulatorPublicHash common.PublicHash
	PublicHash           common.PublicHash
	Timestamp            uint64
}

// BlockGenRequestMessage is a message to request block gen
type BlockGenRequestMessage struct {
	BlockGenRequest *BlockGenRequest
}
