// Copyright 2026 The frp Authors
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
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestV1MessageTypeIDsAreStable(t *testing.T) {
	require.Equal(t, byte('x'), TypeLogin)
	require.Equal(t, byte('a'), TypeLoginResp)
	require.Equal(t, byte('y'), TypeNewProxy)
	require.Equal(t, byte('b'), TypeNewProxyResp)
	require.Equal(t, byte('z'), TypeCloseProxy)
	require.Equal(t, byte('c'), TypeNewWorkConn)
	require.Equal(t, byte('d'), TypeReqWorkConn)
	require.Equal(t, byte('e'), TypeStartWorkConn)
	require.Equal(t, byte('f'), TypeNewVisitorConn)
	require.Equal(t, byte('g'), TypeNewVisitorConnResp)
	require.Equal(t, byte('h'), TypePing)
	require.Equal(t, byte('i'), TypePong)
	require.Equal(t, byte('j'), TypeUDPPacket)
	require.Equal(t, byte('k'), TypeNatHoleVisitor)
	require.Equal(t, byte('l'), TypeNatHoleClient)
	require.Equal(t, byte('m'), TypeNatHoleResp)
	require.Equal(t, byte('n'), TypeNatHoleSid)
	require.Equal(t, byte('o'), TypeNatHoleReport)
}

func TestMessageTypeMapIsCompleteAndUnique(t *testing.T) {
	require.Len(t, msgTypeMap, 18)

	msgTypes := make(map[reflect.Type]struct{}, len(msgTypeMap))

	for _, m := range msgTypeMap {
		msgType := reflect.TypeOf(m)
		require.NotContains(t, msgTypes, msgType)
		msgTypes[msgType] = struct{}{}
	}
}
