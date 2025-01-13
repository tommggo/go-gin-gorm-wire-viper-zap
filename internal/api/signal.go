package api

import (
	"go-gin-gorm-wire-viper-zap/internal/errors"
	"go-gin-gorm-wire-viper-zap/internal/model"
	"go-gin-gorm-wire-viper-zap/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SignalAPI 处理信号相关的 HTTP 请求
type SignalAPI struct {
	signalService service.SignalService
}

// NewSignalAPI 创建 API 处理器
func NewSignalAPI(signalService service.SignalService) *SignalAPI {
	return &SignalAPI{signalService: signalService}
}

// Create 创建信号
func (s *SignalAPI) Create(c *gin.Context) {
	var req CreateSignalReq
	if err := c.ShouldBindJSON(&req); err != nil {
		Error(c, errors.Wrap(errors.ErrInvalidParam, err))
		return
	}

	signal := &model.Signal{
		AssetName:    req.AssetName,
		SignalCode:   req.SignalCode,
		ModelVersion: req.ModelVersion,
		SignalSource: req.SignalSource,
		Priority:     req.Priority,
		Kline:        req.Kline,
		Side:         req.Side,
		AssetType:    req.AssetType,
		PositionSide: req.PositionSide,
		Remark:       req.Remark,
	}

	if err := s.signalService.Create(c.Request.Context(), signal); err != nil {
		Error(c, err)
		return
	}

	Success(c, signal)
}

// Get 获取信号
func (s *SignalAPI) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, errors.Wrap(errors.ErrInvalidParam, err))
		return
	}

	signal, err := s.signalService.Get(c.Request.Context(), id)
	if err != nil {
		Error(c, err)
		return
	}

	Success(c, signal)
}

// Process 处理信号
func (s *SignalAPI) Process(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		Error(c, errors.Wrap(errors.ErrInvalidParam, err))
		return
	}

	if err := s.signalService.ProcessSignal(c.Request.Context(), id); err != nil {
		Error(c, err)
		return
	}

	SuccessMessage(c, "signal processed successfully")
}
