if [[ ! -f "/tmp/count_file" ]]; then
	echo 0 > "/tmp/count_file"
fi

# read IND from count_file
< "/tmp/count_file" read IND

# run program
if [[ "${1}" == "clean" ]]; then
	rm gifs/*
	echo 0 > "/tmp/count_file"
else
	go run lissajous.go > "gifs/out${IND}.gif"
	echo "Saved in gifs/out${IND}.gif"
	gwenview "gifs/out${IND}.gif"

	# increment ind
	IND=$(( $IND + 1 ))
	echo "$IND" > "/tmp/count_file"
fi

