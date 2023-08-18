package calculatePack

type Searcher interface {
	FindLargestByElem(elem int) int
}

type Calculator interface {
	CalculatePacks(order int) map[int]int
}

type BinarySearcher struct {
	sizes []int
}

func NewBinarySearcher(sizes []int) *BinarySearcher {
	return &BinarySearcher{
		sizes: sizes,
	}
}

func (bs *BinarySearcher) FindLargestByElem(elem int) int {
	left, right := 0, len(bs.sizes)-1
	largestFittingPackIndex := -1

	for left <= right {
		mid := left + (right-left)/2
		if bs.sizes[mid] <= elem {
			largestFittingPackIndex = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return largestFittingPackIndex
}

type PackCalculator struct {
	searcher Searcher
	pkgs     []int
}

func NewPackCalculator(pkgs []int, searcher Searcher) *PackCalculator {
	return &PackCalculator{
		searcher: searcher,
		pkgs:     pkgs,
	}
}

func (pc *PackCalculator) CalculatePacks(order int) map[int]int {
	packCounts := make(map[int]int)

	if len(pc.pkgs) > 1 && order > pc.pkgs[0] && order < pc.pkgs[1] {
		packCounts[pc.pkgs[1]] = 1
		return packCounts
	}

	for order > 0 {
		largestFittingPackIndex := pc.searcher.FindLargestByElem(order)

		if largestFittingPackIndex != -1 {
			packSize := pc.pkgs[largestFittingPackIndex]
			packCount := order / packSize
			packCounts[packSize] = packCount
			order %= packSize
		} else {
			packCounts[pc.pkgs[0]] += 1
			break
		}

	}

	return packCounts
}
