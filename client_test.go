package main

import (
	"context"
	"log"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Test_client_redisPing(t *testing.T) {
	tests := []struct {
		clientOptions  []clientoptions
		containerCmd   []string
		containerImage string
		name           string
		want           string
		wantErr        bool
	}{
		{
			clientOptions:  []clientoptions{},
			containerCmd:   []string{"redis-server"},
			containerImage: "redis:7.4",
			name:           "redis_default",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{withPassword("secret")},
			containerCmd:   []string{"redis-server", "--requirepass", "secret"},
			containerImage: "redis:7.4",
			name:           "redis_password",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{},
			containerCmd:   []string{"redis-server"},
			containerImage: "redis:8.4",
			name:           "redis_default",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{withPassword("secret")},
			containerCmd:   []string{"redis-server", "--requirepass", "secret"},
			containerImage: "redis:8.4",
			name:           "redis_password",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{},
			containerCmd:   []string{"valkey-server"},
			containerImage: "valkey/valkey:7.2",
			name:           "valkey_default",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{withPassword("secret")},
			containerCmd:   []string{"valkey-server", "--requirepass", "secret"},
			containerImage: "valkey/valkey:7.2",
			name:           "valkey_password",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{},
			containerCmd:   []string{"valkey-server"},
			containerImage: "valkey/valkey:8.1",
			name:           "valkey_default",
			want:           "PONG",
			wantErr:        false,
		},
		{
			clientOptions:  []clientoptions{withPassword("secret")},
			containerCmd:   []string{"valkey-server", "--requirepass", "secret"},
			containerImage: "valkey/valkey:8.1",
			name:           "valkey_password",
			want:           "PONG",
			wantErr:        false,
		},
	}

	for _, tt := range tests {

		ctx := context.Background()
		request := testcontainers.ContainerRequest{
			Image:        tt.containerImage,
			Cmd:          tt.containerCmd,
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForLog("Ready to accept connections tcp"),
		}

		container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: request,
			Started:          true,
		})
		if err != nil {
			t.Fatal("failed to start container:", err)
		}

		conn, err := container.Endpoint(ctx, "")
		if err != nil {
			log.Fatalln("failed to get connection string:", err)
		}

		defer func() {
			if err := container.Terminate(ctx); err != nil {
				log.Fatalf("failed to terminate container: %s", err)
			}
		}()

		t.Run(tt.name, func(t *testing.T) {
			tt.clientOptions = append(tt.clientOptions, withAddr(conn))
			c = createDefaultClient(tt.clientOptions...)
			got, err := c.redisPing()
			if (err != nil) != tt.wantErr {
				t.Errorf("client.redisPing() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("client.redisPing() = %v, want %v", got, tt.want)
			}
		})
	}
}
