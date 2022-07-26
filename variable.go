package main

var (
	ShiroKeys = []string{
		"2AvVhdsgUs0FSA3SDFAdag==", "kPH+bIxk5D2deZiIxcaaaA==", "3AvVhmFLUs0KTA3Kprsdag==", "4AvVhmFLUs0KTA3Kprsdag==",
		"5aaC5qKm5oqA5pyvAAAAAA==", "6ZmI6I2j5Y+R5aSn5ZOlAA==", "bWljcm9zAAAAAAAAAAAAAA==", "wGiHplamyXlVB11UXWol8g==",
		"Z3VucwAAAAAAAAAAAAAAAA==", "MTIzNDU2Nzg5MGFiY2RlZg==", "zSyK5Kp6PZAAjlT+eeNMlg==", "U3ByaW5nQmxhZGUAAAAAAA==",
		"5AvVhmFLUs0KTA3Kprsdag==", "bXdrXl9eNjY2KjA3Z2otPQ==", "fCq+/xW488hMTCD+cmJ3aQ==", "1QWLxg+NYmxraMoxAXu/Iw==",
		"ZUdsaGJuSmxibVI2ZHc9PQ==", "L7RioUULEFhRyxM7a2R/Yg==", "r0e3c16IdVkouZgk1TKVMg==", "bWluZS1hc3NldC1rZXk6QQ==",
		"a2VlcE9uR29pbmdBbmRGaQ==", "WcfHGU25gNnTxTlmJMeSpw==", "ZAvph3dsQs0FSL3SDFAdag==", "tiVV6g3uZBGfgshesAQbjA==",
		"cmVtZW1iZXJNZQAAAAAAAA==", "ZnJlc2h6Y24xMjM0NTY3OA==", "RVZBTk5JR0hUTFlfV0FPVQ==", "WkhBTkdYSUFPSEVJX0NBVA==",
		"GsHaWo4m1eNbE0kNSMULhg==", "l8cc6d2xpkT1yFtLIcLHCg==", "KU471rVNQ6k7PQL4SqxgJg==", "0AvVhmFLUs0KTA3Kprsdag==",
		"1AvVhdsgUs0FSA3SDFAdag==", "25BsmdYwjnfcWmnhAciDDg==", "3JvYhmBLUs0ETA5Kprsdag==", "6AvVhmFLUs0KTA3Kprsdag==",
		"6NfXkC7YVCV5DASIrEm1Rg==", "7AvVhmFLUs0KTA3Kprsdag==", "8AvVhmFLUs0KTA3Kprsdag==", "8BvVhmFLUs0KTA3Kprsdag==",
		"9AvVhmFLUs0KTA3Kprsdag==", "OUHYQzxQ/W9e/UjiAGu6rg==", "a3dvbmcAAAAAAAAAAAAAAA==", "aU1pcmFjbGVpTWlyYWNsZQ==",
		"bXRvbnMAAAAAAAAAAAAAAA==", "OY//C4rhfwNxCQAQCrQQ1Q==", "5J7bIJIV0LQSN3c9LPitBQ==", "f/SY5TIve5WWzT4aQlABJA==",
		"bya2HkYo57u6fWh5theAWw==", "WuB+y2gcHRnY2Lg9+Aqmqg==", "3qDVdLawoIr1xFd6ietnwg==", "YI1+nBV//m7ELrIyDHm6DQ==",
		"6Zm+6I2j5Y+R5aS+5ZOlAA==", "2A2V+RFLUs+eTA3Kpr+dag==", "6ZmI6I2j3Y+R1aSn5BOlAA==", "SkZpbmFsQmxhZGUAAAAAAA==",
		"2cVtiE83c4lIrELJwKGJUw==", "fsHspZw/92PrS3XrPW+vxw==", "XTx6CKLo/SdSgub+OPHSrw==", "sHdIjUN6tzhl8xZMG3ULCQ==",
		"O4pdf+7e+mZe8NyxMTPJmQ==", "HWrBltGvEZc14h9VpMvZWw==", "rPNqM6uKFCyaL10AK51UkQ==", "Y1JxNSPXVwMkyvES/kJGeQ==",
		"lT2UvDUmQwewm6mMoiw4Ig==", "MPdCMZ9urzEA50JDlDYYDg==", "xVmmoltfpb8tTceuT5R7Bw==", "c+3hFGPjbgzGdrC+MHgoRQ==",
		"ClLk69oNcA3m+s0jIMIkpg==", "Bf7MfkNR0axGGptozrebag==", "1tC/xrDYs8ey+sa3emtiYw==", "ZmFsYWRvLnh5ei5zaGlybw==",
		"cGhyYWNrY3RmREUhfiMkZA==", "IduElDUpDDXE677ZkhhKnQ==", "yeAAo1E8BOeAYfBlm4NG9Q==", "cGljYXMAAAAAAAAAAAAAAA==",
		"2itfW92XazYRi5ltW0M2yA==", "XgGkgqGqYrix9lI6vxcrRw==", "ertVhmFLUs0KTA3Kprsdag==", "5AvVhmFLUS0ATA4Kprsdag==",
		"s0KTA3mFLUprK4AvVhsdag==", "hBlzKg78ajaZuTE0VLzDDg==", "9FvVhtFLUs0KnA3Kprsdyg==", "d2ViUmVtZW1iZXJNZUtleQ==",
		"yNeUgSzL/CfiWw1GALg6Ag==", "NGk/3cQ6F5/UNPRh8LpMIg==", "4BvVhmFLUs0KTA3Kprsdag==", "MzVeSkYyWTI2OFVLZjRzZg==",
		"empodDEyMwAAAAAAAAAAAA==", "A7UzJgh1+EWj5oBFi+mSgw==", "c2hpcm9fYmF0aXMzMgAAAA==", "i45FVt72K2kLgvFrJtoZRw==",
		"U3BAbW5nQmxhZGUAAAAAAA==", "Jt3C93kMR9D5e8QzwfsiMw==", "MTIzNDU2NzgxMjM0NTY3OA==", "vXP33AonIp9bFwGl7aT7rA==",
		"V2hhdCBUaGUgSGVsbAAAAA==", "Q01TX0JGTFlLRVlfMjAxOQ==", "Is9zJ3pzNh2cgTHB4ua3+Q==", "NsZXjXVklWPZwOfkvk6kUA==",
		"GAevYnznvgNCURavBhCr1w==", "66v1O8keKNV3TTcGPK1wzg==", "SDKOLKn2J1j/2BHjeZwAoQ==", "kPH+bIxk5D2deZiIxcabaA==",
		"kPH+bIxk5D2deZiIxcacaA==", "3AvVhdAgUs0FSA4SDFAdBg==", "4AvVhdsgUs0F563SDFAdag==", "FL9HL9Yu5bVUJ0PDU1ySvg==",
		"5RC7uBZLkByfFfJm22q/Zw==", "eXNmAAAAAAAAAAAAAAAAAA==", "fdCEiK9YvLC668sS43CJ6A==", "FJoQCiz0z5XWz2N2LyxNww==",
		"HeUZ/LvgkO7nsa18ZyVxWQ==", "HoTP07fJPKIRLOWoVXmv+Q==", "iycgIIyCatQofd0XXxbzEg==", "m0/5ZZ9L4jjQXn7MREr/bw==",
		"NoIw91X9GSiCrLCF03ZGZw==", "oPH+bIxk5E2enZiIxcqaaA==", "QAk0rp8sG0uJC4Ke2baYNA==", "Rb5RN+LofDWJlzWAwsXzxg==",
		"s2SE9y32PvLeYo+VGFpcKA==", "SrpFBcVD89eTQ2icOD0TMg==", "U0hGX2d1bnMAAAAAAAAAAA==", "Us0KvVhTeasAm43KFLAeng==",
		"Ymx1ZXdoYWxlAAAAAAAAAA==", "YWJjZGRjYmFhYmNkZGNiYQ==", "zIiHplamyXlVB11UXWol8g==", "ZjQyMTJiNTJhZGZmYjFjMQ=="}
	CheckContent string
	UserAgent    string
	UrlFile      string
	Method       string
	SKey         string
	HttpProxy    string
	PostContent  string
	AesMode      string
	Timeout      int
	Interval     int
	t            int
	NRemeberMe   string
	OutPutfile   string
)
