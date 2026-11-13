package br.com.thiagoodev.api.presentation.controllers;

import br.com.thiagoodev.api.domain.entities.SocialContact;
import br.com.thiagoodev.api.domain.exceptions.ResourceNotFoundException;
import br.com.thiagoodev.api.domain.usecases.GetAllSocialContactsUseCase;
import br.com.thiagoodev.api.domain.usecases.GetSocialContactByUuidUseCase;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.webmvc.test.autoconfigure.WebMvcTest;
import org.springframework.http.MediaType;
import org.springframework.test.context.bean.override.mockito.MockitoBean;
import org.springframework.test.web.servlet.MockMvc;

import java.util.List;
import java.util.UUID;

import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.content;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.jsonPath;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

@WebMvcTest(SocialContactController.class)
class SocialContactControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @MockitoBean
    private GetAllSocialContactsUseCase getAllSocialContactsUseCase;

    @MockitoBean
    private GetSocialContactByUuidUseCase getSocialContactByUuidUseCase;

    @Test
    void shouldReturnOkWithSocialContacts_whenContactsExist() throws Exception {
        SocialContact firstContact = new SocialContact(UUID.randomUUID(), "github", "https://github.com/cthiagoodev", "/images/icons/github.svg", null, null);
        when(getAllSocialContactsUseCase.execute()).thenReturn(List.of(firstContact));

        mockMvc.perform(get("/social-contacts"))
                .andExpect(status().isOk())
                .andExpect(content().contentTypeCompatibleWith(MediaType.APPLICATION_JSON))
                .andExpect(jsonPath("$[0].label").value("github"))
                .andExpect(jsonPath("$[0].url").value("https://github.com/cthiagoodev"));
    }

    @Test
    void shouldReturnNotFound_whenNoSocialContactsExist() throws Exception {
        when(getAllSocialContactsUseCase.execute()).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/social-contacts"))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnOkWithSocialContact_whenContactExistsByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        SocialContact contact = new SocialContact(uuid, "github", "https://github.com/cthiagoodev", "", null, null);
        when(getSocialContactByUuidUseCase.execute(uuid)).thenReturn(contact);

        mockMvc.perform(get("/social-contacts/" + uuid))
                .andExpect(status().isOk())
                .andExpect(jsonPath("$.label").value("github"));
    }

    @Test
    void shouldReturnNotFound_whenSocialContactDoesNotExistByUuid() throws Exception {
        UUID uuid = UUID.randomUUID();
        when(getSocialContactByUuidUseCase.execute(uuid)).thenThrow(new ResourceNotFoundException("resource not found"));

        mockMvc.perform(get("/social-contacts/" + uuid))
                .andExpect(status().isNotFound())
                .andExpect(jsonPath("$.error").value("resource not found"));
    }

    @Test
    void shouldReturnInternalServerError_whenUnexpectedExceptionOccurs() throws Exception {
        when(getAllSocialContactsUseCase.execute()).thenThrow(new RuntimeException("database connection failed"));

        mockMvc.perform(get("/social-contacts"))
                .andExpect(status().isInternalServerError())
                .andExpect(jsonPath("$.error").value("internal server error"));
    }
}
