package flameutil

import "flag"

var FLAG_ITERATIONS = flag.Int("iterations", 1000000, "Number of iterations per image")

type FlameConfig struct {
	iterations     int
	variations     []Variation
	flameFunctions []FlameFunction
}

func NewFlameConfig() *FlameConfig {
	return &FlameConfig{*FLAG_ITERATIONS, []Variation{}, []FlameFunction{}}
}

func (config *FlameConfig) AddFlameFunction(prob float64, function func(*Point)) *FlameConfig {
	config.flameFunctions = append(config.flameFunctions,
		FlameFunction{prob, function})
	return config
}

func (config *FlameConfig) AddVariation(weight float64, function func(*Point)) *FlameConfig {
	config.variations = append(config.variations,
		Variation{weight, function})
	return config
}
