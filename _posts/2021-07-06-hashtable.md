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

This way, we map 8 digits phone numbers 6675 2378 and 6874 4483 into up to 3 digits h(6675 2378) = 237 and h(6874 4483) = 336, respectively. Therefore, we only need to prepare array of size M = 997 (or 1000) instead of M = 100 Million.

### 3.2 Hash Table Preview
With hashing, we can now implement the following Table ADT operations using Integer array (instead of Boolean array) as follows:

1. Search(v): Check if A[h(v)] != -1 (we use -1 for an empty cell assuming v ≥ 0),
1. Insert(v): Set A[h(v)] = v (we hash v into h(v) so we need to somehow record key v),
1. Remove(v): Set A[h(v)] = -1 — to be elaborated further.

### 3.3 Hash Table with Satellite Data

If we have keys that map to satellite data and we want to record the original keys too, we can implement the Hash Table using pair of (Integer, satellite-data-type) array as follows:

Search(v): Return A[h(v)], which is a pair (v, satellite-data), possibly empty,
Insert(v, satellite-data): Set A[h(v)] = pair(v, satellite-data),
Remove(v): Set A[h(v)] = (empty pair) — to be elaborated further.
However, by now you should notice that something is incomplete...

### 3.4 Collision
A hash function may, and quite likely, map different keys (Integer or not) into the same Integer slot, i.e. a many-to-one mapping instead of one-to-one mapping.

For example, h(6675 2378) = 237 from three slides earlier and if we want to insert another phone number 6675 4372, we will have a problem as h(6675 4372) = 237 too.

This situation is called a collision, i.e. two (or more) keys have the same hash value.

### 3.5 Probability of Collision

The Birthday (von Mises) Paradox asks: 'How many people (number of keys) must be in a room (Hash Table) of size 365 seats (cells) before the probability that some person share a birthday (collision, two keys are hashed to the same cell), ignoring the leap years (i.e. all years have 365 days), becomes > 50 percent (i.e. more likely than not)?'

The answer, which maybe surprising for some of us, is Reveal(after having just 23 people (keys) in the room (the hash table of size 365 cells), it is more likely (> 50% chance) to have a collision than not... We do not need 365/2 ~= 180+ people.).

Let's do some calculation.

### 3.6 The Calculation
Let Q(n) be the probability of unique birthday for n people in a room.
Q(n) = 365/365 × 364/365 × 363/365 × ... × (365-n+1)/365,
i.e. the first person's birthday can be any of the 365 days, the second person's birthday can be any of the 365 days except the first person's birthday, and so on.

Let P(n) be the probability of same birthday (collision) for n people in a room.
P(n) = 1-Q(n).

We compute that P(23) = 0.507 > 0.5 (50%).

Thus, we only need 23 people (a small amount of keys) in the room (Hash Table) of size 365 seats (cells) for a (more than) 50% chance collision to happen (the birthday of two different people in that room is one of 365 days/slots).

### 3.7 Two Important Issues 
Issue 1: We have seen a simple hash function like the h(v) = v%997 used in Phone Numbers example that maps large range of Integer keys into a smaller range of Integer keys, but how about non Integer keys? How to do such hashing efficiently?

Issue 2: We have seen that by hashing, or mapping, large range into smaller range, there will very likely be a collision. How to deal with them?

## 4. Hash functions

How to create a good hash function with these desirable properties?

1. Fast to compute, i.e. in O(1),
2. Uses as minimum slots/Hash Table size M as possible,
3. Scatter the keys into different base addresses as uniformly as possible ∈ [0..M-1],
4. Experience as minimum collisions as possible.

### 4.1 Preliminaries
Suppose we have a hash table of size M where keys are used to identify the satellite-data and a specific hash function is used to compute a hash value.

A hash value/hash code of key v is computed from the key v with the use of a hash function to get an Integer in the range 0 to M-1. This hash value is used as the base/home index/address of the Hash Table entry for the satellite-data.

### 4.2 example of a bad hash function
Using the Phone Numbers example, if we we define h(v) = floor(v/1 000 000),
i.e. we select the first two digits a phone number.

h(66 75 2378) = 66
h(68 74 4483) = 68
Discuss: What happen when you use that hash function? Hint: See this.

### 4.3 the answer
hidden

### 4.4 Perfect Hash Function

Before discussing the reality, let's discuss the ideal case: perfect hash functions.

A perfect hash function is a one-to-one mapping between keys and hash values, i.e. no collision at all. It is possible if all keys are known beforehand. For example, a compiler/interpreter search for reserved keywords. However, such cases are rare.

A minimal perfect hash function is achieved when the table size is the same as the number of keywords supplied. This case is even rarer.

If you are interested, you can explore GNU gperf, a freely available perfect hash function generator written in C++ that automatically constructs perfect functions (a C++ program) from a user supplied list of keywords.

### 4.5 Hashing Integer - Best Practice
People has tried various ways to hash a large range of Integers into a smaller range of Integers as uniformly as possible. In this e-Lecture, we jump directly to one of the best and most popular version: h(v) = v%M, i.e. map v into Hash Table of size M slots. The (%) is a modulo operator that gives the remainder after division. This is clearly fast, i.e. O(1) assuming that v does not exceed natural Integer data type limit.

The Hash Table size M is set to be a reasonably large prime not near a power of 2, about 2+ times larger than the expected number of keys N that will ever be used in the Hash Table. This way, the load factor α = N/M < 0.5 — we shall see later that having low load factor, thereby sacrificing empty spaces, help improving Hash Table performance.

Discuss: What if we set M to be a power of 10 (decimal) or power of 2 (binary)?

### 4.6 Answer
hidden

### 4.7 Hashing String - Best Practice
People has also tried various ways to hash Strings into a small range of Integers as uniformly as possible. In this e-Lecture, we jump directly to one of the best and most popular version, shown below:

int hash_function(string v) { // assumption 1: v uses ['A'..'Z'] only
    int sum = 0;                // assumption 2: v is a short string
    for (auto &c : v) // for each character c in v
        sum = ((sum*26)%M + (c-'A'+1))%M; // M is table size
    return sum;
}
Discussion: In real life class, discuss the components of the hash function above, e.g. why loop through all characters?, will that be slower than O(1)?, why multiply with 26?, what if the string v uses more than just UPPERCASE chars?, etc

### 4.8 Answer
hidden

## 5 Collision Resolution

There are two major ideas: Open Addressing versus Closed Addressing method.

In Open Addressing, all hashed keys are located in a single array. The hash code of a key gives its base address. Collision is resolved by checking/probing multiple alternative addresses (hence the name open) in the table based on a certain rule.

In Closed Addressing, the Hash Table looks like an Adjacency List (a graph data structure). The hash code of a key gives its fixed/closed base address. Collision is resolved by appending the collided keys inside a (Doubly) Linked List identified by the base address.