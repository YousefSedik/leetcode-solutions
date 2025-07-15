class Solution:
    def isValid(self, word: str) -> bool:
        vowels = {'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u'}
        has_vowel = any([True if letter in vowels else False for letter in word])
        has_consonant = any([True if (letter not in vowels and letter.isalpha()) else False for letter in word])
        if len(word) >= 3 and word.isalnum() and has_vowel and has_consonant:
            return True
        return False #UuE6