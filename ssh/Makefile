bench:
	ragel -Z sshd.rl && go test -test.bench=.
	ragel -G0 -Z sshd.rl && go test -test.bench=.
	ragel -G1 -Z sshd.rl && go test -test.bench=.
	ragel -G2 -Z sshd.rl && go test -test.bench=.
