package bitflyer

type ProductCode int

const (
	Btcjpy ProductCode = iota
	Ethjpy
	Fxbtcjpy
	Ethbtc
	bchbtc
)

// 指値 or 成行
type OrderType int

const (
	Limit OrderType = iota
	Market
)

// 買いか売りか
type Side int

const (
	Buy Side = iota
	Sell
)

// 執行数量条件
type TimeInForce int

const (
	Gtc TimeInForce = iota
	Ioc
	Fok
)

func (orderType OrderType) String() string {
	switch orderType {
	case Limit:
		return "LIMIT"
	case Market:
		return "MARKET"
	default:
		return "LIMIT"
	}
}

func (side Side) String() string {
	switch side {
	case Buy:
		return "BUY"
	case Sell:
		return "SELL"
	default:
		return "BUY"
	}
}

func (tif TimeInForce) String() string {
	switch tif {
	case Gtc:
		return "GTC"
	case Ioc:
		return "IOC"
	case Fok:
		return "FOK"
	default:
		return "GTC"
	}
}

func (code ProductCode) String() string {
	switch code {
	case Btcjpy:
		return "BTC_JPY"
	case Ethjpy:
		return "ETH_JPY"
	case Fxbtcjpy:
		return "FX_BTC_JPY"
	case Ethbtc:
		return "ETH_BTC"
	case bchbtc:
		return "BCH_BTC"
	default:
		return "BTC_JPY"
	}
}
