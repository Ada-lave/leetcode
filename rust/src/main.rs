pub mod problems;
use problems::p1456::solution;

fn main() {
    solution::Solution::max_vowels("abciiidef".to_string(), 3);
}
