//gofst.cpp
#include <fst/fstlib.h>
#include "gofst.h"
#include <string>
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


//ArcSort
void ArcSortInput(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  ArcSort(fst_, ILabelCompare<StdArc>());
  //ArcSort(fst_, StdILabelCompare())
}

void ArcSortOutput(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  ArcSort(fst_, OLabelCompare<StdArc>());
  //ArcSort(fst_, StdOLabelCompare())
}
//operation
void Compose(CFst fst1, CFst fst2, CFst ofst)
{
  StdVectorFst * ifst1_ = (StdVectorFst*)fst1;
  StdVectorFst * ifst2_ = (StdVectorFst*)fst2;
  StdVectorFst * ofst_ = (StdVectorFst*)ofst;
  Compose(*ifst1_, *ifst2_, ofst_);
}

//IO
CFst FstRead(char* filename)
{
  string f_str = filename;
  StdVectorFst *fst_ = StdVectorFst::Read(f_str);
  return (void*)fst_;
}


void FstWrite(CFst fst, char* filename)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  string f_str = filename;
  fst_->Write(f_str);
  return;
}


//怎么实现StateIterator和ArcIterator