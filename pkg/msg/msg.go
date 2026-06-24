// Copyright 2016 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package msg

import (
	"net"
	"reflect"
)

const (
	TypeLogin              byte = 'x'
	TypeLoginResp          byte = 'a'
	TypeNewProxy           byte = 'y'
	TypeNewProxyResp       byte = 'b'
	TypeCloseProxy         byte = 'z'
	TypeNewWorkConn        byte = 'c'
	TypeReqWorkConn        byte = 'd'
	TypeStartWorkConn      byte = 'e'
	TypeNewVisitorConn     byte = 'f'
	TypeNewVisitorConnResp byte = 'g'
	TypePing               byte = 'h'
	TypePong               byte = 'i'
	TypeUDPPacket          byte = 'j'
	TypeNatHoleVisitor     byte = 'k'
	TypeNatHoleClient      byte = 'l'
	TypeNatHoleResp        byte = 'm'
	TypeNatHoleSid         byte = 'n'
	TypeNatHoleReport      byte = 'o'
)

var msgTypeMap = map[byte]any{
	TypeLogin:              Login{},
	TypeLoginResp:          LoginResp{},
	TypeNewProxy:           NewProxy{},
	TypeNewProxyResp:       NewProxyResp{},
	TypeCloseProxy:         CloseProxy{},
	TypeNewWorkConn:        NewWorkConn{},
	TypeReqWorkConn:        ReqWorkConn{},
	TypeStartWorkConn:      StartWorkConn{},
	TypeNewVisitorConn:     NewVisitorConn{},
	TypeNewVisitorConnResp: NewVisitorConnResp{},
	TypePing:               Ping{},
	TypePong:               Pong{},
	TypeUDPPacket:          UDPPacket{},
	TypeNatHoleVisitor:     NatHoleVisitor{},
	TypeNatHoleClient:      NatHoleClient{},
	TypeNatHoleResp:        NatHoleResp{},
	TypeNatHoleSid:         NatHoleSid{},
	TypeNatHoleReport:      NatHoleReport{},
}

var TypeNameNatHoleResp = reflect.TypeFor[NatHoleResp]().Name()

type ClientSpec struct {
	Type           string `json:"a,omitempty"`
	AlwaysAuthPass bool   `json:"b,omitempty"`
}

type Login struct {
	Version      string            `json:"c,omitempty"`
	Hostname     string            `json:"d,omitempty"`
	Os           string            `json:"e,omitempty"`
	Arch         string            `json:"f,omitempty"`
	User         string            `json:"g,omitempty"`
	PrivilegeKey string            `json:"h,omitempty"`
	Timestamp    int64             `json:"i,omitempty"`
	RunID        string            `json:"j,omitempty"`
	ClientID     string            `json:"k,omitempty"`
	Metas        map[string]string `json:"l,omitempty"`
	ClientSpec   ClientSpec        `json:"m,omitempty"`
	PoolCount    int               `json:"n,omitempty"`
}

type LoginResp struct {
	Version string `json:"o,omitempty"`
	RunID   string `json:"p,omitempty"`
	Error   string `json:"q,omitempty"`
}

type NewProxy struct {
	ProxyName          string            `json:"r,omitempty"`
	ProxyType          string            `json:"s,omitempty"`
	UseEncryption      bool              `json:"t,omitempty"`
	UseCompression     bool              `json:"u,omitempty"`
	BandwidthLimit     string            `json:"v,omitempty"`
	BandwidthLimitMode string            `json:"w,omitempty"`
	Group              string            `json:"x,omitempty"`
	GroupKey           string            `json:"y,omitempty"`
	Metas              map[string]string `json:"z,omitempty"`
	Annotations        map[string]string `json:"aa,omitempty"`
	RemotePort         int               `json:"ab,omitempty"`
	CustomDomains      []string          `json:"ac,omitempty"`
	SubDomain          string            `json:"ad,omitempty"`
	Locations          []string          `json:"ae,omitempty"`
	HTTPUser           string            `json:"af,omitempty"`
	HTTPPwd            string            `json:"ag,omitempty"`
	HostHeaderRewrite  string            `json:"ah,omitempty"`
	Headers            map[string]string `json:"ai,omitempty"`
	ResponseHeaders    map[string]string `json:"aj,omitempty"`
	RouteByHTTPUser    string            `json:"ak,omitempty"`
	Sk                 string            `json:"al,omitempty"`
	AllowUsers         []string          `json:"am,omitempty"`
	Multiplexer        string            `json:"an,omitempty"`
}

type NewProxyResp struct {
	ProxyName  string `json:"ao,omitempty"`
	RemoteAddr string `json:"ap,omitempty"`
	Error      string `json:"aq,omitempty"`
}

type CloseProxy struct {
	ProxyName string `json:"ar,omitempty"`
}

type NewWorkConn struct {
	RunID        string `json:"as,omitempty"`
	PrivilegeKey string `json:"at,omitempty"`
	Timestamp    int64  `json:"au,omitempty"`
}

type ReqWorkConn struct{}

type StartWorkConn struct {
	ProxyName string `json:"av,omitempty"`
	SrcAddr   string `json:"aw,omitempty"`
	DstAddr   string `json:"ax,omitempty"`
	SrcPort   uint16 `json:"ay,omitempty"`
	DstPort   uint16 `json:"az,omitempty"`
	Error     string `json:"ba,omitempty"`
}

type NewVisitorConn struct {
	RunID          string `json:"bb,omitempty"`
	ProxyName      string `json:"bc,omitempty"`
	SignKey        string `json:"bd,omitempty"`
	Timestamp      int64  `json:"be,omitempty"`
	UseEncryption  bool   `json:"bf,omitempty"`
	UseCompression bool   `json:"bg,omitempty"`
}

type NewVisitorConnResp struct {
	ProxyName string `json:"bh,omitempty"`
	Error     string `json:"bi,omitempty"`
}

type Ping struct {
	PrivilegeKey string `json:"bj,omitempty"`
	Timestamp    int64  `json:"bk,omitempty"`
}

type Pong struct {
	Error string `json:"bl,omitempty"`
}

type UDPPacket struct {
	Content    []byte       `json:"bm,omitempty"`
	LocalAddr  *net.UDPAddr `json:"bn,omitempty"`
	RemoteAddr *net.UDPAddr `json:"bo,omitempty"`
}

type NatHoleVisitor struct {
	TransactionID string   `json:"bp,omitempty"`
	ProxyName     string   `json:"bq,omitempty"`
	PreCheck      bool     `json:"br,omitempty"`
	Protocol      string   `json:"bs,omitempty"`
	SignKey       string   `json:"bt,omitempty"`
	Timestamp     int64    `json:"bu,omitempty"`
	MappedAddrs   []string `json:"bv,omitempty"`
	AssistedAddrs []string `json:"bw,omitempty"`
}

type NatHoleClient struct {
	TransactionID string   `json:"bp,omitempty"`
	ProxyName     string   `json:"bq,omitempty"`
	Sid           string   `json:"bz,omitempty"`
	MappedAddrs   []string `json:"bv,omitempty"`
	AssistedAddrs []string `json:"bw,omitempty"`
}

type PortsRange struct {
	From int `json:"cc,omitempty"`
	To   int `json:"cd,omitempty"`
}

type NatHoleDetectBehavior struct {
	Role              string       `json:"ce,omitempty"` // sender or receiver
	Mode              int          `json:"cf,omitempty"` // 0, 1, 2...
	TTL               int          `json:"cg,omitempty"`
	SendDelayMs       int          `json:"ch,omitempty"`
	ReadTimeoutMs     int          `json:"ci,omitempty"`
	CandidatePorts    []PortsRange `json:"cj,omitempty"`
	SendRandomPorts   int          `json:"ck,omitempty"`
	ListenRandomPorts int          `json:"cl,omitempty"`
}

type NatHoleResp struct {
	TransactionID  string                `json:"bp,omitempty"`
	Sid            string                `json:"bz,omitempty"`
	Protocol       string                `json:"bs,omitempty"`
	CandidateAddrs []string              `json:"cp,omitempty"`
	AssistedAddrs  []string              `json:"bw,omitempty"`
	DetectBehavior NatHoleDetectBehavior `json:"cr,omitempty"`
	Error          string                `json:"error,omitempty"`
}

type NatHoleSid struct {
	TransactionID string `json:"bp,omitempty"`
	Sid           string `json:"bz,omitempty"`
	Response      bool   `json:"cv,omitempty"`
	Nonce         string `json:"cw,omitempty"`
}

type NatHoleReport struct {
	Sid     string `json:"bz,omitempty"`
	Success bool   `json:"cy,omitempty"`
}
