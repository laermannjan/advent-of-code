"""
--- Day 20: Trench Map ---

With the scanners fully deployed, you turn their attention to mapping the floor of the ocean trench.

When you get back the image from the scanners, it seems to just be random noise. Perhaps you can combine an image enhancement algorithm and the input image (your puzzle input) to clean it up a little.

For example:

..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
#..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###
.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#.
.#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#.....
.#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#..
...####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.....
..##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###
The first section is the image enhancement algorithm. It is normally given on a single line, but it has been wrapped to multiple lines in this example for legibility. The second section is the input image, a two-dimensional grid of light pixels (#) and dark pixels (.).

The image enhancement algorithm describes how to enhance an image by simultaneously converting all pixels in the input image into an output image. Each pixel of the output image is determined by looking at a 3x3 square of pixels centered on the corresponding input image pixel. So, to determine the value of the pixel at (5,10) in the output image, nine pixels from the input image need to be considered: (4,9), (4,10), (4,11), (5,9), (5,10), (5,11), (6,9), (6,10), and (6,11). These nine input pixels are combined into a single binary number that is used as an index in the image enhancement algorithm string.

For example, to determine the output pixel that corresponds to the very middle pixel of the input image, the nine pixels marked by [...] would need to be considered:

# . . # .
#[. . .].
#[# . .]#
.[. # .].
. . # # #
Starting from the top-left and reading across each row, these pixels are ..., then #.., then .#.; combining these forms ...#...#.. By turning dark pixels (.) into 0 and light pixels (#) into 1, the binary number 000100010 can be formed, which is 34 in decimal.

The image enhancement algorithm string is exactly 512 characters long, enough to match every possible 9-bit binary number. The first few characters of the string (numbered starting from zero) are as follows:

0         10        20        30  34    40        50        60        70
|         |         |         |   |     |         |         |         |
..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..##
In the middle of this first group of characters, the character at index 34 can be found: #. So, the output pixel in the center of the output image should be #, a light pixel.

This process can then be repeated to calculate every pixel of the output image.

Through advances in imaging technology, the images being operated on here are infinite in size. Every pixel of the infinite output image needs to be calculated exactly based on the relevant pixels of the input image. The small input image you have is only a small region of the actual infinite input image; the rest of the input image consists of dark pixels (.). For the purposes of the example, to save on space, only a portion of the infinite-sized input and output images will be shown.

The starting input image, therefore, looks something like this, with more dark pixels (.) extending forever in every direction not shown here:

...............
...............
...............
...............
...............
.....#..#......
.....#.........
.....##..#.....
.......#.......
.......###.....
...............
...............
...............
...............
...............
By applying the image enhancement algorithm to every pixel simultaneously, the following output image can be obtained:

...............
...............
...............
...............
.....##.##.....
....#..#.#.....
....##.#..#....
....####..#....
.....#..##.....
......##..#....
.......#.#.....
...............
...............
...............
...............
Through further advances in imaging technology, the above output image can also be used as an input image! This allows it to be enhanced a second time:

...............
...............
...............
..........#....
....#..#.#.....
...#.#...###...
...#...##.#....
...#.....#.#...
....#.#####....
.....#.#####...
......##.##....
.......###.....
...............
...............
...............
Truly incredible - now the small details are really starting to come through. After enhancing the original input image twice, 35 pixels are lit.

Start with the original input image and apply the image enhancement algorithm twice, being careful to account for the infinite size of the images. How many pixels are lit in the resulting image?
--- Part Two ---

You still can't quite make out the details in the image. Maybe you just didn't enhance it enough.

If you enhance the starting input image in the above example a total of 50 times, 3351 pixels are lit in the final output image.

Start again with the original input image and apply the image enhancement algorithm 50 times. How many pixels are lit in the resulting image?
"""

import numpy as np
import numpy.typing as npt


class ImageEnhancer:
    def __init__(self, conversion_table: str):
        self.conversion_table = np.array([c == "#" for c in conversion_table])
        self._surrounding_value = 0

    def enhance(self, image: npt.NDArray) -> npt.NDArray:
        # simulate surrounding that is affecting the image
        img = np.pad(image, 1, constant_values=self._surrounding_value)

        # convert all pixels into conversion_map indices
        kernel_map = (
            img[:-2, :-2] << 8
            | img[:-2, 1:-1] << 7
            | img[:-2, 2:] << 6
            | img[1:-1, :-2] << 5
            | img[1:-1, 1:-1] << 4
            | img[1:-1, 2:] << 3
            | img[2:, :-2] << 2
            | img[2:, 1:-1] << 1
            | img[2:, 2:]
        )

        img = self.conversion_table[kernel_map]

        # as the surroundings of the image start out all 0, they can only ever
        # be 0 or 1, hence they will map to 0b000000000 or 0b111111111
        self._surrounding_value = self.conversion_table[
            0 if self._surrounding_value == 0 else 0b111111111
        ]

        # the image has "grown" through the enhance step
        # the surrounding pixels were affected by the image pixels
        # and thus become part of the image
        return np.pad(img, 1, constant_values=self._surrounding_value)

    @classmethod
    def convert_image(cls, image: list[str]):
        return np.array([[c == "#" for c in row] for row in image])


def parse_input(filename: str):
    with open(filename, "r") as f:
        data = f.read().splitlines()
    image_enhancement_algo = data[0]
    image = data[2:]
    return image_enhancement_algo, image


def day20a(filename: str) -> int:
    algo, image = parse_input(filename)

    # the actual image is infinite, with 0s surrounding our crop
    # we pad 0s with a width 1 to simulate this
    # as the kernel only considers adjacent pixels
    image = np.pad(ImageEnhancer.convert_image(image), 1, constant_values=0)
    enhancer = ImageEnhancer(algo)

    for _ in range(2):
        image = enhancer.enhance(image)

    return int(np.sum(image))


def day20b(filename: str) -> int:
    algo, image = parse_input(filename)

    image = np.pad(ImageEnhancer.convert_image(image), 1, constant_values=0)
    enhancer = ImageEnhancer(algo)

    for _ in range(50):
        image = enhancer.enhance(image)

    return int(np.sum(image))
