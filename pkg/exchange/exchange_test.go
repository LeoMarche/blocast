package exchange

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {

	type test struct {
		IPInit      string
		IPType      int
		IPServ      string
		ShouldError bool
		IsEmpty     bool
	}

	t1 := test{DEFAULT_IPV4, 4, DEFAULT_IPV4, false, false}
	t2 := test{DEFAULT_IPV6, 6, DEFAULT_IPV6, false, false}
	t3 := test{"10.10.10.10", 4, DEFAULT_IPV4, true, true}
	t4 := test{"[::2]", 6, DEFAULT_IPV6, true, true}

	tests := []test{t1, t2, t3, t4}

	for _, v := range tests {
		CONN_LIST = CONN_LIST[:0]

		go func() {
			if v.IPType == 4 {
				DEFAULT_IPV4 = v.IPInit
			} else if v.IPType == 6 {
				DEFAULT_IPV6 = v.IPInit
			}
			err := InitializeList()
			if v.ShouldError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		}()
		l, err := net.Listen("tcp", v.IPServ+":"+DEFAULT_SERVER_PORT)
		assert.NoError(t, err)
		if v.IsEmpty {
			assert.Empty(t, CONN_LIST)
		} else {
			conn, err := l.Accept()
			assert.NoError(t, err)
			assert.NotEmptyf(t, CONN_LIST, "Should not be empty for ip: %s", v.IPInit)
			conn.Close()
		}
		l.Close()
	}
}
