package br.com.thiagoodev.api.domain.exceptions;

public class DomainValidationException extends RuntimeException {

    public DomainValidationException(String message) {
        super(message);
    }
}
