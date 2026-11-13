package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.WorkExperience;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.usecases.GetAllWorkExperiencesUseCase;
import br.com.thiagoodev.api.domain.usecases.GetWorkExperienceByUuidUseCase;
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

@WebMvcTest(WorkExperienceController.class)
class WorkExperienceControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private GetAllWorkExperiencesUseCase getAllWorkExperiencesUseCase;

    @MockitoBean
    private GetWorkExperienceByUuidUseCase getWorkExperienceByUuidUseCase;

    @Test
    void shouldReturnOkWithWorkExperiences_whenExperiencesExist() throws Exception {
        WorkExperience experience = new WorkExperience(UUID.randomUUID(), "Engineer", "Backend", "Company", LocalDate.of(2022, 1, 1), null, null, null);
        when(getAllWorkExperiencesUseCase.execute()).thenReturn(List.of(experience));

        mockMvc.perform(get("/work-experiences"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].role").value("Engineer"))
                .andExpect(jsonPath("$[0].company").value("Company"))
                .andExpect(jsonPath("$[0].start_date").value("2022-01-01"))
                .andExpect(jsonPath("$[0].created_at").exists());
    }

    @Test
    void shouldReturnNotFound_whenNoWorkExperiencesExist() throws Exception {
        when(getAllWorkExperiencesUseCase.execute()).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/work-experiences"))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnOkWithWorkExperience_whenExperienceExistsByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        WorkExperience experience = new WorkExperience(uuid, "Engineer", "Backend", "Company", LocalDate.of(2022, 1, 1), null, null, null);
        when(getWorkExperienceByUuidUseCase.execute(uuid)).thenReturn(experience);

        mockMvc.perform(get("/work-experiences/" + uuid))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.role").value("Engineer"));
    }

    @Test
    void shouldReturnNotFound_whenWorkExperienceDoesNotExistByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        when(getWorkExperienceByUuidUseCase.execute(uuid)).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/work-experiences/" + uuid))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnInternalServerError_whenUnexpectedExceptionOccurs() throws Exception {
        when(getAllWorkExperiencesUseCase.execute()).thenThrow(new RuntimeException("database connection failed"));

        mockMvc.perform(get("/work-experiences"))
                .andExpect(status().isInternalServerError())
                .andExpect(jsonPath("$.error").value("internal server error"));
    }
}
