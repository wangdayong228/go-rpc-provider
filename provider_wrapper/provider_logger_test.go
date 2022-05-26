package providers

import (
	"context"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestLoggerMiddleware(t *testing.T) {
	lp := NewLoggerProvider(MustNewBaseProvider(context.Background(), "http://test.confluxrpc.com"), os.Stdout)

	b := new(hexutil.Big)
	lp.CallContext(context.Background(), b, "eth_blockNumber")
}
