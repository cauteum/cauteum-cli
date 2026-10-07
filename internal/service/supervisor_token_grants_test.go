package service

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/whaleshell/whaleshell-proxy/proxy"
	"github.com/whaleshell/whaleshell-sdk/go/whaleshell"
)

func TestSandboxTokenGrantsPackagesAttachedProfileMetadata(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/sandboxes/sb":
			_, _ = w.Write([]byte(`{"name":"sb","workspace":"team","attached_providers":["corp"]}`))
		case "/v1/providers/corp":
			_, _ = w.Write([]byte(`{"name":"corp","type":"corp-api","workspace":"team"}`))
		case "/v1/profiles/corp-api":
			w.Header().Set("Content-Type", "application/yaml")
			_, _ = w.Write([]byte(`{"source":"custom","profile":{"id":"corp-api","credentials":[{"name":"DYNAMIC_CREDENTIAL","env_vars":["DYNAMIC_TOKEN"],"token_grant":{"grant_type":"token_exchange","token_endpoint":"https://issuer.example/token","audience":"https://api.example.com","jwt_svid_audience":"https://issuer.example","scopes":["read"],"subject_token":{"source":"provider_credential","credential":"UPSTREAM_TOKEN"}}}]}}`))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	c := whaleshell.NewWithToken(server.URL, "test")
	raw := sandboxTokenGrants(context.Background(), c, "sb", server.URL)
	if raw == "" {
		t.Fatal("expected grant metadata")
	}
	var grants map[string]proxy.TokenGrantCredential
	if err := json.Unmarshal([]byte(raw), &grants); err != nil {
		t.Fatal(err)
	}
	g, ok := grants["DYNAMIC_TOKEN"]
	if !ok {
		t.Fatalf("grant keys=%v", grants)
	}
	if g.Provider != "corp" || g.CredentialKey != "DYNAMIC_CREDENTIAL" || g.SubjectTokenCredential != "UPSTREAM_TOKEN" || g.TokenEndpoint != "https://issuer.example/token" {
		t.Fatalf("grant metadata=%+v", g)
	}
}
