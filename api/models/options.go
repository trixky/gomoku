package models

import "fmt"

const (
	SHAPE_SQUARE byte = iota
	SHAPE_STAR
)

type Options struct {
	// Constraints
	Timeout uint16 // ms

	// Depth
	DepthMax     uint8
	DepthMin     uint8
	DepthPruning bool

	// Proximity
	ProximityRadius    uint8
	ProximityThreshold uint8
	ProximityShape     byte
	ProximityEvolution bool

	// Heuristics
	HeuristicsPotentialAlignement uint8
	HeuristicsWinByAlignement     uint8
	HeuristicsPotentialCapture    uint8
	HeuristicsCapture             uint8
	HeuristicsRandom              uint8
}

// Print prints options attributes
func (o *Options) Print() {
	fmt.Println("============================= [OPTIONS]")

	// Prints constraints
	fmt.Println("--------------- [constraints]")
	fmt.Println("timeout:\t", o.Timeout, "ms")

	// Prints depth
	fmt.Println("--------------- [depth]")
	fmt.Println("max:\t", o.DepthMax)
	fmt.Println("min:\t", o.DepthMin)
	fmt.Println("pruning:\t", o.DepthPruning)

	// Prints proximity
	fmt.Println("--------------- [proximity]")
	fmt.Println("radius:\t", o.ProximityRadius)
	fmt.Println("threshold:\t", o.ProximityThreshold)
	fmt.Println("shape:\t", o.ProximityShape)
	fmt.Println("evolution:\t", o.ProximityEvolution)

	// Prints heuristics
	fmt.Println("--------------- [heuristics]")
	fmt.Printf("potential alignement:\t%2d/10\n", o.HeuristicsPotentialAlignement)
	fmt.Printf("win by alignement:\t\t%2d/10\n", o.HeuristicsWinByAlignement)
	fmt.Printf("potential capture:\t\t%2d/10\n", o.HeuristicsPotentialCapture)
	fmt.Printf("capture:\t\t\t%2d/10\n", o.HeuristicsRandom)
	fmt.Printf("random:\t\t\t%2d/10\n", o.HeuristicsRandom)
}
