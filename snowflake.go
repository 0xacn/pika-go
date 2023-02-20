package main\

import (
  "strconv"
  "time"
)

type Snowflake struct {
  epoch uint64
  nodeID, seq uint32
  lastSequenceExhaustion uint64
}

type DecodedSnowflake struct {
  id uint64
  timestamp uint64
  nodeID uint32
  seq uint32
  epoch uint64
}

func NewSnowflakeWithNodeID(epoch uint64, nodeID uint32) *Snowflake {
  return &Snowflake {
    epoch: epoch,
    nodeID: nodeID,
    seq: 0,
    lastSequenceExhaustion: 0,
  }
}

func (sf *Snowflake) gen() string {
  return sf.genWithTs(uint64(time.Now().Unix()))
}
 
func (sf *Snowflake) genWithTs(timestamp uint64) string {
  if sf.seq >= 4095 && timestamp == sf.lastSequenceExhaustion {
    for time.Now().Unix()-int64(timestamp) < 1 {
    }
  }
}
