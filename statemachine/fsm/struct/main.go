package main

import (
	"fmt"

	"github.com/looplab/fsm"
)

const (
	CreatePending      = "CreatePending"
	SSPDCreated        = "SSPDCreated"
	NotConfirmed       = "NotConfirmed"
	FailedSSPDCreate   = "FailedSSPDCreate"
	FailedSqillsCreate = "FailedSqillsCreate"

	CreateBookingEvent = "createBooking"
)

type Booking struct {
	Number string
	Status string
}

func NewBooking(number string) *Booking {
	return &Booking{
		Number: number,
		Status: CreatePending,
	}
}

type BookingEventDispatcher struct {
	fsm *fsm.FSM
}

func NewBookingEventDispatcher() *BookingEventDispatcher {
	eh := BookingEventDispatcher{}

	// can be moved to SSPDBookingCreator
	createSSPD := func(e *fsm.Event) {
		fmt.Println("create SSPD called")
		booking := e.Args[0].(*Booking)

		switch e.FSM.Current() {
		case CreatePending:
			// emulate connection error
			e.Err = fmt.Errorf("connection error")
			booking.Status = FailedSSPDCreate
			return
		default:
			// emulate success creation
			booking.Status = SSPDCreated
		}
	}

	// can be moved to SqillsBookingCreator
	createSqills := func(e *fsm.Event) {
		fmt.Println("create Sqills called")
		booking := e.Args[0].(*Booking)
		booking.Status = NotConfirmed
	}

	eh.fsm = fsm.NewFSM(
		CreatePending,
		fsm.Events{
			{Name: CreateBookingEvent, Src: []string{CreatePending, FailedSSPDCreate}, Dst: SSPDCreated},
			{Name: CreateBookingEvent, Src: []string{SSPDCreated, FailedSqillsCreate}, Dst: NotConfirmed},
		},
		fsm.Callbacks{
			"leave_" + CreatePending:      createSSPD,
			"leave_" + FailedSSPDCreate:   createSSPD,
			"leave_" + SSPDCreated:        createSqills,
			"leave_" + FailedSqillsCreate: createSqills,
		},
	)

	return &eh
}

func (ed *BookingEventDispatcher) execute(booking *Booking, event string) {
	ed.fsm.SetState(booking.Status)
	fmt.Printf("booking status: %s\n", booking.Status)

	if err := ed.fsm.Event(event, booking); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}

func main() {
	fmt.Println("example of processing new booking")

	dispatcher := NewBookingEventDispatcher()
	booking := NewBooking("new booking")
	// dispatching also can be stopped if max attempts count
	for booking.Status != NotConfirmed {
		dispatcher.execute(booking, CreateBookingEvent)
	}

	fmt.Printf("booking status: %s\n", booking.Status)

	fmt.Println("")
	fmt.Println("example of processing existing booking")

	existing := &Booking{Number: "existing booking", Status: FailedSqillsCreate}
	for booking.Status != NotConfirmed {
		dispatcher.execute(existing, CreateBookingEvent)
	}

	fmt.Printf("booking status: %s\n", existing.Status)
}
