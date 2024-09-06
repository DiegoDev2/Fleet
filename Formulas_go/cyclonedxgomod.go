package main

// CyclonedxGomod - Creates CycloneDX Software Bill of Materials (SBOM) from Go modules
// Homepage: https://cyclonedx.org/

import (
	"fmt"
	
	"os/exec"
)

func installCyclonedxGomod() {
	// Método 1: Descargar y extraer .tar.gz
	cyclonedxgomod_tar_url := "https://github.com/CycloneDX/cyclonedx-gomod/archive/refs/tags/v1.7.0.tar.gz"
	cyclonedxgomod_cmd_tar := exec.Command("curl", "-L", cyclonedxgomod_tar_url, "-o", "package.tar.gz")
	err := cyclonedxgomod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cyclonedxgomod_zip_url := "https://github.com/CycloneDX/cyclonedx-gomod/archive/refs/tags/v1.7.0.zip"
	cyclonedxgomod_cmd_zip := exec.Command("curl", "-L", cyclonedxgomod_zip_url, "-o", "package.zip")
	err = cyclonedxgomod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cyclonedxgomod_bin_url := "https://github.com/CycloneDX/cyclonedx-gomod/archive/refs/tags/v1.7.0.bin"
	cyclonedxgomod_cmd_bin := exec.Command("curl", "-L", cyclonedxgomod_bin_url, "-o", "binary.bin")
	err = cyclonedxgomod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cyclonedxgomod_src_url := "https://github.com/CycloneDX/cyclonedx-gomod/archive/refs/tags/v1.7.0.src.tar.gz"
	cyclonedxgomod_cmd_src := exec.Command("curl", "-L", cyclonedxgomod_src_url, "-o", "source.tar.gz")
	err = cyclonedxgomod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cyclonedxgomod_cmd_direct := exec.Command("./binary")
	err = cyclonedxgomod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
