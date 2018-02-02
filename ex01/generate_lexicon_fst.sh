name=$1
fstcompile --isymbols=ascii.syms --osymbols=wotw.syms ${name}.txt ${name}.fst

fstdraw --isymbols=ascii.syms --osymbols=wotw.syms -portrait ${name}.fst|dot -Tjpg -o ${name}.jpg

