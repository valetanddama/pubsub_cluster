package test

import (
	"bytes"
	"github.com/valetanddama/cluster/pkg/config"
	"github.com/valetanddama/cluster/pkg/services"
	"log"
	"testing"
)

func init() {
	var buf bytes.Buffer
	log.SetOutput(&buf)
}

func BenchmarkSendMessageForRoleGenerator(b *testing.B) {
	config.Conn.ServerID = "test"
	config.Conn.ServerRole = config.RoleGeneratorMessages

	for i := 0; i < b.N; i++ {
		services.SendMessage()
	}
}

func BenchmarkGetMessageForRoleHandler(b *testing.B) {
	config.Conn.ServerID = "test"
	config.Conn.ServerRole = config.RoleHandlerMessages

	for i := 0; i < b.N; i++ {
		services.GetMessage()
	}
}

func BenchmarkCheckServerRole(b *testing.B) {
	config.Conn.ServerID = "test"
	config.Conn.ServerRole = config.RoleHandlerMessages

	for i := 0; i < b.N; i++ {
		services.CheckServerRole()
	}
}

func BenchmarkGetRedisConn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		config.Conn.Redis()
	}
}
