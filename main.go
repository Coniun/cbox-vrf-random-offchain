package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func main() {

	nftWinnerList := []int{}

	// get seed from on-chain contract (etherscan)
	var contractRandomSeed = "29775269982517917835738859382887719068435719989660386033959329369793460616679"
	// update nftCount (for week 1, it's 465)
	var nftCount = 465

	// we are converting seed to md5 then int64
	md5Seed := md5.New()
	_, _ = io.WriteString(md5Seed, contractRandomSeed)
	var seed = binary.BigEndian.Uint64(md5Seed.Sum(nil))
	rand.Seed(int64(seed))


	// this ignore list is team/invester owned cboxes.
	ignoreList := []int{
		0,
		742,
		2478,
		2479,
		2654,
		3637,
		3826,
		3972,
		3973,
		4037,
		5046,
		5174,
		5217,
		5285,
		5476,
		5477,
		5478,
		5479,
		5480,
		5481,
		5482,
		5483,
		5484,
		5485,
		5486,
		5487,
		5488,
		5489,
		5490,
		5491,
		5492,
		5493,
		5494,
		5495,
		5496,
		5497,
		5498,
		5499,
		5500,
		5501,
		5502,
		5503,
		5504,
		5505,
		5506,
		5507,
		5508,
		5509,
		5510,
		5511,
		5512,
		5513,
		5514,
		5515,
		5516,
		5517,
		5518,
		5519,
		5520,
		5521,
		5522,
		5523,
		5524,
		5525,
		5526,
		5527,
		5528,
		5529,
		5530,
		5531,
		5532,
		5533,
		5534,
		5535,
		5536,
		5537,
		5538,
		5539,
		5540,
		5541,
		5542,
		5543,
		5544,
		5545,
		5546,
		5547,
		5548,
		5549,
		5550,
		5551,
		5552,
		5553,
		5554,
		5555,
		5556,
		5557,
		5558,
		5559,
		5560,
		5561,
		5562,
		5563,
		5564,
		5565,
		5566,
		5567,
		5568,
		5569,
		5570,
		5571,
		5572,
		5573,
		5574,
		5599,
		5600,
		5601,
		5602,
		5603,
		5604,
		5605,
		5606,
		5607,
		5608,
		5609,
		5610,
		5611,
		5612,
		5613,
		5614,
		5615,
		5616,
		5617,
		5618,
		5619,
		5620,
		5621,
		5622,
		5623,
		5624,
		5625,
		5626,
		5627,
		5628,
		5629,
		5630,
		5631,
		5632,
		5633,
		5634,
		5635,
		5636,
		5642,
		5643,
		5644,
		5645,
		5646,
		5647,
		5648,
		5745,
		5926,
		6104,
		6105,
		6106,
		6107,
		6108,
		6109,
		6110,
		6111,
		6112,
		6113,
		6114,
		6115,
		6116,
		6117,
		6120,
		6124,
		6125,
		6126,
		6127,
		6128,
		6129,
		6130,
		6131,
		6132,
		6133,
		6134,
		6135,
		6136,
		6137,
		6138,
		6139,
		6140,
		6141,
		6142,
		6143,
		6144,
		6145,
		6146,
		6147,
		6148,
		6149,
		6150,
		6151,
		6152,
		6153,
		6154,
		6155,
		6156,
		6157,
		6158,
		6159,
		6160,
		6161,
		6162,
		6163,
		6164,
		6277,
	}



	for len(nftWinnerList) < nftCount {
		// get random cbox id
		randBox := randInt(0, 6303)
		// add this cbox to winnerList if not ignored or not already won from this raffle
		if !contains(nftWinnerList, randBox) && !contains(ignoreList, randBox) {
			nftWinnerList = append(nftWinnerList, randBox)
		}
	}


	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(nftWinnerList)), ","), "[]"))



}


func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}

	return false
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}