pub struct Solution;

fn is_vowel(b: u8) -> bool {
    matches!(b, b'a' | b'e' | b'i' | b'o' | b'u')
}

impl Solution {
    pub fn max_vowels(s: String, k: i32) -> i32 {
        let s_bytes = s.as_bytes();
        let mut left = 0;
        let usize_k = k as usize;
        let mut max_vowels_count = 0;
        let mut current_vowels_count: i32 = 0;

        for right in 0..s.len() {

            if is_vowel(s_bytes[right]) {
                current_vowels_count += 1;
            }

            if right - left + 1 > usize_k {
                if is_vowel(s_bytes[left]) {
                    current_vowels_count -= 1
                }
                left += 1
            }
            
            if right - left + 1 == usize_k {
                max_vowels_count = max_vowels_count.max(current_vowels_count)
            }

            if max_vowels_count == k {
                return max_vowels_count;
            }
        }

        return max_vowels_count
    }
}