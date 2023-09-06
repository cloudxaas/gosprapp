package cxstrconv

import (
	cx "github.com/cloudxaas/gocx"
)


func Atoi(s string) int {

    val := 0
    for i := 0; i < len(s); i++ {
        c := s[i]
        if c < '0' || c > '9' {
            break
        }
        val = val*10 + int(c-'0')
    }
    return val
	
}


func Inttoa(n int) string {
	if n == 0 {
		return "0"
	}

	var buf [20]byte // max int string length is 19 digits plus the sign
	i := len(buf) - 1
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	if neg {
		buf[i] = '-'
		i--
	}

	return cx.B2s(buf[i+1:])
}

func Int8toa(n int8) string {
	if n == 0 {
		return "0"
	}

	var buf [4]byte // max int8 string length is 3 digits plus the sign
	i := len(buf) - 1
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	if neg {
		buf[i] = '-'
	} else {
		i++
	}

	return cx.B2s(buf[i:])
}


func Int16toa(n int16) string {
	if n == 0 {
		return "0"
	}

	var buf [6]byte // max int16 string length is 5 digits plus the sign
	i := len(buf) - 1
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	if neg {
		buf[i] = '-'
	} else {
		i++
	}

	return cx.B2s(buf[i:])
}

func Int32toa(n int32) string {
	if n == 0 {
		return "0"
	}

	var buf [11]byte // max int32 string length is 10 digits plus the sign
	i := len(buf) - 1
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	if neg {
		buf[i] = '-'
	} else {
		i++
	}

	return cx.B2s(buf[i:])
}

func Int64toa(n int64) string {
	if n == 0 {
		return "0"
	}

	var buf [20]byte // max int64 string length is 19 digits plus the sign
	i := len(buf) - 1
	neg := false

	if n < 0 {
		neg = true
		n = -n
	}

	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	if neg {
		buf[i] = '-'
	} else {
		i++
	}

	return cx.B2s(buf[i:])
}

func Uint8toa(n uint8) string {
    if n == 0 {
        return "0"
    }
    buf := [3]byte{}
    i := len(buf)
    for n > 0 {
        i--
        buf[i] = '0' + n%10
        n /= 10
    }
    return cx.B2s(buf[i:])
}

func Uint16toa(n uint16) string {
	if n == 0 {
		return "0"
	}

	var b [5]byte
	i := len(b) - 1
	for n > 0 {
		b[i] = byte(n%10) + '0'
		n /= 10
		i--
	}

	return cx.B2s(b[i+1:])
}
func Uint32toa(n uint32) string {
	if n == 0 {
		return "0"
	}
	buf := [11]byte{} // max uint32 string length is 10 digits
	i := len(buf) - 1
	for n > 0 {
		buf[i] = byte(n%10) + '0'
		n /= 10
		i--
	}
	return cx.B2s(buf[i+1:])
}


func Uint64toa(n uint64) string {
	if n == 0 {
		return "0"
	}
	buf := [20]byte{} // max uint64 string length is 20 digits
	i := len(buf) - 1
	for n > 0 {
		q := n / 10
		r := n % 10
		buf[i] = byte(r + '0')
		i--
		n = q
	}
	return cx.B2s(buf[i+1:])
}
