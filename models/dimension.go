package models
var (
	DimensionList map[string]*Dimension
	DimensionList1 map[string]*DimensionLite
	DimensionList2 map[string]*DimensionLite
	DimensionList3 map[string]*DimensionLite
	DimensionList4 map[string]*DimensionLite
	DimensionList5 map[string]*DimensionLite
	DimensionList6 map[string]*DimensionLite
	DimensionList7 map[string]*DimensionLite
	DimensionList8 map[string]*DimensionLite
	DimensionList9 map[string]*DimensionLite
	DimensionList10 map[string]*DimensionLite
)
func init() {
	DimensionList = make(map[string]*Dimension)
	DimensionList1 = make(map[string]*DimensionLite)
	DimensionList2 = make(map[string]*DimensionLite)
	DimensionList3 = make(map[string]*DimensionLite)
	DimensionList4 = make(map[string]*DimensionLite)
	DimensionList5 = make(map[string]*DimensionLite)
	DimensionList6 = make(map[string]*DimensionLite)
	DimensionList7 = make(map[string]*DimensionLite)
	DimensionList8 = make(map[string]*DimensionLite)
	DimensionList9 = make(map[string]*DimensionLite)
	DimensionList10 = make(map[string]*DimensionLite)
}
type Dimension struct {
	ANL_CAT_ID string
	S_HEAD string
	ANL_CODE string
	NAME string
	STATUS int64
	PROHIBIT_POSTING int64
}

type DimensionLite struct {
	ANL_CAT_ID string
	ANL_CODE string
	ENTRY_NUM int64
	STATUS int64
	PROHIBIT_POSTING int64
	NAME string
	DESCR string
}
func AddDimensionU(u Dimension, index string)  {
	DimensionList[index] = &u	
}
func GetAllDimension() map[string]*Dimension {
	return DimensionList
}
func AddDimensionU1(u DimensionLite, index string)  { DimensionList1[index] = &u	}
func GetAllDimension1() map[string]*DimensionLite { return DimensionList1 }

func AddDimensionU2(u DimensionLite, index string)  { DimensionList2[index] = &u	}
func GetAllDimension2() map[string]*DimensionLite { return DimensionList2 }

func AddDimensionU3(u DimensionLite, index string)  { DimensionList3[index] = &u	}
func GetAllDimension3() map[string]*DimensionLite { return DimensionList3 }

func AddDimensionU4(u DimensionLite, index string)  { DimensionList4[index] = &u	}
func GetAllDimension4() map[string]*DimensionLite { return DimensionList4 }

func AddDimensionU5(u DimensionLite, index string)  { DimensionList5[index] = &u	}
func GetAllDimension5() map[string]*DimensionLite { return DimensionList5 }

func AddDimensionU6(u DimensionLite, index string)  { DimensionList6[index] = &u	}
func GetAllDimension6() map[string]*DimensionLite { return DimensionList6 }

func AddDimensionU7(u DimensionLite, index string)  { DimensionList7[index] = &u	}
func GetAllDimension7() map[string]*DimensionLite { return DimensionList7 }

func AddDimensionU8(u DimensionLite, index string)  { DimensionList8[index] = &u	}
func GetAllDimension8() map[string]*DimensionLite { return DimensionList8 }

func AddDimensionU9(u DimensionLite, index string)  { DimensionList9[index] = &u	}
func GetAllDimension9() map[string]*DimensionLite { return DimensionList9 }

func AddDimensionU10(u DimensionLite, index string)  { DimensionList10[index] = &u	}
func GetAllDimension10() map[string]*DimensionLite { return DimensionList10 }