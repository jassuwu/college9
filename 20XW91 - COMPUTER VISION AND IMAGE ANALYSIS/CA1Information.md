
# 16AUG2024 Internal Test I Revision Class

## Unit I

### Image Digitization
- **What are the steps to convert to a Digital Image? / How to do Image Digitization?**

  Image digitization is a three-step process:

  1. **Sampling**:
     - Denoted as $fs(x, y)$.
     - Involves taking specific instances of the signal.
     - It is crucial to choose the appropriate sampling frequency to avoid **Aliasing**.

  2. **Quantization**:
     - Mapping a large set of values to a smaller set.
     - This is a non-linear and irreversible process.
     - Can be scalar or vector.
     - Irreversible because it's a many-to-one mapping, leading to loss of uniqueness.

  3. **Encoding**:
     - Conversion of quantized values into a digital form.
     - Involves fixed or variable length encoding.

- **Example Problem:**

  For $signal$ defined as $2 \cos(2 \pi (4x + 6y))$, with $\Delta x = 0.2$ and $\Delta y = 0.4$, find the reconstructed signal. (Refer to similar solved problems in class.)

### Distance Measures (mostly for 2 marks)
- **Euclidean Distance**
- **Chessboard Distance**
- **City-Block Distance**

### Adjacency of Pixels
- **4-Adjacency**
- **8-Adjacency**
- **m-Adjacency**

## Unit II

### Image Segmentation
- **Region-based Segmentation - Quad-tree Decomposition**

### Edge Detection
- **Canny Edge Detector:**
  - **Why:** To detect edges in an image with accuracy.
  - **Use Cases:** Object detection, image analysis.
  - **Steps Involved:**
    1. **Gaussian Smoothing**: Smooth the input image to reduce noise.
    2. **Gradient Calculation**: Compute gradients along horizontal and vertical axes.
    3. **Non-Maximum Suppression**: Eliminate false edges by keeping only local maxima.
    4. **Hysteresis Thresholding**: Apply double thresholding to finalize edge detection.

### Convolution and Correlation
- **Tasks:** Questions will be generic, such as "Perform 2-dimensional convolution/correlation on...". Use applicable methods and explain the steps if necessary.

### Image Filtering
- **Example Problem:**

  Given the matrix as:

  $$
  \begin{bmatrix}
  2 & 5 & 6 \\
  10 & 122 & 5 \\
  6 & 2 & 5
  \end{bmatrix}
  $$

  Replace the center pixel value with:

  1. **Averaging Filter** (Default mask: $\frac{1}{16} \begin{bmatrix} 1 & 1 & 1 \\ 1 & 1 & 1 \\ 1 & 1 & 1 \end{bmatrix}$)
  2. **Weighted Averaging (WA) Filter** (Default WA mask: $\frac{1}{16} \begin{bmatrix} 1 & 2 & 1 \\ 2 & 4 & 2 \\ 1 & 2 & 1 \end{bmatrix}$)
  3. **Median Filter**
  4. **Mode Filter**
  5. **Midpoint Filter**: $\frac{\text{max\_val} - \text{min\_val}}{2}$
  6. **Harmonic Mean Filter**

  *Note: Familiarize yourself with different means.*

### Histogram Equalization
- Will be covered in the upcoming class on **21AUG2024**.


