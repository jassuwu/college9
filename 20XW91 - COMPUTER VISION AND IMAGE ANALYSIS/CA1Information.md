# 16AUG2024 CA Test I Infomation Class

## Unit I

### Image Digitization

- **Steps to Convert to a Digital Image / How to Do Image Digitization:**

  Image digitization involves three main steps:

  1. **Sampling**
     - Denoted as `fs(x, y)`.
     - This involves taking discrete samples of the signal.
     - It’s important to choose the right sampling frequency to avoid **Aliasing**.

  2. **Quantization**
     - This maps a large set of values to a smaller set.
     - It’s a non-linear and irreversible process.
     - Quantization can be scalar or vector.
     - It’s irreversible due to many-to-one mapping, leading to a loss of uniqueness.

  3. **Encoding**
     - Converts quantized values into a digital form.
     - Involves either fixed or variable length encoding.

- **Example Problem:**

  For a function defined as `g/n = 2 * cos(2 * pi * (4x + 6y))`, with `Δx = 0.2` and `Δy = 0.4`, find the reconstructed signal. (Refer to similar problems solved in class.)

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

- **Region-Based Segmentation - Quad-tree Decomposition**

### Edge Detection

- **Canny Edge Detector:**
  - **Purpose:** To detect edges in an image with high accuracy.
  - **Use Cases:** Object detection, image analysis.
  - **Steps Involved:**
    1. **Gaussian Smoothing:** Smooth the image to reduce noise.
    2. **Gradient Calculation:** Compute gradients in horizontal and vertical directions.
    3. **Non-Maximum Suppression:** Retain only local maxima to eliminate false edges.
    4. **Hysteresis Thresholding:** Apply double thresholding to finalize edge detection.

### Convolution and Correlation

- **Tasks:** Perform 2-dimensional convolution/correlation as specified. Apply relevant methods and explain the steps if needed.

### Image Filtering

- **Example Problem:**

  Given the matrix `G/n` as:

  ```
  2   5   6
  10 122  5
  6   2   5
  ```

  Replace the center pixel value with the following filters:

  1. **Averaging Filter** (Default mask: `1/16 * [1 1 1; 1 1 1; 1 1 1]`)
  2. **Weighted Averaging (WA) Filter** (Default WA mask: `1/16 * [1 2 1; 2 4 2; 1 2 1]`)
  3. **Median Filter**
  4. **Mode Filter**
  5. **Midpoint Filter:** `(max_val - min_val) / 2`
  6. **Harmonic Mean Filter**

  *Note: Be familiar with different types of means.*

### Histogram Equalization

- Will be covered in the upcoming class on **21AUG2024**.
