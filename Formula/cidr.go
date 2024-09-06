package main

// Cidr - CLI to perform various actions on CIDR ranges
// Homepage: https://github.com/bschaatsbergen/cidr

import (
	"fmt"
	
	"os/exec"
)

func installCidr() {
	// Método 1: Descargar y extraer .tar.gz
	cidr_tar_url := "https://github.com/bschaatsbergen/cidr/archive/refs/tags/v2.2.0.tar.gz"
	cidr_cmd_tar := exec.Command("curl", "-L", cidr_tar_url, "-o", "package.tar.gz")
	err := cidr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cidr_zip_url := "https://github.com/bschaatsbergen/cidr/archive/refs/tags/v2.2.0.zip"
	cidr_cmd_zip := exec.Command("curl", "-L", cidr_zip_url, "-o", "package.zip")
	err = cidr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cidr_bin_url := "https://github.com/bschaatsbergen/cidr/archive/refs/tags/v2.2.0.bin"
	cidr_cmd_bin := exec.Command("curl", "-L", cidr_bin_url, "-o", "binary.bin")
	err = cidr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cidr_src_url := "https://github.com/bschaatsbergen/cidr/archive/refs/tags/v2.2.0.src.tar.gz"
	cidr_cmd_src := exec.Command("curl", "-L", cidr_src_url, "-o", "source.tar.gz")
	err = cidr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cidr_cmd_direct := exec.Command("./binary")
	err = cidr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
