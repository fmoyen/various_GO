package main


import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
	// Go signal notification works by sending os.Signal values on a channel. We’ll create a channel to receive these notifications (we’ll also make one to notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// signal.Notify registers the given channel to receive notifications of the specified signals.
	/*
	Nom	Signification					Comportement
	sighup	Hang-up (fin de connexion)			Terminaison
	sigint	Interruption (ctrl-C)				Terminaison
	sigquit	Interruption forte (ctrl-\)			Terminaison + core dump
	sigfpe	Erreur arithmétique (division par zéro)		Terminaison + core dump
	sigkill	Interruption très forte (ne peut être ignorée)	Terminaison
	sigsegv	Violation des protections mémoire		Terminaison + core dump
	sigpipe	Écriture sur un tuyau sans lecteurs		Terminaison
	sigalrm	Interruption d’horloge				Ignoré
	sigtstp	Arrêt temporaire d’un processus (ctrl-Z)	Suspension
	sigcont	Redémarrage d’un processus arrêté		Ignoré
	sigchld	Un des processus fils est mort ou a été arrêté	Ignoré
	*/
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for signals. When it gets one it’ll print it out and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println("Type of signal received: ",sig)
		done <- true
	}()

	// The program will wait here until it gets the expected signal (as indicated by the goroutine above sending a value on done) and then exit.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("Signal received --> exiting")
}
