package errors

import (
	"encoding/binary"
	"io"
)

type Point struct {
	Longitude     int
	Latitude      int
	Distance      int
	ElevationGain int
	ElevationLoss int
}

func parse(r io.Reader) (*Point, error) {
	var p Point
	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
		return nil, err
	}
	return &p, nil
}

func parseV2(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		// 保证不会覆盖上一个操作的err
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}
	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)
	if err != nil {
		return &p, err
	}
	return &p, nil
}

type Reader struct {
	r io.Reader
	// 使用结构体抽离err变量
	err error
}

// 抽离v2的内部函数
func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

/*
实现参考：

scanner := bufio.NewScanner(input)
for scanner.Scan() {
    token := scanner.Text()
    // process token
}
if err := scanner.Err(); err != nil {
    // process the error
}

*/
func parseV3(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}
	return &p, nil
}
