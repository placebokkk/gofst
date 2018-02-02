//gofst.h



//fst.AddArc(1,StdArc(3,3,2.5,2));
    
//fst.SetFinal(2,3.5);

//fst.Write("binary.fst");

#ifndef GO_FST_H
#define GO_FST_H



#ifdef __cplusplus
extern "C" {
#endif
  typedef void* CFst;

  CFst FstInit(void);
  void FstFree(CFst f);
  void AddState(CFst f);
  void SetStart(CFst f, int startState);
  void Compose(CFst fst1, CFst fst2, CFst ofst);
  void ArcSortInput(CFst fst);
  void ArcSortOutput(CFst fst);

  CFst FstRead(char* filename);
  void FstWrite(CFst fst, char* filename);

  typedef void* CSymbolTable;
#ifdef __cplusplus
}
#endif

#endif
