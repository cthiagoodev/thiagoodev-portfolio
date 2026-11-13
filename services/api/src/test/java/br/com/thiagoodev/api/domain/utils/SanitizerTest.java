package br.com.thiagoodev.api.domain.utils;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class SanitizerTest {

    @Test
    void sanitize_returnsTrimmedString_whenInputHasSurroundingSpaces() {
        String input = "   Thiago Sousa   ";

        String result = Sanitizer.sanitize(input);

        assertEquals("Thiago Sousa", result);
    }

    @Test
    void sanitize_returnsEmptyString_whenInputIsNull() {
        String result = Sanitizer.sanitize(null);

        assertEquals("", result);
    }

    @Test
    void sanitize_returnsEmptyString_whenInputIsOnlyWhitespace() {
        String input = "     ";

        String result = Sanitizer.sanitize(input);

        assertEquals("", result);
    }
}
