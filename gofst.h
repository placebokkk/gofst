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

  /**********Declare**********/
  typedef void* CFst;
  typedef void* CArc;
  typedef void* CSymbolTable;

  /**********Fst**********/
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

  void FstSetInputSymbols(CFst fst, CSymbolTable st);
  CSymbolTable FstInputSymbols(CFst fst);
  void FstSetOutputSymbols(CFst fst, CSymbolTable st);
  CSymbolTable FstOutputSymbols(CFst fst);

  //I/O
  CFst FstRead(char* filename);
  int FstWrite(CFst fst, char* filename);


  /**********Symboltable**********/

  int SymbolTableEqual(CSymbolTable st1, CSymbolTable st2);

  int SymbolTableWrite(CSymbolTable st, char *filename);

  //In C++ openfst, Find use for finding both key and symbol,by parameter overloading.
  //In C wrapper, we have to declare two different interface.
  int SymbolTableFindKey(CSymbolTable st, char *symbol);
  char* SymbolTableFindSymbol(CSymbolTable st, int key);
  int SymbolTableHasKey(CSymbolTable st, int key);
  int SymbolTableHasSymbol(CSymbolTable st, char *symbol);

  void FreeString(char * c);

  CSymbolTable SymbolTableReadText(char *filename);
  CSymbolTable SymbolTableReadBinary(char *filename);

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
