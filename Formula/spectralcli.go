package main

// SpectralCli - JSON/YAML linter and support OpenAPI v3.1/v3.0/v2.0, and AsyncAPI v2.x
// Homepage: https://stoplight.io/open-source/spectral

import (
	"fmt"
	
	"os/exec"
)

func installSpectralCli() {
	// Método 1: Descargar y extraer .tar.gz
	spectralcli_tar_url := "https://registry.npmjs.org/@stoplight/spectral-cli/-/spectral-cli-6.11.1.tgz"
	spectralcli_cmd_tar := exec.Command("curl", "-L", spectralcli_tar_url, "-o", "package.tar.gz")
	err := spectralcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spectralcli_zip_url := "https://registry.npmjs.org/@stoplight/spectral-cli/-/spectral-cli-6.11.1.tgz"
	spectralcli_cmd_zip := exec.Command("curl", "-L", spectralcli_zip_url, "-o", "package.zip")
	err = spectralcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spectralcli_bin_url := "https://registry.npmjs.org/@stoplight/spectral-cli/-/spectral-cli-6.11.1.tgz"
	spectralcli_cmd_bin := exec.Command("curl", "-L", spectralcli_bin_url, "-o", "binary.bin")
	err = spectralcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spectralcli_src_url := "https://registry.npmjs.org/@stoplight/spectral-cli/-/spectral-cli-6.11.1.tgz"
	spectralcli_cmd_src := exec.Command("curl", "-L", spectralcli_src_url, "-o", "source.tar.gz")
	err = spectralcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spectralcli_cmd_direct := exec.Command("./binary")
	err = spectralcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
