package main

type Example struct {
	Title       string
	WasmFile    string
	Description string
}

type PageData struct {
	Examples []Example
}

var data = PageData{
	Examples: []Example{
		{
			Title:       "A Traditional Random Walk",
			WasmFile:    "example0.1.wasm",
			Description: "A traditional random walker will move in a random direction every frame.",
		},
		{
			Title:       "A Random Walk with a Tendency to walk Right",
			WasmFile:    "exercise0.1.wasm",
			Description: "A random walker with a tendency to walk to the right.",
		},
		{
			Title:       "A Uniform Random Distribution",
			WasmFile:    "example0.2.wasm",
			Description: "An example of uniform random distribution.",
		},
		{
			Title:       "A Random Walk with a Tendency to Follow the Mouse",
			WasmFile:    "exercise0.3.wasm",
			Description: "A Random Walk with a Tendency to Follow the Mouse.",
		},
		{
			Title:       "A Gaussian Distribution",
			WasmFile:    "example0.4.wasm",
			Description: "A visual example of Gaussian Distribution.",
		},
		{
			Title:       "A Normal Distribution With Mean and Standard Deviation",
			WasmFile:    "exercise0.4.wasm",
			Description: "A simulation of paint splatter.",
		},
		{
			Title:       "A Gaussian Walker",
			WasmFile:    "exercise0.5.wasm",
			Description: "A random walker where the walker's step size is determined by normal distribution of random numbers.",
		},
		{
			Title:       "An Accept Reject Distribution",
			WasmFile:    "example0.5.wasm",
			Description: "Visual representation of a custom distribution of random numbers.",
		},
		{
			Title:       "An Accept Reject Distribution Walker",
			WasmFile:    "exercise0.6.wasm",
			Description: "Visual representation of a custom distribution of random numbers when applied to a walker.",
		},
		{
			Title:       "A Perlin Noise Walker",
			WasmFile:    "example0.6.wasm",
			Description: "This is a random walker using perlin noise as its step direction.",
		},
		{
			Title:       "A Perlin Noise Walker (Perlin Step Size)",
			WasmFile:    "exercise0.7.wasm",
			Description: "This is a random walker using perlin noise as its step direction and step size.",
		},
	},
}
