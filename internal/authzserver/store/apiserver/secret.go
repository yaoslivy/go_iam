package apiserver

import (
	"context"
	"go_iam/pkg/log"

	pb "go_iam/internal/pkg/proto/apiserver/v1"

	"github.com/AlekSi/pointer"
	"github.com/avast/retry-go"
	"github.com/marmotedu/errors"
)

type secrets struct {
	cli pb.CacheClient
}

func newSecrets(ds *datastore) *secrets {
	return &secrets{ds.cli}
}

// List returns all the authorization secrets.
func (s *secrets) List() (map[string]*pb.SecretInfo, error) {
	secrets := make(map[string]*pb.SecretInfo)

	log.Info("Loading secrets")

	req := &pb.ListSecretsRequest{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	}

	var resp *pb.ListSecretsResponse
	err := retry.Do(
		func() error {
			var listErr error
			resp, listErr = s.cli.ListSecrets(context.Background(), req)
			if listErr != nil {
				return listErr
			}

			return nil
		}, retry.Attempts(3),
	)
	if err != nil {
		return nil, errors.Wrap(err, "list secrets failed")
	}

	log.Infof("Secrets found (%d total):", len(resp.Items))

	for _, v := range resp.Items {
		log.Infof(" - %s:%s", v.Username, v.SecretId)
		secrets[v.SecretId] = v
	}

	return secrets, nil
}
