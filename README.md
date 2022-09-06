# cache-bench

```
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
BenchmarkKodingCache/Set_sm-4         	 1000000	      2057 ns/op	     440 B/op	       5 allocs/op
BenchmarkKodingCache/Get_sm-4         	 1873302	       644.4 ns/op	      23 B/op	       1 allocs/op
BenchmarkKodingCache/Set_md-4         	 1000000	      2162 ns/op	     624 B/op	       5 allocs/op
BenchmarkKodingCache/Get_md-4         	 1974748	       628.8 ns/op	      31 B/op	       1 allocs/op
BenchmarkKodingCache/Set_lg-4         	  113661	     15844 ns/op	   12742 B/op	       5 allocs/op
BenchmarkKodingCache/Get_lg-4         	 1800506	       643.9 ns/op	      55 B/op	       1 allocs/op
BenchmarkHashicorpLRU/Set_sm-4        	 1740830	       646.2 ns/op	     182 B/op	       8 allocs/op
BenchmarkHashicorpLRU/Get_sm-4        	 5374107	       226.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkHashicorpLRU/Set_md-4        	 1446760	       850.4 ns/op	     366 B/op	       8 allocs/op
BenchmarkHashicorpLRU/Get_md-4        	 5685121	       228.0 ns/op	      48 B/op	       3 allocs/op
BenchmarkHashicorpLRU/Set_lg-4        	  113263	     10107 ns/op	   12479 B/op	       8 allocs/op
BenchmarkHashicorpLRU/Get_lg-4        	 4894177	       244.4 ns/op	      72 B/op	       3 allocs/op
BenchmarkCache2Go/Set_sm-4            	 1000000	      1881 ns/op	     386 B/op	       9 allocs/op
BenchmarkCache2Go/Get_sm-4            	 2486300	       473.8 ns/op	      39 B/op	       2 allocs/op
BenchmarkCache2Go/Set_md-4            	 1000000	      2229 ns/op	     571 B/op	       9 allocs/op
BenchmarkCache2Go/Get_md-4            	 2434920	       470.3 ns/op	      47 B/op	       2 allocs/op
BenchmarkCache2Go/Set_lg-4            	   93398	     19940 ns/op	   12552 B/op	       8 allocs/op
BenchmarkCache2Go/Get_lg-4            	 2715637	       394.2 ns/op	      71 B/op	       2 allocs/op
BenchmarkGoCache/Set_sm-4             	 1000000	      1107 ns/op	     279 B/op	       5 allocs/op
BenchmarkGoCache/Get_sm-4             	 3468182	       346.0 ns/op	      23 B/op	       1 allocs/op
BenchmarkGoCache/Set_md-4             	 1000000	      1344 ns/op	     463 B/op	       5 allocs/op
BenchmarkGoCache/Get_md-4             	 3386599	       332.3 ns/op	      31 B/op	       1 allocs/op
BenchmarkGoCache/Set_lg-4             	  127783	     15569 ns/op	   12559 B/op	       5 allocs/op
BenchmarkGoCache/Get_lg-4             	 4393021	       255.9 ns/op	      55 B/op	       1 allocs/op
BenchmarkFreecache/Set_sm-4           	 1473680	       827.9 ns/op	      64 B/op	       3 allocs/op
BenchmarkFreecache/Get_sm-4           	 4988721	       238.1 ns/op	      24 B/op	       2 allocs/op
BenchmarkFreecache/Set_md-4           	 1000000	      1170 ns/op	     465 B/op	       4 allocs/op
BenchmarkFreecache/Get_md-4           	 5745067	       220.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkFreecache/Set_lg-4           	  106264	     15246 ns/op	   24688 B/op	       5 allocs/op
BenchmarkFreecache/Get_lg-4           	 5141546	       240.8 ns/op	     103 B/op	       2 allocs/op
BenchmarkBigCache/Set_sm-4            	 1586569	       792.3 ns/op	      67 B/op	       4 allocs/op
BenchmarkBigCache/Get_sm-4            	 2495007	       422.8 ns/op	      54 B/op	       3 allocs/op
BenchmarkBigCache/Set_md-4            	 1000000	      1173 ns/op	     491 B/op	       5 allocs/op
BenchmarkBigCache/Get_md-4            	 2197276	       488.1 ns/op	     137 B/op	       2 allocs/op
BenchmarkBigCache/Set_lg-4            	   76413	     16387 ns/op	   24818 B/op	       5 allocs/op
BenchmarkBigCache/Get_lg-4            	 2651631	       380.0 ns/op	     175 B/op	       2 allocs/op
BenchmarkGCache/Set_sm-4              	 1994841	       594.6 ns/op	     207 B/op	       7 allocs/op
BenchmarkGCache/Get_sm-4              	 5640460	       222.2 ns/op	      39 B/op	       2 allocs/op
BenchmarkGCache/Set_md-4              	 1514485	       793.2 ns/op	     391 B/op	       7 allocs/op
BenchmarkGCache/Get_md-4              	 5498305	       234.0 ns/op	      47 B/op	       2 allocs/op
BenchmarkGCache/Set_lg-4              	   97742	     10358 ns/op	   12496 B/op	       7 allocs/op
BenchmarkGCache/Get_lg-4              	 5206677	       245.9 ns/op	      72 B/op	       2 allocs/op
BenchmarkSyncMap/Set_sm-4             	 1000000	      1415 ns/op	     258 B/op	       9 allocs/op
BenchmarkSyncMap/Get_sm-4             	 3532458	       335.4 ns/op	      23 B/op	       1 allocs/op
BenchmarkSyncMap/Set_md-4             	 1000000	      1854 ns/op	     442 B/op	       9 allocs/op
BenchmarkSyncMap/Get_md-4             	 3769701	       318.5 ns/op	      31 B/op	       1 allocs/op
BenchmarkSyncMap/Set_lg-4             	  107757	     14111 ns/op	   12556 B/op	       8 allocs/op
BenchmarkSyncMap/Get_lg-4             	 4440552	       258.9 ns/op	      55 B/op	       1 allocs/op
BenchmarkMap/Set_sm-4                 	 2017860	       826.7 ns/op	     142 B/op	       4 allocs/op
BenchmarkMap/Get_sm-4                 	 3364545	       359.2 ns/op	      23 B/op	       1 allocs/op
BenchmarkMap/Set_md-4                 	 1000000	      1169 ns/op	     409 B/op	       4 allocs/op
BenchmarkMap/Get_md-4                 	 3866716	       294.0 ns/op	      31 B/op	       1 allocs/op
BenchmarkMap/Set_lg-4                 	  116473	     13946 ns/op	   12515 B/op	       4 allocs/op
BenchmarkMap/Get_lg-4                 	 4472932	       263.5 ns/op	      55 B/op	       1 allocs/op
```
