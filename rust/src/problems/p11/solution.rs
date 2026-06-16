pub struct Solution;

impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut max_volume = i32::MIN;
        let mut left = 0;
        let mut right = height.len() - 1;

        while left < right {
            let width = right - left;
            let cur_height = if height[left] < height[right] {height[left]} else {height[right]};

            let cur_volume = width as i32 * cur_height;

            max_volume = cur_volume.max(max_volume);

             if height[left] < height[right] {left += 1} else {right -= 1};
        }

        max_volume
    }
}
