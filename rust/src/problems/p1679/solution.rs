pub struct Solution;

impl Solution {
    pub fn max_operations(mut nums: Vec<i32>, k: i32) -> i32 {
        let mut counter = 0;
        nums.sort();
        let (mut idx1, mut idx2) = (0 as usize, nums.len() - 1);

        while idx1 < idx2 {
            let sum_of_idxs = nums[idx1] + nums[idx2];
            if  sum_of_idxs == k {
                counter += 1;
                idx1 += 1;
                idx2 -= 1
            }   else if sum_of_idxs < k {
                idx1 += 1
            } else {
                idx2 -= 1
            }
        }

        counter
    }
}
