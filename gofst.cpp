//gofst.cpp
#include <fst/fstlib.h>
#include "gofst.h"

using namespace fst;
CFst FstInit()
{
  StdVectorFst *fst_ = new StdVectorFst();
  return (void*)fst_;
}
void FstFree(CFst f)
{
  StdVectorFst * fst_ = (StdVectorFst*)f;
  delete fst_;
}

void AddState(CFst f)
{
  StdVectorFst * fst_ = (StdVectorFst*)f;
  fst_->AddState();
}

void SetStart(CFst f, int startState)
{
  StdVectorFst * fst_ = (StdVectorFst*)f;
  fst_->SetStart(startState);
}
