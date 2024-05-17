package main

import (
	"flag"
	"fmt"
	"github.com/kuckjwi0928/prometheus-metric-validator/pkg"
	"os"
)

func main() {
	var metricUrl string

	flag.StringVar(&metricUrl, "metric-url", "", "URL of the metric to validate")
	flag.Parse()

	err := pkg.Validate(metricUrl)

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
	}
}
