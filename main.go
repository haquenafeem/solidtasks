package main

import "fmt"

// Color ...
type Color int

const (
	red Color = iota
	green
	blue
)

// Size ...
type Size int

const (
	small Size = iota
	medium
	large
)

// Product ...
type Product struct {
	name  string
	color Color
	size  Size
}

// Specification ...
type Specification interface {
	IsSatisfied(p *Product) bool
}

// ColorSpecification ...
type ColorSpecification struct {
	color Color
}

// IsSatisfied ...
func (spec ColorSpecification) IsSatisfied(p *Product) bool {
	return p.color == spec.color
}

// SizeSpecification ...
type SizeSpecification struct {
	size Size
}

// IsSatisfied ...
func (spec SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == spec.size
}

// MultipleSpecification ...
type MultipleSpecification struct {
	specsM map[int]Specification
}

// AddSpecification ...
func (specs MultipleSpecification) AddSpecification(index int, s Specification) {
	specs.specsM[index] = s
}

// IsSatisfied ...
func (specs MultipleSpecification) IsSatisfied(p *Product) bool {
	// return spec.first.IsSatisfied(p) &&
	// 	spec.second.IsSatisfied(p)

	for i := range specs.specsM {
		// fmt.Println(v)
		// fmt.Println(specs.specsM[i].IsSatisfied(p))
		if !specs.specsM[i].IsSatisfied(p) {
			return false
		}
	}
	return true
}

// BetterFilter ...
type BetterFilter struct{}

// Filter ...
func (f *BetterFilter) Filter(
	products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", green, large}

	products := []Product{apple, tree, house}

	greenSpec := ColorSpecification{green}
	largeSpec := SizeSpecification{large}

	mSpecification := MultipleSpecification{make(map[int]Specification)}
	mSpecification.AddSpecification(1, greenSpec)
	mSpecification.AddSpecification(7, largeSpec)

	// mSpecification.IsSatisfied(&products[1])
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, mSpecification) {
		fmt.Printf(" - %s is green and large\n", v.name)
	}
}
