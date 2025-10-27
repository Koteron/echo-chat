plugins {
	java
	id("org.springframework.boot") version "3.5.6"
	id("io.spring.dependency-management") version "1.1.7"
	id("org.liquibase.gradle") version "2.2.0"
	id("com.google.protobuf") version "0.9.4"
}

group = "echochat"
version = "0.0.1-SNAPSHOT"
description = "UserService"

java {
	toolchain {
		languageVersion = JavaLanguageVersion.of(17)
	}
}

configurations {
	compileOnly {
		extendsFrom(configurations.annotationProcessor.get())
	}
}

repositories {
	mavenCentral()
}

dependencies {
	implementation("org.springframework.boot:spring-boot-starter-data-jpa")
	implementation("org.springframework.boot:spring-boot-starter-security")
	implementation("org.springframework.boot:spring-boot-starter-web")
	implementation("org.liquibase:liquibase-core")
	implementation("org.mapstruct:mapstruct:1.5.5.Final")
	compileOnly("org.projectlombok:lombok")
	runtimeOnly("org.postgresql:postgresql")
	annotationProcessor("org.projectlombok:lombok")
	annotationProcessor("org.mapstruct:mapstruct-processor:1.5.5.Final")
	testImplementation("org.springframework.boot:spring-boot-starter-test")
	testImplementation("org.springframework.security:spring-security-test")
	testRuntimeOnly("org.junit.platform:junit-platform-launcher")
	implementation("io.grpc:grpc-netty-shaded:1.75.0")
	implementation("io.grpc:grpc-protobuf:1.65.0")
	implementation("io.grpc:grpc-stub:1.65.0")
	implementation("javax.annotation:javax.annotation-api:1.3.2")
	implementation("net.devh:grpc-server-spring-boot-starter:2.15.0.RELEASE")
}

protobuf {
	protoc {
		artifact = "com.google.protobuf:protoc:3.25.3"
	}
	plugins {
		create("grpc") {
			artifact = "io.grpc:protoc-gen-grpc-java:1.65.0"
		}
	}
	generateProtoTasks {
		all().forEach {
			it.plugins {
				create("grpc")
			}
		}
	}
}

sourceSets {
	main {
		proto {
			srcDir("../proto/user_service")
		}
	}
}

tasks.withType<Test> {
	useJUnitPlatform()
}
