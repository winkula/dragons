package model

type rule (func(*World) bool)

var rules = []rule{
	// A dragon can not have other dragons in his territory.
	func(w *World) bool {
		// TODO
		return true
	},
	// A dragon can not be surrounded by more than 50% fire (only counting fields on the grid).
	func(w *World) bool {
		// TODO
		return true
	},
	// Overlapping territories must be fire.
	func(w *World) bool {
		// TODO
		return true
	},
}

// ValidateWorld validates, if a world is in a valid state.
func ValidateWorld(w *World) bool {
	for _, rule := range rules {
		valid := rule(w)
		if !valid {
			return false
		}
	}
	return true
}
