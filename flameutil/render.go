package flameutil

import "math/rand"

/*
	Play's the chaos game (using the provided config) on the provided image.
*/
var Render = func(config *FlameConfig, image *FlameImage) {
	var rng = rand.New(rand.NewSource(0))

	// Staring point selected randomly from bi-unit square
	point := NewPoint(rng.Float64()*2-1, rng.Float64()*2-1)

	// Run the game for a few iterations before actually drawing
	// This gives our point time to get closer into the 'solution'
	for iter := 0; iter < 20; iter++ {
		step(config, randomFlameFuncNumber(config, rng.Float64()), point)
	}

	// Draw the fractal!
	for iter := 0; iter < config.iterations; iter++ {
		id := randomFlameFuncNumber(config, rng.Float64())
		step(config, id, point)
		// Coloring based on id should happen here
		image.addPoint(point /*, color */)
	}
}

func randomFlameFuncNumber(config *FlameConfig, rnd float64) int {
	// Get a function index from the config's flameFunctions
	// This has problems such as: if the probabilities don't add to 1,
	// it can panic from array out of bound.
	// if the sum of probabilities is more than 1, transformations after that
	// wont ever get picked.
	// Using integer weights (think weighted raffle tickets) would solve this.
	j := 0
	for ; rnd < 1; j++ {
		rnd += config.flameFunctions[j].probability
	}
	return j - 1
}

/*
 * Maps a point through the flame function, and weighted variations.
 */
func step(config *FlameConfig, flameNo int, point *Point) {
	config.flameFunctions[flameNo].transformation(point) // map the point.

	// map through each variation, compute a weighted average of the variations
	tmp_point := NewPoint(0, 0) // Weighted point var
	input_point := point.Copy() // Copy of the input we can mutate
	for v := range config.variations {
		w := config.variations[v].weight
		t := config.variations[v].transformation

		t(input_point) // mutates input_point

		// add input_point into tmp_point (weighted)
		tmp_point.X += input_point.X * w
		tmp_point.Y += input_point.Y * w

		// move input_point back to point, to prep for next variation.
		input_point.X = point.X
		input_point.Y = point.Y
	}

	// Done computing weighted location. assign it to point.
	point.X = tmp_point.X
	point.Y = tmp_point.Y
}
