vscode@f210d1c39533:/workspaces/gosl-benchmarks/la/oblas/dgemm$ time ./dgemm 
   size   |     Naive dgemm        (Dt) 
----------|-----------------------------
   2×   2 |  0.04 GFlops (       421ns)
   4×   4 |  0.27 GFlops (       482ns)
   6×   6 |  0.83 GFlops (       521ns)
   8×   8 |  1.86 GFlops (       550ns)
  10×  10 |  2.50 GFlops (       800ns)
  12×  12 |  3.99 GFlops (       866ns)
  14×  14 |  4.62 GFlops (     1.189µs)
  16×  16 |  6.59 GFlops (     1.243µs)
  18×  18 |  7.08 GFlops (     1.648µs)
  20×  20 |  8.91 GFlops (     1.795µs)
  22×  22 |  8.71 GFlops (     2.446µs)
  24×  24 |  9.77 GFlops (      2.83µs)
  26×  26 |  9.37 GFlops (     3.753µs)
  28×  28 | 10.78 GFlops (     4.074µs)
  30×  30 | 10.23 GFlops (     5.278µs)
  32×  32 | 11.92 GFlops (     5.496µs)
  34×  34 | 11.44 GFlops (      6.87µs)
  36×  36 | 12.74 GFlops (     7.326µs)
  38×  38 | 12.02 GFlops (     9.132µs)
  40×  40 | 18.04 GFlops (     7.096µs)
  42×  42 | 28.81 GFlops (     5.144µs)
  44×  44 | 29.73 GFlops (     5.731µs)
  46×  46 | 31.03 GFlops (     6.273µs)
  48×  48 | 24.92 GFlops (     8.877µs)
  50×  50 | 33.86 GFlops (     7.384µs)
  52×  52 | 26.21 GFlops (    10.729µs)
  54×  54 | 27.50 GFlops (    11.453µs)
  56×  56 | 36.11 GFlops (     9.727µs)
  58×  58 | 29.26 GFlops (    13.335µs)
  60×  60 | 31.40 GFlops (    13.759µs)
  62×  62 | 30.48 GFlops (     15.64µs)
  64×  64 | 34.31 GFlops (    15.282µs)
file <results/oblas-dgemm-small-1000samples.res> written

   size   |     Naive dgemm        (Dt) 
----------|-----------------------------
  16×  16 | 15.88 GFlops (       516ns)
  80×  80 | 45.82 GFlops (    22.348µs)
 144× 144 | 52.83 GFlops (   113.051µs)
 208× 208 | 53.54 GFlops (   336.132µs)
 272× 272 | 56.93 GFlops (   707.015µs)
 336× 336 | 59.20 GFlops (  1.281468ms)
 400× 400 | 60.22 GFlops (  2.125441ms)
 464× 464 | 61.82 GFlops (  3.231749ms)
 528× 528 | 59.63 GFlops (  4.936829ms)
 592× 592 | 61.90 GFlops (  6.703769ms)
 656× 656 | 64.11 GFlops (  8.806298ms)
 720× 720 | 65.06 GFlops ( 11.474285ms)
 784× 784 | 64.69 GFlops (  14.89955ms)
 848× 848 | 64.89 GFlops ( 18.794226ms)
 912× 912 | 65.34 GFlops ( 23.220089ms)
 976× 976 | 65.35 GFlops ( 28.451374ms)
1040×1040 | 62.13 GFlops ( 36.212343ms)
1104×1104 | 62.97 GFlops ( 42.740227ms)
1168×1168 | 63.61 GFlops ( 50.103101ms)
1232×1232 | 64.04 GFlops ( 58.402214ms)
1296×1296 | 64.65 GFlops ( 67.335953ms)
1360×1360 | 65.95 GFlops ( 76.280965ms)
file <results/oblas-dgemm-large-100samples.res> written

real    0m45.910s
user    0m46.235s
sys     0m0.433s