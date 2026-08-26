package dev.ixlax.backend;

import io.github.cdimascio.dotenv.Dotenv;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class SocialNetworkBackendApplication {

	public static void main(String[] args) {

		/**
		 * Подключаем dotenv для считывания файла .env при запуске
		 */
		Dotenv dotenv = Dotenv.configure().ignoreIfMissing().load();
		dotenv.entries().forEach(e -> System.setProperty(e.getKey(), e.getValue()));

		SpringApplication.run(SocialNetworkBackendApplication.class, args);
	}

}
