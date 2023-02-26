# leetcode

Templates and test cases for leetcode problem solving in a modular common way across problems

## Intro

Many leetcode problems are very "similar" and yet people try to recreate problems ground up.   Take problems involving BFS for example.   These problems have 2 parts:

1. Traversing neighbors of a node using a queue
2. Processing nodes as they are returned in this order.

Aspect (1) is very standard and Sienna's Algorithm Design Manual provides templates for this. Yet candidates try to reinvent this across problems with subtle changes.
Instead here we solve this by providing generics like a "bfs" method that simply takes helper functions so the interviewee can focus on solving the core of the problem (ie identifying neighbor nodes, handle nodes as they are visited).   The interview can always "drill" down into the generic bfs function when needed.

