import sys

import numpy as np
import numpy.typing as npt
from utils import input


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


def parse_input():
    data = input.lines()
    image_enhancement_algo = data[0]
    image = data[2:]
    return image_enhancement_algo, image


algo, image = parse_input()

image = np.pad(ImageEnhancer.convert_image(image), 1, constant_values=0)
enhancer = ImageEnhancer(algo)

for _ in range(50):
    image = enhancer.enhance(image)

print(int(np.sum(image)), file=sys.stderr)
