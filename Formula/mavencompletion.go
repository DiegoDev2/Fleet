package main

// MavenCompletion - Bash completion for Maven
// Homepage: https://github.com/juven/maven-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installMavenCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	mavencompletion_tar_url := "https://github.com/juven/maven-bash-completion/archive/refs/tags/20200420.tar.gz"
	mavencompletion_cmd_tar := exec.Command("curl", "-L", mavencompletion_tar_url, "-o", "package.tar.gz")
	err := mavencompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mavencompletion_zip_url := "https://github.com/juven/maven-bash-completion/archive/refs/tags/20200420.zip"
	mavencompletion_cmd_zip := exec.Command("curl", "-L", mavencompletion_zip_url, "-o", "package.zip")
	err = mavencompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mavencompletion_bin_url := "https://github.com/juven/maven-bash-completion/archive/refs/tags/20200420.bin"
	mavencompletion_cmd_bin := exec.Command("curl", "-L", mavencompletion_bin_url, "-o", "binary.bin")
	err = mavencompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mavencompletion_src_url := "https://github.com/juven/maven-bash-completion/archive/refs/tags/20200420.src.tar.gz"
	mavencompletion_cmd_src := exec.Command("curl", "-L", mavencompletion_src_url, "-o", "source.tar.gz")
	err = mavencompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mavencompletion_cmd_direct := exec.Command("./binary")
	err = mavencompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
