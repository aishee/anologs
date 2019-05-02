package anologs

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type Engine struct {
	splits Splits
	model  Model
}

func NEngine() Engine {
	x := Engine{}
	x.splits = NewSSplits()
	x.model = &SModel{}
	return x
}

func (x *Engine) Read(text string) *Format {
	logs := NLog(text, x.splits)
	fmt := x.model.read(logs)
	return fmt
}

func (x *Engine) Count() int {
	return x.model.count()
}

func (x *Engine) Save(fpath string) error {
	data := x.DumpModel()
	f, err := os.Create(fpath)
	if err != nil {
		return errors.Wrap(err, "Model file open errors: "+fpath)
	}
	defer f.Close()
	f.Write(data)
	return nil
}

func (x *Engine) DumpModel() []byte {
	data, err := x.model.dump()
	if err != nil {
		panic(err)
	}
	return data
}

func (x *Engine) Load(fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return errors.Wrap(err, "Model file open error: "+fpath)
	}

	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "Fail to read model data: "+fpath)
	}
	model := &SModel{}
	err = json.Unmarshal(data, model)
	if err != nil {
		return errors.Wrap(err, "Fail to load model data: "+fpath)
	}
	x.model = model
	return nil
}

func (x *Engine) Formats() []*Format {
	return []*Format{}
}
