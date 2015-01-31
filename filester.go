package filester

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	OneByte  = 1
	KiloByte = 1024 * OneByte
	MegaByte = 1024 * KiloByte
	GigaByte = 1024 * MegaByte
	TeraByte = 1024 * GigaByte
	PetaByte = 1024 * TeraByte
)

const (
	MaxCountFiles       = 10000
	MaxFileSize   int64 = 10 * GigaByte
)

const DefaultConfigName = "config.json"

const defaultParamsData = `{
	"Path": "./test",
	"CountFiles": 10,
	"File": {
		"Name": {
			"Separator": "_",
			"Prefix": "",
			"RandomPart": {
				"Size": {
					"Min": 8,
					"Max": 15
				}
			},
			"Ext": ""
		},
		"Data": {
			"Size": {
				"Min": 10,
				"Max": 100
			}
		}
	}
}`

type randPair struct {
	Percent int
	Data    []byte
}

type Params struct {
	Path       string
	CountFiles int
	File       struct {
		Name struct {
			Separator  string
			Prefix     string
			RandomPart struct {
				Size struct {
					Min int
					Max int
				}
			}
			Ext string
		}
		Data struct {
			Size struct {
				Min int64
				Max int64
			}
		}
	}
}

func NewParamsDefault() (*Params, error) {

	var p Params

	err := json.Unmarshal([]byte(defaultParamsData), &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (this *Params) pathError() error {

	fi, err := os.Stat(this.Path)
	if err != nil {
		if err = os.MkdirAll(this.Path, os.ModePerm); err != nil {
			return err
		}
	} else {
		if !fi.IsDir() {
			return errors.New("path is not directory")
		}
	}

	return nil
}

func (this *Params) Error() error {

	if err := this.pathError(); err != nil {
		return err
	}

	// File.Name.RandomPart.Size
	{
		size := this.File.Name.RandomPart.Size

		if (size.Min < 3) || (size.Min > 30) {
			return errors.New("Wrong RandomPart.Size.Min")
		}

		if (size.Max < 3) || (size.Max > 30) {
			return errors.New("Wrong RandomPart.Size.Max")
		}

		if size.Min > size.Max {
			return errors.New("RandomPart.Size.Min > RandomPart.Size.Max")
		}
	}

	// File.Data.Size
	{
		size := &(this.File.Data.Size)

		if (this.CountFiles < 0) || (this.CountFiles > MaxCountFiles) {
			return errors.New("Wrong CountFiles")
		}

		if (size.Min < 0) || (size.Min > MaxFileSize) {
			return errors.New("Wrong FileSize.Min")
		}

		if (size.Max < 0) || (size.Max > MaxFileSize) {
			return errors.New("Wrong FileSize.Max")
		}

		if size.Min > size.Max {
			return errors.New("FileSize.Min > FileSize.Max")
		}

		const maxGigaByte = 10
		if (int64(this.CountFiles) * size.Max) > maxGigaByte*int64(GigaByte) {
			return errors.New(fmt.Sprintf("CountFiles * File.Size.Max > %dGB", maxGigaByte))
		}
	}

	return nil
}

func (this *Params) randFileName(r *rand.Rand, rps []randPair) string {

	rs := this.File.Name.RandomPart.Size

	var fileName string

	if filePrefix := this.File.Name.Prefix; len(filePrefix) > 0 {
		fileName = filePrefix + this.File.Name.Separator
	}

	nameSize := randRange(r, rs.Min, rs.Max)
	fileName += string(randFromPairs(r, nameSize, rps))

	if fileExt := this.File.Name.Ext; len(fileExt) > 0 {
		fileName += fileExt
	}

	return filepath.Join(this.Path, fileName)
}

func (this *Params) randFileSize(r *rand.Rand) int64 {
	fs := &(this.File.Data.Size)
	return randRange64(r, fs.Min, fs.Max)
}

func randRange(r *rand.Rand, min, max int) int {

	return (min + r.Intn(max-min+1))
}

func randRange64(r *rand.Rand, min, max int64) int64 {

	return (min + r.Int63n(max-min+1))
}

func randFromPairs(r *rand.Rand, count int, rps []randPair) []byte {

	data := make([]byte, count)
	for i := 0; i < count; i++ {
		sp := 0
		percent := r.Intn(100) // [0 ... 99]
		for j := 0; j < len(rps); j++ {
			sp += rps[j].Percent
			if percent < sp {
				d := rps[j].Data
				data[i] = d[r.Intn(len(d))]
				break
			}
		}
	}

	return data
}

func CreateFiles(p *Params) error {

	if err := p.Error(); err != nil {
		return err
	}

	var rps = []randPair{
		randPair{40, []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")},
		randPair{40, []byte("abcdefghijklmnopqrstuvwxyz")},
		randPair{20, []byte("0123456789")},
	}

	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fillData := make([]byte, MegaByte)

	for i := 0; i < p.CountFiles; i++ {

		fileName := p.randFileName(r, rps)
		fileSize := p.randFileSize(r)

		fillSize := fillBytesSize(fileSize, len(fillData))
		randFillBytes(r, fillData[:fillSize])

		if err := fileCreateAndFill(fileName, fileSize, fillData[:fillSize]); err != nil {
			return err
		}
	}

	return nil
}

func divmod(dividend, divisor int) (quotient, remainder int) {

	quotient = dividend / divisor
	remainder = dividend - quotient*divisor

	return
}

func randFillBytes(r *rand.Rand, dest []byte) {

	src := make([]byte, 4)

	quotient, remainder := divmod(len(dest), 4)

	if quotient > 0 {
		for i := 0; i < quotient; i++ {

			u := r.Uint32()

			src[0] = byte(u)
			src[1] = byte(u >> 8)
			src[2] = byte(u >> 16)
			src[3] = byte(u >> 24)

			copy(dest[(i*4):((i+1)*4)], src)
		}
	}

	if remainder > 0 {

		u := r.Uint32()

		for i := 0; i < remainder; i++ {

			dest[quotient*4+i] = byte(u >> uint(i*8))
		}
	}
}

func fillBytesSize(fileSize int64, maxSize int) int {

	n := int(fileSize / 64)

	if n < KiloByte {
		n = KiloByte
	}

	if n > maxSize {
		n = maxSize
	}

	return n
}

func fileCreateAndFill(fileName string, fileSize int64, fillData []byte) (err error) {

	var wc io.WriteCloser
	wc, err = os.Create(fileName)
	if err != nil {
		return
	}
	defer wc.Close()

	n := int64(0)
	for n < fileSize {

		dn := len(fillData)
		if int64(dn) > (fileSize - n) {
			dn = int(fileSize - n)
		}

		dn, err = wc.Write(fillData[:dn])
		if err != nil {
			return
		}

		n += int64(dn)
	}

	return
}

func readFileBytes(fileName string) ([]byte, error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data := make([]byte, 1024)
	buffer := new(bytes.Buffer)

	for {
		n, err := f.Read(data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		buffer.Write(data[:n])
	}

	return buffer.Bytes(), nil
}

func LoadParamsFromFile(fileName string) (*Params, error) {

	data, err := readFileBytes(fileName)
	if err != nil {
		return nil, err
	}

	var p Params

	if err = json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func CreateDefaultConfigFile() error {

	p, err := NewParamsDefault()
	if err != nil {
		return err
	}

	bs, err := json.MarshalIndent(&p, "", "\t")
	if err != nil {
		return err
	}

	f, err := os.Create(DefaultConfigName)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(bs); err != nil {
		return err
	}

	return nil
}
