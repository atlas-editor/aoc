import strutils

proc p1(input: string): int =
  return 0

proc p2(input: string): int =
  return 0

proc main() =
  var input = readFile("input.txt").strip()
  echo "part1=", p1(input)
  echo "part2=", p2(input)

main()
