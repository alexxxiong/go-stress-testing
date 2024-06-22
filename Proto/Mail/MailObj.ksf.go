// Package Mail comment
// This file was generated by ksf2go 1.3.20
// Generated from Mail.ksf
package Mail

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go.k8sf.cloud/go/KsfGo/ksf"
	m "go.k8sf.cloud/go/KsfGo/ksf/model"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/codec"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/kup"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/res/basef"
	"go.k8sf.cloud/go/KsfGo/ksf/protocol/res/requestf"
	"go.k8sf.cloud/go/KsfGo/ksf/util/current"
	"go.k8sf.cloud/go/KsfGo/ksf/util/tools"
	"go.k8sf.cloud/go/KsfGo/ksf/util/trace"
	"net"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = fmt.Errorf
	_ = codec.FromInt8
	_ = unsafe.Pointer(nil)
	_ = bytes.ErrTooLarge
	_ = net.UDPConn{}
)

// MailObj struct
type MailObj struct {
	servant m.Servant
}

// SendMail is the proxy function for the method defined in the ksf file, with the context
func (obj *MailObj) SendMail(req *Req, err_msg *string, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(*err_msg, 2)
	if err != nil {
		return ret, err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}
	ksfResp := new(requestf.ResponsePacket)
	ksfCtx := context.Background()

	err = obj.servant.KsfInvoke(ksfCtx, 0, "SendMail", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = readBuf.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = readBuf.ReadString(&(*err_msg), 2, true)
	if err != nil {
		return ret, err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range ksfResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// SendMailWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *MailObj) SendMailWithContext(ksfCtx context.Context, req *Req, err_msg *string, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(*err_msg, 2)
	if err != nil {
		return ret, err
	}

	traceData, ok := current.GetTraceData(ksfCtx)
	if ok && traceData.TraceCall {
		traceData.NewSpan()
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCS, uint(buf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value["req"] = req
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCS), trace.TraceAnnotationCS, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "SendMail", 0, traceParam, "")
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	ksfResp := new(requestf.ResponsePacket)
	err = obj.servant.KsfInvoke(ksfCtx, 0, "SendMail", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
	err = readBuf.ReadInt32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = readBuf.ReadString(&(*err_msg), 2, true)
	if err != nil {
		return ret, err
	}

	if ok && traceData.TraceCall {
		var traceParam string
		traceParamFlag := traceData.NeedTraceParam(trace.EstCR, uint(readBuf.Len()))
		if traceParamFlag == trace.EnpNormal {
			value := map[string]interface{}{}
			value[""] = ret
			value["err_msg"] = *err_msg
			p, _ := json.Marshal(value)
			traceParam = string(p)
		} else if traceParamFlag == trace.EnpOverMaxLen {
			traceParam = "{\"trace_param_over_max_len\":true}"
		}
		ksf.Trace(traceData.GetTraceKey(trace.EstCR), trace.TraceAnnotationCR, ksf.GetClientConfig().ModuleName, obj.servant.Name(), "SendMail", int(ksfResp.IRet), traceParam, "")
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range ksfResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// SendMailOneWayWithContext is the proxy function for the method defined in the ksf file, with the context
func (obj *MailObj) SendMailOneWayWithContext(ksfCtx context.Context, req *Req, err_msg *string, opts ...map[string]string) (ret int32, err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	buf := codec.NewBuffer()
	err = req.WriteBlock(buf, 1)
	if err != nil {
		return ret, err
	}

	err = buf.WriteString(*err_msg, 2)
	if err != nil {
		return ret, err
	}

	var statusMap map[string]string
	var contextMap map[string]string
	if len(opts) == 1 {
		contextMap = opts[0]
	} else if len(opts) == 2 {
		contextMap = opts[0]
		statusMap = opts[1]
	}

	ksfResp := new(requestf.ResponsePacket)
	err = obj.servant.KsfInvoke(ksfCtx, 1, "SendMail", buf.ToBytes(), statusMap, contextMap, ksfResp)
	if err != nil {
		return ret, err
	}

	if len(opts) == 1 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
	} else if len(opts) == 2 {
		for k := range contextMap {
			delete(contextMap, k)
		}
		for k, v := range ksfResp.Context {
			contextMap[k] = v
		}
		for k := range statusMap {
			delete(statusMap, k)
		}
		for k, v := range ksfResp.Status {
			statusMap[k] = v
		}
	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

// SetServant sets servant for the service.
func (obj *MailObj) SetServant(servant m.Servant) {
	obj.servant = servant
}

// GetServant gets servant for the service.
func (obj *MailObj) GetServant() (servant *m.Servant) {
	return &obj.servant
}

// SetOnConnectCallback
func (obj *MailObj) SetOnConnectCallback(callback func(string)) {
	obj.servant.SetOnConnectCallback(callback)
}

// SetOnCloseCallback
func (obj *MailObj) SetOnCloseCallback(callback func(string)) {
	obj.servant.SetOnCloseCallback(callback)
}

// SetKsfCallback
func (obj *MailObj) SetKsfCallback(callback MailObjKsfCallback) {
	var push MailObjPushCallback
	push.Cb = callback
	obj.servant.SetKsfCallback(&push)
}

// SetPushCallback
func (obj *MailObj) SetPushCallback(callback func([]byte)) {
	obj.servant.SetPushCallback(callback)
}
func (obj *MailObj) req2Byte(rsp *requestf.ResponsePacket) []byte {
	req := requestf.RequestPacket{}
	req.IVersion = rsp.IVersion
	req.IRequestId = rsp.IRequestId
	req.IMessageType = rsp.IMessageType
	req.CPacketType = rsp.CPacketType
	req.Context = rsp.Context
	req.Status = rsp.Status
	req.SBuffer = rsp.SBuffer

	os := codec.NewBuffer()
	req.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	length := sbuf.Len()
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(length))
	return sbuf.Bytes()
}

func (obj *MailObj) rsp2Byte(rsp *requestf.ResponsePacket) []byte {
	if rsp.IVersion == basef.KUPVERSION {
		return obj.req2Byte(rsp)
	}
	os := codec.NewBuffer()
	rsp.WriteTo(os)
	bs := os.ToBytes()
	sbuf := bytes.NewBuffer(nil)
	sbuf.Write(make([]byte, 4))
	sbuf.Write(bs)
	length := sbuf.Len()
	binary.BigEndian.PutUint32(sbuf.Bytes(), uint32(length))
	return sbuf.Bytes()
}

// KsfPing
func (obj *MailObj) KsfPing() {
	ksfCtx := context.Background()
	obj.servant.KsfPing(ksfCtx)
}

// KsfSetTimeout sets the timeout for the servant which is in ms.
func (obj *MailObj) KsfSetTimeout(timeout int) {
	obj.servant.KsfSetTimeout(timeout)
}

// KsfSetVersion default as KSFVERSION,you can set JSONVERSION.
func (obj *MailObj) KsfSetVersion(version int16) {
	obj.servant.KsfSetVersion(version)
}

// KsfSetProtocol sets the protocol for the servant.
func (obj *MailObj) KsfSetProtocol(p m.Protocol) {
	obj.servant.KsfSetProtocol(p)
}

// AddServant adds servant  for the service.
func (obj *MailObj) AddServant(imp MailObjServant, servantObj string) {
	ksf.AddServant(obj, imp, servantObj)
}

// AddServantWithContext adds servant  for the service with context.
func (obj *MailObj) AddServantWithContext(imp MailObjServantWithContext, servantObj string) {
	ksf.AddServantWithContext(obj, imp, servantObj)
}

type MailObjServant interface {
	SendMail(req *Req, err_msg *string) (ret int32, err error)
}
type MailObjServantWithContext interface {
	SendMail(ksfCtx context.Context, req *Req, err_msg *string) (ret int32, err error)

	DoClose(ctx context.Context)
}

// Dispatch is used to call the server side implement for the method defined in the ksf file. withContext shows using context or not.
func (obj *MailObj) Dispatch(ksfCtx context.Context, val interface{}, ksfReq *requestf.RequestPacket, ksfResp *requestf.ResponsePacket, withContext bool) (err error) {
	var (
		length int32
		have   bool
		ty     byte
	)
	readBuf := codec.NewReader(tools.Int8ToByte(ksfReq.SBuffer))
	buf := codec.NewBuffer()
	switch ksfReq.SFuncName {
	case "SendMail":
		var req Req
		var err_msg string

		if ksfReq.IVersion == basef.KSFVERSION {

			err = req.ReadBlock(readBuf, 1, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			reqKup := kup.NewUniAttribute()
			reqKup.Decode(readBuf)

			var kupBuffer []byte

			reqKup.GetBuffer("req", &kupBuffer)
			readBuf.Reset(kupBuffer)
			err = req.ReadBlock(readBuf, 0, true)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.JSONVERSION {
			var jsonData map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(readBuf.ToBytes()))
			decoder.UseNumber()
			err = decoder.Decode(&jsonData)
			if err != nil {
				return fmt.Errorf("decode reqpacket failed, error: %+v", err)
			}
			{
				jsonStr, _ := json.Marshal(jsonData["req"])
				req.ResetDefault()
				if err = json.Unmarshal(jsonStr, &req); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("decode reqpacket fail, error version: %d", ksfReq.IVersion)
			return err
		}

		traceData, ok := current.GetTraceData(ksfCtx)
		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSR, uint(readBuf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value["req"] = req
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSR), trace.TraceAnnotationSR, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "SendMail", 0, traceParam, "")
		}

		var funRet int32
		if !withContext {
			imp := val.(MailObjServant)
			funRet, err = imp.SendMail(&req, &err_msg)
		} else {
			imp := val.(MailObjServantWithContext)
			funRet, err = imp.SendMail(ksfCtx, &req, &err_msg)
		}

		if err != nil {
			return err
		}

		if ksfReq.IVersion == basef.KSFVERSION {
			buf.Reset()

			err = buf.WriteInt32(funRet, 0)
			if err != nil {
				return err
			}

			err = buf.WriteString(err_msg, 2)
			if err != nil {
				return err
			}

		} else if ksfReq.IVersion == basef.KUPVERSION {
			rspKup := kup.NewUniAttribute()

			err = buf.WriteInt32(funRet, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("", buf.ToBytes())
			rspKup.PutBuffer("ksf_ret", buf.ToBytes())

			buf.Reset()
			err = buf.WriteString(err_msg, 0)
			if err != nil {
				return err
			}

			rspKup.PutBuffer("err_msg", buf.ToBytes())

			buf.Reset()
			err = rspKup.Encode(buf)
			if err != nil {
				return err
			}
		} else if ksfReq.IVersion == basef.JSONVERSION {
			rspJson := map[string]interface{}{}
			rspJson["ksf_ret"] = funRet
			rspJson["err_msg"] = err_msg

			var rspByte []byte
			if rspByte, err = json.Marshal(rspJson); err != nil {
				return err
			}

			buf.Reset()
			err = buf.WriteSliceUint8(rspByte)
			if err != nil {
				return err
			}
		}

		if ok && traceData.TraceCall {
			var traceParam string
			traceParamFlag := traceData.NeedTraceParam(trace.EstSS, uint(buf.Len()))
			if traceParamFlag == trace.EnpNormal {
				value := map[string]interface{}{}
				value[""] = funRet
				value["err_msg"] = err_msg
				p, _ := json.Marshal(value)
				traceParam = string(p)
			} else if traceParamFlag == trace.EnpOverMaxLen {
				traceParam = "{\"trace_param_over_max_len\":true}"
			}
			ksf.Trace(traceData.GetTraceKey(trace.EstSS), trace.TraceAnnotationSS, ksf.GetClientConfig().ModuleName, ksfReq.SServantName, "SendMail", 0, traceParam, "")
		}

	default:
		if ksfReq.SFuncName == "DoClose" {
			if withContext {
				imp := val.(MailObjServantWithContext)
				imp.DoClose(ksfCtx)
			}
			return nil
		}
		return fmt.Errorf("func mismatch")
	}
	var statusMap map[string]string
	if status, ok := current.GetResponseStatus(ksfCtx); ok && status != nil {
		statusMap = status
	}
	var contextMap map[string]string
	if ctx, ok := current.GetResponseContext(ksfCtx); ok && ctx != nil {
		contextMap = ctx
	}
	*ksfResp = requestf.ResponsePacket{
		IVersion:     ksfReq.IVersion,
		CPacketType:  0,
		IRequestId:   ksfReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(buf.ToBytes()),
		Status:       statusMap,
		SResultDesc:  "",
		Context:      contextMap,
	}

	_ = readBuf
	_ = buf
	_ = length
	_ = have
	_ = ty
	return nil
}

type MailObjKsfCallback interface {
	SendMail_Callback(ret *int32, err_msg *string, opt ...map[string]string)
	SendMail_ExceptionCallback(err error)
}

// MailObjPushCallback struct
type MailObjPushCallback struct {
	Cb MailObjKsfCallback
}

func (cb *MailObjPushCallback) Ondispatch(ksfResp *requestf.ResponsePacket) {
	switch ksfResp.SResultDesc {
	case "SendMail":
		err := func() error {
			var (
				length int32
				have   bool
				ty     byte
			)
			var err error
			readBuf := codec.NewReader(tools.Int8ToByte(ksfResp.SBuffer))
			var ret int32
			err = readBuf.ReadInt32(&ret, 0, true)
			if err != nil {
				return err
			}

			var err_msg string

			err = readBuf.ReadString(&err_msg, 2, true)
			if err != nil {
				return err
			}

			_ = length
			_ = have
			_ = ty
			if ksfResp.Context != nil {
				cb.Cb.SendMail_Callback(&ret, &err_msg, ksfResp.Context)
				return nil
			} else {
				cb.Cb.SendMail_Callback(&ret, &err_msg)
				return nil
			}
		}()
		if err != nil {
			cb.Cb.SendMail_ExceptionCallback(err)
		}
	}
}
func (obj *MailObj) AsyncSendResponse_SendMail(ctx context.Context, ret int32, err_msg *string, opt ...map[string]string) (err error) {

	conn, udpAddr, ok := current.GetRawConn(ctx)
	if !ok {
		return fmt.Errorf("connection not found")
	}
	buf := codec.NewBuffer()

	err = buf.WriteInt32(ret, 0)
	if err != nil {
		return err
	}

	err = buf.WriteString(*err_msg, 2)
	if err != nil {
		return err
	}

	resp := &requestf.ResponsePacket{
		SBuffer: tools.ByteToInt8(buf.ToBytes()),
	}
	resp.IVersion = basef.KSFVERSION
	if resp.Status == nil {
		resp.Status = make(map[string]string)
	}
	resp.Status["KSF_FUNC"] = "SendMail"
	resp.SResultDesc = "SendMail"
	if len(opt) > 0 {
		if opt[0] != nil {
			resp.Context = opt[0]
		}
	}
	rspData := obj.rsp2Byte(resp)
	if udpAddr != nil {
		udpConn, _ := conn.(*net.UDPConn)
		_, err = udpConn.WriteToUDP(rspData, udpAddr)
	} else {
		_, err = conn.Write(rspData)
	}
	return err
}