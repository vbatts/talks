// +build ignore

package main

// START1 OMIT
const (
	DeviceCreate TaskType = iota
	DeviceReload
	DeviceRemove
	DeviceRemoveAll
	DeviceSuspend
	DeviceResume
	DeviceInfo
	DeviceDeps
)

// STOP1 OMIT

type TaskType int

func main() {
}
