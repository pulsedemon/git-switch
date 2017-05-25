package main

import "testing"

func TestGit_getCurrentGitUser(t *testing.T) {
  currentUserName, currentUserEmail := getCurrentGitUser()

  if len(currentUserName) == 0 {
    t.Errorf("git username is not set")
  }

  if len(currentUserEmail) == 0 {
    t.Errorf("git email is not set")
  }
}
