use std::collections::HashSet;

pub struct Solution;

impl Solution {
    pub fn reverse_vowels(s: String) -> String {
        let mut result_chars: Vec<char> = s.chars().collect();
        let vowels: HashSet<char> = HashSet::from(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']);
        let (mut idx1, mut idx2) = (0, s.len() - 1);
        while idx1 < idx2 {
            let exists1 = vowels.contains(&result_chars[idx1]);
            if !exists1 {
                idx1 += 1;
            }

            let exists2 = vowels.contains(&result_chars[idx2]);
            if !exists2 {
                idx2 -= 1;
            }
            if exists1 && exists2 {
                (result_chars[idx1], result_chars[idx2]) = (result_chars[idx2], result_chars[idx1]);
                idx1 += 1;
                idx2 -= 1;
            }
        }
        
        return String::from_iter(result_chars);
    }
}
