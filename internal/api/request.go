package api

// CreateSignalReq 创建信号请求
type CreateSignalReq struct {
	AssetName    string `json:"asset_name" binding:"required"`                     // 资产名称
	SignalCode   string `json:"signal_code" binding:"required"`                    // 信号代码
	ModelVersion string `json:"model_version" binding:"required"`                  // 模型版本
	SignalSource string `json:"signal_source" binding:"required"`                  // 信号来源
	Priority     string `json:"priority" binding:"required,oneof=high medium low"` // 优先级
	Kline        int    `json:"kline" binding:"required,min=1"`                    // K线周期
	Side         string `json:"side" binding:"required,oneof=buy sell"`            // 方向
	AssetType    string `json:"asset_type" binding:"required"`                     // 资产类型
	PositionSide string `json:"position_side" binding:"required"`                  // 持仓方向
	Remark       string `json:"remark"`                                            // 备注
}

// 其他请求模型...
