# LZF Compressor
A thin wrapper that adds/removes LZF payloads with LZF headers

This library makes heavy use of zhuyie's [golzf](https://github.com/zhuyie/golzf). This is just a tool for handling payloads with the LZF header.

## Usage

### Import:

```golang
    import "github.com/silverbeak/lzfcompressor"
``` 

### To compress a payload:

```golang
    content := []byte(strings.Repeat("lzf", 199))

    // N.B. Compressed will contain the LZF header + compressed payload
	compressed, err := lzfcompressor.CompressLZF(content)
	if err != nil {
		log.Fatalf("Error when compressing: %v", err)
	}
```

### To decompress:

```golang
    // N.B. Compressed must contain the LZF header
    var compressed []byte = ...

	decompressed, err := lzfcompressor.DecompressLZF(compressed)
	if err != nil {
		log.Fatalf("Could not decompress lzf %v", err)
	}
```

Still some testing to be done, but I'm gonna leave this here for now.