pub struct Solution;

impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        if nums.len() <= 1 {
            return;
        } 

        let mut writer_ptr = 0;

        for non_zero_reader_ptr in 0..nums.len()  {

            if nums[non_zero_reader_ptr] != 0 {
                (nums[writer_ptr], nums[non_zero_reader_ptr]) = (nums[non_zero_reader_ptr], nums[writer_ptr]);
                writer_ptr += 1;
            }       
        }
    }
}

