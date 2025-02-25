package framework

import (
	"fmt"
	"net/http"
	"os"

	"frp_lib/test/e2e/framework/consts"
	"frp_lib/test/e2e/mock/server"
	"frp_lib/test/e2e/mock/server/httpserver"
	"frp_lib/test/e2e/mock/server/streamserver"
	"frp_lib/test/e2e/pkg/port"
)

const (
	TCPEchoServerPort    = "TCPEchoServerPort"
	UDPEchoServerPort    = "UDPEchoServerPort"
	UDSEchoServerAddr    = "UDSEchoServerAddr"
	HTTPSimpleServerPort = "HTTPSimpleServerPort"
)

type MockServers struct {
	tcpEchoServer    server.Server
	udpEchoServer    server.Server
	udsEchoServer    server.Server
	httpSimpleServer server.Server
}

func NewMockServers(portAllocator *port.Allocator) *MockServers {
	s := &MockServers{}
	tcpPort := portAllocator.Get()
	udpPort := portAllocator.Get()
	httpPort := portAllocator.Get()
	s.tcpEchoServer = streamserver.New(streamserver.TCP, streamserver.WithBindPort(tcpPort))
	s.udpEchoServer = streamserver.New(streamserver.UDP, streamserver.WithBindPort(udpPort))
	s.httpSimpleServer = httpserver.New(httpserver.WithBindPort(httpPort),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			_, _ = w.Write([]byte(consts.TestString))
		})),
	)

	udsIndex := portAllocator.Get()
	udsAddr := fmt.Sprintf("%s/frp_echo_server_%d.sock", os.TempDir(), udsIndex)
	os.Remove(udsAddr)
	s.udsEchoServer = streamserver.New(streamserver.Unix, streamserver.WithBindAddr(udsAddr))
	return s
}

func (m *MockServers) Run() error {
	if err := m.tcpEchoServer.Run(); err != nil {
		return err
	}
	if err := m.udpEchoServer.Run(); err != nil {
		return err
	}
	if err := m.udsEchoServer.Run(); err != nil {
		return err
	}
	if err := m.httpSimpleServer.Run(); err != nil {
		return err
	}
	return nil
}

func (m *MockServers) Close() {
	m.tcpEchoServer.Close()
	m.udpEchoServer.Close()
	m.udsEchoServer.Close()
	m.httpSimpleServer.Close()
	os.Remove(m.udsEchoServer.BindAddr())
}

func (m *MockServers) GetTemplateParams() map[string]interface{} {
	ret := make(map[string]interface{})
	ret[TCPEchoServerPort] = m.tcpEchoServer.BindPort()
	ret[UDPEchoServerPort] = m.udpEchoServer.BindPort()
	ret[UDSEchoServerAddr] = m.udsEchoServer.BindAddr()
	ret[HTTPSimpleServerPort] = m.httpSimpleServer.BindPort()
	return ret
}

func (m *MockServers) GetParam(key string) interface{} {
	params := m.GetTemplateParams()
	if v, ok := params[key]; ok {
		return v
	}
	return nil
}
