package main

import "testing"

type fakeUpdaterRunner struct {
	runCalls     int
	runOnceCalls int
	lastOptions  CheckOptions
	result       CheckResult
}

func (u *fakeUpdaterRunner) Run() {
	u.runCalls++
}

func (u *fakeUpdaterRunner) RunOnce(options CheckOptions) CheckResult {
	u.runOnceCalls++
	u.lastOptions = options
	return u.result
}

func TestExecuteRunDefaultMode(t *testing.T) {
	updater := &fakeUpdaterRunner{}

	exitCode := executeRun(updater, false)
	if exitCode != 0 {
		t.Fatalf("exitCode = %d, want 0", exitCode)
	}
	if updater.runCalls != 1 {
		t.Fatalf("runCalls = %d, want 1", updater.runCalls)
	}
	if updater.runOnceCalls != 0 {
		t.Fatalf("runOnceCalls = %d, want 0", updater.runOnceCalls)
	}
}

func TestExecuteRunForceModeSuccess(t *testing.T) {
	updater := &fakeUpdaterRunner{
		result: CheckResult{},
	}

	exitCode := executeRun(updater, true)
	if exitCode != 0 {
		t.Fatalf("exitCode = %d, want 0", exitCode)
	}
	if updater.runCalls != 0 {
		t.Fatalf("runCalls = %d, want 0", updater.runCalls)
	}
	if updater.runOnceCalls != 1 {
		t.Fatalf("runOnceCalls = %d, want 1", updater.runOnceCalls)
	}
	if !updater.lastOptions.Force {
		t.Fatal("CheckOptions.Force = false, want true")
	}
}

func TestExecuteRunForceModeFailure(t *testing.T) {
	updater := &fakeUpdaterRunner{
		result: CheckResult{Failures: 1},
	}

	exitCode := executeRun(updater, true)
	if exitCode != 1 {
		t.Fatalf("exitCode = %d, want 1", exitCode)
	}
	if updater.runOnceCalls != 1 {
		t.Fatalf("runOnceCalls = %d, want 1", updater.runOnceCalls)
	}
}
