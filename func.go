package loe

// Partial returns new function that, when called, has its first argument set to the provided value.
func Partial[T1, T2, R any](f func(a T1, b T2) (R, error), arg1 T1) func(T2) (R, error) {
	return func(t2 T2) (R, error) {
		return f(arg1, t2)
	}
}

// Partial1 returns new function that, when called, has its first argument set to the provided value.
func Partial1[T1, T2, R any](f func(T1, T2) (R, error), arg1 T1) func(T2) (R, error) {
	return Partial(f, arg1)
}

// Partial2 returns new function that, when called, has its first argument set to the provided value.
func Partial2[T1, T2, T3, R any](f func(T1, T2, T3) (R, error), arg1 T1) func(T2, T3) (R, error) {
	return func(t2 T2, t3 T3) (R, error) {
		return f(arg1, t2, t3)
	}
}

// Partial3 returns new function that, when called, has its first argument set to the provided value.
func Partial3[T1, T2, T3, T4, R any](f func(T1, T2, T3, T4) (R, error), arg1 T1) func(T2, T3, T4) (R, error) {
	return func(t2 T2, t3 T3, t4 T4) (R, error) {
		return f(arg1, t2, t3, t4)
	}
}

// Partial4 returns new function that, when called, has its first argument set to the provided value.
func Partial4[T1, T2, T3, T4, T5, R any](f func(T1, T2, T3, T4, T5) (R, error), arg1 T1) func(T2, T3, T4, T5) (R, error) {
	return func(t2 T2, t3 T3, t4 T4, t5 T5) (R, error) {
		return f(arg1, t2, t3, t4, t5)
	}
}

// Partial5 returns new function that, when called, has its first argument set to the provided value
func Partial5[T1, T2, T3, T4, T5, T6, R any](f func(T1, T2, T3, T4, T5, T6) (R, error), arg1 T1) func(T2, T3, T4, T5, T6) (R, error) {
	return func(t2 T2, t3 T3, t4 T4, t5 T5, t6 T6) (R, error) {
		return f(arg1, t2, t3, t4, t5, t6)
	}
}
