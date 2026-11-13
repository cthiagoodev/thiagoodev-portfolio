package br.com.thiagoodev.api.presentation.exception;

import br.com.thiagoodev.api.domain.exceptions.DomainValidationException;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import org.junit.jupiter.api.Test;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException;

import java.util.Map;
import java.util.UUID;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertNotNull;

class GlobalExceptionHandlerTest {

    private final GlobalExceptionHandler globalExceptionHandler = new GlobalExceptionHandler();

    @Test
    void shouldReturnNotFound_whenResourceNotFoundExceptionIsHandled() {
        ResourceNotFoundException exception = new ResourceNotFoundException("resource not found");

        ResponseEntity<Map<String, String>> response = globalExceptionHandler.handleResourceNotFound(exception);

        assertEquals(HttpStatus.NOT_FOUND, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("resource not found", response.getBody().get("error"));
    }

    @Test
    void shouldReturnBadRequest_whenDomainValidationExceptionIsHandled() {
        DomainValidationException exception = new DomainValidationException("domain entity validation error: name is required");

        ResponseEntity<Map<String, String>> response = globalExceptionHandler.handleDomainValidation(exception);

        assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("domain entity validation error: name is required", response.getBody().get("error"));
    }

    @Test
    void shouldReturnBadRequest_whenMethodArgumentTypeMismatchExceptionIsHandled() {
        MethodArgumentTypeMismatchException exception = new MethodArgumentTypeMismatchException("invalid-uuid", UUID.class, "uuid", null, new IllegalArgumentException());

        ResponseEntity<Map<String, String>> response = globalExceptionHandler.handleMethodArgumentTypeMismatch(exception);

        assertEquals(HttpStatus.BAD_REQUEST, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("invalid parameter format", response.getBody().get("error"));
    }

    @Test
    void shouldReturnInternalServerError_whenGenericExceptionIsHandled() {
        RuntimeException exception = new RuntimeException("sensitive database error");

        ResponseEntity<Map<String, String>> response = globalExceptionHandler.handleGenericException(exception);

        assertEquals(HttpStatus.INTERNAL_SERVER_ERROR, response.getStatusCode());
        assertNotNull(response.getBody());
        assertEquals("internal server error", response.getBody().get("error"));
    }
}
