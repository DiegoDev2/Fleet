package main

// Gradle - Open-source build automation tool based on the Groovy and Kotlin DSL
// Homepage: https://www.gradle.org/

import (
	"fmt"
	
	"os/exec"
)

func installGradle() {
	// Método 1: Descargar y extraer .tar.gz
	gradle_tar_url := "https://services.gradle.org/distributions/gradle-8.10-all.zip"
	gradle_cmd_tar := exec.Command("curl", "-L", gradle_tar_url, "-o", "package.tar.gz")
	err := gradle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gradle_zip_url := "https://services.gradle.org/distributions/gradle-8.10-all.zip"
	gradle_cmd_zip := exec.Command("curl", "-L", gradle_zip_url, "-o", "package.zip")
	err = gradle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gradle_bin_url := "https://services.gradle.org/distributions/gradle-8.10-all.zip"
	gradle_cmd_bin := exec.Command("curl", "-L", gradle_bin_url, "-o", "binary.bin")
	err = gradle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gradle_src_url := "https://services.gradle.org/distributions/gradle-8.10-all.zip"
	gradle_cmd_src := exec.Command("curl", "-L", gradle_src_url, "-o", "source.tar.gz")
	err = gradle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gradle_cmd_direct := exec.Command("./binary")
	err = gradle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
