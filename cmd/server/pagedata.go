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
	},
}
