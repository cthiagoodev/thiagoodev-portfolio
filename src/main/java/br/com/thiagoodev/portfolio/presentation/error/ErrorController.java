package br.com.thiagoodev.portfolio.presentation.error;

import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.thymeleaf.exceptions.TemplateInputException;
import org.thymeleaf.exceptions.TemplateProcessingException;

@ControllerAdvice
public class ErrorController {
    @ExceptionHandler({TemplateInputException.class, TemplateProcessingException.class})
    public String handleThymeleadError(Exception exception, Model model) {
        System.err.println(exception.getMessage());
        model.addAttribute("errorMessage", "Ops! Houve um problema ao carregar parte do conteúdo visual.");
        return "error";
    }

    @ExceptionHandler(Exception.class)
    public String handleGeneralError(Exception exception, Model model) {
        System.err.println(exception.getMessage());
        model.addAttribute("errorMessage", "Erro interno no servidor.");
        return "error";
    }
}
