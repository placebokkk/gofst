//gofst.cpp

#include <fst/fstlib.h>
#include "gofst.h"
#include <string>

using namespace fst;

CFst FstInit()
{
  StdVectorFst *fst_ = new StdVectorFst();
  return (CFst)fst_;
}

void FstFree(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  delete fst_;
}

void FstAddState(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  fst_->AddState();
}

void FstSetStart(CFst fst, int start_state)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  fst_->SetStart(start_state);
}

void FstAddArc(CFst fst, int state, CArc arc)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  fst_->AddArc(state, * (StdArc *)(arc) );
}


//operation
void FstCompose(CFst fst1, CFst fst2, CFst ofst)
{
  StdVectorFst * ifst1_ = (StdVectorFst*)fst1;
  StdVectorFst * ifst2_ = (StdVectorFst*)fst2;
  StdVectorFst * ofst_ = (StdVectorFst*)ofst;
  Compose(*ifst1_, *ifst2_, ofst_);
}

void FstDeterminize(CFst fst,  CFst ofst)
{
  StdVectorFst * ifst_ = (StdVectorFst*)fst;
  StdVectorFst * ofst_ = (StdVectorFst*)ofst;
  Determinize(*ifst_, ofst_);
}
//ArcSort
void FstArcSortInput(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  ArcSort(fst_, ILabelCompare<StdArc>());
  //ArcSort(fst_, StdILabelCompare())
}

void FstArcSortOutput(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  ArcSort(fst_, OLabelCompare<StdArc>());
  //ArcSort(fst_, StdOLabelCompare())
}
//IO
CFst FstRead(char* filename)
{
  string f_str = filename;
  StdVectorFst *fst_ = StdVectorFst::Read(f_str);
  return (CFst)fst_;
}


void FstWrite(CFst fst, char* filename)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  string f_str = filename;
  fst_->Write(f_str);
  return;
}


//怎么实现StateIterator

//   for (StateIterator<StdFst> siter(fst);
//        !siter.Done();
//        siter.Next()) {
//     StateId s = siter.Value();
//    }

//StateIterator<StdFst>可以改为StateIterator<Fst>吗?
CStateIterator StateIteratorInit(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  StateIterator<StdFst> * siter = new StateIterator<StdFst>(*fst_);
  return (CStateIterator)siter;
}

void StateIteratorNext(CStateIterator si)
{
  StateIterator<StdFst> * siter = (StateIterator<StdFst> *)si;
  siter->Next();
  return;
}

CStateId StateIteratorValue(CStateIterator si)
{
  StateIterator<StdFst> * siter = (StateIterator<StdFst> *)si;
  return (CStateId)siter->Value();
  
}
int StateIteratorDone(CStateIterator si)
{
  StateIterator<StdFst> * siter = (StateIterator<StdFst> *)si;
  return (int) ((CStateId)siter->Done());
  
}


CArc ArcInit(int ilabel,int olabel,float weight,int state_id)
{
  StdArc * arc = new StdArc(ilabel, olabel, weight, state_id);
  return (CArc*) arc;
}

int ArcGetILabel(CArc arc) {
  return ((StdArc *)arc)->ilabel;
}
int ArcGetOLabel(CArc arc) {
  return ((StdArc *)arc)->olabel;
}
float ArcGetWeight(CArc arc) {
  return ((StdArc *)arc)->weight.Value();
}
int ArcGetNextState(CArc arc) {
  return ((StdArc *)arc)->nextstate;
}

//怎么实现ArcIterator

// for (StateIterator<Fst<Arc>> siter(fst); !siter.Done(); siter.Next()) {
//     ++nstates_;
//     const auto s = siter.Value();
//     for (ArcIterator<Fst<Arc>> aiter(fst, s); !aiter.Done(); aiter.Next()) {
//       const auto &arc = aiter.Value();
// }


//需要用StdArc吗?
CArcIterator ArcIteratorInit(CFst fst, CStateId state_id)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  ArcIterator<Fst<StdArc>> * aiter = new ArcIterator<Fst<StdArc>>(*fst_, state_id);
  return (CArcIterator)aiter;
}

void ArcIteratorNext(CArcIterator ai)
{
  ArcIterator<Fst<StdArc>> * aiter = (ArcIterator<Fst<StdArc>> *)ai;
  aiter->Next();
  return;
}

CArc ArcIteratorValue(CArcIterator ai)
{
  ArcIterator<Fst<StdArc>> * aiter = (ArcIterator<Fst<StdArc>> *)ai;
  return (CArc) &(aiter->Value());
  
}

int ArcIteratorDone(CArcIterator ai)
{
  ArcIterator<Fst<StdArc>> * aiter = (ArcIterator<Fst<StdArc>> *)ai;
  return (int) aiter->Done();
}