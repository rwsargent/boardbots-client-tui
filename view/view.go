package view

type (
	// X and Y for placement on the terminal
	View struct {
		X, Y int
	}

	/**
	 * Set the required cells on the terminal window.
	 */
	Drawable interface {
		Draw()
	}

	/**
	 * Dimensional information for a view.
	 */
	Widget interface {
		Width() int
		Height() int
	}
)