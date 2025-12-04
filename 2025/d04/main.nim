import strutils

iterator nbrs(r, c, R, C: int): (int, int) =
  for dr in [-1, 0, 1]:
    for dc in [-1, 0, 1]:
      if dr == 0 and dc == 0:
        continue
      var rr = r + dr
      var cc = c + dc
      if 0 <= rr and rr < R and 0 <= cc and cc < C:
        yield (rr, cc)

proc rolls(m: seq[string]): seq[(int, int)] =
  var (R, C) = (m.len(), m[0].len())

  var rolls: seq[(int, int)]

  for r in 0..<R:
    for c in 0..<C:
      if m[r][c] != '@':
        continue

      var rcnt = 0
      for (rr, cc) in nbrs(r, c, R, C):
        if m[rr][cc] == '@':
          rcnt += 1

      if rcnt < 4:
        rolls.add((r, c))

  return rolls

proc p1(input: string): int =
  return rolls(input.splitLines()).len()


proc p2(input: string): int =
  var m = input.splitLines()

  var res = 0

  while true:
    var rmRolls = rolls(m)

    if rmRolls.len() == 0:
      break

    for (r, c) in rmRolls:
      m[r][c] = '.'

    res += rmRolls.len()

  return res

proc main() =
  var input = readFile("input.txt").strip()
  echo "part1=", p1(input)
  echo "part2=", p2(input)

main()
