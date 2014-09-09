TWISTED_STEEL=`pwd`
for folder in `ls -d -- */`                                                                                                                                  ✭
do
    package=`echo "$folder" | rev | cut -c 2- | rev`
    echo "running go fmt in $TWISTED_STEEL/$package"
    cd "$TWISTED_STEEL/$package"
    go fmt
    cd $TWISTED_STEEL
done
