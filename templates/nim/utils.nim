import re, strutils, sequtils, sugar

proc ints*(s: string): seq[int] =
  return findAll(s, re(r"-?\d+")).map(x => parseInt(x))
