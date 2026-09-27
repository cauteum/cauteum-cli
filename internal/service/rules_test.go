package service

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestRuleApproveAllGatewayApprovesPendingWithoutClearing(t *testing.T) {
	approved := map[string]bool{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/sandboxes/demo/proposals":
			if r.URL.Query().Get("status") != "pending" {
				t.Errorf("status filter = %q", r.URL.Query().Get("status"))
			}
			fmt.Fprint(w, `{"proposals":[{"id":"ordinary","status":"pending"},{"id":"flagged","status":"pending","security_flagged":true}]}`)
		case "/v1/sandboxes/demo/proposals/ordinary/approve":
			approved["ordinary"] = true
			fmt.Fprint(w, `{"id":"ordinary","status":"approved"}`)
		case "/v1/sandboxes/demo/proposals/flagged/approve":
			approved["flagged"] = true
			fmt.Fprint(w, `{"id":"flagged","status":"approved"}`)
		case "/v1/sandboxes/demo/policy":
			fmt.Fprint(w, "version: 1\nnetwork_policies: {}\n")
		default:
			http.Error(w, "unexpected path", http.StatusNotFound)
		}
	}))
	defer server.Close()
	a := &App{GatewayURLOverride: server.URL}
	if err := a.RuleApproveAll("demo", false); err != nil {
		t.Fatal(err)
	}
	if !approved["ordinary"] || approved["flagged"] {
		t.Fatalf("approved = %v", approved)
	}
	if err := a.RuleApproveAll("demo", true); err != nil {
		t.Fatal(err)
	}
	if !approved["flagged"] {
		t.Fatal("explicit opt-in did not approve flagged proposal")
	}
}

func TestRuleApproveAllLocalPreservesHistoryAndFlaggedRules(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	dir, err := localParityDir()
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(dir, "rules.json")
	initial := map[string]any{
		"ordinary": map[string]any{"sandbox": "demo", "status": "pending"},
		"flagged":  map[string]any{"sandbox": "demo", "status": "pending", "security_flagged": true},
		"other":    map[string]any{"sandbox": "other", "status": "pending"},
		"history":  map[string]any{"sandbox": "demo", "status": "rejected"},
	}
	if err := writeJSONMap(path, initial); err != nil {
		t.Fatal(err)
	}
	a := &App{}
	if err := a.ruleApproveAllLocal("demo", false); err != nil {
		t.Fatal(err)
	}
	got, err := readJSONMap(path)
	if err != nil {
		t.Fatal(err)
	}
	for id, want := range map[string]string{
		"ordinary": "approved", "flagged": "pending", "other": "pending", "history": "rejected",
	} {
		row := got[id].(map[string]any)
		if row["status"] != want {
			t.Errorf("%s status = %v, want %s", id, row["status"], want)
		}
	}
	if len(got) != len(initial) {
		t.Fatalf("proposal history lost: got %d rows, want %d", len(got), len(initial))
	}
	if err := a.ruleApproveAllLocal("demo", true); err != nil {
		t.Fatal(err)
	}
	got, err = readJSONMap(path)
	if err != nil {
		t.Fatal(err)
	}
	if got["flagged"].(map[string]any)["status"] != "approved" {
		t.Fatal("explicit opt-in did not approve flagged proposal")
	}
}
