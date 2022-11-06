package main

import (
	"context"
	"fmt"

	"github.com/qmuntal/stateless"
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
	sm *stateless.StateMachine
}

func NewBookingEventDispatcher(booking *Booking) *BookingEventDispatcher {
	eh := BookingEventDispatcher{}

	// can be moved to SSPDBookingCreator
	createSSPD := func(ctx context.Context, args ...interface{}) error {
		fmt.Printf("create SSPD called, args: %v\n", args)
		//booking := args[1].(*Booking)

		switch args[0].(string) {
		case CreatePending:
			// emulate connection error

			// booking.Status = FailedSSPDCreate
			return fmt.Errorf("connection error")
		default:
			// emulate success creation
			// booking.Status = SSPDCreated
			return nil
		}
	}

	// can be moved to SqillsBookingCreator
	createSqills := func(ctx context.Context, args ...interface{}) error {
		fmt.Println("create Sqills called")
		//booking := e.Args[0].(*Booking)
		//booking.Status = NotConfirmed

		return nil
	}

	eh.sm = stateless.NewStateMachineWithExternalStorage(func(_ context.Context) (stateless.State, error) {
		return booking.Status, nil
	}, func(_ context.Context, state stateless.State) error {
		booking.Status = state.(string)
		return nil
	}, stateless.FiringQueued)

	eh.sm.Configure(CreatePending).Permit(CreateBookingEvent, SSPDCreated).OnEntry(createSSPD)
	eh.sm.Configure(FailedSSPDCreate).Permit(CreateBookingEvent, SSPDCreated).OnEntry(createSSPD)
	eh.sm.Configure(SSPDCreated).Permit(CreateBookingEvent, NotConfirmed).OnEntry(createSqills)
	eh.sm.Configure(FailedSqillsCreate).Permit(CreateBookingEvent, NotConfirmed).OnEntry(createSqills)

	return &eh
}

func (ed *BookingEventDispatcher) execute(booking *Booking, event string) {
	//ed.sm.SetState(booking.Status)
	fmt.Printf("booking status: %s\n", booking.Status)

	if err := ed.sm.Fire(event, booking); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}

func main() {
	fmt.Println("example of processing new booking")

	booking := NewBooking("new booking")

	dispatcher := NewBookingEventDispatcher(booking)
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
