package service

import (
	"context"
	"testing"

	"github.com/whaleshell/whaleshell-core/defaults"
	"github.com/whaleshell/whaleshell-sdk/go/whaleshell"
)

type providerEndpointStub struct{ record whaleshell.ProviderRecord }

func (s providerEndpointStub) GetProvider(context.Context, string) (whaleshell.ProviderRecord, error) {
	return s.record, nil
}
func TestInferenceUpstreamConfiguration(t *testing.T) {
	cases := []struct {
		name   string
		record whaleshell.ProviderRecord
		want   string
	}{
		{"deepinfra", whaleshell.ProviderRecord{Type: "deepinfra"}, defaults.InferenceDeepInfra},
		{"unknown", whaleshell.ProviderRecord{Type: "custom"}, ""},
		{"override", whaleshell.ProviderRecord{Type: "custom", Config: whaleshell.ProviderConfig{"base_url": "https://models.example/v1"}}, "https://models.example/v1"},
		{"invalid", whaleshell.ProviderRecord{Type: "openai", Config: whaleshell.ProviderConfig{"base_url": "file:///credentials"}}, ""},
		{"userinfo", whaleshell.ProviderRecord{Type: "openai", Config: whaleshell.ProviderConfig{"base_url": "https://secret@models.example"}}, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := inferenceUpstreamForType("test", providerEndpointStub{tc.record}, context.Background())
			if got != tc.want {
				t.Fatalf("got %q want %q", got, tc.want)
			}
		})
	}
}
