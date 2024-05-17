package pkg

import (
	"fmt"
	ioprometheusclient "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"io"
	"net/http"
)

var parser expfmt.TextParser

func Validate(metricUrl string) error {
	if metricUrl == "" {
		return fmt.Errorf("metric-url is required")
	}

	reader, err := fetchMetric(metricUrl)

	defer func() {
		if reader != nil {
			_ = reader.Close()
		}
	}()

	if err != nil {
		return err
	}

	_, err = parseMetricFamily(reader)

	if err != nil {
		return err
	}

	return nil
}

func parseMetricFamily(reader io.Reader) (map[string]*ioprometheusclient.MetricFamily, error) {
	mf, err := parser.TextToMetricFamilies(reader)
	if err != nil {
		return nil, err
	}
	return mf, nil
}

func fetchMetric(metricUrl string) (io.ReadCloser, error) {
	resp, err := http.Get(metricUrl)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
