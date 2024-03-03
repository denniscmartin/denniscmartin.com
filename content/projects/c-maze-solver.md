+++
title = 'A C program to solve PNG mazes'
date = 2022-10-16
thumbnail = 'projects-c-maze-solver-thumbnail.png'
draft = false
tags = ['c-language', 'maze', 'programming']
+++

I recently saw the 
[Maze solving youtube video](https://www.youtube.com/watch?v=rop0W4QDOUI) from 
Computerphile and I find it very interesting. So I decided to build my own maze solver program.

It's a C program that takes a maze PNG and outputs the solution using the [Wall Follower algorithm](https://en.wikipedia.org/wiki/Maze-solving_algorithm)

The program has a few constraints:

-   It only accepts PNG files
-   Mazes should be square
-   Walls should be black `rgb(0, 0, 0)` and path white `rg(255, 255, 255)`
-   Walls and path should be `1px` width
-   The starting point must be at `(x: 0, y: 1)`
-   The ending point should be at `(x: width, y: height - 1)`

You can find more information to build an use the program in my [Github](https://github.com/denniscmartin/maze-solver)

## Usage

Input file:

![Unsolved maze]({{< cdn >}}projects-c-maze-solver-input.png{{< /cdn >}})

```bash
$ ./maze_solver maze.png
Filename: maze.png
Width: 101
Height: 101
Algorithm duration: 0.000121 seconds
```

Output file:

![Solved maze]({{< cdn >}}projects-c-maze-solver-output.png{{< /cdn >}})