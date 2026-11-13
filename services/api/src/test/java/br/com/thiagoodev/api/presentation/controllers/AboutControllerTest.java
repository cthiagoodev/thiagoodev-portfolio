package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.About;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.usecases.GetAboutUseCase;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.webmvc.test.autoconfigure.WebMvcTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import java.util.UUID;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(AboutController.class)
class AboutControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private GetAboutUseCase getAboutUseCase;

    @Test
    void shouldReturnOkWithAbout_whenAboutExists() throws Exception {
        About expectedAbout = new About(
                UUID.randomUUID(),
                "Thiago Sousa",
                "+5511999999999",
                "Software Engineer",
                "Bio completa",
                null,
                null
        );
        when(getAboutUseCase.execute()).thenReturn(expectedAbout);

        mockMvc.perform(get("/about"))
                .andExpect(status().isOk())
                .andExpect(content().contentTypeCompatibleWith(MediaType.APPLICATION_JSON))
                .andExpect(jsonPath("$.name").value("Thiago Sousa"));
    }

    @Test
    void shouldReturnNotFound_whenAboutDoesNotExist() throws Exception {
        when(getAboutUseCase.execute()).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/about"))
                .andExpect(status().isNotFound())
                .andExpect(content().contentTypeCompatibleWith(MediaType.APPLICATION_JSON))
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnInternalServerError_whenUnexpectedExceptionOccurs() throws Exception {
        when(getAboutUseCase.execute()).thenThrow(new RuntimeException("database connection failed"));

        mockMvc.perform(get("/about"))
                .andExpect(status().isInternalServerError())
                .andExpect(content().contentTypeCompatibleWith(MediaType.APPLICATION_JSON))
                .andExpect(jsonPath("$.error").value("internal server error"));
    }
}
