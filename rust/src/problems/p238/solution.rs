pub struct Solution;

impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let mut result:Vec<i32> = vec![0; nums.len()];
        
        let mut left:Vec<i32> = vec![0; nums.len()];
        let mut right:Vec<i32> = vec![0; nums.len()];

        let mut suffix = 1;
        let mut prefix = 1;

        for i in 0..nums.len() {
            left[i] = prefix;
            prefix *= nums[i]
        }

        for i in (0..nums.len()).rev() {
            right[i] = suffix;
            suffix *= nums[i]
        }

        for i in 0..nums.len() {
            result[i] = left[i] * right[i];
        }

        return result;
    }
}