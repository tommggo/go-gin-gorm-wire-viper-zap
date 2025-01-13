package model

import "time"

type Signal struct {
	ID           uint64
	AssetName    string
	SignalCode   string
	ModelVersion string
	Processed    bool
	SignalSource string
	Priority     string
	Kline        int
	Side         string
	AssetType    string
	PositionSide string
	Remark       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName 指定表名
func (Signal) TableName() string {
	return "signal"
}
