package users

import (
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"testing"

	api "github.com/rodkevich/go-course/homework/hw007/users"
)

func TestServer(t *testing.T) {
	for scenario, fn := range map[string]func(
		t *testing.T,
		client api.LogClient,
		config *Config,
	){
		"produce/consume a message ":      testProduceConsume,
		"produce/consume stream succeeds": testProduceConsumeStream,
	} {
		t.Run(scenario, func(t *testing.T) {
			client, config, teardown := setupTest(t, nil)
			defer teardown()
			fn(t, client, config)
		})
	}
}

func testProduceConsume(t *testing.T, client interface{}, config *interface{}) {

}

func testProduceConsumeStream(t *testing.T, client interface{}, config *interface{}) {

}
//
// func setupTest(t *testing.T, fn func(*Config)) (
// 	client api.LogClient,
// 	cfg *Config,
// 	teardown func(),
// ) {
// 	t.Helper()
//
// 	l, err := net.Listen("tcp", ":0")
// 	require.NoError(t, err)
//
// 	clientOptions := []grpc.DialOption{grpc.WithInsecure()}
// 	cc, err := grpc.Dial(l.Addr().String(), clientOptions...)
// 	require.NoError(t, err)
//
// 	dir, err := ioutil.TempDir("", "server-test")
// 	require.NoError(t, err)
//
// 	cfg = &Config{
// 		CommitLog: clog,
// 	}
// 	if fn != nil {
// 		fn(cfg)
// 	}
// 	server, err := NewGRPCServer(cfg)
// 	require.NoError(t, err)
//
// 	go func() {
// 		server.Serve(l)
// 	}()
//
// 	client = api.NewLogClient(cc)
//
// 	return client, cfg, func() {
// 		server.Stop()
// 		cc.Close()
// 		l.Close()
// 	}
// }
