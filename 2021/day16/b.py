import sys
from abc import ABC, abstractmethod
from dataclasses import dataclass

import numpy as np
from utils import input


@dataclass
class Packet(ABC):
    version: int
    packet_type: int

    @abstractmethod
    def version_sum(self) -> int:
        pass

    @property
    @abstractmethod
    def value(self) -> int:
        pass

    def __int__(self):
        return self.value


@dataclass
class LiteralValuePacket(Packet):
    _value: int

    def version_sum(self) -> int:
        return self.version

    @property
    def value(self) -> int:
        return self._value


@dataclass
class OperatorPacket(Packet):
    subpackets: list[Packet]

    def version_sum(self) -> int:
        return sum(packet.version_sum() for packet in self.subpackets) + self.version

    @property
    def value(self) -> int:
        if self.packet_type == 0:
            return sum([p.value for p in self.subpackets])
        if self.packet_type == 1:
            return int(np.product([p.value for p in self.subpackets]))
        if self.packet_type == 2:
            return min([p.value for p in self.subpackets])
        if self.packet_type == 3:
            return max([p.value for p in self.subpackets])
        if self.packet_type == 5:
            return 1 if self.subpackets[0].value > self.subpackets[1].value else 0
        if self.packet_type == 6:
            return 1 if self.subpackets[0].value < self.subpackets[1].value else 0
        if self.packet_type == 7:
            return 1 if self.subpackets[0].value == self.subpackets[1].value else 0


class BITS:
    def __init__(self, bit_string: str):
        self.data = bit_string

    def __len__(self):
        return len(self.data)

    def __parse_packet(self) -> tuple[Packet, "BITS"]:
        """
        Parse the bits as a packet and recursively parse the sub-packets
        Returns:
            A tuple containing the parsed packet and the remaining bits.
        """
        data = self.data
        version = int(data[0:3], 2)
        packet_type = int(data[3:6], 2)
        if packet_type == 4:
            num = ""
            for i in range(6, len(data), 5):
                num += data[i + 1 : i + 5]
                if data[i] == "0":
                    break
            value = int(num, 2)
            return LiteralValuePacket(version, packet_type, value), BITS(data[i + 5 :])
        length_type = int(data[6])
        subpackets = []
        if length_type == 0:
            # next 15 bits are a number that represents the total length in bits of the sub-packets contained by this packet
            length = int(data[7:22], 2)
            subpacket_bits = BITS(data[22 : 22 + length])
            while sum(int(x) for x in subpacket_bits.data):
                subpacket, subpacket_bits = subpacket_bits.__parse_packet()
                subpackets.append(subpacket)
            data = data[22 + length :]
        else:
            # next 11 bits are a number that represents the number of sub-packets immediately contained by this packet
            length = int(data[7:18], 2)
            data = data[18:]
            for i in range(0, length):
                subpacket, bits = BITS(data).__parse_packet()
                subpackets.append(subpacket)
                data = bits.data

        return (
            OperatorPacket(version, packet_type, subpackets),
            BITS(data),
        )

    def parse_packet(self) -> Packet:
        """
        Parse the bits as a packet.
        Returns:
            The parsed packet.
        """
        return self.__parse_packet()[0]

    def __repr__(self):
        return self.data

    @classmethod
    def from_hex(cls, hex_string: str) -> "BITS":
        """
        Convert a hex string to binary and left-pad with zeroes to make it a multiple of 4 bits
        """
        return cls(bin(int(hex_string, 16))[2:].zfill(len(hex_string) * 4))

    @classmethod
    def from_file(cls) -> "BITS":
        hex_data = input.lines()[0]
        return cls.from_hex(hex_data)


bits = BITS.from_file()
packet = bits.parse_packet()
print(packet.value, file=sys.stderr)
