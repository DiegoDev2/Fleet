package main

// Nuclei - HTTP/DNS scanner configurable via YAML templates
// Homepage: https://nuclei.projectdiscovery.io/

import (
	"fmt"
	
	"os/exec"
)

func installNuclei() {
	// Método 1: Descargar y extraer .tar.gz
	nuclei_tar_url := "https://github.com/projectdiscovery/nuclei/archive/refs/tags/v3.3.2.tar.gz"
	nuclei_cmd_tar := exec.Command("curl", "-L", nuclei_tar_url, "-o", "package.tar.gz")
	err := nuclei_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nuclei_zip_url := "https://github.com/projectdiscovery/nuclei/archive/refs/tags/v3.3.2.zip"
	nuclei_cmd_zip := exec.Command("curl", "-L", nuclei_zip_url, "-o", "package.zip")
	err = nuclei_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nuclei_bin_url := "https://github.com/projectdiscovery/nuclei/archive/refs/tags/v3.3.2.bin"
	nuclei_cmd_bin := exec.Command("curl", "-L", nuclei_bin_url, "-o", "binary.bin")
	err = nuclei_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nuclei_src_url := "https://github.com/projectdiscovery/nuclei/archive/refs/tags/v3.3.2.src.tar.gz"
	nuclei_cmd_src := exec.Command("curl", "-L", nuclei_src_url, "-o", "source.tar.gz")
	err = nuclei_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nuclei_cmd_direct := exec.Command("./binary")
	err = nuclei_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
