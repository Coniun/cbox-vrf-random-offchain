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

	// CBOXRandomSeedGenerator: https://etherscan.io/address/0xaf8BFFf3962E49afaEA9e49BbaFAb57F4daa77E0#readContract
	var contractRandomSeed = "GET_SEED_FROM_CBOXRandomSeedGenerator_CONTRACT"
	var nftCount = 16

	// we are converting seed to md5 then int64
	md5Seed := md5.New()
	_, _ = io.WriteString(md5Seed, contractRandomSeed)
	var seed = binary.BigEndian.Uint64(md5Seed.Sum(nil))
	rand.Seed(int64(seed))


	// this ignore list is team/investor owned cboxes.
	// this list can be updated (please watch this repo)
	ignoreList := []int{
		// partner passes, excluded
		981, 4402, 1744, 1796, 1089, 2693, 530, 1743, 3560, 3257, 2360, 3608, 3939, 1877, 754, 3663, 195, 4102, 1476,
		2422, 3874, 1088, 5823, 4026, 3281, 1404, 1482, 3015, 3722, 4539, 1224, 3479, 1892, 5824, 246, 2311, 2558,
		// - end of partner passes
		// cbox coniun passes excluded (only for current week)
		5931, 3100, 2880, 5408, 3214, 2475, 4736, 5643, 136,
		// - end cbox coniun passes
		0,
		678,
		679,
		742,
		2478,
		2479,
		2654,
		3217,
		3637,
		3826,
		3972,
		3973,
		4037,
		5714,
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
		5646,
		5745,
		5926,
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
		// get random cbox id between 0, 6302
		randBox := randInt(0, 6302)
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
