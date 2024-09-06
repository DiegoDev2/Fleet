package main

// Trivy - Vulnerability scanner for container images, file systems, and Git repos
// Homepage: https://aquasecurity.github.io/trivy/

import (
	"fmt"
	
	"os/exec"
)

func installTrivy() {
	// Método 1: Descargar y extraer .tar.gz
	trivy_tar_url := "https://github.com/aquasecurity/trivy/archive/refs/tags/v0.55.0.tar.gz"
	trivy_cmd_tar := exec.Command("curl", "-L", trivy_tar_url, "-o", "package.tar.gz")
	err := trivy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trivy_zip_url := "https://github.com/aquasecurity/trivy/archive/refs/tags/v0.55.0.zip"
	trivy_cmd_zip := exec.Command("curl", "-L", trivy_zip_url, "-o", "package.zip")
	err = trivy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trivy_bin_url := "https://github.com/aquasecurity/trivy/archive/refs/tags/v0.55.0.bin"
	trivy_cmd_bin := exec.Command("curl", "-L", trivy_bin_url, "-o", "binary.bin")
	err = trivy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trivy_src_url := "https://github.com/aquasecurity/trivy/archive/refs/tags/v0.55.0.src.tar.gz"
	trivy_cmd_src := exec.Command("curl", "-L", trivy_src_url, "-o", "source.tar.gz")
	err = trivy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trivy_cmd_direct := exec.Command("./binary")
	err = trivy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
