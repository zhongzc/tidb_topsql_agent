/*
Copyright © 2021 Li Yilong <liyilongko@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"testing"
	"time"

	pb "github.com/dragonly/tidb_topsql_agent/internal/app/protobuf"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

const (
	maxSQLNum = 5000
)

func testPlanBinaryDecoderFunc(plan string) (string, error) {
	return plan, nil
}

func populateCache(ts *TopSQLCollector, begin, end int, timestamp uint64) {
	// register normalized sql
	for i := begin; i < end; i++ {
		key := "sqlDigest" + strconv.Itoa(i+1)
		value := "sqlNormalized" + strconv.Itoa(i+1)
		ts.RegisterNormalizedSQL(key, value)
	}
	// register normalized plan
	for i := begin; i < end; i++ {
		key := "planDigest" + strconv.Itoa(i+1)
		value := "planNormalized" + strconv.Itoa(i+1)
		ts.RegisterNormalizedPlan(key, value)
	}
	// collect
	var records []TopSQLRecord
	for i := begin; i < end; i++ {
		records = append(records, TopSQLRecord{
			SQLDigest:  "sqlDigest" + strconv.Itoa(i+1),
			PlanDigest: "planDigest" + strconv.Itoa(i+1),
			CPUTimeMs:  uint32(i + 1),
		})
	}
	ts.Collect(timestamp, records)
}

func populateCache1(ts *TopSQLCollector, begin, end int, timestamp uint64) {
	// register normalized sql
	for i := begin; i < end; i++ {
		key := "sqlDigest" + strconv.Itoa(i+1)
		value := "sqlNormalized" + strconv.Itoa(i+1)
		ts.RegisterNormalizedSQL(key, value)
	}
	// register normalized plan
	for i := begin; i < end; i++ {
		key := "planDigest" + strconv.Itoa(i+1)
		value := "planNormalized" + strconv.Itoa(i+1)
		ts.RegisterNormalizedPlan(key, value)
	}
	// collect
	var records []TopSQLRecord
	for i := begin; i < end; i++ {
		records = append(records, TopSQLRecord{
			SQLDigest:  "sqlDigest" + strconv.Itoa(i+1),
			PlanDigest: "planDigest" + strconv.Itoa(i+1),
			CPUTimeMs:  uint32(i + 1),
		})
	}
	ts.Collect1(timestamp, records)
}

func initializeCache(maxSQLNum int, addr string) *TopSQLCollector {
	config := &TopSQLCollectorConfig{
		PlanBinaryDecoder:   testPlanBinaryDecoderFunc,
		MaxSQLNum:           maxSQLNum,
		SendToAgentInterval: time.Minute,
		AgentGRPCAddress:    addr,
		InstanceID:          "tidb-server",
	}
	ts := NewTopSQLCollector(config)
	populateCache(ts, 0, maxSQLNum, 1)
	return ts
}

func initializeCache1(maxSQLNum int, addr string) *TopSQLCollector {
	config := &TopSQLCollectorConfig{
		PlanBinaryDecoder:   testPlanBinaryDecoderFunc,
		MaxSQLNum:           maxSQLNum,
		SendToAgentInterval: time.Minute,
		AgentGRPCAddress:    addr,
		InstanceID:          "tidb-server",
	}
	ts := NewTopSQLCollector(config)
	populateCache1(ts, 0, maxSQLNum, 1)
	return ts
}

type testAgentServer struct {
	pb.UnimplementedTopSQLAgentServer
	batch []*pb.CollectCPUTimeRequest
}

func (svr *testAgentServer) CollectCPUTime(stream pb.TopSQLAgent_CollectCPUTimeServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		svr.batch = append(svr.batch, req)
	}
	resp := &pb.CollectCPUTimeResponse{}
	stream.SendAndClose(resp)
	return nil
}

func startTestServer(t *testing.T) (*grpc.Server, *testAgentServer, int) {
	addr := ":0"
	lis, err := net.Listen("tcp", addr)
	assert.NoError(t, err, "failed to listen to address %s", addr)
	server := grpc.NewServer()
	agentServer := &testAgentServer{}
	pb.RegisterTopSQLAgentServer(server, agentServer)
	go func() {
		err := server.Serve(lis)
		assert.NoError(t, err, "failed to start server")
	}()
	return server, agentServer, lis.Addr().(*net.TCPAddr).Port
}

func TestTopSQL_CollectAndGet(t *testing.T) {
	ts := initializeCache(maxSQLNum, ":23333")
	for i := 0; i < maxSQLNum; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		key := encodeCacheKey(sqlDigest, planDigest)
		entry := ts.topSQLCache.Get(key).(*TopSQLDataPoints)
		assert.Equal(t, uint32(i+1), entry.CPUTimeMsList[0])
		assert.Equal(t, uint64(1), entry.TimestampList[0])
	}
}

func TestTopSQL_CollectAndGet1(t *testing.T) {
	ts := initializeCache1(maxSQLNum, ":23333")
	for i := 0; i < maxSQLNum; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		key := encodeCacheKey(sqlDigest, planDigest)
		entry := ts.topSQLMap[key]
		assert.Equal(t, uint32(i+1), entry.CPUTimeMsList[0])
		assert.Equal(t, uint64(1), entry.TimestampList[0])
	}
}

func TestTopSQL_CollectAndVerifyFrequency(t *testing.T) {
	ts := initializeCache(maxSQLNum, ":23333")
	// traverse the frequency list, and check frequency/item content
	elem := ts.topSQLCache.freqList.Front()
	for i := 0; i < maxSQLNum; i++ {
		elem = elem.Next()
		entry := elem.Value.(*freqEntry)
		assert.Equal(t, uint64(i+1), entry.freq)
		assert.Equal(t, 1, len(entry.items))
		for item := range entry.items {
			point := item.value.(*TopSQLDataPoints)
			assert.Equal(t, uint32(i+1), point.CPUTimeMsList[0])
			assert.Equal(t, uint64(1), point.TimestampList[0])
		}
	}
}

func TestTopSQL_CollectAndVerifyFrequency1(t *testing.T) {
	ts := initializeCache1(maxSQLNum, ":23333")
	// traverse the map, and check CPU time and content
	for i := 0; i < maxSQLNum; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		encodedKey := encodeCacheKey(sqlDigest, planDigest)
		value, exist := ts.topSQLMap[encodedKey]
		assert.Equal(t, true, exist)
		assert.Equal(t, uint64(i+1), value.CPUTimeMsTotal)
		assert.Equal(t, 1, len(value.CPUTimeMsList))
		assert.Equal(t, 1, len(value.TimestampList))
		assert.Equal(t, uint32(i+1), value.CPUTimeMsList[0])
		assert.Equal(t, uint64(1), value.TimestampList[0])
	}
}

func TestTopSQL_CollectAndEvict(t *testing.T) {
	ts := initializeCache(maxSQLNum, ":23333")
	// Collect maxSQLNum records with timestamp 2 and sql plan digest from maxSQLNum/2 to maxSQLNum/2*3.
	populateCache(ts, maxSQLNum/2, maxSQLNum/2*3, 2)
	// The first maxSQLNum/2 sql plan digest should have been evicted
	for i := 0; i < maxSQLNum/2; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		key := encodeCacheKey(sqlDigest, planDigest)
		_, exist := ts.topSQLCache.items[key]
		assert.Equal(t, false, exist, "cache key '%' should be evicted", key)
		_, exist = ts.normalizedSQLMap[sqlDigest]
		assert.Equal(t, false, exist, "normalized SQL with digest '%s' should be evicted", sqlDigest)
		_, exist = ts.normalizedPlanMap[planDigest]
		assert.Equal(t, false, exist, "normalized plan with digest '%s' should be evicted", planDigest)
	}
	// Because CPU time is populated as i+1,
	// we should expect digest maxSQLNum/2+1 - maxSQLNum to have CPU time maxSQLNum+2, maxSQLNum+4, ..., maxSQLNum*2
	// and digest maxSQLNum+1 - maxSQLNum/2*3 to have CPU time maxSQLNum+1, maxSQLNum+2, ..., maxSQLNum/2*3.
	for i := maxSQLNum / 2; i < maxSQLNum/2*3; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		key := encodeCacheKey(sqlDigest, planDigest)
		item, exist := ts.topSQLCache.items[key]
		assert.Equal(t, true, exist, "cache key '%s' should exist", exist)
		entry := item.freqElement.Value.(*freqEntry)
		if i < maxSQLNum {
			assert.Equal(t, uint64((i+1)*2), entry.freq)
		} else {
			assert.Equal(t, uint64(i+1), entry.freq)
		}
	}
}

func TestTopSQL_CollectAndEvict1(t *testing.T) {
	ts := initializeCache1(maxSQLNum, ":23333")
	// Collect maxSQLNum records with timestamp 2 and sql plan digest from maxSQLNum/2 to maxSQLNum/2*3.
	populateCache1(ts, maxSQLNum/2, maxSQLNum/2*3, 2)
	// The first maxSQLNum/2 sql plan digest should have been evicted
	for i := 0; i < maxSQLNum/2; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		key := encodeCacheKey(sqlDigest, planDigest)
		_, exist := ts.topSQLMap[key]
		assert.Equal(t, false, exist, "cache key '%' should be evicted", key)
		_, exist = ts.normalizedSQLMap[sqlDigest]
		assert.Equal(t, false, exist, "normalized SQL with digest '%s' should be evicted", sqlDigest)
		_, exist = ts.normalizedPlanMap[planDigest]
		assert.Equal(t, false, exist, "normalized plan with digest '%s' should be evicted", planDigest)
	}
	// Because CPU time is populated as i+1,
	// we should expect digest maxSQLNum/2+1 - maxSQLNum to have CPU time maxSQLNum+2, maxSQLNum+4, ..., maxSQLNum*2
	// and digest maxSQLNum+1 - maxSQLNum/2*3 to have CPU time maxSQLNum+1, maxSQLNum+2, ..., maxSQLNum/2*3.
	for i := maxSQLNum / 2; i < maxSQLNum/2*3; i++ {
		sqlDigest := "sqlDigest" + strconv.Itoa(i+1)
		planDigest := "planDigest" + strconv.Itoa(i+1)
		key := encodeCacheKey(sqlDigest, planDigest)
		value, exist := ts.topSQLMap[key]
		assert.Equal(t, true, exist, "cache key '%s' should exist", exist)
		if i < maxSQLNum {
			assert.Equal(t, uint64((i+1)*2), value.CPUTimeMsTotal)
		} else {
			assert.Equal(t, uint64(i+1), value.CPUTimeMsTotal)
		}
	}
}

func TestTopSQL_CollectAndSnapshot(t *testing.T) {
	ts := initializeCache1(maxSQLNum, ":23333")
	batch := ts.snapshot()
	for _, req := range batch {
		sqlDigest := req.SqlDigest
		planDigest := req.PlanDigest
		key := encodeCacheKey(sqlDigest, planDigest)
		value, exist := ts.topSQLMap[key]
		assert.Equal(t, true, exist, "key '%s' should exist")
		assert.Equal(t, len(value.CPUTimeMsList), len(req.CpuTimeMsList))
		for i, ct := range value.CPUTimeMsList {
			assert.Equal(t, ct, req.CpuTimeMsList[i])
		}
		assert.Equal(t, len(value.TimestampList), len(req.TimestampList))
		for i, ts := range value.TimestampList {
			assert.Equal(t, ts, req.TimestampList[i])
		}
	}
}

func TestTopSQL_CollectAndSendBatch(t *testing.T) {
	server, agentServer, port := startTestServer(t)
	t.Logf("server is listening on :%d", port)
	defer server.Stop()

	ts := initializeCache1(maxSQLNum, fmt.Sprintf(":%d", port))
	batch := ts.snapshot()

	conn, stream, err := newAgentClient(ts.agentGRPCAddress, 30*time.Second)
	assert.NoError(t, err, "failed to create agent client")
	err = ts.sendBatch(stream, batch)
	assert.NoError(t, err, "failed to send batch to server")
	err = conn.Close()
	assert.NoError(t, err, "failed to close connection")

	// check for equality of server received batch and the original data
	for _, req := range agentServer.batch {
		key := encodeCacheKey(req.SqlDigest, req.PlanDigest)
		value, exist := ts.topSQLMap[key]
		assert.Equal(t, true, exist, "key '%s' should exist in topSQLMap", key)
		for i, ct := range value.CPUTimeMsList {
			assert.Equal(t, ct, req.CpuTimeMsList[i])
		}
		for i, ts := range value.TimestampList {
			assert.Equal(t, ts, req.TimestampList[i])
		}
		normalizedSQL, exist := ts.normalizedSQLMap[req.SqlDigest]
		assert.Equal(t, true, exist, "key '%s' should exist in normalizedSQLMap", req.SqlDigest)
		assert.Equal(t, normalizedSQL, req.NormalizedSql)
		normalizedPlan, exist := ts.normalizedPlanMap[req.PlanDigest]
		assert.Equal(t, true, exist, "key '%s' should exist in normalizedPlanMap", req.PlanDigest)
		assert.Equal(t, normalizedPlan, req.NormalizedPlan)
	}
}

func BenchmarkTopSQL_CollectAndIncrementFrequency(b *testing.B) {
	ts := initializeCache(maxSQLNum, ":23333")
	for i := 0; i < b.N; i++ {
		populateCache(ts, 0, maxSQLNum, uint64(i))
	}
}

func BenchmarkTopSQL_CollectAndIncrementFrequency1(b *testing.B) {
	ts := initializeCache1(maxSQLNum, ":23333")
	for i := 0; i < b.N; i++ {
		populateCache1(ts, 0, maxSQLNum, uint64(i))
	}
}

func BenchmarkTopSQL_CollectAndEvict(b *testing.B) {
	ts := initializeCache(maxSQLNum, ":23333")
	begin := 0
	end := maxSQLNum
	for i := 0; i < b.N; i++ {
		begin += maxSQLNum
		end += maxSQLNum
		populateCache(ts, begin, end, uint64(i))
	}
}

func BenchmarkTopSQL_CollectAndEvict1(b *testing.B) {
	ts := initializeCache1(maxSQLNum, ":23333")
	begin := 0
	end := maxSQLNum
	for i := 0; i < b.N; i++ {
		begin += maxSQLNum
		end += maxSQLNum
		populateCache1(ts, begin, end, uint64(i))
	}
}
