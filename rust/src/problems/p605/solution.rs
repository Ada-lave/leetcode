pub struct Solution;

impl Solution {
    pub fn can_place_flowers(flowerbed: Vec<i32>, n: i32) -> bool {
        let mut flowerbed_n = n;
        let mut i = 0;
        while i < flowerbed.len() {
            if (i == 0 || flowerbed[i - 1] != 1) && (flowerbed[i] == 0) && (i == flowerbed.len() - 1 || flowerbed[i + 1] != 1) {
                flowerbed_n -= 1;
                i += 1;
            }
            i += 1;
        } 

        return flowerbed_n <= 0;
    }
}