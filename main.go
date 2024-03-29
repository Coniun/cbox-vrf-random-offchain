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
	var nftCount = 10

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
		1232,1591,2674,5807,3888,5855,5776,879,1524,2613,4702,3550,4403,2361,5,437,1197,2007,2910,5845,4271,5011,136,870,1537,2623,5733,3726,451,1497,472,1486,2040,6004,3457,4279,5749,291,
		6295,6293,5825,5752,4670,3981,2817,2188,1685,1637,1571,1134,1038,659,154,4739,1547,1548,
		2756,48,4039,713,5042,1324,3673,622,5028,5920,3112,189,4431,1049,5092,1692,3754,712,5029,5959,3130,231,4467,5763,
		// - end of partner passes
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
