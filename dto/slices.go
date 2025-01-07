package dto

// List makes []O using []I.
func List[O any, L ~[]I, I any](inputs L, convert func(I) O) []O {
	if convert == nil || inputs == nil || len(inputs) == 0 {
		return nil
	}
	outputs := make([]O, len(inputs))
	for i := range inputs {
		outputs[i] = convert(inputs[i])
	}
	return outputs
}

//// List1 makes []*O using []I.
//func List1[O any, L ~[]I, I any](inputs L, convert func(I) O) []*O {
//	if convert == nil || inputs == nil || len(inputs) == 0 {
//		return nil
//	}
//	outputs := make([]*O, len(inputs))
//	for i := range inputs {
//		result := convert(inputs[i])
//		outputs[i] = &result
//	}
//	return outputs
//}
//
//// List2 makes []O using []*I.
//func List2[O any, L ~[]*I, I any](inputs L, convert func(I) O) []O {
//	if convert == nil || inputs == nil || len(inputs) == 0 {
//		return nil
//	}
//	outputs := make([]O, len(inputs))
//	for i := range inputs {
//		if inputs[i] != nil {
//			outputs[i] = convert(*inputs[i])
//		} else {
//			var zero I
//			outputs[i] = convert(zero)
//		}
//	}
//	return outputs
//}
//
//// List3 makes []*O using []*I.
//func List3[O any, L ~[]*I, I any](inputs L, convert func(I) O) []*O {
//	if convert == nil || inputs == nil || len(inputs) == 0 {
//		return nil
//	}
//	outputs := make([]*O, len(inputs))
//	for i := range inputs {
//		if inputs[i] != nil {
//			output := convert(*inputs[i])
//			outputs[i] = &output
//		} else {
//			var zero I
//			output := convert(zero)
//			outputs[i] = &output
//		}
//	}
//	return outputs
//}
