# robinhood

A go implementation of Robin Hood hashmap. Has 30% improvment of normal linear probing hashmap.

A detailed introduction can be found [here](https://code.mforever78.com/security/2016/09/24/robin-hood-hashmap/) (Chinese).

## Reproduce testing result

1. Generate queries from dictionary.

```bash
" use generator to generate proper inputs.
" the three number represent inserts, deletes, quries respectively.
cd generator
go run main.go > ../in
330000 2000 300000
```

2. Run robinhood hashmap.

```bash
cd robinhood
cp ../in .
/usr/bin/time go run main.go
```

3. Run linear hashmap.

```bash
cd linear
cp ../in .
/usr/bin/time go run main.go
```

4. (optional) Run std hashmap.

```bash
cd stdmap
cp ../in .
/usr/bin/time go run main.go
```

### Example test result

```bash
➜  robinhood /usr/bin/time ./main > out
7.05user 2.97system 0:09.03elapsed 110%CPU (0avgtext+0avgdata 48404maxresident)k
0inputs+19272outputs (0major+9894minor)pagefaults 0swaps
➜  robinhood /usr/bin/time ./linear > out 
11.54user 3.10system 0:13.25elapsed 110%CPU (0avgtext+0avgdata 48624maxresident)k
0inputs+19272outputs (0major+9906minor)pagefaults 0swaps
➜  robinhood /usr/bin/time ./stdmap > out 
3.75user 2.67system 0:05.75elapsed 111%CPU (0avgtext+0avgdata 59816maxresident)k
0inputs+19272outputs (0major+10906minor)pagefaults 0swaps
```
