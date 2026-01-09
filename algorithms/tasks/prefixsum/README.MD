Explaining the prefix sum concept with examples. Checking the codebase for related examples:

## The Core Idea

**Prefix sum** precomputes cumulative sums so you can answer range sum queries in O(1) instead of O(n).

### Basic Concept

Given an array `nums = [1, 2, 3, 4, 5]`, the prefix sum array is:

```
Original:  [1,  2,  3,  4,  5]
Prefix:    [1,  3,  6, 10, 15]
           ↑   ↑   ↑   ↑   ↑
           0   1   2   3   4
```

Where `prefix[i]` = sum of all elements from `nums[0]` to `nums[i]` inclusive.

### Why It's Useful

**Without prefix sum** (O(n) per query):
```go
// To get sum from index 2 to 4:
sum := 0
for i := 2; i <= 4; i++ {
    sum += nums[i]  // 3 + 4 + 5 = 12
}
```

**With prefix sum** (O(1) per query):
```go
// prefix = [1, 3, 6, 10, 15]
// Sum from index 2 to 4:
sum = prefix[4] - prefix[1]  // 15 - 3 = 12
```

### The Formula

For range sum `[left, right]`:
```
sumRange(left, right) = prefix[right] - prefix[left-1]
```

If `left == 0`, then `sumRange(0, right) = prefix[right]`.

### Example: Range Sum Query

For `nums = [-2, 0, 3, -5, 2, -1]`:

```
Original: [-2,  0,  3, -5,  2, -1]
Prefix:   [-2, -2,  1, -4, -2, -3]
           ↑    ↑   ↑   ↑   ↑   ↑
           0    1   2   3   4   5
```

- `sumRange(0, 2)` = `prefix[2]` = `1` ✓
- `sumRange(2, 5)` = `prefix[5] - prefix[1]` = `-3 - (-2)` = `-1` ✓
- `sumRange(0, 5)` = `prefix[5]` = `-3` ✓

### Implementation Pattern

```go
// Build prefix sum array
prefix := make([]int, len(nums))
prefix[0] = nums[0]
for i := 1; i < len(nums); i++ {
    prefix[i] = prefix[i-1] + nums[i]
}

// Query range sum [left, right]
func sumRange(left, right int) int {
    if left == 0 {
        return prefix[right]
    }
    return prefix[right] - prefix[left-1]
}
```

### Key Benefits

1. Precomputation: O(n) to build, then O(1) per query
2. Multiple queries: efficient when you need many range sums
3. Pattern: useful for subarray problems, sliding windows, and cumulative calculations

### In Your Pivot Index Problem

You're using a prefix sum concept implicitly:
- `leftSum` is the prefix sum up to (but not including) index `i`
- `rightSum` = `totalSum - leftSum - nums[i]`

This avoids building a full prefix array since you only need one value at a time.

The prefix sum technique is a common optimization for range queries and cumulative calculations.

Checking the README for context, then listing real-world use cases:

## Real-world use cases for prefix sum

## 1. **Financial Systems & Banking**

**Stock Price Analysis:**
- Calculate cumulative returns over time periods
- Find maximum profit in a time window
- Track portfolio value changes

**Transaction Processing:**
- Calculate account balance at any point in time
- Sum transactions within date ranges
- Generate monthly/yearly financial reports

```go
// Example: Calculate account balance at any point in time
type Account struct {
    transactions []int // positive = deposits, negative = withdrawals
    prefix       []int // prefix[i] = balance after transaction i
}

func NewAccount(transactions []int) *Account {
    prefix := make([]int, len(transactions))
    prefix[0] = transactions[0]
    for i := 1; i < len(transactions); i++ {
        prefix[i] = prefix[i-1] + transactions[i]
    }
    return &Account{transactions: transactions, prefix: prefix}
}

// Get balance after transaction at index i
func (a *Account) BalanceAfter(i int) int {
    return a.prefix[i]
}

// Get balance change between two transaction indices
func (a *Account) BalanceChange(from, to int) int {
    if from == 0 {
        return a.prefix[to]
    }
    return a.prefix[to] - a.prefix[from-1]
}

// Example: Calculate total sales in a date range
// prefix[i] = total sales from day 0 to day i
// Query: sales from day 5 to day 10 = prefix[10] - prefix[4]
```

## 2. **Analytics & Data Processing**

**Web Analytics:**
- Page views per time period
- User engagement metrics over ranges
- Cumulative user growth

**Time Series Data:**
- Sensor readings (temperature, pressure) over intervals
- Network traffic analysis
- System performance metrics

```go
// Example: Track page views per time period
type PageViewTracker struct {
    hourlyViews []int // views per hour
    prefix       []int // prefix[i] = total views from hour 0 to hour i
}

func NewPageViewTracker(hourlyViews []int) *PageViewTracker {
    prefix := make([]int, len(hourlyViews))
    prefix[0] = hourlyViews[0]
    for i := 1; i < len(hourlyViews); i++ {
        prefix[i] = prefix[i-1] + hourlyViews[i]
    }
    return &PageViewTracker{hourlyViews: hourlyViews, prefix: prefix}
}

// Get total views from hour start to hour end (inclusive)
func (p *PageViewTracker) ViewsInRange(start, end int) int {
    if start == 0 {
        return p.prefix[end]
    }
    return p.prefix[end] - p.prefix[start-1]
}

// Example usage:
// tracker := NewPageViewTracker([]int{100, 150, 200, 180, 120})
// views := tracker.ViewsInRange(1, 3) // views from hour 1 to 3 = 530
```

## 3. **Image Processing & Computer Vision**

**2D Prefix Sum (Integral Image):**
- Fast calculation of sum of pixels in rectangular regions
- Used in face detection algorithms (Haar cascades)
- Image filtering and feature extraction

```go
// Example: 2D Prefix Sum for image processing
type ImageSum struct {
    pixels [][]int
    prefix [][]int // prefix[i][j] = sum of all pixels in rectangle [0,0] to [i,j]
}

func NewImageSum(pixels [][]int) *ImageSum {
    rows, cols := len(pixels), len(pixels[0])
    prefix := make([][]int, rows)
    
    // Build 2D prefix sum
    for i := 0; i < rows; i++ {
        prefix[i] = make([]int, cols)
        for j := 0; j < cols; j++ {
            prefix[i][j] = pixels[i][j]
            if i > 0 {
                prefix[i][j] += prefix[i-1][j]
            }
            if j > 0 {
                prefix[i][j] += prefix[i][j-1]
            }
            if i > 0 && j > 0 {
                prefix[i][j] -= prefix[i-1][j-1]
            }
        }
    }
    return &ImageSum{pixels: pixels, prefix: prefix}
}

// Get sum of pixels in rectangle from (r1, c1) to (r2, c2) inclusive
func (img *ImageSum) SumRegion(r1, c1, r2, c2 int) int {
    sum := img.prefix[r2][c2]
    if r1 > 0 {
        sum -= img.prefix[r1-1][c2]
    }
    if c1 > 0 {
        sum -= img.prefix[r2][c1-1]
    }
    if r1 > 0 && c1 > 0 {
        sum += img.prefix[r1-1][c1-1]
    }
    return sum
}

// 2D prefix sum allows O(1) sum of any rectangle
// Instead of O(width * height) per query
```

## 4. **Gaming & Game Development**

**Score Tracking:**
- Cumulative scores over levels
- Leaderboard calculations
- Achievement progress tracking

**Resource Management:**
- Total resources collected in time periods
- Experience points accumulation
- Inventory value calculations

```go
// Example: Cumulative score tracking in games
type ScoreTracker struct {
    levelScores []int // score earned per level
    prefix      []int // prefix[i] = total score after level i
}

func NewScoreTracker(levelScores []int) *ScoreTracker {
    prefix := make([]int, len(levelScores))
    prefix[0] = levelScores[0]
    for i := 1; i < len(levelScores); i++ {
        prefix[i] = prefix[i-1] + levelScores[i]
    }
    return &ScoreTracker{levelScores: levelScores, prefix: prefix}
}

// Get total score after completing level i
func (s *ScoreTracker) TotalScoreAfter(level int) int {
    return s.prefix[level]
}

// Get score earned between two levels
func (s *ScoreTracker) ScoreBetween(fromLevel, toLevel int) int {
    if fromLevel == 0 {
        return s.prefix[toLevel]
    }
    return s.prefix[toLevel] - s.prefix[fromLevel-1]
}
```

## 5. **Database Systems**

**Query Optimization:**
- Pre-computed aggregations for fast range queries
- Materialized views with cumulative data
- Time-series databases (InfluxDB, TimescaleDB)

**Index Structures:**
- B+ trees with prefix sums for range queries
- Columnar databases (OLAP) for analytical queries

## 6. **Network & System Monitoring**

**Bandwidth Monitoring:**
- Total data transferred in time windows
- Network usage statistics
- Rate limiting calculations

**System Metrics:**
- CPU usage over time periods
- Memory consumption tracking
- Request count aggregations

```go
// Example: Network bandwidth monitoring
type BandwidthMonitor struct {
    hourlyBytes []int64 // bytes transferred per hour
    prefix      []int64 // prefix[i] = total bytes from hour 0 to hour i
}

func NewBandwidthMonitor(hourlyBytes []int64) *BandwidthMonitor {
    prefix := make([]int64, len(hourlyBytes))
    prefix[0] = hourlyBytes[0]
    for i := 1; i < len(hourlyBytes); i++ {
        prefix[i] = prefix[i-1] + hourlyBytes[i]
    }
    return &BandwidthMonitor{hourlyBytes: hourlyBytes, prefix: prefix}
}

// Get total data transferred from hour start to hour end
func (b *BandwidthMonitor) DataTransferred(start, end int) int64 {
    if start == 0 {
        return b.prefix[end]
    }
    return b.prefix[end] - b.prefix[start-1]
}

// Check if data limit exceeded in time window
func (b *BandwidthMonitor) ExceedsLimit(start, end int, limit int64) bool {
    return b.DataTransferred(start, end) > limit
}
```

## 7. **Geographic Information Systems (GIS)**

**Population Density:**
- Cumulative population in regions
- Resource distribution analysis
- Spatial range queries

## 8. **E-commerce & Retail**

**Sales Analytics:**
- Revenue in date ranges
- Product sales over periods
- Inventory value calculations

**Customer Analytics:**
- Total purchases per customer
- Spending patterns over time
- Loyalty program points

```go
// Example: E-commerce sales analytics
type SalesAnalytics struct {
    dailyRevenue []float64 // revenue per day
    prefix       []float64 // prefix[i] = total revenue from day 0 to day i
}

func NewSalesAnalytics(dailyRevenue []float64) *SalesAnalytics {
    prefix := make([]float64, len(dailyRevenue))
    prefix[0] = dailyRevenue[0]
    for i := 1; i < len(dailyRevenue); i++ {
        prefix[i] = prefix[i-1] + dailyRevenue[i]
    }
    return &SalesAnalytics{dailyRevenue: dailyRevenue, prefix: prefix}
}

// Get total revenue from day start to day end
func (s *SalesAnalytics) RevenueInRange(start, end int) float64 {
    if start == 0 {
        return s.prefix[end]
    }
    return s.prefix[end] - s.prefix[start-1]
}

// Get average daily revenue in range
func (s *SalesAnalytics) AverageDailyRevenue(start, end int) float64 {
    days := end - start + 1
    return s.RevenueInRange(start, end) / float64(days)
}
```

## 9. **Scientific Computing**

**Signal Processing:**
- Cumulative energy in signals
- Frequency domain analysis
- Moving averages

**Statistics:**
- Cumulative distribution functions
- Running statistics (mean, variance)
- Probability calculations

```go
// Example: Moving average using prefix sum
func MovingAverage(nums []float64, windowSize int) []float64 {
    // Build prefix sum
    prefix := make([]float64, len(nums))
    prefix[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        prefix[i] = prefix[i-1] + nums[i]
    }
    
    // Calculate moving averages
    result := make([]float64, len(nums)-windowSize+1)
    for i := 0; i <= len(nums)-windowSize; i++ {
        end := i + windowSize - 1
        var sum float64
        if i == 0 {
            sum = prefix[end]
        } else {
            sum = prefix[end] - prefix[i-1]
        }
        result[i] = sum / float64(windowSize)
    }
    return result
}

// Example: Calculate 3-day moving average
// nums = [1.0, 2.0, 3.0, 4.0, 5.0]
// MovingAverage(nums, 3) = [2.0, 3.0, 4.0]
// [0:2] = (1+2+3)/3 = 2.0
// [1:3] = (2+3+4)/3 = 3.0
// [2:4] = (3+4+5)/3 = 4.0
```

## 10. **Competitive Programming & Interviews**

Common patterns:
- Subarray sum problems
- Sliding window optimizations
- Range query problems
- Kadane's algorithm variants

## Why It Matters in Production

1. Performance: O(1) queries instead of O(n) when you need many range sums
2. Scalability: Handles high query volumes efficiently
3. Cost: Reduces database load and computation costs
4. Real-time: Enables fast analytics dashboards and reporting

## Example: Real-Time Analytics Dashboard

```go
// Pre-compute prefix sums for hourly sales
// When user requests "sales from 2pm to 5pm":
// O(1) lookup: prefix[5pm] - prefix[1pm]
// Instead of summing thousands of transactions on-the-fly
```

The pattern appears whenever you need to answer many range sum queries quickly. It's a fundamental optimization technique in systems that process time-series data, analytics, and any scenario with frequent range queries.