package mycalendar2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type booking struct {
	start  int
	end    int
	expect bool
	desc   string
}

func TestBook(t *testing.T) {
	tests := []struct {
		name     string
		bookings []booking
	}{
		{
			name: "ExampleFromProblem",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{50, 60, true, "book(50, 60) should return true"},
				{10, 40, true, "book(10, 40) should return true (double booking allowed)"},
				{5, 15, false, "book(5, 15) should return false (would cause triple booking)"},
				{5, 10, true, "book(5, 10) should return true"},
				{25, 55, true, "book(25, 55) should return true"},
			},
		},
		{
			name: "SingleBookings",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{30, 40, true, "book(30, 40) should return true"},
				{50, 60, true, "book(50, 60) should return true"},
			},
		},
		{
			name: "DoubleBookings",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{15, 25, true, "book(15, 25) should return true (double booking allowed)"},
			},
		},
		{
			name: "TripleBookings_Rejected",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{15, 25, true, "book(15, 25) should return true"},
				{18, 22, false, "book(18, 22) should return false (triple booking)"},
			},
		},
		{
			name: "AdjacentBookings",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{20, 30, true, "book(20, 30) should return true (adjacent, no overlap)"},
			},
		},
		{
			name: "ExactOverlap",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{10, 20, true, "book(10, 20) again should return true (double booking allowed)"},
				{10, 20, false, "book(10, 20) third time should return false (triple booking)"},
			},
		},
		{
			name: "PartialOverlaps",
			bookings: []booking{
				{10, 30, true, "book(10, 30) should return true"},
				{5, 20, true, "book(5, 20) should return true (double booking)"},
				{15, 35, false, "book(15, 35) should return false (triple booking in [15, 20))"},
			},
		},
		{
			name: "ContainedBooking",
			bookings: []booking{
				{10, 30, true, "book(10, 30) should return true"},
				{15, 25, true, "book(15, 25) should return true (double booking)"},
				{18, 22, false, "book(18, 22) should return false (triple booking)"},
			},
		},
		{
			name: "OverlappingAtStart",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{10, 15, true, "book(10, 15) should return true (double booking)"},
				{10, 12, false, "book(10, 12) should return false (triple booking)"},
			},
		},
		{
			name: "OverlappingAtEnd",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{15, 20, true, "book(15, 20) should return true (double booking)"},
				{18, 20, false, "book(18, 20) should return false (triple booking)"},
			},
		},
		{
			name: "MultipleNonOverlapping",
			bookings: []booking{
				{1, 5, true, "book(1, 5) should return true"},
				{10, 15, true, "book(10, 15) should return true"},
				{20, 25, true, "book(20, 25) should return true"},
				{30, 35, true, "book(30, 35) should return true"},
			},
		},
		{
			name: "ComplexScenario",
			bookings: []booking{
				{10, 20, true, "book(10, 20) should return true"},
				{15, 25, true, "book(15, 25) should return true (overlaps with first: [15, 20))"},
				{30, 40, true, "book(30, 40) should return true"},
				{35, 45, true, "book(35, 45) should return true (overlaps with third: [35, 40))"},
				{16, 19, false, "book(16, 19) should return false (triple booking)"},
				{36, 39, false, "book(36, 39) should return false (triple booking)"},
				{12, 14, true, "book(12, 14) should return true (only overlaps with first)"},
			},
		},
		{
			name: "BoundaryConditions_MinimumValues",
			bookings: []booking{
				{0, 1, true, "book(0, 1) should return true"},
				{0, 2, true, "book(0, 2) should return true (double booking)"},
				{0, 1, false, "book(0, 1) again should return false (triple booking)"},
			},
		},
		{
			name: "BoundaryConditions_LargeValues",
			bookings: []booking{
				{1000000000 - 10, 1000000000 - 5, true, "book with large values should return true"},
				{1000000000 - 8, 1000000000 - 3, true, "book with large values should return true"},
				{1000000000 - 7, 1000000000 - 4, false, "book with large values should return false (triple booking)"},
			},
		},
		{
			name: "CustomEventSequence",
			bookings: []booking{
				{28, 46, true, "book(28, 46) should return true"},
				{9, 21, true, "book(9, 21) should return true"},
				{21, 39, true, "book(21, 39) should return true"},
				{37, 48, false, "book(37, 48) should return false"},
				{38, 50, false, "book(38, 50) should return false"},
				{22, 39, false, "book(22, 39) should return false"},
				{45, 50, true, "book(45, 50) should return true"},
				{1, 12, true, "book(1, 12) should return true"},
				{40, 50, false, "book(40, 50) should return false"},
				{31, 44, false, "book(31, 44) should return false"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cal := Constructor()
			for _, b := range tt.bookings {
				result := cal.Book(b.start, b.end)
				if b.expect {
					assert.True(t, result, b.desc)
				} else {
					assert.False(t, result, b.desc)
				}
			}
		})
	}
}
