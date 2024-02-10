// based on [pageant](https://github.com/kbolino/pageant) of Kristian Bolino
package main

import (
	"log"
	"os"

	sshagent "github.com/xanzy/ssh-agent"
	"golang.org/x/crypto/ssh"
)

// This example requires all of the following to work:
//   - environment variable PAGEANT_TEST_SSH_ADDR is set to a valid SSH
//     server address (host:port)
//   - environment variable PAGEANT_TEST_SSH_USER is set to a user name
//     that the SSH server recognizes
//   - Pageant is running on the local machine
//   - Pageant has a key that is authorized for the user on the server
func main() {
	sshAgent, pageantConn, err := sshagent.New()
	if err != nil {
		log.Fatalf("error on New: %s", err)
	}
	defer pageantConn.Close()
	keys, err := sshAgent.List()
	if err != nil {
		log.Fatalf("error on agent.List: %s", err)
	}
	if len(keys) == 0 {
		log.Fatalf("no keys listed by Pagent")
	}
	for i, key := range keys {
		log.Printf("key %d: %s %s\n", i, key.Comment, ssh.FingerprintSHA256(key))
	}

	signers, err := sshAgent.Signers()
	if err != nil {
		log.Fatalf("cannot obtain signers from SSH agent: %s", err)
	}
	sshUser := os.Getenv("PAGEANT_TEST_SSH_USER")
	config := ssh.ClientConfig{
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signers...)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		User:            sshUser,
	}
	sshAddr := os.Getenv("PAGEANT_TEST_SSH_ADDR")
	sshConn, err := ssh.Dial("tcp", sshAddr, &config)
	if err != nil {
		log.Fatalf("failed to connect to %s@%s due to error: %s", sshUser, sshAddr, err)
	}
	sshConn.Close()
}
