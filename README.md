# Bloom Filter Implementation in Go

## Bloom Filters - No means No, Yes means Maybe

A Bloom Filter is a space-efficient probabilistic data structure used for membership query operations in sets. It uses multiple hash functions and a bit array to determine if an element is possibly in the set or definitely not in the set. Bloom filters have applications in various fields where fast set membership checks are required, such as caching, spell checkers, and network routers.

## Hash Function for Bloom Filters

Hash functions are a critical component of Bloom filters, and choosing the right one is essential for optimal performance. Instead of using cryptographic hash functions like MD5, which are slow, Bloom filters usually use faster non-cryptographic hash functions like [MurmurHash](https://en.wikipedia.org/wiki/MurmurHash) or [FNV](https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function).

Independence of hash functions is crucial to ensure accurate results. It refers to the property that the output of hash functions should be as uncorrelated as possible. If hash functions aren't independent, there's a high chance they'll produce the same result for a given input, leading to more collisions and false positives.

In this implementation, we use multiple hash functions to map input elements to different positions in the bit array. The choice of hash functions can impact the performance and accuracy of the Bloom Filter.
The likelihood of false positives can be reduced by using multiple hash functions and a larger bit array. Sidenote, that doing so will increase the space requirements of your Bloom filter.

## Pros

- Space-efficient: Bloom Filters require significantly less memory compared to storing the actual elements in the set.

- Fast membership queries: Checking if an element is in the set is usually faster compared to other data structures.

- No false negatives: If the Bloom Filter reports that an element is not in the set, it's guaranteed to be absent.

- Suitable for large datasets: Bloom Filters can handle large sets with a relatively small memory footprint.

## Cons and Mitigations

- False positives: Bloom Filters can have false positive results, meaning they might incorrectly report an element as in the set when it's not.
  - To mitigate false positives, one approach is to tune the number of hash functions and the size of the bit array based on the expected number of elements and the acceptable false positive rate.
  - Dynamic resizing: Implement dynamic resizing techniques to adjust the size of the Bloom Filter as the number of elements increases.

- One of the downsides of Bloom filters is that they can be difficult to resize effectively. Unlike other data structures that store elements directly, Bloom filters don't keep a record of the elements they've seen. So, if you want to build a larger Bloom filter, you'll need to retrieve all of the original keys from permanent storage.

- However, there is a solution to this challenge. Scalable Bloom filters, as the name implies, are Bloom filters that can automatically resize as the number of elements in the filter grows. This is achieved through the use of a series of smaller Bloom filters, each with a fixed size, that are combined into a chain. As the number of elements in the filter increases, more Bloom filters are added to the chain, effectively increasing its size.

## Other Variants of Bloom Filters

### Counting Bloom Filter

A Counting Bloom Filter enhances the basic Bloom Filter by allowing multiple hash functions to increment counters in the bit array instead of just setting bits. This enables tracking the frequency of elements rather than just their presence.

### Scalable Bloom Filter

The Scalable Bloom Filter addresses the issue of fixed size in the classic Bloom Filter by using a hierarchy of filters with increasing sizes. New filters are added as needed, and elements are promoted to higher-level filters when they become more prevalent.

## Usage

To use this Bloom Filter implementation in your Go code, follow these steps:

1. Clone the repository.
2. Import the `bloomfilter` package in your code.
3. Initialize a Bloom Filter with the desired parameters.
4. Add elements to the Bloom Filter using the `Add` method.
5. Check for membership using the `Contains` method.

## Contributing    
Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

Before contributing, please read the [code of conduct](CODE_OF_CONDUCT.md) & [contributing guidelines](CONTRIBUTING.md).
        
## License
Distributed under the MIT License. See [LICENSE](LICENSE) for more information.