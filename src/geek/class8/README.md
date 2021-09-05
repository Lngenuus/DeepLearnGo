# class8

## 1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能

redis get 10 bytes

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 10 -t get
====== GET ======
  10000 requests completed in 0.29 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.39% <= 0.2 milliseconds
0.95% <= 0.3 milliseconds
1.63% <= 0.4 milliseconds
2.31% <= 0.5 milliseconds
3.04% <= 0.6 milliseconds
9.68% <= 0.7 milliseconds
91.60% <= 0.8 milliseconds
93.07% <= 0.9 milliseconds
94.21% <= 1.0 milliseconds
95.12% <= 1.1 milliseconds
95.76% <= 1.2 milliseconds
96.29% <= 1.3 milliseconds
96.80% <= 1.4 milliseconds
97.23% <= 1.5 milliseconds
97.69% <= 1.6 milliseconds
98.08% <= 1.7 milliseconds
98.50% <= 1.8 milliseconds
98.98% <= 1.9 milliseconds
99.31% <= 2 milliseconds
99.73% <= 3 milliseconds
99.98% <= 4 milliseconds
100.00% <= 4 milliseconds
34246.57 requests per second
```

redis get 10 bytes

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 10 -t set
====== SET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.38% <= 0.2 milliseconds
0.90% <= 0.3 milliseconds
1.48% <= 0.4 milliseconds
2.14% <= 0.5 milliseconds
2.90% <= 0.6 milliseconds
69.01% <= 0.7 milliseconds
92.98% <= 0.8 milliseconds
94.21% <= 0.9 milliseconds
94.86% <= 1.0 milliseconds
95.48% <= 1.1 milliseconds
95.99% <= 1.2 milliseconds
96.74% <= 1.3 milliseconds
97.45% <= 1.4 milliseconds
97.98% <= 1.5 milliseconds
98.49% <= 1.6 milliseconds
98.96% <= 1.7 milliseconds
99.39% <= 1.8 milliseconds
99.65% <= 1.9 milliseconds
99.81% <= 2 milliseconds
100.00% <= 2 milliseconds
35714.29 requests per second
```

redis get 20 bytes

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 20 -t get
====== GET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.68% <= 0.2 milliseconds
1.43% <= 0.3 milliseconds
2.29% <= 0.4 milliseconds
3.12% <= 0.5 milliseconds
4.14% <= 0.6 milliseconds
62.58% <= 0.7 milliseconds
90.46% <= 0.8 milliseconds
92.14% <= 0.9 milliseconds
93.31% <= 1.0 milliseconds
94.19% <= 1.1 milliseconds
95.02% <= 1.2 milliseconds
95.83% <= 1.3 milliseconds
96.69% <= 1.4 milliseconds
97.48% <= 1.5 milliseconds
98.24% <= 1.6 milliseconds
98.97% <= 1.7 milliseconds
99.52% <= 1.8 milliseconds
99.83% <= 1.9 milliseconds
99.89% <= 2 milliseconds
100.00% <= 2 milliseconds
35714.29 requests per second
```

redis 20 bytes set

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 20 -t set
====== SET ======
  10000 requests completed in 0.30 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.32% <= 0.2 milliseconds
0.75% <= 0.3 milliseconds
1.28% <= 0.4 milliseconds
2.85% <= 0.5 milliseconds
5.50% <= 0.6 milliseconds
51.59% <= 0.7 milliseconds
74.76% <= 0.8 milliseconds
77.98% <= 0.9 milliseconds
80.62% <= 1.0 milliseconds
82.88% <= 1.1 milliseconds
85.02% <= 1.2 milliseconds
87.27% <= 1.3 milliseconds
89.41% <= 1.4 milliseconds
91.54% <= 1.5 milliseconds
93.59% <= 1.6 milliseconds
95.65% <= 1.7 milliseconds
97.44% <= 1.8 milliseconds
98.14% <= 1.9 milliseconds
98.60% <= 2 milliseconds
99.76% <= 3 milliseconds
100.00% <= 3 milliseconds
33003.30 requests per second
```

redis 50 bytes get

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 50 -t get
====== GET ======
  10000 requests completed in 0.31 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.49% <= 0.2 milliseconds
1.04% <= 0.3 milliseconds
1.69% <= 0.4 milliseconds
3.42% <= 0.5 milliseconds
6.15% <= 0.6 milliseconds
47.26% <= 0.7 milliseconds
74.23% <= 0.8 milliseconds
77.54% <= 0.9 milliseconds
80.24% <= 1.0 milliseconds
82.96% <= 1.1 milliseconds
85.58% <= 1.2 milliseconds
87.97% <= 1.3 milliseconds
90.32% <= 1.4 milliseconds
92.60% <= 1.5 milliseconds
94.64% <= 1.6 milliseconds
96.56% <= 1.7 milliseconds
98.21% <= 1.8 milliseconds
98.65% <= 1.9 milliseconds
98.92% <= 2 milliseconds
99.51% <= 3 milliseconds
99.75% <= 4 milliseconds
99.99% <= 5 milliseconds
100.00% <= 5 milliseconds
32786.88 requests per second
```

redis 50 bytes set

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 50 -t set
====== SET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.55% <= 0.2 milliseconds
1.29% <= 0.3 milliseconds
2.09% <= 0.4 milliseconds
2.91% <= 0.5 milliseconds
3.73% <= 0.6 milliseconds
69.06% <= 0.7 milliseconds
91.13% <= 0.8 milliseconds
93.08% <= 0.9 milliseconds
94.11% <= 1.0 milliseconds
94.97% <= 1.1 milliseconds
95.65% <= 1.2 milliseconds
96.26% <= 1.3 milliseconds
96.82% <= 1.4 milliseconds
97.41% <= 1.5 milliseconds
97.93% <= 1.6 milliseconds
98.53% <= 1.7 milliseconds
99.02% <= 1.8 milliseconds
99.37% <= 1.9 milliseconds
99.64% <= 2 milliseconds
100.00% <= 2 milliseconds
35587.19 requests per second
```

redis 100 bytes get

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 100 -t get
====== GET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.39% <= 0.2 milliseconds
0.88% <= 0.3 milliseconds
1.49% <= 0.4 milliseconds
2.10% <= 0.5 milliseconds
2.87% <= 0.6 milliseconds
57.08% <= 0.7 milliseconds
92.78% <= 0.8 milliseconds
94.01% <= 0.9 milliseconds
94.77% <= 1.0 milliseconds
95.37% <= 1.1 milliseconds
95.90% <= 1.2 milliseconds
96.42% <= 1.3 milliseconds
96.86% <= 1.4 milliseconds
97.28% <= 1.5 milliseconds
97.68% <= 1.6 milliseconds
98.14% <= 1.7 milliseconds
98.54% <= 1.8 milliseconds
98.89% <= 1.9 milliseconds
99.11% <= 2 milliseconds
99.64% <= 3 milliseconds
99.94% <= 4 milliseconds
100.00% <= 4 milliseconds
35335.69 requests per second
```

redis 100 bytes set

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 100 -t set
====== SET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.25% <= 0.2 milliseconds
0.64% <= 0.3 milliseconds
1.03% <= 0.4 milliseconds
1.52% <= 0.5 milliseconds
1.99% <= 0.6 milliseconds
68.91% <= 0.7 milliseconds
94.87% <= 0.8 milliseconds
96.06% <= 0.9 milliseconds
96.63% <= 1.0 milliseconds
97.17% <= 1.1 milliseconds
97.59% <= 1.2 milliseconds
97.99% <= 1.3 milliseconds
98.29% <= 1.4 milliseconds
98.58% <= 1.5 milliseconds
98.89% <= 1.6 milliseconds
99.15% <= 1.7 milliseconds
99.44% <= 1.8 milliseconds
99.61% <= 1.9 milliseconds
99.69% <= 2 milliseconds
99.98% <= 3 milliseconds
100.00% <= 3 milliseconds
35714.29 requests per second
```

redis 200 bytes get

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 200 -t get
====== GET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.36% <= 0.2 milliseconds
0.80% <= 0.3 milliseconds
1.31% <= 0.4 milliseconds
1.83% <= 0.5 milliseconds
2.41% <= 0.6 milliseconds
56.45% <= 0.7 milliseconds
93.15% <= 0.8 milliseconds
94.36% <= 0.9 milliseconds
94.96% <= 1.0 milliseconds
95.51% <= 1.1 milliseconds
95.95% <= 1.2 milliseconds
96.43% <= 1.3 milliseconds
96.88% <= 1.4 milliseconds
97.29% <= 1.5 milliseconds
97.75% <= 1.6 milliseconds
98.16% <= 1.7 milliseconds
98.51% <= 1.8 milliseconds
98.85% <= 1.9 milliseconds
99.05% <= 2 milliseconds
99.69% <= 3 milliseconds
99.96% <= 4 milliseconds
100.00% <= 4 milliseconds
35211.27 requests per second
```

redis 200 bytes set

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 200 -t set
====== SET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.08% <= 0.2 milliseconds
0.21% <= 0.3 milliseconds
0.34% <= 0.4 milliseconds
0.46% <= 0.5 milliseconds
0.61% <= 0.6 milliseconds
69.57% <= 0.7 milliseconds
98.09% <= 0.8 milliseconds
98.37% <= 0.9 milliseconds
98.57% <= 1.0 milliseconds
98.74% <= 1.1 milliseconds
98.88% <= 1.2 milliseconds
98.99% <= 1.3 milliseconds
99.08% <= 1.4 milliseconds
99.20% <= 1.5 milliseconds
99.34% <= 1.6 milliseconds
99.50% <= 1.7 milliseconds
99.62% <= 1.8 milliseconds
99.73% <= 1.9 milliseconds
99.83% <= 2 milliseconds
100.00% <= 2 milliseconds
35714.29 requests per second
```

redis 1000 bytes get

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 1000 -t get
====== GET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.59% <= 0.2 milliseconds
1.19% <= 0.3 milliseconds
1.89% <= 0.4 milliseconds
2.63% <= 0.5 milliseconds
3.39% <= 0.6 milliseconds
56.16% <= 0.7 milliseconds
92.51% <= 0.8 milliseconds
93.88% <= 0.9 milliseconds
94.89% <= 1.0 milliseconds
95.69% <= 1.1 milliseconds
96.30% <= 1.2 milliseconds
96.82% <= 1.3 milliseconds
97.34% <= 1.4 milliseconds
97.84% <= 1.5 milliseconds
98.28% <= 1.6 milliseconds
98.73% <= 1.7 milliseconds
99.27% <= 1.8 milliseconds
99.61% <= 1.9 milliseconds
99.90% <= 2 milliseconds
100.00% <= 2 milliseconds
35714.29 requests per second
```

redis 1000 bytes set

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 1000 -t set
====== SET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.13% <= 0.2 milliseconds
0.36% <= 0.3 milliseconds
0.66% <= 0.4 milliseconds
0.99% <= 0.5 milliseconds
1.38% <= 0.6 milliseconds
60.55% <= 0.7 milliseconds
95.40% <= 0.8 milliseconds
96.07% <= 0.9 milliseconds
96.46% <= 1.0 milliseconds
96.90% <= 1.1 milliseconds
97.22% <= 1.2 milliseconds
97.51% <= 1.3 milliseconds
97.81% <= 1.4 milliseconds
98.14% <= 1.5 milliseconds
98.41% <= 1.6 milliseconds
98.64% <= 1.7 milliseconds
98.92% <= 1.8 milliseconds
99.17% <= 1.9 milliseconds
99.34% <= 2 milliseconds
100.00% <= 2 milliseconds
35211.27 requests per second
```

redis 5000 bytes get

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 5000 -t get
====== GET ======
  10000 requests completed in 0.28 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.14% <= 0.2 milliseconds
0.35% <= 0.3 milliseconds
0.59% <= 0.4 milliseconds
0.82% <= 0.5 milliseconds
1.05% <= 0.6 milliseconds
46.39% <= 0.7 milliseconds
97.20% <= 0.8 milliseconds
98.08% <= 0.9 milliseconds
98.49% <= 1.0 milliseconds
98.74% <= 1.1 milliseconds
98.92% <= 1.2 milliseconds
99.09% <= 1.3 milliseconds
99.25% <= 1.4 milliseconds
99.42% <= 1.5 milliseconds
99.61% <= 1.6 milliseconds
99.74% <= 1.7 milliseconds
99.87% <= 1.8 milliseconds
99.93% <= 1.9 milliseconds
99.96% <= 2 milliseconds
100.00% <= 2 milliseconds
35460.99 requests per second
```

redis 5000 bytes set

```bash
root@6ae02fa439aa:/data# redis-benchmark -h 127.0.0.1 -p 6379 -c 50 -n 10000 -d 5000 -t set
====== SET ======
  10000 requests completed in 0.29 seconds
  50 parallel clients
  5000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

0.01% <= 0.1 milliseconds
0.09% <= 0.2 milliseconds
0.41% <= 0.3 milliseconds
0.77% <= 0.4 milliseconds
1.15% <= 0.5 milliseconds
1.67% <= 0.6 milliseconds
27.70% <= 0.7 milliseconds
94.20% <= 0.8 milliseconds
95.97% <= 0.9 milliseconds
96.50% <= 1.0 milliseconds
96.93% <= 1.1 milliseconds
97.41% <= 1.2 milliseconds
97.75% <= 1.3 milliseconds
98.04% <= 1.4 milliseconds
98.29% <= 1.5 milliseconds
98.57% <= 1.6 milliseconds
98.79% <= 1.7 milliseconds
99.00% <= 1.8 milliseconds
99.29% <= 1.9 milliseconds
99.46% <= 2 milliseconds
99.86% <= 3 milliseconds
100.00% <= 3 milliseconds
34843.21 requests per second
```

## 2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间

对不同value大小的数据添加 观测info memory 得出结论 相同长度的value在写入数量越多情况下，平均每个value占用内存更多
