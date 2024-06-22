// Package ksTrade comment
// This file was generated by ksf2go 1.3.20
// Generated from HostingObj.ksf
package ksTrade

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// HostingReq struct implement
type HostingReq struct {
	Request_id int32             `json:"request_id"`
	Func       string            `json:"func"`
	Buf        []int8            `json:"buf"`
	Extends    map[string]string `json:"extends"`
}

func (st *HostingReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *HostingReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Func, 1, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(2, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Buf = make([]int8, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadInt8(&st.Buf[i0], 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SimpleList {

			_, err = readBuf.SkipTo(codec.BYTE, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadSliceInt8(&st.Buf, length, true)
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

	have, err = readBuf.SkipTo(codec.MAP, 3, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Extends = make(map[string]string)
		for i1, e1 := int32(0), length; i1 < e1; i1++ {
			var k1 string
			var v1 string

			err = readBuf.ReadString(&k1, 0, false)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&v1, 1, false)
			if err != nil {
				return err
			}

			st.Extends[k1] = v1
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *HostingReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require HostingReq, but not exist. tag %d", tag)
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
func (st *HostingReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Func, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.SimpleList, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Buf)), 0)
	if err != nil {
		return err
	}

	err = buf.WriteSliceInt8(st.Buf)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extends)), 0)
	if err != nil {
		return err
	}

	for k2, v2 := range st.Extends {

		err = buf.WriteString(k2, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v2, 1)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *HostingReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// HostingRsp struct implement
type HostingRsp struct {
	Request_id int32             `json:"request_id"`
	Func       string            `json:"func"`
	Param1     []int8            `json:"param1"`
	Param2     []int8            `json:"param2"`
	Is_last    bool              `json:"is_last"`
	Extends    map[string]string `json:"extends"`
}

func (st *HostingRsp) ResetDefault() {
	st.Is_last = true
}

// ReadFrom reads  from readBuf and put into struct.
func (st *HostingRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Request_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Func, 1, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(2, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Param1 = make([]int8, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadInt8(&st.Param1[i0], 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SimpleList {

			_, err = readBuf.SkipTo(codec.BYTE, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadSliceInt8(&st.Param1, length, true)
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

	have, ty, err = readBuf.SkipToNoCheck(3, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Param2 = make([]int8, length)
			for i1, e1 := int32(0), length; i1 < e1; i1++ {

				err = readBuf.ReadInt8(&st.Param2[i1], 0, false)
				if err != nil {
					return err
				}

			}
		} else if ty == codec.SimpleList {

			_, err = readBuf.SkipTo(codec.BYTE, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			err = readBuf.ReadSliceInt8(&st.Param2, length, true)
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

	err = readBuf.ReadBool(&st.Is_last, 4, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 10, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Extends = make(map[string]string)
		for i2, e2 := int32(0), length; i2 < e2; i2++ {
			var k2 string
			var v2 string

			err = readBuf.ReadString(&k2, 0, false)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&v2, 1, false)
			if err != nil {
				return err
			}

			st.Extends[k2] = v2
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *HostingRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require HostingRsp, but not exist. tag %d", tag)
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
func (st *HostingRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Request_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Func, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.SimpleList, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Param1)), 0)
	if err != nil {
		return err
	}

	err = buf.WriteSliceInt8(st.Param1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.SimpleList, 3)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.BYTE, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Param2)), 0)
	if err != nil {
		return err
	}

	err = buf.WriteSliceInt8(st.Param2)
	if err != nil {
		return err
	}

	err = buf.WriteBool(st.Is_last, 4)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 10)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extends)), 0)
	if err != nil {
		return err
	}

	for k3, v3 := range st.Extends {

		err = buf.WriteString(k3, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v3, 1)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *HostingRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// RegistReq struct implement
type RegistReq struct {
	Private_topic int32 `json:"private_topic"`
	Public_topic  int32 `json:"public_topic"`
}

func (st *RegistReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *RegistReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadInt32(&st.Private_topic, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Public_topic, 1, false)
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
func (st *RegistReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require RegistReq, but not exist. tag %d", tag)
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
func (st *RegistReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteInt32(st.Private_topic, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Public_topic, 1)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *RegistReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// RegistRsp struct implement
type RegistRsp struct {
	Msg     string            `json:"msg"`
	Extends map[string]string `json:"extends"`
}

func (st *RegistRsp) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *RegistRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Msg, 0, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 1, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.Extends = make(map[string]string)
		for i0, e0 := int32(0), length; i0 < e0; i0++ {
			var k0 string
			var v0 string

			err = readBuf.ReadString(&k0, 0, false)
			if err != nil {
				return err
			}

			err = readBuf.ReadString(&v0, 1, false)
			if err != nil {
				return err
			}

			st.Extends[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *RegistRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require RegistRsp, but not exist. tag %d", tag)
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
func (st *RegistRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Msg, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extends)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Extends {

		err = buf.WriteString(k1, 0)
		if err != nil {
			return err
		}

		err = buf.WriteString(v1, 1)
		if err != nil {
			return err
		}

	}

	return err
}

// WriteBlock encode struct
func (st *RegistRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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