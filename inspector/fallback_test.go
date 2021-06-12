package inspector_test

// func TestInspectFallback(t *testing.T) {
// 	t.Run("It inspects bool", func(t *testing.T) {
// 		var output bytes.Buffer

// 		trueValue := true
// 		vType := reflect.TypeOf(trueValue)
// 		vValue := reflect.ValueOf(trueValue)

// 		ioP := iop.New()
// 		ioP.SetOutput(&output)

// 		sliceInspector := inspector.NewFallbackInspector()
// 		sliceInspector.Inspect(ioP, vType, vValue, 0)

// 		expected := "\ttrue,\n"
// 		got := output.String()
// 		if got != expected {
// 			t.Errorf("Expect: %s, but got: %s", expected, got)
// 		}
// 	})
// }
