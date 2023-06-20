package v1

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const ServiceName = "third-party-notice"

func NewClient() (ThirdPartyNoticeClient, func(), error) {
	if os.Getenv("ETCD") != "" {
		points := os.Getenv("ETCD_ENDPOINTS")
		endpoints := strings.Split(points, ",")
		log.Debugf("Etcd Endpoints %v", endpoints)
		cli, err := clientv3.New(clientv3.Config{
			Endpoints: endpoints,
		})
		if err != nil {
			log.Errorf("New ThirdPartyNotice Grpc Client Connect Etcd err %v", err)
			return nil, nil, err
		}
		r := etcd.New(cli)
		conn, err := grpc.DialInsecure(
			context.Background(),
			grpc.WithEndpoint(fmt.Sprintf("discovery:///%s", ServiceName)),
			grpc.WithDiscovery(r),
		)
		if err != nil {
			log.Errorf("New ThirdPartyNotice Grpc DialInsecure err %v", err)
			return nil, nil, err
		}
		return NewThirdPartyNoticeClient(conn), func() {
			conn.Close()
			cli.Close()
		}, nil
	}

	return nil, nil, nil
}
