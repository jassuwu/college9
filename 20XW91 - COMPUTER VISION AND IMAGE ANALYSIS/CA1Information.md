# 16AUG2024  Internal Test I Revision class
## Unit I
### Image digitization
- **What are the steps to convert to Digital Image? / How to do Image Digitization?**
- It's a three step process.
`Analog Image -> Sampling -> Quantization -> Encoding -> Digital Image`
	1. **Sampling** `fs(x, y)`
		- Taking specific instance of the signal.
		- Utmost care for choosing sampling freq.
		- Improper choice of sampling freq lead to **Aliasing**.
	2.  **Quantization** :
		- Mapping a large set of values to a smaller set of values.
		- Non-linear and Irreversible process
		- Scalar, Vector
		- Irreversible because it's a many to one mapping, leading to loss of uniqueness
	3. **Encoding**:
		- Conversion of text.
		- Fixed, Variable length encoding.
- **g/n f(x, y) = 2cos 2 pi (4x + 6y), ∆x = 0.2, ∆y = 0.4, find the reconstructed signal. (Solved similar problems in class.)**
### Distance measures (mostly for 2 marks)
- Euclidean distance
- Chessboard distance
- City-block distance
###  Adjacency of Pixel
- 4-adjacency
- 8-adjacency
- m-adjacency
## Unit II
### Image Segmentation
- **Region based - Quad-tree Decomposition**
### Edge detection
- **Canny Edge Detector - why, usecases and steps involved**
	1.  Gaussian Smoothing of input images
	2.  Gradient calculation along horizontal and vertical axis
	3.  Non-maximum suppression of false edges
	4.  Applying hysteresis thresholding
### Convolution and Correlation
- Questions will be generic in terms of methods to use like: **Perform 2-dimensional convolution/correlation on ...**. Can use any applicable method. Explain the steps if necessary.
### G/n [[2, 5, 6], [10, 122, 5], [6, 2, 5]] -- Replace the center pixel value with ... (10 marks)
**1. Averaging Filter (Default A mask: 1/16 [[1, 1, 1], [1, 1, 1], [1, 1, 1]] ),
2. WA Filter (Default WA mask: 1/16 * [[1, 2, 1], [2, 4, 2], [1, 2, 1]] ),
3. Median Filter,
4. Mode Filter,
5. Mid point `(max_val - min_val) / 2`,
6. Harmonic Mean,
7. etc...**

P.S. Know the different means.
### Histogram Equalization - Will be covered upcoming Wednesday (21AUG2024).
