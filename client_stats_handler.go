package ocgorpc

import (
	"context"
	"time"

	"github.com/gxed/opencensus-go/tag"
	rpcstats "github.com/libp2p/go-libp2p-gorpc/stats"
)

type statsKey string

// statsTagRPC gets the tag.Map populated by the application code, serializes
// its tags into the gorpc metadata in order to be sent to the server.
func (h *ClientHandler) statsTagRPC(ctx context.Context, info *rpcstats.RPCTagInfo) context.Context {
	startTime := time.Now()
	if info == nil {
		logger.Infof("clientHandler.TagRPC called with nil info.", info.FullMethodName)
		return ctx
	}

	d := &rpcData{
		startTime: startTime,
		method:    info.FullMethodName,
	}
	ts := tag.FromContext(ctx)
	if ts != nil {
		encoded := tag.Encode(ts)
		ctx = context.WithValue(ctx, statsKey("gorpc-stats"), encoded)
	}

	return context.WithValue(ctx, rpcDataKey, d)
}
