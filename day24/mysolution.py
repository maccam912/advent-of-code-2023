from dataclasses import dataclass
from z3 import *

@dataclass
class Hailstone:
    x: int
    y: int
    z: int
    vx: int
    vy: int
    vz: int

hailstones = []
with open("input.txt") as f:
    lines = f.readlines()
    for line in lines[:20]:
        parts = line.strip().split("@")
        pos = parts[0]
        vel = parts[1]
        pos_coords = parts[0].split(",")
        vel_coords = parts[1].split(",")
        x = int(pos_coords[0].strip())
        y = int(pos_coords[1].strip())
        z = int(pos_coords[2].strip())
        vx = int(vel_coords[0].strip())
        vy = int(vel_coords[1].strip())
        vz = int(vel_coords[2].strip())
        hailstone = Hailstone(x, y, z, vx, vy, vz)
        hailstones += [hailstone]

x = Real('x')
y = Real('y')
z = Real('z')
vx = Real('vx')
vy = Real('vy')
vz = Real('vz')

s = Solver()
for i, hailstone in enumerate(hailstones):
    t = Real(f"t_{i}")
    s.add(x + vx*t == float(hailstone.x) + float(hailstone.vx)*t)
    s.add(y + vy*t == float(hailstone.y) + float(hailstone.vy)*t)
    s.add(z + vz*t == float(hailstone.z) + float(hailstone.vz)*t)

s.check()
print(s.model())
# model = s.model()
# for var in s.model():
#     if "t" not in var.name():
#         model.evaluate()
#         print(f"{var} = {}")