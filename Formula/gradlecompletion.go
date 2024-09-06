package main

// GradleCompletion - Bash and Zsh completion for Gradle
// Homepage: https://gradle.org/

import (
	"fmt"
	
	"os/exec"
)

func installGradleCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	gradlecompletion_tar_url := "https://github.com/gradle/gradle-completion/archive/refs/tags/v1.4.1.tar.gz"
	gradlecompletion_cmd_tar := exec.Command("curl", "-L", gradlecompletion_tar_url, "-o", "package.tar.gz")
	err := gradlecompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gradlecompletion_zip_url := "https://github.com/gradle/gradle-completion/archive/refs/tags/v1.4.1.zip"
	gradlecompletion_cmd_zip := exec.Command("curl", "-L", gradlecompletion_zip_url, "-o", "package.zip")
	err = gradlecompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gradlecompletion_bin_url := "https://github.com/gradle/gradle-completion/archive/refs/tags/v1.4.1.bin"
	gradlecompletion_cmd_bin := exec.Command("curl", "-L", gradlecompletion_bin_url, "-o", "binary.bin")
	err = gradlecompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gradlecompletion_src_url := "https://github.com/gradle/gradle-completion/archive/refs/tags/v1.4.1.src.tar.gz"
	gradlecompletion_cmd_src := exec.Command("curl", "-L", gradlecompletion_src_url, "-o", "source.tar.gz")
	err = gradlecompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gradlecompletion_cmd_direct := exec.Command("./binary")
	err = gradlecompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
