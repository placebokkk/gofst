//gofst.h
//This file only supports StdVectorFst, StdArc


//fst.AddArc(1, StdArc(3, 3, 2.5, 2));
    
//fst.SetFinal(2, 3.5);

//fst.Write("binary.fst");

#ifndef GO_FST_H
#define GO_FST_H


#ifdef __cplusplus
extern "C" {
#endif

  /**********Fst**********/
  typedef void* CFst;
  typedef void* CArc;
  //fst init/free
  CFst FstInit(void);
  void FstFree(CFst fst);

  //fst basic 
  void FstAddState(CFst fst);
  void FstSetStart(CFst fst, int start_state);
  void FstAddArc(CFst fst, int start_state, CArc arc);

  //operation
  void FstCompose(CFst fst1, CFst fst2, CFst ofst);
  void FstDeterminize(CFst fst1,  CFst ofst);
  
  void FstArcSortInput(CFst fst);
  void FstArcSortOutput(CFst fst);

  //I/O
  CFst FstRead(char* filename);
  void FstWrite(CFst fst, char* filename);


  /**********Symboltable**********/
  typedef void* CSymbolTable;


  /**********State Iterator**********/
  //Use CStateId or just use int or int64?
  typedef int CStateId;
  typedef void* CStateIterator;
  CStateIterator StateIteratorInit(CFst fst);
  void StateIteratorNext(CStateIterator si);
  CStateId StateIteratorValue(CStateIterator si);
  int StateIteratorDone(CStateIterator si);

  //Do not export below.
  //  void StateIteratorReset() {}
  //  void StateIteratorSeek(size_t a) {}
  //  size_t StateIteratorPosition() {  }


  /**********Arc**********/
  CArc ArcInit(int ilabel,int olabel,float weight,int state_id);

  int ArcGetILabel(CArc arc);
  int ArcGetOLabel(CArc arc);
  float ArcGetWeight(CArc arc);
  int ArcGetNextState(CArc arc);
  /**********Arc Iterator**********/
  
  typedef void* CArcIterator;
  CArcIterator ArcIteratorInit(CFst fst, int state_id);
  void ArcIteratorNext(CArcIterator ai);
  CArc ArcIteratorValue(CArcIterator ai);
  int ArcIteratorDone(CArcIterator ai);

  //Do not export below.
  //  void ArcIteratorReset() {}
  //  void ArcIteratorSeek(size_t a) {}
  //  size_t ArcIteratorPosition() {  }


#ifdef __cplusplus
}
#endif

#endif
