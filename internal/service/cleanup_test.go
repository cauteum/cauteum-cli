package service

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCleanupDockerTestResourcesIsScopedAndDryRunByDefault(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("fake Docker executable uses a Unix shell script")
	}

	binDir := t.TempDir()
	logPath := filepath.Join(t.TempDir(), "docker.log")
	docker := filepath.Join(binDir, "docker")
	script := `#!/bin/sh
printf '%s\n' "$*" >> "$WHALESHELL_FAKE_DOCKER_LOG"
case "$1 $2" in
  "volume ls") printf 'anonymous-a\nanonymous-b\n' ;;
  "ps -aq") printf 'container-a\n' ;;
esac
`
	if err := os.WriteFile(docker, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	t.Setenv("WHALESHELL_FAKE_DOCKER_LOG", logPath)

	a := New()
	if err := a.CleanupDockerTestResources(context.Background(), false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	log := string(data)
	if strings.Contains(log, "volume rm") || strings.Contains(log, "rm -f") || strings.Contains(log, "prune") {
		t.Fatalf("dry-run performed destructive docker operation: %q", log)
	}

	if err := a.CleanupDockerTestResources(context.Background(), true); err != nil {
		t.Fatal(err)
	}
	data, err = os.ReadFile(logPath)
	if err != nil {
		t.Fatal(err)
	}
	log = string(data)
	for _, want := range []string{
		"volume rm anonymous-a",
		"volume rm anonymous-b",
		"rm -f container-a",
	} {
		if !strings.Contains(log, want) {
			t.Fatalf("docker operation %q missing from log %q", want, log)
		}
	}
	if strings.Contains(log, "system prune") || strings.Contains(log, "volume rm named") {
		t.Fatalf("cleanup escaped scoped resource set: %q", log)
	}
}
