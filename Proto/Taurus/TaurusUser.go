// Package Taurus comment
// This file was generated by ksf2go 1.3.21
// Generated from TaurusUser.ksf
package Taurus

import (
	"fmt"

	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8

// const as define in ksf file
const (
	ENCRYPT_METHOD_SM4 string = "SM4"
)

// TradeLoginInfo struct implement
type TradeLoginInfo struct {
	SecureInfo  []SecureInfo      `json:"secureInfo"`
	ClientName  string            `json:"clientName"`
	ClientId    string            `json:"clientId"`
	Pass        string            `json:"pass"`
	Date        int32             `json:"date"`
	Quick_token string            `json:"quick_token"`
	Mac         string            `json:"mac"`
	ExtInfo     map[string]string `json:"extInfo"`
}

func (st *TradeLoginInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *TradeLoginInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

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

			st.SecureInfo = make([]SecureInfo, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = st.SecureInfo[i0].ReadBlock(readBuf, 0, false)
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

	err = readBuf.ReadString(&st.ClientName, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.ClientId, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Pass, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Date, 5, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Quick_token, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Mac, 7, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 8, false)
	if err != nil {
		return err
	}

	if have {
		err = readBuf.ReadInt32(&length, 0, true)
		if err != nil {
			return err
		}

		st.ExtInfo = make(map[string]string)
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

			st.ExtInfo[k1] = v1
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *TradeLoginInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require TradeLoginInfo, but not exist. tag %d", tag)
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
func (st *TradeLoginInfo) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteHead(codec.LIST, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.SecureInfo)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.SecureInfo {

		err = v.WriteBlock(buf, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.ClientName, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.ClientId, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Pass, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Date, 5)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Quick_token, 6)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Mac, 7)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 8)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.ExtInfo)), 0)
	if err != nil {
		return err
	}

	for k2, v2 := range st.ExtInfo {

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
func (st *TradeLoginInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// UserAccLogin struct implement
type UserAccLogin struct {
	User_id           string            `json:"user_id"`
	Account_id        string            `json:"account_id"`
	Login_status      string            `json:"login_status"`
	Connection_status string            `json:"connection_status"`
	Last_login        int64             `json:"last_login"`
	Login_expire_time int64             `json:"login_expire_time"`
	Extends           map[string]string `json:"extends"`
}

func (st *UserAccLogin) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *UserAccLogin) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.User_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_id, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Login_status, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Connection_status, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Last_login, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Login_expire_time, 5, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 6, false)
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
func (st *UserAccLogin) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require UserAccLogin, but not exist. tag %d", tag)
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
func (st *UserAccLogin) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.User_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_id, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Login_status, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Connection_status, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Last_login, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Login_expire_time, 5)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 6)
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
func (st *UserAccLogin) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// UserAccInfo struct implement
type UserAccInfo struct {
	User_id      string            `json:"user_id"`
	Account_id   string            `json:"account_id"`
	Account_name string            `json:"account_name"`
	Account_pwd  string            `json:"account_pwd"`
	Effective    int32             `json:"effective"`
	Extends      map[string]string `json:"extends"`
}

func (st *UserAccInfo) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *UserAccInfo) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.User_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_id, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_name, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Account_pwd, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Effective, 4, false)
	if err != nil {
		return err
	}

	have, err = readBuf.SkipTo(codec.MAP, 5, false)
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
func (st *UserAccInfo) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require UserAccInfo, but not exist. tag %d", tag)
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
func (st *UserAccInfo) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.User_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_id, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_name, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Account_pwd, 3)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Effective, 4)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 5)
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
func (st *UserAccInfo) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// RbacPermission struct implement
type RbacPermission struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Parent_id   string   `json:"parent_id"`
	Effective   string   `json:"effective"`
	Actions     []string `json:"actions"`
	Objects     []string `json:"objects"`
	Description string   `json:"description"`
	Sequence    int32    `json:"sequence"`
	Create_time string   `json:"create_time"`
	Update_time string   `json:"update_time"`
}

func (st *RbacPermission) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *RbacPermission) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Name, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Parent_id, 2, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Effective, 3, false)
	if err != nil {
		return err
	}

	have, ty, err = readBuf.SkipToNoCheck(4, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Actions = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Actions[i0], 0, false)
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

	have, ty, err = readBuf.SkipToNoCheck(5, false)
	if err != nil {
		return err
	}

	if have {
		if ty == codec.LIST {
			err = readBuf.ReadInt32(&length, 0, true)
			if err != nil {
				return err
			}

			st.Objects = make([]string, length)
			for i1, e1 := int32(0), length; i1 < e1; i1++ {

				err = readBuf.ReadString(&st.Objects[i1], 0, false)
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

	err = readBuf.ReadString(&st.Description, 6, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt32(&st.Sequence, 7, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Create_time, 8, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Update_time, 9, false)
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
func (st *RbacPermission) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require RbacPermission, but not exist. tag %d", tag)
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
func (st *RbacPermission) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Name, 1)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Parent_id, 2)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Effective, 3)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 4)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Actions)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Actions {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteHead(codec.LIST, 5)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Objects)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Objects {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Description, 6)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(st.Sequence, 7)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Create_time, 8)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Update_time, 9)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *RbacPermission) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// RbacRole struct implement
type RbacRole struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	Description string   `json:"description"`
	Create_time string   `json:"create_time"`
	Update_time string   `json:"update_time"`
}

func (st *RbacRole) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *RbacRole) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.Id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Name, 1, false)
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

			st.Permissions = make([]string, length)
			for i0, e0 := int32(0), length; i0 < e0; i0++ {

				err = readBuf.ReadString(&st.Permissions[i0], 0, false)
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

	err = readBuf.ReadString(&st.Description, 3, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Create_time, 4, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Update_time, 5, false)
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
func (st *RbacRole) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require RbacRole, but not exist. tag %d", tag)
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
func (st *RbacRole) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.Id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Name, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.LIST, 2)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Permissions)), 0)
	if err != nil {
		return err
	}

	for _, v := range st.Permissions {

		err = buf.WriteString(v, 0)
		if err != nil {
			return err
		}

	}

	err = buf.WriteString(st.Description, 3)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Create_time, 4)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Update_time, 5)
	if err != nil {
		return err
	}

	return err
}

// WriteBlock encode struct
func (st *RbacRole) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// GenLicenseReq struct implement
type GenLicenseReq struct {
	User_id       string            `json:"user_id"`
	Pwd           string            `json:"pwd"`
	Duration      int64             `json:"duration"`
	Extend_params map[string]string `json:"extend_params"`
}

func (st *GenLicenseReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *GenLicenseReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.User_id, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadString(&st.Pwd, 1, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Duration, 2, false)
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

		st.Extend_params = make(map[string]string)
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

			st.Extend_params[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *GenLicenseReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require GenLicenseReq, but not exist. tag %d", tag)
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
func (st *GenLicenseReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.User_id, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(st.Pwd, 1)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Duration, 2)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 10)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extend_params)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Extend_params {

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
func (st *GenLicenseReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// GenLicenseRsp struct implement
type GenLicenseRsp struct {
	License       string            `json:"license"`
	Expire_time   int64             `json:"expire_time"`
	Extend_params map[string]string `json:"extend_params"`
}

func (st *GenLicenseRsp) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *GenLicenseRsp) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.License, 0, false)
	if err != nil {
		return err
	}

	err = readBuf.ReadInt64(&st.Expire_time, 1, false)
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

		st.Extend_params = make(map[string]string)
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

			st.Extend_params[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *GenLicenseRsp) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require GenLicenseRsp, but not exist. tag %d", tag)
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
func (st *GenLicenseRsp) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.License, 0)
	if err != nil {
		return err
	}

	err = buf.WriteInt64(st.Expire_time, 1)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 10)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extend_params)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Extend_params {

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
func (st *GenLicenseRsp) WriteBlock(buf *codec.Buffer, tag byte) error {
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

// CheckLicenseLegalReq struct implement
type CheckLicenseLegalReq struct {
	License       string            `json:"license"`
	Extend_params map[string]string `json:"extend_params"`
}

func (st *CheckLicenseLegalReq) ResetDefault() {
}

// ReadFrom reads  from readBuf and put into struct.
func (st *CheckLicenseLegalReq) ReadFrom(readBuf *codec.Reader) error {
	var (
		err    error
		length int32
		have   bool
		ty     byte
	)
	st.ResetDefault()

	err = readBuf.ReadString(&st.License, 0, false)
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

		st.Extend_params = make(map[string]string)
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

			st.Extend_params[k0] = v0
		}
	}

	_ = err
	_ = length
	_ = have
	_ = ty
	return nil
}

// ReadBlock reads struct from the given tag , require or optional.
func (st *CheckLicenseLegalReq) ReadBlock(readBuf *codec.Reader, tag byte, require bool) error {
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
			return fmt.Errorf("require CheckLicenseLegalReq, but not exist. tag %d", tag)
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
func (st *CheckLicenseLegalReq) WriteTo(buf *codec.Buffer) (err error) {

	err = buf.WriteString(st.License, 0)
	if err != nil {
		return err
	}

	err = buf.WriteHead(codec.MAP, 10)
	if err != nil {
		return err
	}

	err = buf.WriteInt32(int32(len(st.Extend_params)), 0)
	if err != nil {
		return err
	}

	for k1, v1 := range st.Extend_params {

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
func (st *CheckLicenseLegalReq) WriteBlock(buf *codec.Buffer, tag byte) error {
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