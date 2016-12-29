package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/shanemhansen/countish"
)

var support, errorTolerance, failureProb, threshold float64

var impl string // naive/sticky/lossy

func main() {
	flag.Float64Var(&support, "support", 0.0001, "support")
	flag.Float64Var(&errorTolerance, "error-tolerance", 0.0001, "tolerance")
	flag.Float64Var(&failureProb, "failureProb", 0.0001, "failure probability (delta)")
	flag.Float64Var(&threshold, "threshold", 0.05, "Frequency threshold")
	flag.StringVar(&impl, "impl", "sticky", "counting implementation")
	flag.Parse()

	var s countish.Counter
	switch impl {
	case "sticky":
		s = countish.NewSampler(support, errorTolerance, failureProb)
	case "naive":
		s = countish.NewNaiveSampler()
	case "lossy":
		s = countish.NewLossyCounter(support, errorTolerance)
	default:
		panic("unknown sampler")
	}
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		i++
		s.Observe(scanner.Text())
	}
	vals := s.ItemsAboveThreshold(threshold)
	for _, val := range vals {
		fmt.Printf("%f %s\n", val.Frequency, val.Key)
	}
}
