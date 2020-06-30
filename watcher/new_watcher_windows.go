// +build windows

package watcher

func NewWatcher(absoluteDirectory string) Watcher {
	return NewINotifyWatcher(absoluteDirectory)
}
