np-hard
================================================================================

A collection of UNIX utilities that solve various NP hard problems

Similar to the way grep solves search in a composable manner, these
utilities should be repositories of exact and heuristic solutions to
hard problems that have shell-composable line-oriented interfaces.

Currently the only utility included is a Travelling Salesman Problem
solver in the tsp/ directory, but the idea is to eventually cover
several popular NP-hard problems.

COMPILE
--------------------------------------------------------------------------------

To compile any of these utilities:

    go build name.go

TSP
--------------------------------------------------------------------------------

Usage:

    cat points.txt | tsp -nearest-neighbor

We offer a selection of exact and heuristic algorithms:

 - Exact
   - Brute Force (`-brute`)
 - Heuristic
   - Nearest Neighbor (`-nearest-neighbor`)

The result of `tsp --help`:

    Usage of ./tsp:
      -brute=false: Brute force searcher: slow but exact
       Time: O(n!) Space: O(n)
             Multithreaded by default, turn off with -single-core
      -nearest-neighbor=false: Nearest Neighbor heuristic: fast and approximate
       Time: O(n**3) Space: O(n)
             Path produced is on average 25% longer than optimal:
              [Johnson, D.S. and McGeoch, L.A.. "The traveling salesman problem:
               A case study in local optimization", Local search in combinatorial
                optimization, 1997, 215-310]
      -single-core=false: Disable multithreading

LICENSE
--------------------------------------------------------------------------------
Copyright (c) 2013 thenoviceoof

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.