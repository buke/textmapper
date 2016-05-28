package tm

var tmNumClasses = 36

var tmRuneClass = []uint8{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 35, 4, 1, 1, 12, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	35, 17, 9, 13, 31, 11, 30, 2, 22, 23, 8, 28, 21, 10, 20, 5,
	34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 14, 19, 27, 15, 18, 29,
	32, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 6, 3, 7, 1, 33,
	1, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 24, 16, 26, 25,
}

var tmStateMap = []int{
	0, 0, 1,
}

var tmToken = []Token{
	39, 1, 2, 3, 0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	31, 32, 33, 34, 35, 36, 37, 40, 41, 42, 43, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78,
}

var tmLexerAction = []int32{
	-2, -1, 2, -1, 3, 4, 5, 6, 7, 8, 9, 10, 3, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 3, -1, -1, 2, -1, 3, 4, 5, 6, 7, 8, 9, 10,
	3, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 33, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 3, -1, 2, 34, 35, -1, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, -8, -8, -8, -8,
	3, -8, -8, -8, -8, -8, -8, -8, 3, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, -8, 3,
	-1, 36, 36, 37, -1, -1, 38, 36, 39, 36, 36, 36, -1, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-1, 8, 8, 40, -1, 8, 8, 8, 8, 41, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, 32, -1, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, 42, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -9, 11, 11, 11,
	43, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, 44, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, 45, -16, -16, 46, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -14, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, 47, -14, -14, -14, -14, -14, -14, -14,
	-14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -14, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, 48, -36, -36, -36, -36,
	-36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36, -36,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -79, -79, -79, -79, -79, -79, -79, -79,
	-79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79,
	-79, 49, -79, -79, -79, -79, -79, -79, -79, -79, -79, -79, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, 50, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35, -35,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, 51, -38,
	-38, -38, -38, -38, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, 52, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, 31, 31, -3,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6, -6,
	-6, -6, 32, -6, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80,
	-80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, -80, 49, -80, -80,
	-80, -80, -80, -80, -80, -80, -80, -80, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3,
	-3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -3, -1, 2, 2, 2,
	-1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	-1, 36, 36, 53, -1, 54, 55, 36, 36, 36, 36, 36, -1, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, -1, 36, 36, 36, -1, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, -1, 38, 38, 56, -1, 38, 38, 36,
	38, 38, 38, 38, -1, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, -1, 39, 39, 39,
	39, 39, 39, 39, 57, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	-1, 8, 8, 8, -1, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -7, 42, 42, 42, 58, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, 42, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, 59,
	-13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17, -17,
	-17, -17, -17, -17, -17, -17, -17, -17, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19,
	-19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -19, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15, -15,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18, -18,
	-18, -18, -18, -18, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 52, -1, -1, -1, -1, -1,
	-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1,
	-1, 31, 31, -1, -1, 36, 36, 36, -1, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36, 36,
	36, 36, 36, 36, 36, 36, 36, 36, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4,
	-4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -4, -1, 55, 55, 60,
	-1, 55, 55, 36, 55, 55, 55, 55, -1, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	-1, 38, 38, 38, -1, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, -1, 39, 39, 39, 39, 61, 39, 39, 57, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-1, 55, 55, 55, -1, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10,
}
