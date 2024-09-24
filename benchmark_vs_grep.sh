#!/usr/bin/bash
(go install \
&& date \
&& git rev-parse --abbrev-ref HEAD \
&& git rev-parse HEAD \
&& hyperfine 'grep a test_data/3M' \
&& hyperfine 'go-grep a test_data/3M' \
&& echo '===============================') \
&>> benchmark_vs_grep_stats.log