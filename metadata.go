/*package main

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	Sample     int           `md:"sample,required,allowed(GET,POST,PUT,PATCH,DELETE)"`
	ListOfKeys []interface{} `md:"listOfKeys"`
}

type HandlerSettings struct {
	Method string `md:"method,required,allowed(GET,POST,PUT,PATCH,DELETE)"`
	Path   string `md:"path,required"`
}

type Input struct {
	InputSample int `md:"inputSample"`
}

type Output struct {
	OutputSample int `md:"outputSample"`
}

func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"inputSample": i.InputSample,
	}
}

func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	i.InputSample, err = coerce.ToInt(values["inputSample"])
	return err
}
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"outputSample": o.OutputSample,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.OutputSample, err = coerce.ToInt(values["outputSample"])
	return err
}
8*/