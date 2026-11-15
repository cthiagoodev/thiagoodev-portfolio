package br.com.thiagoodev.api.presentation.scheduler;

import br.com.thiagoodev.api.domain.usecases.SyncProjectsUseCase;
import lombok.extern.slf4j.Slf4j;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class SyncProjectsScheduler {

    private final SyncProjectsUseCase syncProjectsUseCase;

    public SyncProjectsScheduler(SyncProjectsUseCase syncProjectsUseCase) {
        this.syncProjectsUseCase = syncProjectsUseCase;
    }

//    @Scheduled(cron = "0 0 3 * * *")
//    public void execute() {
//        try {
//            syncProjectsUseCase.execute();
//        } catch (Exception exception) {
//            log.error("Failed to sync portfolio projects from GitHub", exception);
//        }
//    }
}
