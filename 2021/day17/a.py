import re
import sys
from dataclasses import dataclass
from itertools import chain, dropwhile, product, takewhile

import numpy as np
from utils import input
from utils.misc import sum_to, sum_to_inverse


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
        last_step_size = abs(self.target_y_min)
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
            vy_min = int(np.ceil((self.target_y_min + sum_to(t - 1)) / t))
            vy_max = int(np.floor((self.target_y_max + sum_to(t - 1)) / t))
            vy_range = range(vy_min, vy_max + 1)

            vx_min = int(
                np.ceil((self.target_x_min + sum_to(t - 1)) / t)
            )  # discard vx that stop in less than t steps
            vx_max = int(np.floor((self.target_x_max + sum_to(t - 1)) / t))

            vx_range = chain(
                dropwhile(lambda v: v < t, range(vx_min, vx_max + 1)),
                takewhile(lambda v: v < t, range(vx_min_stopping, vx_max_stopping + 1)),
            )

            valid_velocities.update(product(vx_range, vy_range))
        return valid_velocities

    @classmethod
    def from_file(cls) -> "ProbeLauncher":
        data = input.lines()[0]
        x_min, x_max, y_min, y_max = map(int, re.findall(r"(-?\d+)", data))
        return cls(x_min, x_max, y_min, y_max, x_init=0, y_init=0)


launcher = ProbeLauncher.from_file()
print(launcher.max_y_position, file=sys.stderr)
