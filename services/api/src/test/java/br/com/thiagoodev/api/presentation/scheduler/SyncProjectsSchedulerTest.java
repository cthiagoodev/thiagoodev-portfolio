package br.com.thiagoodev.api.presentation.scheduler;

import br.com.thiagoodev.api.domain.usecases.SyncProjectsUseCase;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.List;

import static org.junit.jupiter.api.Assertions.assertDoesNotThrow;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@ExtendWith(MockitoExtension.class)
class SyncProjectsSchedulerTest {

    @Mock
    private SyncProjectsUseCase syncProjectsUseCase;

    @InjectMocks
    private SyncProjectsScheduler syncProjectsScheduler;

    @Test
    void shouldInvokeSyncProjectsUseCase_whenSchedulerExecutes() {
        when(syncProjectsUseCase.execute()).thenReturn(List.of());

        syncProjectsScheduler.execute();

        verify(syncProjectsUseCase, times(1)).execute();
    }

    @Test
    void shouldNotPropagateException_whenSyncProjectsUseCaseFails() {
        when(syncProjectsUseCase.execute()).thenThrow(new RuntimeException("github timeout"));

        assertDoesNotThrow(() -> syncProjectsScheduler.execute());
        verify(syncProjectsUseCase, times(1)).execute();
    }
}
