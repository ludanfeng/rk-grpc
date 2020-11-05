// Copyright (c) 2020 rookie-ninja
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
package rk_grpc_log

import (
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
)

var (
	realm         = zap.String("realm", getEnvValueOrDefault("REALM", "unknown"))
	region        = zap.String("region", getEnvValueOrDefault("REGION", "unknown"))
	az            = zap.String("az", getEnvValueOrDefault("AZ", "unknown"))
	domain        = zap.String("domain", getEnvValueOrDefault("DOMAIN", "unknown"))
	appVersion    = zap.String("app_version", getEnvValueOrDefault("APP_VERSION", "latest"))
	localIP       = zap.String("local.IP", getLocalIp())
	localHostname = zap.String("local.hostname", getLocalHostname())
	appName       = "Unknown"
)

const (
	// We will add three dot after truncation
	maxRequestStrLen  = 1021
	maxResponseStrLen = 1021
)

func getEnvValueOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)

	if len(value) < 1 {
		return defaultValue
	}

	return value
}

// Get remote endpoint information set including IP, Port, NetworkType
// We will do as best as we can to determine it
// If fails, then just return default ones
func getRemoteAddressSet(ctx context.Context) []zap.Field {
	remoteIP := "0.0.0.0"
	remotePort := "0"
	remoteNetworkType := ""
	if peer, ok := peer.FromContext(ctx); ok {
		remoteNetworkType = peer.Addr.Network()

		// Here is the tricky part
		// We only try to parse IPV4 style Address
		// Rest of peer.Addr implementations are not well formatted string
		// and in this case, we leave port as zero and IP as the returned
		// String from Addr.String() function
		//
		// BTW, just skip the error since it would not impact anything
		// Operators could observe this error from monitor dashboards by
		// validating existence of IP & PORT fields
		remoteIP, remotePort, _ = net.SplitHostPort(peer.Addr.String())
	}

	headers, ok := metadata.FromIncomingContext(ctx)

	if ok {
		forwardedRemoteIPList := headers["x-forwarded-for"]

		// Deal with forwarded remote ip
		if len(forwardedRemoteIPList) > 0 {
			forwardedRemoteIP := forwardedRemoteIPList[0]

			if forwardedRemoteIP == "::1" {
				forwardedRemoteIP = "localhost"
			}

			remoteIP = forwardedRemoteIP
		}
	}

	if remoteIP == "::1" {
		remoteIP = "localhost"
	}

	return []zap.Field{
		zap.String("remote.IP", remoteIP),
		zap.String("remote.port", remotePort),
		zap.String("remote.net_type", remoteNetworkType),
	}
}

// This is a tricky function
// We will iterate through all the network interfaces
// but will choose the first one since we are assuming that
// eth0 will be the default one to use in most of the case.
//
// Currently, we do not have any interfaces for selecting the network
// interface yet.
func getLocalIp() string {
	localIP := "localhost"

	// skip the error since we don't want to break RPC calls because of it
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		items := strings.Split(addr.String(), "/")
		if len(items) < 2 || items[0] == "127.0.0.1" {
			continue
		}
		if match, err := regexp.MatchString(`\d+\.\d+\.\d+\.\d+`, items[0]); err == nil && match {
			localIP = items[0]
		}
	}

	return localIP
}

func getLocalHostname() string {
	hostname, err := os.Hostname()
	if err != nil || len(hostname) < 1 {
		hostname = "unknown"
	}

	return hostname
}

func interfaceToString(input interface{}, max int) string {
	if input == nil {
		return "nil"
	}
	pbMarshaler := &runtime.JSONPb{
		EmitDefaults: true,
		OrigName:     true,
		EnumsAsInts:  true,
	}

	bytes, err := pbMarshaler.Marshal(input)
	if err != nil {
		// Error occurred
		return "parse_failure"
	}

	str := string(bytes)

	if len(str) > max {
		str = str[0:max] + "..."
	}

	return str
}