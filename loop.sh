count=$1
if [[ "$count" -gt 100 ]]; then
  count=100
fi
i=1
while [[ "$i" -le "$count" ]]; do
  echo "This is loop number $i"
  i=$((i + 1))
done