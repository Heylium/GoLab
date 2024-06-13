package client

// setup go path to replace import path
import (
	"goLab/go_learning/src/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))

	t.Log(series.Square(2))
}
