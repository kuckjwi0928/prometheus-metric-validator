package pkg

import (
	"strings"
	"testing"
)

func TestParseMetricFamily(t *testing.T) {
	tests := []struct {
		name         string
		metricString string
		wantErr      bool
	}{
		{
			name: "Valid metric",
			metricString: `# HELP net_conntrack_dialer_conn_attempted_total
# TYPE net_conntrack_dialer_conn_attempted_total untyped
net_conntrack_dialer_conn_attempted_total{dialer_name="federate",instance="localhost:9090",job="prometheus"} 1 1608520832877
`,
			wantErr: false,
		},
		{
			name:         "Invalid metric",
			metricString: "# HELP net_conntrack_dialer_conn_attempted_total",
			wantErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.metricString)
			_, err := parseMetricFamily(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseMetricFamily() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
