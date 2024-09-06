package main

// Range2cidr - Converts IP ranges to CIDRs
// Homepage: https://ipinfo.io

import (
	"fmt"
	
	"os/exec"
)

func installRange2cidr() {
	// Método 1: Descargar y extraer .tar.gz
	range2cidr_tar_url := "https://github.com/ipinfo/cli/archive/refs/tags/range2cidr-1.3.0.tar.gz"
	range2cidr_cmd_tar := exec.Command("curl", "-L", range2cidr_tar_url, "-o", "package.tar.gz")
	err := range2cidr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	range2cidr_zip_url := "https://github.com/ipinfo/cli/archive/refs/tags/range2cidr-1.3.0.zip"
	range2cidr_cmd_zip := exec.Command("curl", "-L", range2cidr_zip_url, "-o", "package.zip")
	err = range2cidr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	range2cidr_bin_url := "https://github.com/ipinfo/cli/archive/refs/tags/range2cidr-1.3.0.bin"
	range2cidr_cmd_bin := exec.Command("curl", "-L", range2cidr_bin_url, "-o", "binary.bin")
	err = range2cidr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	range2cidr_src_url := "https://github.com/ipinfo/cli/archive/refs/tags/range2cidr-1.3.0.src.tar.gz"
	range2cidr_cmd_src := exec.Command("curl", "-L", range2cidr_src_url, "-o", "source.tar.gz")
	err = range2cidr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	range2cidr_cmd_direct := exec.Command("./binary")
	err = range2cidr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
