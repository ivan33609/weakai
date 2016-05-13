package neuralnet

import (
	"fmt"
	"math"
)

// An ActivationFunc is a function designed to
// introduce non-linearity into a neural net.
type ActivationFunc interface {
	Serializer

	// Eval evaluates the activation function for
	// an input x.
	Eval(x float64) float64

	// Deriv evaluates the derivative of the
	// activation function for an input x.
	Deriv(x float64) float64
}

type Sigmoid struct{}

func (_ Sigmoid) Eval(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

func (s Sigmoid) Deriv(x float64) float64 {
	v := s.Eval(x)
	return v * (1.0 - v)
}

func (_ Sigmoid) Serialize() []byte {
	return []byte{}
}

func (_ Sigmoid) SerializerType() string {
	return "sigmoid"
}

func deserializeActivation(data []byte, serializerType string) (ActivationFunc, error) {
	activationDes, ok := Deserializers[serializerType]
	if !ok {
		return nil, fmt.Errorf("unknown activation type: %s", serializerType)
	}
	activation, err := activationDes(data)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize activation: %s", err.Error())
	} else if _, ok := activation.(ActivationFunc); !ok {
		return nil, fmt.Errorf("expected ActivationFunc but got %T", activation)
	}
	return activation.(ActivationFunc), nil
}
