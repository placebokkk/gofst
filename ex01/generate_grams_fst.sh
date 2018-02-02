name=$1
fstcompile --isymbols=wotw.syms --osymbols=wotw.syms ${name}.txt ${name}.fst

fstdraw --isymbols=wotw.syms --osymbols=wotw.syms -portrait ${name}.fst|dot -Tjpg -o ${name}.jpg

