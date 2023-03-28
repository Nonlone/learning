package per.nonlone.h2d;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.apache.dubbo.config.spring.context.annotation.EnableDubbo;

@SpringBootApplication
@EnableDubbo
public class H2dApplication {

	public static void main(String[] args) {
		SpringApplication.run(H2dApplication.class, args);
	}

}
