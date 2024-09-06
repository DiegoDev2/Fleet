package main

// SonarqubeLts - Manage code quality
// Homepage: https://www.sonarqube.org/

import (
	"fmt"
	
	"os/exec"
)

func installSonarqubeLts() {
	// Método 1: Descargar y extraer .tar.gz
	sonarqubelts_tar_url := "https://binaries.sonarsource.com/Distribution/sonarqube/sonarqube-9.9.6.92038.zip"
	sonarqubelts_cmd_tar := exec.Command("curl", "-L", sonarqubelts_tar_url, "-o", "package.tar.gz")
	err := sonarqubelts_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sonarqubelts_zip_url := "https://binaries.sonarsource.com/Distribution/sonarqube/sonarqube-9.9.6.92038.zip"
	sonarqubelts_cmd_zip := exec.Command("curl", "-L", sonarqubelts_zip_url, "-o", "package.zip")
	err = sonarqubelts_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sonarqubelts_bin_url := "https://binaries.sonarsource.com/Distribution/sonarqube/sonarqube-9.9.6.92038.zip"
	sonarqubelts_cmd_bin := exec.Command("curl", "-L", sonarqubelts_bin_url, "-o", "binary.bin")
	err = sonarqubelts_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sonarqubelts_src_url := "https://binaries.sonarsource.com/Distribution/sonarqube/sonarqube-9.9.6.92038.zip"
	sonarqubelts_cmd_src := exec.Command("curl", "-L", sonarqubelts_src_url, "-o", "source.tar.gz")
	err = sonarqubelts_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sonarqubelts_cmd_direct := exec.Command("./binary")
	err = sonarqubelts_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}
