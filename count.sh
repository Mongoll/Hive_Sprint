count=$(find . | wc -l)
result=$((count*5))
printf "\t\vTotal files * 5: %d\v\n" "$result"