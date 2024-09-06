package main

// GradleAT6 - Open-source build automation tool based on the Groovy and Kotlin DSL
// Homepage: https://www.gradle.org/

import (
	"fmt"
	
	"os/exec"
)

func installGradleAT6() {
	// Método 1: Descargar y extraer .tar.gz
	gradleat6_tar_url := "https://services.gradle.org/distributions/gradle-6.9.4-all.zip"
	gradleat6_cmd_tar := exec.Command("curl", "-L", gradleat6_tar_url, "-o", "package.tar.gz")
	err := gradleat6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gradleat6_zip_url := "https://services.gradle.org/distributions/gradle-6.9.4-all.zip"
	gradleat6_cmd_zip := exec.Command("curl", "-L", gradleat6_zip_url, "-o", "package.zip")
	err = gradleat6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gradleat6_bin_url := "https://services.gradle.org/distributions/gradle-6.9.4-all.zip"
	gradleat6_cmd_bin := exec.Command("curl", "-L", gradleat6_bin_url, "-o", "binary.bin")
	err = gradleat6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gradleat6_src_url := "https://services.gradle.org/distributions/gradle-6.9.4-all.zip"
	gradleat6_cmd_src := exec.Command("curl", "-L", gradleat6_src_url, "-o", "source.tar.gz")
	err = gradleat6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gradleat6_cmd_direct := exec.Command("./binary")
	err = gradleat6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
