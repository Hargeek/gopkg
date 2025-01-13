package utils

import "github.com/shopspring/decimal"

// Float32ToDecimalPtr *float32 -> *decimal.Decimal
func Float32ToDecimalPtr(f *float32) *decimal.Decimal {
	if f == nil {
		return nil
	}
	d := decimal.NewFromFloat32(*f)
	return &d
}

// Float64ToDecimalPtr *float64 -> *decimal.Decimal
func Float64ToDecimalPtr(f *float64) *decimal.Decimal {
	if f == nil {
		return nil
	}
	d := decimal.NewFromFloat(*f)
	return &d
}

// DecimalPtrToFloat32 *decimal.Decimal -> *float32
func DecimalPtrToFloat32(d *decimal.Decimal) *float32 {
	if d == nil {
		return nil
	}
	f := float32(d.InexactFloat64())
	return &f
}

// DecimalPtrToFloat64 *decimal.Decimal -> *float64
func DecimalPtrToFloat64(d *decimal.Decimal) *float64 {
	if d == nil {
		return nil
	}
	f := d.InexactFloat64()
	return &f
}
