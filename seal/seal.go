package seal

import (
	"time"

	"github.com/celer-network/goutils/log"
)

func NewTransactorLog(sender string) *TransactorLog {
	msgtypes := make(map[string]uint32)
	now := time.Now().UnixNano()
	return &TransactorLog{
		MsgType:         msgtypes,
		Sender:          sender,
		ExecutionTimeMs: (float64)(now),
	}
}

func CommitTransactorLog(entry *TransactorLog) {
	now := time.Now().UnixNano()
	entry.ExecutionTimeMs = ((float64)(now) - entry.ExecutionTimeMs) / 1000000
	if len(entry.Error) > 0 {
		log.Errorln("TransactorLog:", entry)
	} else {
		log.Infoln("TransactorLog:", entry)
	}
}

func NewMsgLog() *MsgLog {
	now := time.Now().UnixNano()
	return &MsgLog{
		ChanInfo:        &ChannelInfo{},
		Penalty:         &Penalty{},
		Change:          &Change{},
		ExecutionTimeMs: (float64)(now),
	}
}

func CommitMsgLog(entry *MsgLog) {
	now := time.Now().UnixNano()
	entry.ExecutionTimeMs = ((float64)(now) - entry.ExecutionTimeMs) / 1000000
	if len(entry.ChanInfo.ChanId) == 0 {
		entry.ChanInfo = nil
	}
	if entry.Penalty.Nonce == 0 && len(entry.Penalty.Validator) == 0 {
		entry.Penalty = nil
	}

	if entry.Change.Id == 0 {
		entry.Change = nil
	}

	if len(entry.Error) > 0 {
		log.Errorln("MsgLog:", entry)
	} else if len(entry.Warn) > 0 {
		log.Warnln("MsgLog:", entry)
	} else {
		log.Infoln("MsgLog:", entry)
	}
}
