package main
//Stack Implementation where user can perform GetMin,GetMax, Peek ,Push and all the operations are O(1)
type stack struct {
	data []int
	MinMax []minMax
}
type minMax struct {
	min int
	max int
}
func(s *stack) Push(value int){
	if len(s.data)==0{
		 s.MinMax = append(s.MinMax, minMax{
			 min: value,
			 max: value,
		 })
	}else{
		if value > s.MinMax[len(s.MinMax)-1].max{
			  s.MinMax = append(s.MinMax, minMax{
				  min: s.MinMax[len(s.MinMax)-1].min,
				  max: value,
			  })
		}
		if value < s.MinMax[len(s.MinMax)-1].min{
			s.MinMax = append(s.MinMax, minMax{
				min: value,
				max: s.MinMax[len(s.MinMax)-1].max,
			})
		}
	}
	 s.data = append(s.data, value)
}
func (s *stack) Pop() int{
	result := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if result == s.MinMax[len(s.MinMax)-1].max || result == s.MinMax[len(s.MinMax)-1].min{
		s.MinMax = s.MinMax[:len(s.MinMax)-1]
	}
	return result
}
func(s *stack) GetMin() int{
	return s.MinMax[len(s.MinMax)-1].min
}
func(s *stack) GetMax() int{
	return s.MinMax[len(s.MinMax)-1].max
}