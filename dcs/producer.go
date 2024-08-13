package dcs

import "mpp/layout"

func ProducerFac(in chan *layout.Input) func(put func(mp *layout.MpData)) error {
	return func(put func(mp *layout.MpData)) error {
		for r := range in {
			product := &layout.MpData{
				Input:  r,
				Output: &layout.Output{},
			}
			put(product)
		}
		return nil
	}
}
