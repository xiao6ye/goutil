package goutil

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

// BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

// ByteSwap 字节数组前后调换
func ByteSwap(b []byte) (newB []byte) {
	middleLen := len(b) / 2
	newB = BytesCombine(newB, b[middleLen:])
	newB = BytesCombine(newB, b[:middleLen])
	return
}

// IntToBytes 整形转换成字节
func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

// BytesToBinaryString 字节数组转二进制字符串
func BytesToBinaryString(bs []byte) string {
	buf := bytes.NewBuffer([]byte{})
	for _, v := range bs {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String()
}

// ByteToBinaryString byte变量的二进制字符串
func ByteToBinaryString(data byte) (str string) {
	var a byte
	for i := 0; i < 8; i++ {
		a = data
		data <<= 1
		data >>= 1
		switch a {
		case data:
			str += "0"
		default:
			str += "1"
		}
		data <<= 1
	}
	return str
}

// BytesToIntU 字节数(大端)组转成int(无符号的)
func BytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// BytesToInt isSymbol表示有无符号
func BytesToInt(b []byte, isSymbol bool) (int, error) {
	if isSymbol {
		return BytesToIntS(b)
	}
	return bytesToIntU(b)
}

//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

// BytesToIntS 字节数(大端)组转成int(有符号)
func BytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

func parseFromRead(s string) (f float32, err error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return
	}
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.BigEndian, &f)
	return
}

func parseFromMath(s string) (f float32, err error) {
	i, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return
	}
	f = math.Float32frombits(uint32(i))
	return
}

func BytesToFloat32(b []byte) (f float32, err error) {
	buf := bytes.NewReader(b)
	err = binary.Read(buf, binary.BigEndian, &f)
	return
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)

	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

// SplitArray 数组平分
func SplitArray(arr []byte, num int32) [][]byte {
	max := int32(len(arr))
	if max < num {
		return nil
	}
	var segments = make([][]byte, 0)
	quantity := max / num
	end := int32(0)
	for i := int32(1); i <= num; i++ {
		qu := i * quantity
		if i != num {
			segments = append(segments, arr[i-1+end:qu])
		} else {
			segments = append(segments, arr[i-1+end:])
		}
		end = qu - i
	}
	return segments
}

// BytesToArray 字节数组转二维int 数组
func BytesToArray(data []byte) (mArr [][]int) {
	var dataInt []int
	a2 := SplitArray(data, int32(len(data))/2)
	for _, v := range a2 {
		v = []byte{v[1], v[0]}
		i, _ := BytesToIntU(v)
		dataInt = append(dataInt, i)
	}

	for k, v := range dataInt {
		var i []int
		i = append(i, k+1)
		i = append(i, v)
		mArr = append(mArr, i)
	}
	return
}
