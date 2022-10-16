package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/winebarrel/cronplan/utils"
)

func TestList(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    int
		end      int
		origin   int
		max      int
		expected []int
	}{
		{2, 11, 1, 12, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{1, 12, 1, 12, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{11, 2, 1, 12, []int{11, 12, 1, 2}},
		{12, 1, 1, 12, []int{12, 1}},
		{1, 22, 0, 23, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}},
		{22, 1, 0, 23, []int{22, 23, 0, 1}},
		{0, 23, 0, 23, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}},
		{23, 0, 0, 23, []int{23, 0}},
	}

	for _, t := range tt {
		actual, err := utils.List(t.start, t.end, t.origin, t.max)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}

func TestListError(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    int
		end      int
		origin   int
		max      int
		expected string
	}{
		{0, 11, 1, 12, "'start' must be >= 'origin'"},
		{1, 13, 1, 12, "'end' must be <= 'max'"},
		{13, 2, 1, 12, "'start' must be <= 'max'"},
		{11, 0, 1, 12, "'end' must be >= 'origin'"},
	}

	for _, t := range tt {
		_, err := utils.List(t.start, t.end, t.origin, t.max)
		assert.EqualError(err, t.expected)
	}
}

func TestListMinute(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    int
		end      int
		expected []int
	}{
		{1, 58, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58}},
		{58, 1, []int{58, 59, 0, 1}},
		{0, 59, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59}},
		{59, 0, []int{59, 0}},
	}

	for _, t := range tt {
		actual, err := utils.ListMinute(t.start, t.end)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}

func TestListHour(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    int
		end      int
		expected []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}},
		{22, 1, []int{22, 23, 0, 1}},
		{0, 23, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23}},
		{23, 0, []int{23, 0}},
	}

	for _, t := range tt {
		actual, err := utils.ListHour(t.start, t.end)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}

func TestListDayOfMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		tm       time.Time
		start    int
		end      int
		expected []int
	}{
		{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 2, 30, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}},
		{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 30, 2, []int{30, 31, 1, 2}},
		{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 1, 31, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}},
		{time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 31, 1, []int{31, 1}},
		{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC), 2, 27, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}},
		{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC), 27, 2, []int{27, 28, 1, 2}},
		{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC), 1, 28, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28}},
		{time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC), 28, 1, []int{28, 1}},
	}

	for _, t := range tt {
		actual, err := utils.ListDayOfMonth(t.tm, t.start, t.end)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}

func TestListMonth(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    time.Month
		end      time.Month
		expected []time.Month
	}{
		{time.February, time.November, []time.Month{time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November}},
		{time.January, time.December, []time.Month{time.January, time.February, time.March, time.April, time.May, time.June, time.July, time.August, time.September, time.October, time.November, time.December}},
		{time.November, time.February, []time.Month{time.November, time.December, time.January, time.February}},
		{time.December, time.January, []time.Month{time.December, 1}},
	}

	for _, t := range tt {
		actual, err := utils.ListMonth(t.start, t.end)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}

func TestListWeekday(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    time.Weekday
		end      time.Weekday
		expected []time.Weekday
	}{
		{time.Tuesday, time.Friday, []time.Weekday{time.Tuesday, time.Wednesday, time.Thursday, time.Friday}},
		{time.Friday, time.Tuesday, []time.Weekday{time.Friday, time.Saturday, time.Sunday, time.Monday, time.Tuesday}},
		{time.Monday, time.Sunday, []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday}},
		{time.Monday, time.Saturday, []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}},
		{time.Sunday, time.Monday, []time.Weekday{time.Sunday, time.Monday}},
		{time.Sunday, time.Saturday, []time.Weekday{time.Sunday, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday}},
	}

	for _, t := range tt {
		actual, err := utils.ListWeekday(t.start, t.end)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}

func TestLisYear(t *testing.T) {
	assert := assert.New(t)

	tt := []struct {
		start    int
		end      int
		expected []int
	}{
		{1971, 2198, []int{1971, 1972, 1973, 1974, 1975, 1976, 1977, 1978, 1979, 1980, 1981, 1982, 1983, 1984, 1985, 1986, 1987, 1988, 1989, 1990, 1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029, 2030, 2031, 2032, 2033, 2034, 2035, 2036, 2037, 2038, 2039, 2040, 2041, 2042, 2043, 2044, 2045, 2046, 2047, 2048, 2049, 2050, 2051, 2052, 2053, 2054, 2055, 2056, 2057, 2058, 2059, 2060, 2061, 2062, 2063, 2064, 2065, 2066, 2067, 2068, 2069, 2070, 2071, 2072, 2073, 2074, 2075, 2076, 2077, 2078, 2079, 2080, 2081, 2082, 2083, 2084, 2085, 2086, 2087, 2088, 2089, 2090, 2091, 2092, 2093, 2094, 2095, 2096, 2097, 2098, 2099, 2100, 2101, 2102, 2103, 2104, 2105, 2106, 2107, 2108, 2109, 2110, 2111, 2112, 2113, 2114, 2115, 2116, 2117, 2118, 2119, 2120, 2121, 2122, 2123, 2124, 2125, 2126, 2127, 2128, 2129, 2130, 2131, 2132, 2133, 2134, 2135, 2136, 2137, 2138, 2139, 2140, 2141, 2142, 2143, 2144, 2145, 2146, 2147, 2148, 2149, 2150, 2151, 2152, 2153, 2154, 2155, 2156, 2157, 2158, 2159, 2160, 2161, 2162, 2163, 2164, 2165, 2166, 2167, 2168, 2169, 2170, 2171, 2172, 2173, 2174, 2175, 2176, 2177, 2178, 2179, 2180, 2181, 2182, 2183, 2184, 2185, 2186, 2187, 2188, 2189, 2190, 2191, 2192, 2193, 2194, 2195, 2196, 2197, 2198}},
		{2198, 1971, []int{2198, 2199, 1970, 1971}},
		{1970, 2199, []int{1970, 1971, 1972, 1973, 1974, 1975, 1976, 1977, 1978, 1979, 1980, 1981, 1982, 1983, 1984, 1985, 1986, 1987, 1988, 1989, 1990, 1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029, 2030, 2031, 2032, 2033, 2034, 2035, 2036, 2037, 2038, 2039, 2040, 2041, 2042, 2043, 2044, 2045, 2046, 2047, 2048, 2049, 2050, 2051, 2052, 2053, 2054, 2055, 2056, 2057, 2058, 2059, 2060, 2061, 2062, 2063, 2064, 2065, 2066, 2067, 2068, 2069, 2070, 2071, 2072, 2073, 2074, 2075, 2076, 2077, 2078, 2079, 2080, 2081, 2082, 2083, 2084, 2085, 2086, 2087, 2088, 2089, 2090, 2091, 2092, 2093, 2094, 2095, 2096, 2097, 2098, 2099, 2100, 2101, 2102, 2103, 2104, 2105, 2106, 2107, 2108, 2109, 2110, 2111, 2112, 2113, 2114, 2115, 2116, 2117, 2118, 2119, 2120, 2121, 2122, 2123, 2124, 2125, 2126, 2127, 2128, 2129, 2130, 2131, 2132, 2133, 2134, 2135, 2136, 2137, 2138, 2139, 2140, 2141, 2142, 2143, 2144, 2145, 2146, 2147, 2148, 2149, 2150, 2151, 2152, 2153, 2154, 2155, 2156, 2157, 2158, 2159, 2160, 2161, 2162, 2163, 2164, 2165, 2166, 2167, 2168, 2169, 2170, 2171, 2172, 2173, 2174, 2175, 2176, 2177, 2178, 2179, 2180, 2181, 2182, 2183, 2184, 2185, 2186, 2187, 2188, 2189, 2190, 2191, 2192, 2193, 2194, 2195, 2196, 2197, 2198, 2199}},
		{2199, 1970, []int{2199, 1970}},
	}

	for _, t := range tt {
		actual, err := utils.ListYear(t.start, t.end)
		assert.NoError(err)
		assert.Equal(t.expected, actual, t)
	}
}
