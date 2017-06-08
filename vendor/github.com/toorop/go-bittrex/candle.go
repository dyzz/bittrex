package bittrex

type Candle struct {
	TimeStamp string  `json:"T"`
	Open      float64 `json:"O"`
	Close     float64 `json:"C"`
	High      float64 `json:"H"`
	Low       float64 `json:"L"`
	Volume    float64 `json:"V"`
}

type Candle_Resp struct {
	Success bool     `json:"success"`
	Message string   `json:"message"`
	Result  []Candle `json:"result"`
}

type NewCandles struct {
	Ticks []Candle `json:"ticks"`
}
