# Given a string, find the length of the longest substring without repeating characters.
#
# Examples:
#
# Given "abcabcbb", the answer is "abc", which the length is 3.
#
# Given "bbbbb", the answer is "b", with the length of 1.
#
# Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

class Solution(object):
    def lengthOfLongestSubstring(self, s):
        # edge case s is shorter than 2 length
        t = len(s)
        if t < 2:
            return t 
            
        longest = 1
        currLen = 1
        
        startIndex = 0
        for endIndex in range(1, t):
            
            # if a repeat character is found
            if s[endIndex] in s[startIndex:endIndex]:
                
                # assign longest
                if currLen > longest:
                    longest = currLen
                    
                # move start index forward
                while (s[startIndex] != s[endIndex]):
                    startIndex += 1
                startIndex += 1
                
                #if (t - startIndex < longest):
                #    return longest
				
                # reset counter
                currLen = endIndex - startIndex
                
            currLen = currLen + 1
            
        # assign longest
        if currLen > longest:
            longest = currLen
 
        return longest
        
        
        """
        :type s: str
        :rtype: int
        """