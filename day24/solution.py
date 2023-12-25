import re
import numpy as np

data = open('input.txt').read().split('\n')

orig_left = 200000000000000
orig_right = 400000000000000

pts = [[int(y) for y in re.findall(r'-?[0-9]+', x) ]for x in data]

#outputs one row for the matrix A as described, and one entry of the const column vector
def calcs(x1,y1,dx1,dy1,x2,y2,dx2,dy2):
    return([dy2 - dy1, dx1 - dx2, y2 - y1, x2 - x1], y1 * dx1 - y2 * dx2 + x2 * dy2 - x1 * dy1 )


#define a shift for the first 8 vectors to help with numpy roundoff errors
shift = min([ x for  y in pts[:8] for x in y[:3] ])

#shift the first 8 vectors
for i in range(8):
    for j in range(3):
        pts[i][j] -= shift

rows1, rows2 = [], []
col1, col2 = [], []
ans = []

for i in range(0,8,2):
    # get rows of A in solving for 
    row, num = calcs(*(pts[i][:2]+pts[i][3:5]), *((pts[i + 1][:2]+pts[i + 1][3:5])))
    rows1.append(row)
    col1.append(num)
    # populate solving for a, c, d, f
    row, num = calcs(*([pts[i][0], pts[i][2], pts[i][3], pts[i][5]]), *([pts[i + 1][0], pts[i + 1][2], pts[i + 1][3], pts[i + 1][5]]))
    rows2.append(row)
    col2.append(num)

A = np.array(rows1)
col = np.array(col1)
# (a, b, e, f)
ans1 = np.linalg.solve(A, col)

A = np.array(rows2)
col = np.array(col2)
# (a, c, d, f) : 
ans2 = np.linalg.solve(A, col) 

#add up a, b, c, and shift back (3 * shift) to help with the roundoff errors
print( round(ans1[0]) + round(ans1[1]) + round(ans2[1]) + 3 * shift)
