// Copyright (c) 2022-2024 Focela Technologies, All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import "errors"

func Add(x, y int) (int, error) {
	return x + y, nil
}

func Subtract(x, y int) (int, error) {
	return x - y, nil
}

func Multiply(x, y int) (int, error) {
	return x * y, nil
}

func Divide(x, y int) (float64, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return float64(x) / float64(y), nil
}
