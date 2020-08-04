# Advent of Code 2019
Language: GoLang.
[https://adventofcode.com/2019](https://adventofcode.com/2019)

---
## Summary of Days
Day | Name | Type of Algo & Notes
--- | --- | ---
1 | The Tyranny of the Rocket Equation | - Simple math problem
2 | Program Alarm | - Intro to the crazy Intcode problems that are half the AoC days... <br> - Array (slice...) manipulation <br> - I used recursion
3 | Crossed Wires | - Geometry kind of algo, finding intersections of lines on a grid
4 | Secure Container | - May appear math-y, but it's really a string manipulation problem
5 | Sunny with a Chance of Asteroids | - Yay more Intcode!........ <br> - This gave me fits... <br> - Good application for recursion (in my opinion)
6 | Universal Orbit Map | - __Tree traversal__ and depth calculations. It's not quite a Graph, but it has a __directed graph__ algo feel too
7 | Amplification Circuit | - More Intcode... Piping together multiple Intcode computers 😳😳😳 <br> - Refactored Intcode computer to an OOP approach so a single computer maintains its data <br> - Also requires making __permutations generator__ <br> - Some gymnastics to make this circular, but its easier with this OOP approach and the "objects"/instances of a struct maintaining their own data <br> - Concurrency could be used to sync these Amps together...
8 | Space Image Format | 3D Array manipulation, pretty straight forward
9 | Sensor Boost | __MORE INTCODE. YAYY__ 🙃 <br> - A new parameter mode and opcode. <br> - __Really feeling the (tech) debt of some earlier design choices here, went back to refactor day07 before jumping into this one__, then it was a small bit of code for the relative param/opcode & resizing computer memory if necessary
10 | Monitoring Station | - This (part2) is my favorite algo... Yes I have a favorite algo <br> Fundamentally it's a geometry problem, angles and trig <br> - Part 1: Calculated via slopes <br> - Part 2: Using Arctangent to find angles an asteroid makes against a vertical line from the home base asteroid. Then those angles can be used to determine if Asteroids are covering each other, AND iterating through all of them can find the next angle Asteroid to vaporize
11 | Space Police | - More Intcode stuff... <br> - __2D Array/Slice manipulation__ and a bit of maths/graphing <br> - Implemented a __RotateGrid__ algo
12 | The N-Body Problem | I like to call this a _(harmonic) frequency_ algo. Finding the harmonic frequency of multiple bodies/items and then finding the Least Common Multiple of those frequencies will tell you when all the bodies have returned to their initial state. <br> - I've used this approach for a leetcode problem about prisoners in jail cells too
13 | Care Package | Intcode again! It's basically every other day... <br> - part1: 2D array manipulation again <br> - part2: holy algo, some logic to basically play Bricks game. <br> - This is more of a design question for how you manage state
14 | Space Stoichiometry | __Weighted Graph and Breadth First Traversals__ <br> - Because not all of the products have a quantity of 1, it complicates the graph's data model. I ended up mapping the product/chemical name to a map of its stoichiometry where the product's number is positive & reactants were negative. <br> - part2: not just a simple division because of "extra byproducts" of generating each piece of fuel. I just let this brute force thing run for ~30 seconds...
15 | Oxygen System | YAY INTCODE 🙄 <br> - Combination of __searching algo__, __backtracking algo__ and the the Intcode... <br> - I've realized that I really need to stop using x and y for 2D grids and start using row and col because mathematically x is horizontal and y is vertical... My brain is all jumbled up <br> - Created a Robot struct/class that has a computer inside of it. It goes and searches around, collecting data on the floor types at various coordinates. That data is transformed into a 2D grid/array, and then finally fed into a backtracking, searching algorithm to determine the shortest path (turns out there's only one path to the O2 tank...) <br> - part2 is fairly straight forward 2D grid traversing and tagging a spread of oxygen to valid tiles/hallway spaces
