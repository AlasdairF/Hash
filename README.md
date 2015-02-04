## Hash

Herein are very highly optimized 32-bit and 64-bit Fnv-1a hashing functions. This implemention is faster than the native Go implementation and this particular algorithm has very few collisions while executing at very high speeds.

From my dictionary of 5,000,000 words (a-z only, 3-20 bytes each), the 32-bit hash gave 72 collisions and the 64-bit hash gave 2 collisions.
