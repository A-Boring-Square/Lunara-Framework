package lunara

import (
	"errors"
	"math"
	"math/cmplx"
)

func PowerOf(base float64, exponent float64) (float64, error) {
	if base == 0 && exponent == 0 {
		return math.NaN(), errors.New("error: 0 to the power of 0 is undefined")
	}

	result := math.Pow(base, exponent)

	if math.IsInf(result, 0) || math.IsNaN(result) {
		return math.NaN(), errors.New("error: Result is infinite or NaN")
	}

	return result, nil
}

func RealRoot(root int64, radicand float64, rootIsFloat bool) (float64, error) {
	if root <= 0 {
		return math.NaN(), errors.New("error: Root must be a positive integer or a positive float")
	}

	if radicand < 0 && root%2 == 0 {
		return math.NaN(), errors.New("error: No real root exists for the given input")
	}

	if radicand < 0 && !rootIsFloat {
		return math.NaN(), errors.New("error: Radicand cannot be negative for integer root")
	}

	return math.Pow(radicand, 1.0/float64(root)), nil
}

func ComplexRoot(root interface{}, radicand float64) (complex128, error) {
	switch r := root.(type) {
	case int:
		if r <= 0 {
			return 0, errors.New("error: Integer root must be a positive integer")
		}
		return cmplx.Pow(complex(radicand, 0), 1.0/complex(float64(r), 0)), nil
	default:
		return 0, errors.New("error: Unsupported root type")
	}
}

func FastInvirseSqrt32(x float32) (float32, error) {
	if x < 0 {
		return 0, errors.New("error: Input can't be less than zero")
	}

	xhalf := 0.5 * x
	i := math.Float32bits(x)
	i = 0x5f3759df - (i >> 1)
	x = math.Float32frombits(i)
	x = x * (1.5 - xhalf*x*x)
	return x, nil
}

func AreaOfCircleUsingRadius(r float64) (float64, error) {
	if r < 0 {
		return math.NaN(), errors.New("error: Radius can't be less than zero")
	}
	return math.Pi * math.Pow(r, 2), nil
}

func AreaOfCircleUsingCircumference(c float64) (float64, error) {
	if c < 0 {
		return math.NaN(), errors.New("error: Circumference can't be less than zero")
	}
	return math.Pow(c, 2) / (4 * math.Pi), nil
}
