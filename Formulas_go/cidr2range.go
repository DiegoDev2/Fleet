package main

// Cidr2range - Converts CIDRs to IP ranges
// Homepage: https://ipinfo.io

import (
	"fmt"
	
	"os/exec"
)

func installCidr2range() {
	// Método 1: Descargar y extraer .tar.gz
	cidr2range_tar_url := "https://github.com/ipinfo/cli/archive/refs/tags/cidr2range-1.2.0.tar.gz"
	cidr2range_cmd_tar := exec.Command("curl", "-L", cidr2range_tar_url, "-o", "package.tar.gz")
	err := cidr2range_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cidr2range_zip_url := "https://github.com/ipinfo/cli/archive/refs/tags/cidr2range-1.2.0.zip"
	cidr2range_cmd_zip := exec.Command("curl", "-L", cidr2range_zip_url, "-o", "package.zip")
	err = cidr2range_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cidr2range_bin_url := "https://github.com/ipinfo/cli/archive/refs/tags/cidr2range-1.2.0.bin"
	cidr2range_cmd_bin := exec.Command("curl", "-L", cidr2range_bin_url, "-o", "binary.bin")
	err = cidr2range_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cidr2range_src_url := "https://github.com/ipinfo/cli/archive/refs/tags/cidr2range-1.2.0.src.tar.gz"
	cidr2range_cmd_src := exec.Command("curl", "-L", cidr2range_src_url, "-o", "source.tar.gz")
	err = cidr2range_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cidr2range_cmd_direct := exec.Command("./binary")
	err = cidr2range_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
