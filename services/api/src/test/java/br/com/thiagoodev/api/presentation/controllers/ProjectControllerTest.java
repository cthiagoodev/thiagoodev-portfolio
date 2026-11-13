package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Project;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.usecases.GetAllProjectsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetProjectByUuidUseCase;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.webmvc.test.autoconfigure.WebMvcTest;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;
import java.util.UUID;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(ProjectController.class)
class ProjectControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private GetAllProjectsUseCase getAllProjectsUseCase;

    @MockitoBean
    private GetProjectByUuidUseCase getProjectByUuidUseCase;

    @Test
    void shouldReturnOkWithProjects_whenProjectsExist() throws Exception {
        Project project = new Project(UUID.randomUUID(), 1L, "Portfolio", "Desc", "https://app.dev", "https://github.com/repo", List.of("Java"), null, null, null, null);
        when(getAllProjectsUseCase.execute()).thenReturn(List.of(project));

        mockMvc.perform(get("/projects"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].name").value("Portfolio"));
    }

    @Test
    void shouldReturnNotFound_whenNoProjectsExist() throws Exception {
        when(getAllProjectsUseCase.execute()).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/projects"))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnOkWithProject_whenProjectExistsByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        Project project = new Project(uuid, 1L, "Portfolio", "Desc", "", "", null, null, null, null, null);
        when(getProjectByUuidUseCase.execute(uuid)).thenReturn(project);

        mockMvc.perform(get("/projects/" + uuid))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.name").value("Portfolio"));
    }

    @Test
    void shouldReturnNotFound_whenProjectDoesNotExistByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        when(getProjectByUuidUseCase.execute(uuid)).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/projects/" + uuid))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnInternalServerError_whenUnexpectedExceptionOccurs() throws Exception {
        when(getAllProjectsUseCase.execute()).thenThrow(new RuntimeException("database connection failed"));

        mockMvc.perform(get("/projects"))
                .andExpect(status().isInternalServerError())
                .andExpect(jsonPath("$.error").value("internal server error"));
    }
}
