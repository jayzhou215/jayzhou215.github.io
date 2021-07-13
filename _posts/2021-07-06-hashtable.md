---
layout: post
title: data struct - hash table
tags: [data struct]
readtime: true
comments: true
---
## 1. Hash Table
Hash Table is a data structure to map key to values (also called Table or Map Abstract Data Type/ADT). It uses a hash function to map large or even non-Integer keys into a small range of Integer indices (typically [0..hash_table_size-1]).

The probability of two distinct keys colliding into the same index is relatively high and each of this potential collision needs to be resolved to maintain data integrity.

There are several collision resolution strategies that will be highlighted in this visualization: Open Addressing (Linear Probing, Quadratic Probing, and Double Hashing) and Closed Addressing (Separate Chaining). Try clicking Search(8) for a sample animation of searching a value in a Hash Table using Separate Chaining technique.

## 2. Motivation
Hashing is an algorithm (via a hash function) that maps large data sets of variable length, called keys, not necessarily Integers, into smaller Integer data sets of a fixed length.

A Hash Table is a data structure that uses a hash function to efficiently map keys to values (Table or Map ADT), for efficient search/retrieval, insertion, and/or removals.

Hash Table is widely used in many kinds of computer software, particularly for associative arrays, database indexing, caches, and sets.

In this e-Lecture, we will digress to Table ADT, the basic ideas of Hashing, the discussion of Hash Functions before going into the details of Hash Table data structure itself.

### 2.1 Table ADT 
A Table ADT must support at least the following three operations as efficient as possible:

* Search(v) — determine if v exists in the ADT or not,
* Insert(v) — insert v into the ADT,
* Remove(v) — remove v from the ADT.
Hash Table is one possible good implementation for this Table ADT (the other one is this).

PS1: For two weaker implementations of Table ADT, you can click the respective link: unsorted array or a sorted array to read the detailed discussions.

PS2: In live class, you may want to compare the requirements of Table ADT vs List ADT.

### 2.2 Direct Addressing Table (DAT)
When the range of the Integer keys is small, e.g. [0..M-1], we can use an initially empty (Boolean) array A of size M and implement the following Table ADT operations directly:

* Search(v): Check if A[v] is true (filled) or false (empty),
* Insert(v): Set A[v] to be true (filled),
* Remove(v): Set A[v] to be false (empty).
That's it, we use the small Integer key itself to determine the address in array A, hence the name Direct Addressing. It is clear that all three major Table ADT operations are O(1).

### 2.3 Example Of ADT
In Singapore (as of Mar 2018), bus routes are numbered from [2..990].

Not all integers between [2..990] are currently used, e.g. there is no bus route 989 — Search(989) should return false. A new bus route x may be introduced, i.e. Insert(x) or an existing bus route y may be discontinued, i.e. Remove(y).

As the range of possible bus routes is small, to record the data whether a bus route number exists or not, we can use a DAT with a Boolean array of size 1 000.

Discussion: In real life class, we may discuss on why we use 1 000 instead of 990 (or 991).

### 2.4 Example Of DAT With Satellite Data
Notice that we can always add satellite data instead of just using a Boolean array to record the existence of the keys.

For example, we can use an associative String array A instead to map a bus route number to its operator name, e.g.

A[2] = "Go-Ahead Singapore",
A[10] = "SBS Transit",
A[183] = "Tower Transit Singapore",
A[188] = "SMRT Buses", etc.
Discussion: Can you think of a few other real-life DAT examples?

### 2.5 The Answer 
Hidden

### 2.6 DAT Limitations
The keys must be (or can be easily mapped to) non-negative Integer values.
Basic DAT has problem in the full version of the example in the previous two slides as there are actually variations of bus route numbers in Singapore, e.g. 96B, 151A, NR10, etc.

The range of keys must be small.
The memory usage will be (insanely) large if we have (insanely) large range.

The keys must be dense, i.e. not many gaps in the key values.
DAT will contain too many empty cells otherwise.

We will overcome these restrictions with hashing.

## 3. Hashing: ideas 
Using hashing, we can:

1. Map (some) non-Integer keys to Integers keys,
1. Map large Integers to smaller Integers,
1. Influence the density, or load factor α = N/M, of the Hash Table where N is the number of keys and M is the size of the Hash Table.

### 3.1 Phone numbers Example 
For example, we have N = 400 Singapore phone numbers (Singapore phone number has 8 digits, so there are up to 10^8 = 100M possible phone numbers in Singapore).

Instead of using a DAT and use a gigantic array up to size M = 100 Million, we can use the following simple hash function h(v) = v%997.

This way, we map 8 digits phone numbers 6675 2378 and 6874 4483 into up to 3 digits h(6675 2378) = 237 and h(6874 4483) = 336, respectively. Therefore, we only need to prepare array of size M = 997 (or 1000) instead of M = 100 Million.

### 3.2 Hash Table Preview
With hashing, we can now implement the following Table ADT operations using Integer array (instead of Boolean array) as follows:

1. Search(v): Check if A[h(v)] != -1 (we use -1 for an empty cell assuming v ≥ 0),
1. Insert(v): Set A[h(v)] = v (we hash v into h(v) so we need to somehow record key v),
1. Remove(v): Set A[h(v)] = -1 — to be elaborated further.