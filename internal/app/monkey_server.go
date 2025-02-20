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
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pingcap/kvproto/pkg/resource_usage_agent"
	"github.com/pingcap/tipb/go-tipb"
	"google.golang.org/grpc"
)

// gRPC server
var _ tipb.TopSQLAgentServer = &monkeyServer{}
var _ resource_usage_agent.ResourceUsageAgentServer = &monkeyServer{}

type monkeyServer struct{}

func (*monkeyServer) ReportPlanMeta(stream tipb.TopSQLAgent_ReportPlanMetaServer) error {
	log.Println("TiDB called ReportPlanMeta()")
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		count += 1;
	}
	log.Printf("received # of tidb plan meta: %v\n", count)
	resp := &tipb.EmptyResponse{}
	stream.SendAndClose(resp)
	return nil
}

func (*monkeyServer) ReportSQLMeta(stream tipb.TopSQLAgent_ReportSQLMetaServer) error {
	log.Println("TiDB called ReportSQLMeta()")
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		count += 1;
	}
	log.Printf("received # of tidb sql meta: %v\n", count)
	resp := &tipb.EmptyResponse{}
	stream.SendAndClose(resp)
	return nil
}

func (*monkeyServer) ReportCPUTimeRecords(stream tipb.TopSQLAgent_ReportCPUTimeRecordsServer) error {
	log.Println("TiDB called ReportCPUTimeRecords()")
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		count += 1;
	}
	log.Printf("received # of tidb cpu time: %v\n", count)
	resp := &tipb.EmptyResponse{}
	stream.SendAndClose(resp)
	return nil
}

func (s *monkeyServer) ReportCpuTime(stream resource_usage_agent.ResourceUsageAgent_ReportCpuTimeServer) error {
	log.Println("TiKV called ReportCpuTime()")
	count := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		count += 1;
	}
	log.Printf("received # of tikv cpu time: %v\n", count)

	resp := &resource_usage_agent.ReportCpuTimeResponse{}
	stream.SendAndClose(resp)
	return nil
}

var (
	dropStartTime atomic.Value // time.Time
	dropEndTime   atomic.Value //time.Time
)

// startStdinReader reads commands from stdin, and do things
func startStdinReader() {
	dropStartTime.Store(time.Now().Add(-time.Second))
	dropEndTime.Store(time.Now().Add(-time.Second))
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Trim(text, "\n")
		if len(text) > 4 && text[:4] == "drop" {
			seconds, err := strconv.Atoi(text[5:])
			if err != nil {
				log.Printf("[stdin] wrong command: %s, %v\n", text, err)
				continue
			}
			log.Printf("[stdin] drop TCP traffic for %d seconds\n", seconds)
			now := time.Now()
			dropStartTime.Store(now)
			dropEndTime.Store(now.Add(time.Second * time.Duration(seconds)))
		} else {
			log.Printf("[stdin] echo: %s\n", text)
		}
	}
}

// startProxyServer proxy TCP traffic to gRPC server, with some interference
func startProxyServer(lisProxy net.Listener, toPort int) error {
	toAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("localhost:%d", toPort))
	if err != nil {
		return err
	}
	log.Printf("top addr: %+v\n", toAddr)
	for {
		connFrom, err := lisProxy.Accept()
		log.Printf("new client connection: %+v\n", connFrom)
		if err != nil {
			log.Printf("[proxy] failed to create connection from client, %v", err)
			continue
		}
		go func() {
			now := time.Now()
			if now.After(dropStartTime.Load().(time.Time)) && now.Before(dropEndTime.Load().(time.Time)) {
				// now the TCP traffic should be dropped
				log.Print("[proxy] drop TCP traffic")
				return
			}
			defer connFrom.Close()
			connTo, err := net.DialTCP("tcp", nil, toAddr)
			log.Printf("new upstream connection: %+v\n", connTo)
			if err != nil {
				log.Printf("[proxy] failed to create connection to upstream, %v", err)
				return
			}
			defer connTo.Close()
			wg := sync.WaitGroup{}
			wg.Add(2)
			go copy(&wg, connTo, connFrom)
			go copy(&wg, connFrom, connTo)
			wg.Wait()
		}()
	}
}

// TODO: support dropping in-flight TCP packets, loop with CopyN
func copy(wg *sync.WaitGroup, dst io.Writer, src io.Reader) {
	n, err := io.Copy(dst, src)
	log.Printf("copied %d bytes from %v to %v, err: %v", n, src, dst, err)
	wg.Done()
}

func StartMonkeyServer(proxyAddress string) {
	// proxy
	lisProxy, err := net.Listen("tcp", proxyAddress)
	addrProxy := lisProxy.Addr().(*net.TCPAddr)
	if err != nil {
		log.Fatalf("[proxy] failed to listen on TCP address %s:%d, %v", addrProxy.IP, addrProxy.Port, err)
	}
	defer lisProxy.Close()
	log.Printf("[proxy] start listening on %s:%d", addrProxy.IP, addrProxy.Port)

	// gRPC
	lisGRPC, err := net.Listen("tcp", ":0")
	addrGRPC := lisGRPC.Addr().(*net.TCPAddr)
	if err != nil {
		log.Fatalf("[gRPC] failed to listen on TCP address %s:%d, %v", addrGRPC.IP, addrGRPC.Port, err)
	}
	defer lisGRPC.Close()
	log.Printf("[gRPC] start listening on %s:%d", addrGRPC.IP, addrGRPC.Port)

	server := grpc.NewServer()
	service := monkeyServer{}
	tipb.RegisterTopSQLAgentServer(server, &service)
	resource_usage_agent.RegisterResourceUsageAgentServer(server, &service)
	log.Printf("[gRPC] start gRPC server")

	go func() {
		if err := server.Serve(lisGRPC); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	go startStdinReader()

	log.Println("[proxy] start proxying TCP traffic")
	startProxyServer(lisProxy, addrGRPC.Port)
}
