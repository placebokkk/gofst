#no need to use this
g++ -g -O2 -std=c++11 -o gofst.o -c gofst.cpp
ar r libgofst.so  gofst.o
