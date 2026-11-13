package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Education;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.usecases.GetAllEducationsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetEducationByUuidUseCase;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.webmvc.test.autoconfigure.WebMvcTest;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import java.time.LocalDate;
import java.util.List;
import java.util.UUID;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(EducationController.class)
class EducationControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private GetAllEducationsUseCase getAllEducationsUseCase;

    @MockitoBean
    private GetEducationByUuidUseCase getEducationByUuidUseCase;

    @Test
    void shouldReturnOkWithEducations_whenEducationsExist() throws Exception {
        Education education = new Education(UUID.randomUUID(), "Ciência da Computação", "Bacharelado", "Universidade", LocalDate.of(2018, 2, 1), null, null, null);
        when(getAllEducationsUseCase.execute()).thenReturn(List.of(education));

        mockMvc.perform(get("/educations"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].course").value("Ciência da Computação"));
    }

    @Test
    void shouldReturnNotFound_whenNoEducationsExist() throws Exception {
        when(getAllEducationsUseCase.execute()).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/educations"))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnOkWithEducation_whenEducationExistsByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        Education education = new Education(uuid, "Ciência da Computação", "Bacharelado", "Universidade", LocalDate.of(2018, 2, 1), null, null, null);
        when(getEducationByUuidUseCase.execute(uuid)).thenReturn(education);

        mockMvc.perform(get("/educations/" + uuid))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.course").value("Ciência da Computação"));
    }

    @Test
    void shouldReturnNotFound_whenEducationDoesNotExistByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        when(getEducationByUuidUseCase.execute(uuid)).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/educations/" + uuid))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnInternalServerError_whenUnexpectedExceptionOccurs() throws Exception {
        when(getAllEducationsUseCase.execute()).thenThrow(new RuntimeException("database connection failed"));

        mockMvc.perform(get("/educations"))
                .andExpect(status().isInternalServerError())
                .andExpect(jsonPath("$.error").value("internal server error"));
    }
}
