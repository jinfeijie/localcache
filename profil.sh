#!/bin/bash
go test -bench=. -benchtime="3s" -cpuprofile profile_cpu.out
go tool pprof -pdf profile_cpu.out > profile_cpu.pdf