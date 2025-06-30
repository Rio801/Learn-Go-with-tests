package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	wanted := 40.0

	if got != wanted {
		t.Errorf("got: %.2f, wanted: %.2f", got, wanted)
	}
}

func TestArea(t *testing.T) {
	//Table Driven Test
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12.0, 6.0}, 36.0},
	}

	for _, test := range areaTests {
		t.Run("Testing Rectangle,Circle,Triangle Areas", func(t *testing.T) {
			got := test.shape.Area()

			if got != test.want {
				t.Errorf("got: %.2f, wanted: %.2f", got, test.want)
			}
		})
	}

}
