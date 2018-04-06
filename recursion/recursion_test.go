package recursion

func ExampleLookForKey1() {
	var mainBox = Box{
		Items: Items{
			Box{},
			Box{
				Items: Items{
					Box{},
					Box{},
				},
			},
			Box{
				Items: Items{
					Key(1),
				},
			},
			// Key(2),
		},
	}

	LookForKey1(mainBox)

	// Output: Found The Key!, 1
}

func ExampleLookForKey2() {
	var mainBox = Box{
		Items: Items{
			Box{},
			Box{
				Items: Items{
					Box{},
					Box{},
				},
			},
			Box{
				Items: Items{
					Key(1),
				},
			},
			// Key(2),
		},
	}

	LookForKey2(&mainBox)

	// Output: Found The Key!, 1
}

func ExampleCountdown() {
	Countdown(4)

	// Output: 4
	// 3
	// 2
	// 1
	// 0
}
