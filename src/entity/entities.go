package entity

import "time"

type USDToCNH struct {
	ID      int64
	Price   float32
	Change  string
	Percent string
}

func NewUSDToCNH(price float32, change string, percent string) USDToCNH {
	return USDToCNH{
		ID:      time.Now().UnixMilli(),
		Price:   price,
		Change:  change,
		Percent: percent,
	}
}

func (t USDToCNH) TableName() string {
	return "t_usd_cnh"
}

type BitcoinToUSD struct {
	ID      int64
	Price   float32
	Change  string
	Percent string
}

func (t BitcoinToUSD) TableName() string {
	return "t_bitcoin_usd"
}

func NewBitcoinToUSD(price float32, change string, percent string) BitcoinToUSD {
	return BitcoinToUSD{
		ID:      time.Now().UnixMilli(),
		Price:   price,
		Change:  change,
		Percent: percent,
	}
}

type Shangzheng struct {
	ID      int64
	Price   float32
	Change  string
	Percent string
}

func (t Shangzheng) TableName() string {
	return "t_shangzheng"
}

func NewShangzheng(price float32, change string, percent string) Shangzheng {
	return Shangzheng{
		ID:      time.Now().UnixMilli(),
		Price:   price,
		Change:  change,
		Percent: percent,
	}
}
