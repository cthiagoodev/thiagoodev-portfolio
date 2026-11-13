package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.Skill;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.usecases.GetAllSkillsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetSkillByUuidUseCase;
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

@WebMvcTest(SkillController.class)
class SkillControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private GetAllSkillsUseCase getAllSkillsUseCase;

    @MockitoBean
    private GetSkillByUuidUseCase getSkillByUuidUseCase;

    @Test
    void shouldReturnOkWithSkills_whenSkillsExist() throws Exception {
        Skill skill = new Skill(UUID.randomUUID(), "Java", "https://openjdk.org", "/icons/java.svg", null, null);
        when(getAllSkillsUseCase.execute()).thenReturn(List.of(skill));

        mockMvc.perform(get("/skills"))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$[0].label").value("Java"));
    }

    @Test
    void shouldReturnNotFound_whenNoSkillsExist() throws Exception {
        when(getAllSkillsUseCase.execute()).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/skills"))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnOkWithSkill_whenSkillExistsByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        Skill skill = new Skill(uuid, "Java", "https://openjdk.org", "", null, null);
        when(getSkillByUuidUseCase.execute(uuid)).thenReturn(skill);

        mockMvc.perform(get("/skills/" + uuid))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.label").value("Java"));
    }

    @Test
    void shouldReturnNotFound_whenSkillDoesNotExistByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        when(getSkillByUuidUseCase.execute(uuid)).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/skills/" + uuid))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnInternalServerError_whenUnexpectedExceptionOccurs() throws Exception {
        when(getAllSkillsUseCase.execute()).thenThrow(new RuntimeException("database connection failed"));

        mockMvc.perform(get("/skills"))
                .andExpect(status().isInternalServerError())
                .andExpect(jsonPath("$.error").value("internal server error"));
    }
}
