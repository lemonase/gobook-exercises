if [[ ! -f "count_file" ]]; then
	echo 0 > count_file
fi

# read IND from count_file
< count_file read IND

# run program
if [[ "${1}" == "clean" ]]; then
	rm gifs/*
else
	go run lissajous.go > gifs/out${IND}.gif
	echo "Saved in gifs/out${IND}.gif"
	gwenview gifs/out${IND}.gif
fi

# increment ind
IND=$(( $IND + 1 ))
echo $IND > count_file
echo $IND
