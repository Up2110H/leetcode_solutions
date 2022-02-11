package hashTable

/*func findAnagrams(s string, p string) []int {
	hash := make([]int, 26)
	dynamicHash := make([]int, 26)
	var ans []int
	c, l, k := 0, -1, len(p)

	for _, v := range p {
		if hash[v-'a'] == 0 {
			c++
		}
		hash[v-'a']++
	}

	counter := c

	for i, v := range s {
		if hash[v-'a'] == 0 {
			dynamicHash = make([]int, 26)
			l = i
			counter = c
			continue
		}
		if l < i - k {
			dynamicHash[s[i-k]-'a']--

			if tmp := hash[s[i-k]-'a'] - dynamicHash[s[i-k]-'a']; tmp == 0 {
				counter--
			} else if tmp == 1 {
				counter++
			}
		}
		dynamicHash[v-'a']++

		if tmp := hash[v-'a'] - dynamicHash[v-'a']; tmp == 0 {
			counter--
		} else if tmp == -1 {
			counter++
		}

		if counter == 0 {
			ans = append(ans, i - k + 1)
		}
	}

	return ans
}*/

func findAnagrams(s string, p string) []int {
	var hash, dynamicHash [26]int
	var ans []int
	k := len(p)

	for _, v := range p {
		hash[v-'a']++
	}

	for i, v := range s {
		if i >= k {
			dynamicHash[s[i-k]-'a']--
		}
		dynamicHash[v-'a']++
		if hash == dynamicHash {
			ans = append(ans, i-k+1)
		}
	}

	return ans
}

/* 23:11 11.02.2022 - 23:45 11.02.2022 */
/* 23:45 11.02.2022 - 23:53 11.02.2022 */
