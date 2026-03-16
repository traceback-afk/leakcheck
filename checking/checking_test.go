package checking

import "testing"

func TestSecrets(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
	}{
		// True positives
		{"AWS Access Key", `aws_access_key = "AKIA1234567890EXAMPLE"`, true},
		{"AWS Secret Key", `aws_secret_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"`, true},
		{"Stripe Live Key", `key = "sk_live_1234567890abcdef12345678"`, true},
		{"GitHub Token", `GITHUB_TOKEN = "ghp_abcdefghijklmnopqrstuvwx1234567890"`, true},
		{"Slack Token", `SLACK_TOKEN="xoxb-123456789012-1234567890123-abcdef123456"`, true},
		{"JWT Secret", `jwt_secret = "supersecretjwtkey"`, true},
		{"Database Password", `db_password = "Pa$$w0rd!"`, true},
		{"Private Key", `private_key = "-----BEGIN PRIVATE KEY-----ABC123"`, true},
		{"Complex Secret", `key = "3c1@2y@#s33g=pr@vwjxr+p+0=px3+fg@(*)_1+v@tx6v!e55="`, true},
		{"Google API Key", `google_api_key = "AIzaSyA-BCDEFGHIJKLMNOPQRSTUVWX"`, true},
		{"Heroku API Key", `HEROKU_API_KEY="0123456789abcdef0123456789abcdef"`, true},
		{"Long Random Secret", `secret = "a8f7c2d91b4e6f3a8d2e5c9f7b1a3d6f9e4c2b1a"`, true},
		{"Bearer Token", `Authorization = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.abc.def"`, true},
		{"Password Assignment", `password="admin12345!"`, true},
		{"Database Connection String", `DATABASE_URL="postgres://user:pass123@localhost:5432/db"`, true},
		{"Base64 Secret", `token="QWxhZGRpbjpPcGVuU2VzYW1lQWxhZGRpbjpPcGVuU2VzYW1l"`, true},
		{"RSA Private Key", `key="-----BEGIN RSA PRIVATE KEY-----MIIEpAIBAAKCAQEA"`, true},
		{"SSH Private Key", `ssh_key="-----BEGIN OPENSSH PRIVATE KEY-----b3BlbnNzaC1rZXktdjE="`, true},
		{"API Token Variable", `api_token = "1234567890abcdef1234567890abcdef"`, true},


		// False positives
		{"Random Text", `hello world`, false},
		{"Comment Line", `# this is a comment`, false},
		{"Config Variable", `max_retries = 5`, false},
		{"Username", `username = "admin"`, false},
		{"Empty String", `line = ""`, false},
		{"URL", `endpoint = "https://example.com/api"`, false},
		{"Number Only", `port = 8080`, false},
		{"Email", `email = "user@example.com"`, false},
		{"Log Message", `log.Println("starting server")`, false},
		{"UUID", `request_id = "550e8400-e29b-41d4-a716-446655440000"`, false},
		{"MD5 Hash", `checksum = "5d41402abc4b2a76b9719d911017c592"`, false},
		{"SHA256 Hash", `sha = "9e107d9d372bb6826bd81d3542a419d6d5f2c8c7c9c6b1a5a6d8e4c7b9c1f0d2"`, false},
		{"Hex Color", `color = "#ff5733"`, false},
		{"Simple Base64", `data = "SGVsbG8gd29ybGQ="`, false},
		{"Short Random String", `value = "a8f7c2d9"`, false},
		{"File Path", `path = "/usr/local/bin/app"`, false},
		{"Hostname", `host = "api.internal.local"`, false},
		{"Docker Image", `image = "nginx:latest"`, false},
		{"Date String", `date = "2026-03-16"`, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsSecret(tc.line)
			if got != tc.expected {
				t.Errorf("line: %q, got: %v, expected: %v", tc.line, got, tc.expected)
			}
		})
	}
}