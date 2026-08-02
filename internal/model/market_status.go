package model

import "time"

type MarketStatus struct {
	MsID                int64
	MsStatus            int
	MsCurrentRound      int
	MsCurrentMonth      int
	MsCurrentSeason     int
	MsShutdownDay       int
	MsShutdownMonth     int
	MsShutdownYear      int
	MsShutdownHour      int
	MsShutdownMinute    int
	MsShutdownTimestamp time.Time
	MsPostRound         bool
	MsNewMonth          bool
	MsRunning           bool
}
