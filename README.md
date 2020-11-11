# Searching Networks Notes
Oftentimes we want to find out if we can find a particular item in a graph,
whether that graph is a tree, a social network, a map of cities, or an obstacle course.

What's the best strategy to find it? Do we want the smallest number of edges separating us from our target?
Do we know that our node is nearby and don't want to get stuck search far away edges? Are we working with a very wide tree?
Is our network vast? 

We will discuss in particular two search algorithsm: breadth first, and depth first. We'll then touch on a combination of the two that is used in AI.

This depends on the structure of the search tree and the number and location of solutions (aka searched-for items).

If you know a solution is not far from the root of the tree, a breadth first search (BFS) might be better.

If the tree is very deep and solutions are rare, depth first search (DFS) might take an extremely long time... so BFS could be faster.

If the tree is very wide, a BFS might need too much memory for its queue, so it might be impractical.

If solutions are frequent but located deep in the tree, BFS could be impractical.

If the search tree is very deep you will need to restrict the search depth for depth first search (DFS), anyway (for example with iterative deepening).

Before we move on, we'll discuss traversals. A traversal is where you visit every single node in a graph. While BFS will explore nodes in essentially the same order every time (always creating the same tree), a depth-first search can create various trees. 
We'll talk about the 3 different ways you can do a depth-first traversal: in-order, pre-order, or post-order. Each of these results in enumerating the vertices in different ways.

So far we've covered DFS and BFS and have talked about some benefits of each. How can we get some benefits of both? With iterative deepening! It takes DFS and adds a twist: only search up to a certain depth. 