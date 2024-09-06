package main

// Naabu - Fast port scanner
// Homepage: https://github.com/projectdiscovery/naabu

import (
	"fmt"
	
	"os/exec"
)

func installNaabu() {
	// Método 1: Descargar y extraer .tar.gz
	naabu_tar_url := "https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.3.1.tar.gz"
	naabu_cmd_tar := exec.Command("curl", "-L", naabu_tar_url, "-o", "package.tar.gz")
	err := naabu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	naabu_zip_url := "https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.3.1.zip"
	naabu_cmd_zip := exec.Command("curl", "-L", naabu_zip_url, "-o", "package.zip")
	err = naabu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	naabu_bin_url := "https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.3.1.bin"
	naabu_cmd_bin := exec.Command("curl", "-L", naabu_bin_url, "-o", "binary.bin")
	err = naabu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	naabu_src_url := "https://github.com/projectdiscovery/naabu/archive/refs/tags/v2.3.1.src.tar.gz"
	naabu_cmd_src := exec.Command("curl", "-L", naabu_src_url, "-o", "source.tar.gz")
	err = naabu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	naabu_cmd_direct := exec.Command("./binary")
	err = naabu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
