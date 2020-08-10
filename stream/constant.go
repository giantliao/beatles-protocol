package stream

const(
	bufferBitsSize int32 = 18
	int32BitsSize  int32 = 32
	int32Size      int32 = 4

	bufferSize int32 = (1<<bufferBitsSize) -1
	lengthMagicMax int32 = (1<<(int32BitsSize-bufferBitsSize)) - 1

)
