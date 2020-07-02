# File Integrity Check

## Generate a Hash from a file to verify its integrity

Execute it like this: (all flags can also be left out)
`file-integrity-check -byteCount=8 -seed=0 path/to/file`

- The hash function is calculated by a randomized algorithm which is correct (1/2)^(byteCount*8) of the time (the higher the byteCount, the higher the probability).
- Since the file size is not a part of the probabilty, this algorithm will work equally good on small, average and large file sizes (only the execution time goes up)
- This algorithm is implemented to have a runtime of O(poly(n)), which means it can be calculated efficiently
- It should however not be used for a usecase where security is important, since when the seed is known, you can easily find out which bits are checked and which are not, which means that you would have the knowledge on which bits to modify without the hash changing.
