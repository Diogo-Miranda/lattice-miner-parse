package main

import (
	"fmt"

	"github.com/fatih/structs"
)

type Object struct {
	CODE1 string `csv:"CODE1"`

	HADS1_2017 string `csv:"HADS1_2017"`
	HADS1_2018 string `csv:"HADS1_2018"`
	HADS1_2019 string `csv:"HADS1_2019"`

	HADS2_2017 string `csv:"HADS2_2017"`
	HADS2_2018 string `csv:"HADS2_2018"`
	HADS2_2019 string `csv:"HADS2_2019"`

	HADS3_2017 string `csv:"HADS3_2017"`
	HADS3_2018 string `csv:"HADS3_2018"`
	HADS3_2019 string `csv:"HADS3_2019"`

	HADS4_2017 string `csv:"HADS4_2017"`
	HADS4_2018 string `csv:"HADS4_2018"`
	HADS4_2019 string `csv:"HADS4_2019"`

	HADS6_2017 string `csv:"HADS6_2017"`
	HADS6_2018 string `csv:"HADS6_2018"`
	HADS6_2019 string `csv:"HADS6_2019"`

	HADS9_2017 string `csv:"HADS9_2017"`
	HADS9_2018 string `csv:"HADS9_2018"`
	HADS9_2019 string `csv:"HADS9_2019"`

	HADS11_2017 string `csv:"HADS11_2017"`
	HADS11_2018 string `csv:"HADS11_2018"`
	HADS11_2019 string `csv:"HADS11_2019"`

	HADS12_2017 string `csv:"HADS12_2017"`
	HADS12_2018 string `csv:"HADS12_2018"`
	HADS12_2019 string `csv:"HADS12_2019"`

	HADS14_2017 string `csv:"HADS14_2017"`
	HADS14_2018 string `csv:"HADS14_2018"`
	HADS14_2019 string `csv:"HADS14_2019"`

	SOMS1_2017 string `csv:"SOMS1_2017"`
	SOMS1_2018 string `csv:"SOMS1_2018"`
	SOMS1_2019 string `csv:"SOMS1_2019"`
	
	SOMS2_2017 string `csv:"SOMS2_2017"`
	SOMS2_2018 string `csv:"SOMS2_2018"`
	SOMS2_2019 string `csv:"SOMS2_2019"`

	SOMS3_2017 string `csv:"SOMS3_2017"`
	SOMS3_2018 string `csv:"SOMS3_2018"`
	SOMS3_2019 string `csv:"SOMS3_2019"`

	SOMS4_2017 string `csv:"SOMS4_2017"`
	SOMS4_2018 string `csv:"SOMS4_2018"`
	SOMS4_2019 string `csv:"SOMS4_2019"`

	SOMS8_2017 string `csv:"SOMS8_2017"`
	SOMS8_2018 string `csv:"SOMS8_2018"`
	SOMS8_2019 string `csv:"SOMS8_2019"`
	
	SOMS13_2017 string `csv:"SOMS13_2017"`
	SOMS13_2018 string `csv:"SOMS13_2018"`
	SOMS13_2019 string `csv:"SOMS13_2019"`

	SOMS14_2017 string `csv:"SOMS14_2017"`
	SOMS14_2018 string `csv:"SOMS14_2018"`
	SOMS14_2019 string `csv:"SOMS14_2019"`

	SOMS29_2017 string `csv:"SOMS29_2017"`
	SOMS29_2018 string `csv:"SOMS29_2018"`
	SOMS29_2019 string `csv:"SOMS29_2019"`

	SOMS30_2017 string `csv:"SOMS30_2017"`
	SOMS30_2018 string `csv:"SOMS30_2018"`
	SOMS30_2019 string `csv:"SOMS30_2019"`

	SOMS54_2017 string `csv:"SOMS54_2017"`
	SOMS54_2018 string `csv:"SOMS54_2018"`
	SOMS54_2019 string `csv:"SOMS54_2019"`
}

type LatticeMinerInput struct {
	Name string `json:"name"`
	Objects []string `json:"objects"`
	Attributes []string `json:"attributes"`
	Conditions []string `json:"conditions"`
	Relations [][][]string `json:"relations"`
}

func (o *Object) ContextMap() map[string]interface{} {
	m := structs.Map(o)

	for f, v := range m {
		if f != "CODE1" {
			m[f] = decodeField(v.(string))
		}
	}

	return m
}

func (o *Object) VerifyContextMap() bool {
	m := o.ContextMap()

	resp := true

	for _, v := range m {
		if v == "" {
			resp = false
		}
	}

	return resp
}

func decodeField(v string) string {	
	if v == "1" || v == "2" {
		return "False"
	} else if v == "3" || v == "4" {
		return "True"
	} else {
		return ""
	}
}

func (o *Object) Print() {
	m := o.ContextMap()
	fmt.Println(m)
}
