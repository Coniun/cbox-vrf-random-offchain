# CBOX vrf random offchain


This code will generate an array of winners like:

```js
winners = [105, 503, 203, ...]
```

We also have an array of nfts like:

```js
nfts = ['CONTRACT_ADDRESS0-TOKEN_ID0', "CONTRACT_ADDRESS1-TOKEN_ID1", 'CONTRACT_ADDRESS2-TOKEN_ID2']

// we are sorting nft's with alphabetical order
nfts = sort(nfts)
```

After you have both list; the cbox which is located on `index 0` of `winners` array will eligible for `index 0` of `nfts` array.

So for this example, winners and nft's will be distrubuted like below:
```
CBOX105 -> CONTRACT_ADDRESS0 # 0
CBOX503 -> CONTRACT_ADDRESS1 # 1
CBOX203 -> CONTRACT_ADDRESS2 # 2
``` 


## How to obtain 'contractRandomSeed'

Every week we will use our contract (verified on etherscan, check links below) and create a CBOX RANDOM SEED for that week.

This transaction cannot be manipulated and it uses [ChainLink VRF](https://docs.chain.link/docs/chainlink-vrf/).

You can use Etherscan to obtain weekly random seed.

## Notes

We believe that, this system %100 transparent and can be noterized by everyone.



## Links

- CBOXRandomSeedGenerator contract (TBA)