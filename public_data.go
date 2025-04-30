package okex

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetInstrumentsService
type GetInstrumentsService struct {
	c        *Client
	instType InstType
	uly      *string
	instId   *string
}

// Set instrument type
func (s *GetInstrumentsService) InstrumentType(instType InstType) *GetInstrumentsService {
	s.instType = instType
	return s
}

// Set underlying
func (s *GetInstrumentsService) Underlying(uly string) *GetInstrumentsService {
	s.uly = &uly
	return s
}

// Set instrument id
func (s *GetInstrumentsService) InstrumentId(instId string) *GetInstrumentsService {
	s.instId = &instId
	return s
}

// Do send request
func (s *GetInstrumentsService) Do(ctx context.Context, opts ...RequestOption) (res *GetInstrumentsServiceRespone, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/instruments",
	}

	r.setParam("instType", string(s.instType))

	if s.uly != nil {
		r.setParam("uly", *s.uly)
	}
	if s.instId != nil {
		r.setParam("instId", *s.instId)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetInstrumentsServiceRespone)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if res.Code != codeSucc {
		return nil, fmt.Errorf("code [%v] msg [%s]", res.Code, res.Msg)
	}
	return res, nil
}

// Response to GetInstrumentsService
type GetInstrumentsServiceRespone struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data []*InstrumentDetail `json:"data"`
}

type InstrumentDetail struct {
	InstType     string `json:"instType"`     // 产品类型：永续合约
	InstId       string `json:"instId"`       // 合约ID（格式：基础货币-计价货币-类型）
	Uly          string `json:"uly"`          // 标的资产（如PEPE-USDT指数）
	Category     string `json:"category"`     // 手续费档位（1表示普通档）
	BaseCcy      string `json:"baseCcy"`      // 基础货币（如PEPE）
	QuoteCcy     string `json:"quoteCcy"`     // 计价货币（如USDT）
	SettleCcy    string `json:"settleCcy"`    // 结算货币（通常与计价货币相同）
	CtVal        string `json:"ctVal"`        // 单张合约面值（例如：PEPE-USDT价格是0.000006 CtValue=10000000 那么一张的价值=60USDT)
	CtMult       string `json:"ctMult"`       // 合约乘数（通常为1）
	CtValCcy     string `json:"ctValCcy"`     // 面值计价货币（USD与USDT 1:1锚定）
	OptType      string `json:"optType"`      // 期权类型（永续合约为空）
	Stk          string `json:"stk"`          // 行权价（期权专用，永续合约为空）
	ListTime     string `json:"listTime"`     // 合约上线时间（Unix时间戳，毫秒）
	ExpTime      string `json:"expTime"`      // 到期时间（永续合约为空）
	Lever        string `json:"lever"`        // 当前最大杠杆倍数（如75倍）
	TickSz       string `json:"tickSz"`       // 价格最小变动单位（0.0000001 USDT）
	LotSz        string `json:"lotSz"`        // 数量最小单位（1张）
	MinSz        string `json:"minSz"`        // 最小下单数量（1张）
	CtType       string `json:"ctType"`       // 合约类型（linear=正向，inverse=反向）
	Alias        string `json:"alias"`        // 合约别名（永续合约通常为"swap"）
	State        string `json:"state"`        // 合约状态（live=交易中）
	MaxLmtSz     string `json:"maxLmtSz"`     // 单笔限价单最大张数
	MaxMktSz     string `json:"maxMktSz"`     // 单笔市价单最大张数
	MaxTwapSz    string `json:"maxTwapSz"`    // 单笔TWAP订单最大张数
	MaxIcebergSz string `json:"maxIcebergSz"` // 单笔冰山订单最大张数
	MaxTriggerSz string `json:"maxTriggerSz"` // 单笔条件单最大张数
	MaxStopSz    string `json:"maxStopSz"`    // 单笔止损单最大张数
}

// GetDeliveryExerciseHistoryService
type GetDeliveryExerciseHistoryService struct {
	c        *Client
	instType string
	uly      string
	after    *string
	before   *string
	limit    *string
}

// Set instrument type
func (s *GetDeliveryExerciseHistoryService) InstrumentType(instType string) *GetDeliveryExerciseHistoryService {
	s.instType = instType
	return s
}

// Set underlying
func (s *GetDeliveryExerciseHistoryService) Underlying(uly string) *GetDeliveryExerciseHistoryService {
	s.uly = uly
	return s
}

// Set after
func (s *GetDeliveryExerciseHistoryService) After(after string) *GetDeliveryExerciseHistoryService {
	s.after = &after
	return s
}

// Set before
func (s *GetDeliveryExerciseHistoryService) Before(before string) *GetDeliveryExerciseHistoryService {
	s.before = &before
	return s
}

// Set limit
func (s *GetDeliveryExerciseHistoryService) Limit(limit string) *GetDeliveryExerciseHistoryService {
	s.limit = &limit
	return s
}

// Do send request
func (s *GetDeliveryExerciseHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *GetDeliveryExerciseHistoryServiceResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/api/v5/public/delivery-exercise-history",
	}

	r.setParam("instType", s.instType)
	r.setParam("uly", s.uly)

	if s.after != nil {
		r.setParam("after", *s.after)
	}
	if s.before != nil {
		r.setParam("before", *s.before)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetDeliveryExerciseHistoryServiceResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if res.Code != codeSucc {
		return nil, fmt.Errorf("code [%v] msg [%s]", res.Code, res.Msg)
	}
	return res, nil
}

// Response to GetInstrumentsService
type GetDeliveryExerciseHistoryServiceResponse struct {
	Code string               `json:"code"`
	Msg  string               `json:"msg"`
	Data []*DeliveryExcercise `json:"data"`
}

type DeliveryExcercise struct {
	Ts      string                     `json:"timestamp"`
	Details []*DeliveryExcerciseDetail `json:"details"`
}

type DeliveryExcerciseDetail struct {
	Type   string `json:"type"`
	InstId string `json:"instId"`
	Px     string `json:"px"`
}
