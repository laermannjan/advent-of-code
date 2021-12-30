"""
--- Day 17: Trick Shot ---
You finally decode the Elves' message. HI, the message says. You continue searching for the sleigh keys.
Ahead of you is what appears to be a large ocean trench. Could the keys have fallen into it? You'd better send a probe to investigate.
The probe launcher on your submarine can fire the probe with any integer velocity in the x (forward) and y (upward, or downward if negative) directions. For example, an initial x,y velocity like 0,10 would fire the probe straight up, while an initial velocity like 10,-1 would fire the probe forward at a slight downward angle.
The probe's x,y position starts at 0,0. Then, it will follow some trajectory by moving in steps. On each step, these changes occur in the following order:
The probe's x position increases by its x velocity.
The probe's y position increases by its y velocity.
Due to drag, the probe's x velocity changes by 1 toward the value 0; that is, it decreases by 1 if it is greater than 0, increases by 1 if it is less than 0, or does not change if it is already 0.
Due to gravity, the probe's y velocity decreases by 1.
For the probe to successfully make it into the trench, the probe must be on some trajectory that causes it to be within a target area after any step. The submarine computer has already calculated this target area (your puzzle input). For example:
target area: x=20..30, y=-10..-5
This target area means that you need to find initial x,y velocity values such that after any step, the probe's x position is at least 20 and at most 30, and the probe's y position is at least -10 and at most -5.
Given this target area, one initial velocity that causes the probe to be within the target area after any step is 7,2:
.............#....#............
.......#..............#........
...............................
S........................#.....
...............................
...............................
...........................#...
...............................
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTT#TT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
In this diagram, S is the probe's initial position, 0,0. The x coordinate increases to the right, and the y coordinate increases upward. In the bottom right, positions that are within the target area are shown as T. After each step (until the target area is reached), the position of the probe is marked with #. (The bottom-right # is both a position the probe reaches and a position in the target area.)
Another initial velocity that causes the probe to be within the target area after any step is 6,3:
...............#..#............
...........#........#..........
...............................
......#..............#.........
...............................
...............................
S....................#.........
...............................
...............................
...............................
.....................#.........
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................T#TTTTTTTTT
....................TTTTTTTTTTT
Another one is 9,0:
S........#.....................
.................#.............
...............................
........................#......
...............................
....................TTTTTTTTTTT
....................TTTTTTTTTT#
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
....................TTTTTTTTTTT
One initial velocity that doesn't cause the probe to be within the target area after any step is 17,-4:
S..............................................................
...............................................................
...............................................................
...............................................................
.................#.............................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT................................
....................TTTTTTTTTTT..#.............................
....................TTTTTTTTTTT................................
...............................................................
...............................................................
...............................................................
...............................................................
................................................#..............
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
...............................................................
..............................................................#
The probe appears to pass through the target area, but is never within it after any step. Instead, it continues down and to the right - only the first few steps are shown.
If you're going to fire a highly scientific probe out of a super cool probe launcher, you might as well do it with style. How high can you make the probe go while still reaching the target area?
In the above example, using an initial velocity of 6,9 is the best you can do, causing the probe to reach a maximum y position of 45. (Any higher initial y velocity causes the probe to overshoot the target area entirely.)
Find the initial velocity that causes the probe to reach the highest y position and still eventually be within the target area after any step. What is the highest y position it reaches on this trajectory?
--- Part Two ---
Maybe a fancy trick shot isn't the best idea; after all, you only have one probe, so you had better not miss.
To get the best idea of what your options are for launching the probe, you need to find every initial velocity that causes the probe to eventually be within the target area after any step.
In the above example, there are 112 different initial velocity values that meet these criteria:
23,-10  25,-9   27,-5   29,-6   22,-6   21,-7   9,0     27,-7   24,-5
25,-7   26,-6   25,-5   6,8     11,-2   20,-5   29,-10  6,3     28,-7
8,0     30,-6   29,-8   20,-10  6,7     6,4     6,1     14,-4   21,-6
26,-10  7,-1    7,7     8,-1    21,-9   6,2     20,-7   30,-10  14,-3
20,-8   13,-2   7,3     28,-8   29,-9   15,-3   22,-5   26,-8   25,-8
25,-6   15,-4   9,-2    15,-2   12,-2   28,-9   12,-3   24,-6   23,-7
25,-10  7,8     11,-3   26,-7   7,1     23,-9   6,0     22,-10  27,-6
8,1     22,-8   13,-4   7,6     28,-6   11,-4   12,-4   26,-9   7,4
24,-10  23,-8   30,-8   7,0     9,-1    10,-1   26,-5   22,-9   6,5
7,5     23,-6   28,-10  10,-2   11,-1   20,-9   14,-2   29,-7   13,-3
23,-5   24,-8   27,-9   30,-7   28,-5   21,-10  7,9     6,6     21,-5
27,-10  7,2     30,-9   21,-8   22,-7   24,-9   20,-6   6,9     29,-5
8,-2    27,-8   30,-5   24,-7
How many distinct initial velocity values cause the probe to be within the target area after any step?
"""


import re
from dataclasses import dataclass
from adventofcode.lib import sum_to, sum_to_inverse
import numpy as np
from itertools import product, chain, takewhile


@dataclass
class ProbeLauncher:
    target_x_min: int
    target_x_max: int
    target_y_min: int
    target_y_max: int
    x_init: int
    y_init: int

    @property
    def max_y_position(self) -> int:
        """
        Returns maximum reachable height of the probe that will still hit the target area.

        Realize that because of the way gravity decreases y_velocity every step by 1,
        after being shot upwards from (0, 0), the probe will always travel through a position
        with y=0.
        The largest step the probe can take after reaching that point to still hit the target
        area is from 0 -> target_area.y_min = abs(target_area.y_min).
        The step it took right before was one less and the one before one less still.
        Therefore, we can simply sum up the integers from 1 to the step_size before last.

        """
        last_step_size = abs(self.target_area.y_min)
        return int(sum_to(last_step_size - 1))

    @property
    def valid_trajectories(self):
        """
        List of all initial velocity vectors (vx, vy) that - at any step - hit the target area

        The problem is solved by determining independent vx/vy bounds for any given number steps t.
        For example, given t=1 and target area parameters
            - y_min=30, y_max=20
            - x_min= 5, x_max=10
            => vx_min = x_min = 5, vx_max = x_max = 10
        as probes launched with vx_min <= vx <= vx_max will hit the target within 1 step.
        vx and vy bounds are determined independently; the result is the cartesian product of both ranges.

        Determining these bounds without drag or gravity would be trivial as the velocities are then constant.
            . x[t] = x[0] + t * vx
            . y[t] = y[0] + t * vy
        With x[0] = 0; y[0] = 0
            - x[t] = t * vx
            - y[t] = t * vy
            => vy_min[t] = y_min / t
            => vy_max[t] = y_max / t
            => ...

        Gravity decreases vy after every step by 1. Hence, we need to increase our initial vy to compensate the pull.
        Notice, that gravity does not decrease the initial y-velocity vy.
        As such, we can describe vx with respect to t and adjust y[t].
            . vy[t] = vy[t-1] - 1
            . vy[1] = vy_init
            - vy[2] = vy[1] - 1 = vy_init - 1
            - vy[3] = vy[2] - 1 = (vy_init - 1) - 1
            => vx[t] = vx_init - (t-1)
        It follows for y[t]
            . y[t] = y[t-1] + vy[t]
            . y[0] = y_init
            - y[1] = y_init + vy_init
            - y[2] = (y_init + vy_init) + vy_init - 1
            - y[3] = ((y_init + vy_init) + vy_init - 1) - 2 = y_init + 2 * vy_init - (1 + 2)
            => y[t] = y_init + t * vy_init - sum[1...(t-1)]

        With y_init = 0
            - y[t] = t * vy_init - sum[1...(t-1)]
            => vy_init = (y[t] + sum[1...(t-1)]) / t
        An more intuitive interpretation is to consider the following
        at any step t > 0
            . y[t] = t * vy_init - sum[1...(t-1)]  # from above; point that will be reached due to gravity
            . y'[t] = t * vy_init                  # point that would have been reached without gravity; constant vy
            - d[t] = y[t] - y'[t]                  # the delta we missed y'[t] by due to gravity
            - d'[t] = -d[t] = + sum[1...(t-1)]                 # the "aim correction" needed to not hit y[t]
            => y~[t] = t * vy_init + d'[t]
            => vy_init = y~[t] / t                 # y~[t] is the point to aim for, pretending to shoot straight
                                                   # but hitting y[t] at step t.


        Similarly, drag reduces vx after every step by 1, without altering the initial vx, but does not reduce below 0.
        First, similar math applies
            . vx[t] = max(vx[t-1] - 1, 0)
            . vx[1] = vx_init
            - vx[t] =    if vx_init >= t;   vx_init - (t-1)
                         else;               <unknown yet>
            => vx_init = if vx_init >= t;   (x[t] + sum[1...(t-1)]) / t
                         else;               <unknown yet>

        This allows us to determine vx that reach the target area after t steps.
        We observe that
            . vx_init > t  -->  vx > 0 after t steps; probe x-trajectory passes through target area
            . vx_init = t  -->  vx = 0 after t steps; probe x-trajectory stops within target area

        vx_init < t would cause the probe to reach vx = 0 and stay stationary before step t.
        If such a probe has reached the target area before reaching vx = 0, it will stay in the target area
        (only with respect to the x-axis).
        Rather than dealing with the else clause from above, we realize that
            - x = sum_to(vx_init)           # x is the stationary point of a probe launched with vx_init
            => vx_init = sum_to_inverse(x)  # vx_init, the velocity that lets the probe stop at position x
        Luckily the inversion is trivial:
                    sum_to := S = 1+2+...+n = (n+1)(n/2)
                              2*S = n(n+1) = n^2 + n = (n + 1/2)^2 - 1/4   [from binomial formula; (a+b)**2 = a**2 + 2ab + b**2]
                              2*S + 1/4 = (n + 1/2)^2
            sum_to_inverse := sprt(2*S + 1/4) - 1/2 = n
        Hence, we complete above's formula
            - vx_init = if   vx_init >= t;   (x[t] + sum[1...(t-1)]) / t
                        elif vx_init < t;     sum_to_inverse(x[t])
                         else;               <unknown yet>

        The result is now a list produced by the cartesian product over all (vx_init[t], vx_init[t]) for all time
        (necessary) time steps t.
        """

        vx_min_stopping = int(np.ceil(sum_to_inverse(self.target_x_min)))
        vx_max_stopping = int(np.floor(sum_to_inverse(self.target_x_max)))

        # from max_y_position we know that target_y_min is largest step we can take in any y direction without
        # overshooting. Computing backwards, we start with vy = target_y_min, the step before was one less, etc. until
        # we reach 0 and then invert vy to come back down to y_init.
        max_steps = abs(self.target_y_min) * 2
        valid_velocities = set()
        for t in range(1, max_steps + 1):
            vy_min = int(np.ceil((self.target_y_min + sum_to(t-1)) / t))
            vy_max = int(np.floor((self.target_y_max + sum_to(t-1)) / t))
            vy_range = range(vy_min, vy_max + 1)

            vx_min = int(np.ceil((self.target_x_min + sum_to(t-1)) / t))  # discard vx that stop in less than t steps
            vx_max = int(np.floor((self.target_x_max + sum_to(t-1)) / t))
            vx_range = chain(
                takewhile(lambda v: v >= t, range(vx_min, vx_max + 1)),
                takewhile(lambda v: v < t, range(vx_min_stopping, vx_max_stopping + 1))
            )
            valid_velocities.update(product(vx_range, vy_range))
        return valid_velocities

    @classmethod
    def from_file(cls, filename: str) -> "ProbeLauncher":
        with open(filename, 'r') as f:
            data = f.read().rstrip()
        x_min, x_max, y_min, y_max = map(int, re.findall(r"(-?\d+)", data))
        return cls(x_min, x_max, y_min, y_max, x_init=0, y_init=0)


def day17a(filename: str) -> int:
    launcher = ProbeLauncher.from_file(filename)
    return launcher.max_y_position


def day17b(filename: str) -> int:
    launcher = ProbeLauncher.from_file(filename)
    return len(launcher.valid_trajectories)
