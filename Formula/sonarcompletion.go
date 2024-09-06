package main

// SonarCompletion - Bash completion for Sonar
// Homepage: https://github.com/a1dutch/sonarqube-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installSonarCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	sonarcompletion_tar_url := "https://github.com/a1dutch/sonarqube-bash-completion/archive/refs/tags/1.1.tar.gz"
	sonarcompletion_cmd_tar := exec.Command("curl", "-L", sonarcompletion_tar_url, "-o", "package.tar.gz")
	err := sonarcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sonarcompletion_zip_url := "https://github.com/a1dutch/sonarqube-bash-completion/archive/refs/tags/1.1.zip"
	sonarcompletion_cmd_zip := exec.Command("curl", "-L", sonarcompletion_zip_url, "-o", "package.zip")
	err = sonarcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sonarcompletion_bin_url := "https://github.com/a1dutch/sonarqube-bash-completion/archive/refs/tags/1.1.bin"
	sonarcompletion_cmd_bin := exec.Command("curl", "-L", sonarcompletion_bin_url, "-o", "binary.bin")
	err = sonarcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sonarcompletion_src_url := "https://github.com/a1dutch/sonarqube-bash-completion/archive/refs/tags/1.1.src.tar.gz"
	sonarcompletion_cmd_src := exec.Command("curl", "-L", sonarcompletion_src_url, "-o", "source.tar.gz")
	err = sonarcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sonarcompletion_cmd_direct := exec.Command("./binary")
	err = sonarcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
