#!/usr/bin/env bash
BC_LINE_LENGTH=0 bc -q <<end
0;1; for (i=1;i<100000;i++) {g=last;last+=f;f=g;}
last
end
