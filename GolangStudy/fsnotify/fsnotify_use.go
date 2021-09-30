package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {

	// Create a File watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// A channel just used for keeping rolling indefinitely
	done := make(chan bool)

	// The go routine below is running an infinite loop in order to continue monitoring the fs watcher
	// (the only way it'll stop is when the caller main function will end)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("--> event:", event)
				log.Println("   > event.Op:", event.Op)
				log.Println("   > fsnotify.Write:", fsnotify.Write)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("--> modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("--> error:", err)
			}
		}
	}()

	// Adding a file to the fs watcher
	// (It is possible to use the following command several times to add several files to the watcher)
	err = watcher.Add("/tmp/foo")
	if err != nil {
		log.Fatal(err)
	}

	// Look at the script: We're never writing anything to "done" channel
	//   => This script will never end as with the following command it is waiting for something to come from the channel (something which will never come)
	//   => The go-routine above will never stop (for-loop is infinite in order to continue monitoring the file "/tmp/foo" indefinitely) as long as the caller is running
	<-done
}
