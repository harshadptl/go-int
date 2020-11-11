package goint

type Uint128 struct {
	
	High uint64

	Low  uint64
}

func NewInt128(n uint) *Uint128 {

	return &Uint128{
		High: 0,
		Low:  uint64(n),
	}
}

func (n *Uint128) Add(a *Uint128) error {

	var carry uint64 = 0

	// add the lows
	low := n.Low + a.Low

	// check for overflow
	if low < n.Low ||
		low < a.Low {

			carry = 1
		}

	// add the highs
	high := n.High + a.High + carry

	n.Low = low
	n.High = high

	// TODO: handle high overflow
	return nil
}

func (n *Uint128) AddUint64(a uint64) error {
	var carry uint64 = 0
	
	// add the lows
	low := n.Low + a

	// check for overflow
	if low < n.Low ||
		low < a {

			carry = 1
		}

	// add the highs
	high := n.High + carry

	n.Low = low
	n.High = high

	// TODO: handle high overflow
	return nil
}
