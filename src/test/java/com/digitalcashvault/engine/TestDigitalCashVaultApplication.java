package com.digitalcashvault.engine;

import org.springframework.boot.SpringApplication;

public class TestDigitalCashVaultApplication {

	public static void main(String[] args) {
		SpringApplication.from(DigitalCashVaultApplication::main).with(TestcontainersConfiguration.class).run(args);
	}

}
