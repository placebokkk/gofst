#编译Mars的char->Word
fstcompile --isymbols=ascii.syms --osymbols=wotw.syms Mars.txt Mars.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait Mars.fst|dot -Tjpg -o Mars.jpg

#编译man的char->Word
fstcompile --isymbols=ascii.syms --osymbols=wotw.syms man.txt man.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait man.fst|dot -Tjpg -o man.jpg

#编译Martian的char->Word
fstcompile --isymbols=ascii.syms --osymbols=wotw.syms Martian.txt Martian.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait Martian.fst|dot -Tjpg -o Martian.jpg

#编译标点符号
fstcompile --isymbols=ascii.syms --osymbols=wotw.syms punct.txt punct.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait punct.fst|dot -Tjpg -o punct.jpg

#将三个单词的fst Union起来,增加闭包(返回初始节点)
fstunion man.fst Mars.fst | fstunion - Martian.fst |fstconcat - punct.fst|fstclosure > lexicon.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait lexicon.fst|dot -Tjpg -o lexicon.jpg

#去除epsilon,determinize, minimize化
fstrmepsilon lexicon.fst|fstdeterminize|fstminimize > lexicon_opt.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait lexicon_opt.fst|dot -Tjpg -o lexicon_opt.jpg

#编译Marsman的char->char的fst
fstcompile --isymbols=ascii.syms --osymbols=ascii.syms Marsman_t.txt Marsman_t.fst

#下面的compose+project操作,完成了一个向某个fst(lexicon_opt)喂一个输入串得到对应输出串的操作.

#compose操作,char->char + char->word,得到char->word.
fstcompose Marsman_t.fst lexicon_opt.fst  > Marsman.fst
fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait Marsman.fst|dot -Tjpg -o Marsman.jpg
#取output label
fstproject --project_output Marsman.fst|fstrmepsilon > tokens.fst
fstdraw --isymbols=wotw.syms --osymbols=wotw.syms -portrait tokens.fst|dot -Tjpg -o tokens.jpg 
