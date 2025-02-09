package graphs

// A type representing the weight/cost of an edge

// use the below if you want to use the exp package:
// import "golang.org/x/exp/constraints"
// type EdgeCostExp interface {
// constraints.Integer | constraints.Float
// }

type EdgeCost interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}
