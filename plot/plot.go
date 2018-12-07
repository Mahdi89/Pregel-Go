package main

import (
	ps "golang.org/x/tools/benchmark/parse"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"math/rand"
	"os"
	"strconv"
)

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}

func main() {

	f, err := os.Open("../bench.out")
	if err != nil {
		panic(err)
	}
	set, err := ps.ParseSet(f)
	if err != nil {
		panic(err)
	}

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Performance Graph"
	p.X.Label.Text = "Graph Size"
	p.Y.Label.Text = "Operations/s"

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	pts0 := make(plotter.XYs, n)
	pts1 := make(plotter.XYs, n)
	pts2 := make(plotter.XYs, n)

	for i := 0; i < n; i++ {

		pts0[i].X = float64(i)
		pts1[i].X = float64(i)
		pts2[i].X = float64(i)
	}

	for i := 0; i < n; i++ {

		pts0[i].Y = float64(set["BenchmarkPageRank_Pregel-4"][i].NsPerOp)
		pts1[i].Y = float64(set["BenchmarkPageRank_Matrix-4"][i].NsPerOp)
		pts2[i].Y = float64(set["BenchmarkPageRank_Stream-4"][i].NsPerOp)

	}

	err = plotutil.AddLinePoints(p,
		"Pregel", pts0,
		"Matrix", pts1,
		"Stream", pts2)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(16*vg.Inch, 16*vg.Inch, "bench.out.png"); err != nil {
		panic(err)
	}

}
