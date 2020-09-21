# gosl benchmarks

Computer time measurement with perf
```
perf stat mpirun -np 4 /tmp/gosl/test
```

Requirements
```
sudo apt install linux-tools-common linux-tools-generic
sudo echo "kernel.perf_event_paranoid = -1" >> /etc/sysctl.conf
sudo sysctl -p
```
