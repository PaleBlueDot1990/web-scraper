package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string 
		inputURL  string 
		inputBody string 
		expected  []string 
	}{
		{
			name:     "Get_URLs_From_HTML_TC1",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", 
					           "https://other.com/path/one"},
		},
		{
			name:      "Get_URLs_From_HTML_TC2",
			inputURL:  "http://wordpress.bhuwan.com/",
			inputBody: `
<html>
	<head>
		<title>Sample Links Page</title>
	</head>
	<body>
		<h1>Welcome to My Site</h1>
		<p>Here are some useful links:</p>
		<ul>
			<li><a href="/about">About Me</a></li>
			<li><a href="/articles/movies">Movie Articles</a></li>
		</ul>
		<ol>
			<li><a href="https://google.com/bhuwan">My Google Link</a></li>
			<li><a href="http://facebook.com/bhuwan">My Facebook Link</a></li>
		</ol>
	</body>
</html>
`,
			expected: []string{"http://wordpress.bhuwan.com/about", 
			                   "http://wordpress.bhuwan.com/articles/movies", 
			                   "https://google.com/bhuwan", 
							   "http://facebook.com/bhuwan"},
		},
		{
			name:       "Get_URLs_From_HTML_TC3",
			inputURL:   "https://deeplynestedlinks.com",
			inputBody:  `
<html>
	<head>
		<title>Deeply Nested Link</title>
	</head>
	<body>
		<div>
			<section>
				<article>
					<p>
					  	Here is a deeply nested link: 
					  	<span>
							<em>
						  		<a href="/example">Click Me!</a>
							</em>
					  	</span>
					</p>
				</article>
			</section>
		</div>
	</body>
</html>
`,
			expected: []string{"https://deeplynestedlinks.com/example"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return 
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}