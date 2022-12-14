package models

import "fmt"

const (
	SHAPE_SQUARE byte = iota
	SHAPE_STAR
)

type PreComputedOptions struct {
	ReversedDepthPruningPercentage int
	ReversedWidthPruningPercentage int
}

type Options struct {
	// Constraints
	TimeOut int64 // ms

	// Depth
	DepthMax               uint8
	DepthMin               uint8
	DepthPruningPercentage int

	// Width
	WidthMax               int
	WidthMultiThreading    bool
	WidthPruningPercentage int

	// Proximity
	ProximityRadius    uint8
	ProximityThreshold uint8
	ProximityShape     byte
	ProximityEvolution bool

	// Heuristics
	HeuristicAggro            int
	HeuristicDepthDivisor     int
	HeuristicAlignementWeight int
	HeuristicCaptureWeight    int
	HeuristicRandomWeight     int

	// Suspicion
	SuspicionRadius int

	// Pre computed options
	PreComputedOptions PreComputedOptions
}

func (o *Options) PreCompute() {
	o.PreComputedOptions.ReversedDepthPruningPercentage = 100 - o.DepthPruningPercentage
	o.PreComputedOptions.ReversedWidthPruningPercentage = 100 - o.WidthPruningPercentage
}

// Print prints options attributes
func (o *Options) Print() {
	fmt.Println("============================= [OPTIONS]")

	// Prints constraints
	fmt.Println("--------------- [constraints]")
	fmt.Println("timeout:\t", o.TimeOut, "ms")

	// Prints depth
	fmt.Println("--------------- [depth]")
	fmt.Println("max:\t", o.DepthMax)
	fmt.Println("min:\t", o.DepthMin)
	fmt.Println("pruning:\t", o.DepthPruningPercentage, " %")

	// Prints width
	fmt.Println("--------------- [width]")
	fmt.Println("max:\t\t", o.WidthMax)
	fmt.Println("multi-threading:\t", o.WidthMultiThreading)
	fmt.Println("pruning:\t\t", o.WidthPruningPercentage, " %")

	// Prints proximity
	fmt.Println("--------------- [proximity]")
	fmt.Println("radius:\t", o.ProximityRadius)
	fmt.Println("threshold:\t", o.ProximityThreshold)
	fmt.Println("shape:\t", o.ProximityShape)
	fmt.Println("evolution:\t", o.ProximityEvolution)

	// Prints heuristics
	fmt.Println("--------------- [heuristics]")
	fmt.Printf("aggro:\t\t\t%5d\n", o.HeuristicAggro)
	fmt.Printf("depth-divisor:\t\t%5d\n", o.HeuristicDepthDivisor)
	fmt.Printf("alignement:\t\t\t%2d/10\n", o.HeuristicAlignementWeight)
	fmt.Printf("capture:\t\t\t%2d/10\n", o.HeuristicCaptureWeight)
	fmt.Printf("random:\t\t\t%2d/10\n", o.HeuristicRandomWeight)

	// Prints suspicion
	fmt.Println("--------------- [suspicion]")
	if o.SuspicionRadius > 0 {
		fmt.Printf("radius:\t%2d\n", o.SuspicionRadius)
	} else {
		fmt.Println("radius:\tdisable")
	}
}
