package main

type Pooltest struct {
	ID   string
	Name string
	Age  string
	Address string
}
//var Pool = sync.Pool{New: func() interface{}{return new(Pooltest)}}

func GetPool() Pooltest{
	//var poolData = Pool.Get().(*Pooltest)
	poolData := Pooltest{
		ID:      "fjjwkw",
		Name:    "nwnnbfwe",
		Age:     "njefsnjws",
		Address: "nfannnasnmn",
	}
	return poolData
  //  defer Pool.Put(poolData)
}
