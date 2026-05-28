pub struct Solution;

impl Solution {
    pub fn reverse_words(s: String) -> String {
        let mut result: Vec<String> = Vec::new();
        let mut current_string = String::new();
        for char in  s.chars() {
            if char != ' ' {
                std::fmt::Write::write_char(&mut current_string, char).unwrap();
            } else {
                if !current_string.is_empty() {
                    result.push(current_string.clone());
                }
                current_string.clear();
            }
        }

        if !current_string.is_empty() {
            result.push(current_string.clone());
        }

        result.reverse();
        return result.join(" ");
    }
}
