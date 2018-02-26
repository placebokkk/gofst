//gofst.cpp

#include <fst/fstlib.h>
#include "gofst.h"
#include <string>

using namespace fst;

//FST API
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

int FstGetStart(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  return fst_->Start();
}  

int FstIsFinal(CFst fst, int state)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  if (fst_->Final(state) != StdArc::Weight::Zero())
  {
    return 1;
  }
  else
  {
    return 0;
  }
}

void FstSetFinal(CFst fst, int state, float weight)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  fst_->SetFinal(state, weight);
}

void FstAddArc(CFst fst, int state, CArc arc)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  fst_->AddArc(state, * (StdArc *)(arc) );
}

CFst FstCopy(CFst ifst)
{
  StdVectorFst * ifst_ = (StdVectorFst*)ifst;
  return (CFst)((*ifst_).Copy());
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

void FstRmEpsilon(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  RmEpsilon(fst_);
}

void FstInvert(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  Invert(fst_);
}

void FstMinimize(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  Minimize(fst_);
}

void FstShortestPath(CFst fst,  CFst ofst, int n)
{
  StdVectorFst * ifst_ = (StdVectorFst*)fst;
  StdVectorFst * ofst_ = (StdVectorFst*)ofst;
  ShortestPath(*ifst_, ofst_, n);
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


int FstWrite(CFst fst, char* filename)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  string f_str = filename;
  return fst_->Write(f_str);
}

void FstSetInputSymbols(CFst fst, CSymbolTable st){
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  SymbolTable * st_ = (SymbolTable * ) st;
  fst_->SetInputSymbols(st_);
}

CSymbolTable FstInputSymbols(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  return (CSymbolTable) fst_->InputSymbols();
}

void FstSetOutputSymbols(CFst fst, CSymbolTable st)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  SymbolTable * st_ = (SymbolTable * ) st;
  fst_->SetOutputSymbols(st_);
}

CSymbolTable FstOutputSymbols(CFst fst)
{
  StdVectorFst * fst_ = (StdVectorFst*)fst;
  return (CSymbolTable) fst_->OutputSymbols();
}

//SymbolTable API
CSymbolTable SymbolTableInit()
{
  SymbolTable *st = new SymbolTable();
  return (CSymbolTable)st;
}


int SymbolTableEqual(CSymbolTable st1, CSymbolTable st2)
{
  SymbolTable * st1_ = (SymbolTable *) st1;
  SymbolTable * st2_ = (SymbolTable *) st2;
  if( st1_->LabeledCheckSum() == st2_->LabeledCheckSum() )
  {
    return 1;
  }
  else
  {
    return 0;
  }
}

int SymbolTableWrite(CSymbolTable st, char* filename)
{
  SymbolTable * st_ = (SymbolTable*) st;
  string f_str = filename;
  return st_->Write(f_str);
}

int SymbolTableFindKey(CSymbolTable st, char *symbol)
{
  SymbolTable * st_ = (SymbolTable*) st;
  return st_->Find(symbol);
}
char* SymbolTableFindSymbol(CSymbolTable st, int key)
{
  
  SymbolTable * st_ = (SymbolTable*) st;
  string symbol = st_->Find(key);
  int l = symbol.length() + 1;
  char *c =  new char[l];
  strcpy(c, symbol.c_str());
  return c;
}

int SymbolTableHasKey(CSymbolTable st, int key){
  SymbolTable * st_ = (SymbolTable*) st;
  if( st_->Member(key) )
  {
    return 1;
  }
  else
  {
    return 0;
  }
}
int SymbolTableHasSymbol(CSymbolTable st, char *symbol){
  SymbolTable * st_ = (SymbolTable*) st;
  if( st_->Member(symbol) )
  {
    return 1;
  }
  else
  {
    return 0;
  }
}


int SymbolTableAddSymbol(CSymbolTable st, char *symbol)
{
   SymbolTable * st_ = (SymbolTable*) st;
   return st_->AddSymbol(symbol);
}

int SymbolTableAddSymbolKey(CSymbolTable st, char *symbol, int key)
{
   SymbolTable * st_ = (SymbolTable*) st;
   st_->AddSymbol(symbol, key);
   return key;
}

CSymbolTable SymbolTableReadText(char *filename)
{
  string f_str = filename;
  SymbolTable * st_ = SymbolTable::ReadText(f_str);
  return (CSymbolTable)st_;
}

CSymbolTable SymbolTableReadBinary(char *filename)
{
  string f_str = filename;
  SymbolTable * st_ = SymbolTable::Read(f_str);
  return (CSymbolTable)st_;
}

void FreeString(char * c){
  delete [] c;
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