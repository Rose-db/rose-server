package roseServer

import "rose/rose"

type sdk struct {
	Rose *rose.Rose
}

func newSdk() (*sdk, rose.Error) {
	r, err := rose.New(true)

	if err != nil {
		return nil, err
	}

	return &sdk{Rose: r}, nil
}

func (s *sdk) Write(id string, data []uint8) {

}

func (s *sdk) Read(id string) {

}

func (s *sdk) Delete(id string) {

}
