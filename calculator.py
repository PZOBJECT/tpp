#!/usr/bin/env python3
"""简单的计算器模块 - 由 Hermes + Claude Code 协同完成"""

from typing import Union

Number = Union[int, float]


def add(a: Number, b: Number) -> Number:
    """加法"""
    return a + b


def subtract(a: Number, b: Number) -> Number:
    """减法"""
    return a - b


def multiply(a: Number, b: Number) -> Number:
    """乘法"""
    return a * b


def divide(a: Number, b: Number) -> Number:
    """除法（含除零保护）"""
    if b == 0:
        raise ValueError("除数不能为0")
    return a / b


def main():
    """演示"""
    x, y = 10, 3
    print(f"{x} + {y} = {add(x, y)}")
    print(f"{x} - {y} = {subtract(x, y)}")
    print(f"{x} x {y} = {multiply(x, y)}")
    print(f"{x} / {y} = {divide(x, y):.2f}")


if __name__ == "__main__":
    main()
