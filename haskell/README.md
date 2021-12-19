# Advent of Code - Haskell Implementation

## Haskell

Install the `haskell-platform` via ghcup and make sure to install `stack` (see the [docs](https://www.haskell.org/downloads/)):

```bash
curl --proto '=https' --tlsv1.2 -sSf https://get-ghcup.haskell.org | sh
```

VSCode seems to have the easiest and most accessible [plugin for Haskell including the Haskell-Language-Server](https://marketplace.visualstudio.com/items?itemName=haskell.haskell).

[Haskelly](https://marketplace.visualstudio.com/items?itemName=UCL.haskelly) allows to easily run and haskell files in interpreter mode and may be installed additionally.

## Usage

Make sure that `stack` is in your `$PATH`. Then from the this directory build the project and run it:

```bash
stack build
stack exec adventofcode  # prints usage info
stack exec adventofcode -- 2021 4a  # solves Problem 4a of AoC 2021
```

## Caveats

I am using `System.TimeIt` to time the program execution. The results seem highly unstable. Run each problem a few times to get a better idea.
I might try `Criterion` in the future.
