package br.com.thiagoodev.api.domain.utils;

public final class Sanitizer {

    private Sanitizer() {
    }

    public static String sanitize(String value) {
        return value == null ? "" : value.trim();
    }
}
