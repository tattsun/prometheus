package url

import (
	"context"
	"net/http"

	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/config"
)

type UrlDiscovery struct {
	urls []string
}

func NewDiscovery(conf *config.UrlSDConfig) *UrlDiscovery {
	return &UrlDiscovery{}
}

func (ud *UrlDiscovery) Run(ctx context.Context, ch chan<- []*config.TargetGroup) {
}

func (ud *UrlDiscovery) refresh(ctx context.Context, ch chan<- []*config.TargetGroup) error {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	for _, url := range ud.urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		req = req.WithContext(ctx)
		res, err := client.Do(req)
		if err != nil {
			return errors.Wrapf(err, "Failed to fetch url %s", url)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.Wrapf(err, "Failed to read response body %s", url)
		}
		tgroups, err := parseTargetGroup(body)
		if err != nil {
			return errors.Wrapf(err, "Failed to parse response body %s", url)
		}

	}
}

func parseTargetGroup(content []byte) ([]*config.TargetGroup, error) {
}
