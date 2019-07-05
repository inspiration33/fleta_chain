package pof

import (
	"github.com/fletaio/fleta/common"
	"github.com/fletaio/fleta/core/types"
)

// consts
const (
	EmptyState        = iota
	RoundVoteState    = iota
	RoundVoteAckState = iota
	BlockWaitState    = iota
	BlockVoteState    = iota
)

// VoteRound is data for the voting round
type VoteRound struct {
	RoundState                 int
	VoteTargetHeight           uint32
	VoteFailCount              int
	RemainBlocks               uint32
	RoundVoteMessageMap        map[common.PublicHash]*RoundVoteMessage
	MinVotePublicHash          common.PublicHash
	RoundVoteAckMessageMap     map[common.PublicHash]*RoundVoteAckMessage
	MinRoundVoteAck            *RoundVoteAck
	BlockRound                 *BlockRound
	RoundVoteWaitMap           map[common.PublicHash]*RoundVoteMessage
	RoundVoteAckMessageWaitMap map[common.PublicHash]*RoundVoteAckMessage
}

// NewVoteRound returns a VoteRound
func NewVoteRound(TargetHeight uint32, MaxBlocksPerFormulator uint32) *VoteRound {
	vr := &VoteRound{
		RoundState:                 RoundVoteState,
		VoteTargetHeight:           TargetHeight,
		RemainBlocks:               MaxBlocksPerFormulator,
		RoundVoteMessageMap:        map[common.PublicHash]*RoundVoteMessage{},
		RoundVoteAckMessageMap:     map[common.PublicHash]*RoundVoteAckMessage{},
		RoundVoteWaitMap:           map[common.PublicHash]*RoundVoteMessage{},
		RoundVoteAckMessageWaitMap: map[common.PublicHash]*RoundVoteAckMessage{},
		BlockRound:                 NewBlockRound(TargetHeight),
	}
	return vr
}

// BlockRound is data for the block round
type BlockRound struct {
	TargetHeight            uint32
	BlockVoteMap            map[common.PublicHash]*BlockVote
	BlockGenMessage         *BlockGenMessage
	Context                 *types.Context
	BlockVoteMessageWaitMap map[common.PublicHash]*BlockVoteMessage
	BlockGenMessageWait     *BlockGenMessage
}

// NewBlockRound returns a VoteRound
func NewBlockRound(TargetHeight uint32) *BlockRound {
	vr := &BlockRound{
		TargetHeight:            TargetHeight,
		BlockVoteMap:            map[common.PublicHash]*BlockVote{},
		BlockVoteMessageWaitMap: map[common.PublicHash]*BlockVoteMessage{},
	}
	return vr
}

type voteSortItem struct {
	PublicHash common.PublicHash
	Priority   uint64
}

type voteSorter []*voteSortItem

func (s voteSorter) Len() int {
	return len(s)
}

func (s voteSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s voteSorter) Less(i, j int) bool {
	a := s[i]
	b := s[j]
	if a.Priority == b.Priority {
		return a.PublicHash.Less(b.PublicHash)
	} else {
		return a.Priority < b.Priority
	}
}
