package hash

import (
	"github.com/speps/go-hashids"
)

func (h *hash) HashidsEncode(params []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	hashNum, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	hashStr, err := hashNum.Encode(params)
	if err != nil {
		return "", err
	}
	return hashStr, nil
}

func (h *hash) HashidsDecode(hash string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	idsNum, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}
	ids, err := idsNum.DecodeWithError(hash)
	if err != nil {
		return ids, err
	}
	return ids, nil
}
