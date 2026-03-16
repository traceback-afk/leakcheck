package parsing

import "testing"

func TestParseLineDoubleQuote(t *testing.T){
	line := "key = \"password\""

	expectedKey := "key"
	expectedValue := "password"

	returnedKey, returnedValue := ParseLine(line)

	if returnedKey != expectedKey || returnedValue != expectedValue {
		t.Errorf("retKey: %s, expKey: %s - retValue: %s, expValue: %s", returnedKey, expectedKey, returnedValue, expectedValue)
	}
}

func TestParseLineSingleQuote(t *testing.T){
	line := "key = 'password'"

	expectedKey := "key"
	expectedValue := "password"

	returnedKey, returnedValue := ParseLine(line)

	if returnedKey != expectedKey || returnedValue != expectedValue {
		t.Errorf("retKey: %s, expKey: %s - retValue: %s, expValue: %s", returnedKey, expectedKey, returnedValue, expectedValue)
	}
}

func TestParseLinePlain(t *testing.T){
	line := "key = password"

	expectedKey := "key"
	expectedValue := "password"

	returnedKey, returnedValue := ParseLine(line)

	if returnedKey != expectedKey || returnedValue != expectedValue {
		t.Errorf("retKey: %s, expKey: %s - retValue: %s, expValue: %s", returnedKey, expectedKey, returnedValue, expectedValue)
	}
}

