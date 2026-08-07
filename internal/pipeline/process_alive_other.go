//go:build !unix && !windows

package pipeline

func processAlive(int) bool { return false }
