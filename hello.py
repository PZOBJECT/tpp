#!/usr/bin/env python3
"""Hermes-test: AI协同编程演示"""

def greet(name: str) -> str:
    """问候函数"""
    return f"Hello, {name}! 这是 Hermes + Claude Code 协同写的代码。"

def main():
    print(greet("World"))

if __name__ == "__main__":
    main()

