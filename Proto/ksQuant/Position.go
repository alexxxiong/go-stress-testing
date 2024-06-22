// Package ksQuant comment
// This file was generated by ksf2go 1.3.20
// Generated from Position.ksf
package ksQuant

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// Position struct implement
type Position struct {
	Inst_id           string  `json:"inst_id"`
	Side              string  `json:"side"`
	Qty               float64 `json:"qty"`
	Price             float64 `json:"price"`
	High              float64 `json:"high"`
	Low               float64 `json:"low"`
	Pnl               float64 `json:"pnl"`
	Urpnl             float64 `json:"urpnl"`
	Avail_qty         float64 `json:"avail_qty"`
	Freeze_qty        float64 `json:"freeze_qty"`
	Daily_overall_pnl float64 `json:"daily_overall_pnl"`
	Daily_pnl         float64 `json:"daily_pnl"`
	Daily_urpnl       float64 `json:"daily_urpnl"`
	Daily_qty         float64 `json:"daily_qty"`
	Daily_price       float64 `json:"daily_price"`
	His_qty           float64 `json:"his_qty"`
	His_price         float64 `json:"his_price"`
	His_ur_pnl        float64 `json:"his_ur_pnl"`
	Asset             float64 `json:"asset"`
	Close_price       float64 `json:"close_price"`
	Margin            float64 `json:"margin"`
	Dynamic_margin    float64 `json:"dynamic_margin"`
	Commission        float64 `json:"commission"`
	Market_value      float64 `json:"market_value"`
	Pre_price         float64 `json:"pre_price"`
	Last_price        float64 `json:"last_price"`
	Profit            float64 `json:"profit"`
	Date              int32   `json:"date"`
	Account_id        string  `json:"account_id"`
	Strategy_id       string  `json:"strategy_id"`
	Create            int64   `json:"create"`
	Update            int64   `json:"update"`
	M2m               float64 `json:"m2m"`
}

func (st *Position) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *Position) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Inst_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Side, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Qty, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Price, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.High, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Low, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Pnl, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Urpnl, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Avail_qty, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Freeze_qty, 9, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Daily_overall_pnl, 10, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Daily_pnl, 11, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Daily_urpnl, 12, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Daily_qty, 13, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Daily_price, 14, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.His_qty, 15, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.His_price, 16, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.His_ur_pnl, 17, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Asset, 18, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Close_price, 19, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Margin, 20, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Dynamic_margin, 21, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Commission, 22, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Market_value, 23, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Pre_price, 24, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Last_price, 25, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Profit, 26, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Date, 27, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_id, 28, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Strategy_id, 29, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Create, 30, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Update, 31, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.M2m, 32, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *Position) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require Position, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *Position) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Inst_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Side, 1)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Qty, 2)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Price, 3)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.High, 4)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Low, 5)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Pnl, 6)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Urpnl, 7)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Avail_qty, 8)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Freeze_qty, 9)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Daily_overall_pnl, 10)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Daily_pnl, 11)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Daily_urpnl, 12)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Daily_qty, 13)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Daily_price, 14)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.His_qty, 15)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.His_price, 16)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.His_ur_pnl, 17)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Asset, 18)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Close_price, 19)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Margin, 20)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Dynamic_margin, 21)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Commission, 22)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Market_value, 23)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Pre_price, 24)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Last_price, 25)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Profit, 26)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Date, 27)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_id, 28)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Strategy_id, 29)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Create, 30)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Update, 31)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.M2m, 32)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *Position) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// InstrumentPositions struct implement
type InstrumentPositions struct {
	Pos map[string]Position `json:"pos"`
}

func (st *InstrumentPositions) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *InstrumentPositions) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.MAP, 0, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Pos = make(map[string]Position)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {
			var k0 string
			var v0 Position

			err = readBuf.ReadString(&k0, 0, false)
			if err != nil {
				return err
			}

			err = v0.ReadBlock(readBuf, 1, false)
			if err != nil {
				return err
			}

			st.Pos[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *InstrumentPositions) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require InstrumentPositions, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *InstrumentPositions) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.MAP, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Pos)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Pos {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = v1.WriteBlock(buf, 1)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *InstrumentPositions) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}

// DailyPositions struct implement
type DailyPositions struct {
	Date         int32      `json:"date"`
	Positions    []Position `json:"positions"`
	Pos_num      int32      `json:"pos_num"`
	Market_value float64    `json:"market_value"`
	Urpnl        float64    `json:"urpnl"`
	Pnl          float64    `json:"pnl"`
}

func (st *DailyPositions) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *DailyPositions) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Date, 0, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(1, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Positions = make([]Position, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.Positions[i0].ReadBlock(readBuf, 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SimpleList {
			err = fmt.Errorf("not support SimpleList type")
			if err != nil {
				return err
			}

		} else {
			err = fmt.Errorf("require vector, but not")
			if err != nil {
				return err
			}

		}
	}

	err = readBuf.ReadInt32(&st.Pos_num, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Market_value, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Urpnl, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadFloat64(&st.Pnl, 5, false)
	if err != nil {
		return err
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *DailyPositions) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
	var (
		err  error
		have bool
	)
	st.ResetDefault()

	have, err = readBuf.SkipTo(codec.StructBegin, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require DailyPositions, but not exist. tag %d", tag)
		}
		return nil
	}

	err = st.ReadFrom(readBuf)
	if err != nil {
		return err
	}

	err = readBuf.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

// WriteTo encode struct to buffer
func (st *DailyPositions) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Date, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Positions)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Positions {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteInt32(st.Pos_num, 2)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Market_value, 3)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Urpnl, 4)
	if err != nil {
		return err
	}

	err = buf.WriteFloat64(st.Pnl, 5)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *DailyPositions) WriteBlock(buf *codec.Buffer, tag byte) error {
	var err error
	err = buf.WriteHead(codec.StructBegin, tag)
	if err != nil {
		return err
	}

	err = st.WriteTo(buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.StructEnd, 0)
	if err != nil {
		return err
	}
	return nil
}