package main

// SonarScanner - Launcher to analyze a project with SonarQube
// Homepage: https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/

import (
	"fmt"
	
	"os/exec"
)

func installSonarScanner() {
	// Método 1: Descargar y extraer .tar.gz
	sonarscanner_tar_url := "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.1.0.4477.zip"
	sonarscanner_cmd_tar := exec.Command("curl", "-L", sonarscanner_tar_url, "-o", "package.tar.gz")
	err := sonarscanner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sonarscanner_zip_url := "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.1.0.4477.zip"
	sonarscanner_cmd_zip := exec.Command("curl", "-L", sonarscanner_zip_url, "-o", "package.zip")
	err = sonarscanner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sonarscanner_bin_url := "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.1.0.4477.zip"
	sonarscanner_cmd_bin := exec.Command("curl", "-L", sonarscanner_bin_url, "-o", "binary.bin")
	err = sonarscanner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sonarscanner_src_url := "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-6.1.0.4477.zip"
	sonarscanner_cmd_src := exec.Command("curl", "-L", sonarscanner_src_url, "-o", "source.tar.gz")
	err = sonarscanner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sonarscanner_cmd_direct := exec.Command("./binary")
	err = sonarscanner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
