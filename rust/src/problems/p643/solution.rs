pub struct Solution;

impl Solution {
    pub fn find_max_average(nums: Vec<i32>, k: i32) -> f64 {
        let mut max_average = f64::MIN;
        let mut sum = 0;
        let mut left = 0;
        let k_usize = k as usize;
        for right in 0..nums.len() {
            sum += nums[right];

            if right - left + 1 > k_usize {
                sum -= nums[left];
                left += 1;
            }

            if right - left + 1 == k_usize {
                max_average = max_average.max((sum as f64 / k as f64).into())
            }
        }
        max_average
    }
}
