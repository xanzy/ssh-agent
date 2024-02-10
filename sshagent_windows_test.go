// based on [pageant](https://github.com/kbolino/pageant) of Kristian Bolino

package sshagent

import (
	"testing"
)

// Pageant must be running for this test to work.
func TestNew(t *testing.T) {
	_, conn, err := New()
	if err != nil {
		t.Fatalf("error on New: %s", err)
	} else if conn == nil {
		t.Fatalf("New returned nil")
	}
	err = conn.Close()
	if err != nil {
		t.Fatalf("error on Conn.Close: %s", err)
	}
}

// Pageant must be running and have at least 1 key loaded for this test to work.
func TestSSHAgentList(t *testing.T) {
	sshAgent, conn, err := New()
	if err != nil {
		t.Fatalf("error on New: %s", err)
	}
	defer conn.Close()
	keys, err := sshAgent.List()
	if err != nil {
		t.Fatalf("error on agent.List: %s", err)
	}
	if len(keys) == 0 {
		t.Fatalf("no keys listed by Pagent")
	}
	for i, key := range keys {
		t.Logf("key %d: %s", i, key.Comment)
	}
}
