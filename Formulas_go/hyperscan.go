package main

// Hyperscan - High-performance regular expression matching library
// Homepage: https://www.intel.com/content/www/us/en/developer/articles/technical/introduction-to-hyperscan.html

import (
	"fmt"
	
	"os/exec"
)

func installHyperscan() {
	// Método 1: Descargar y extraer .tar.gz
	hyperscan_tar_url := "https://github.com/intel/hyperscan/archive/refs/tags/v5.4.2.tar.gz"
	hyperscan_cmd_tar := exec.Command("curl", "-L", hyperscan_tar_url, "-o", "package.tar.gz")
	err := hyperscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hyperscan_zip_url := "https://github.com/intel/hyperscan/archive/refs/tags/v5.4.2.zip"
	hyperscan_cmd_zip := exec.Command("curl", "-L", hyperscan_zip_url, "-o", "package.zip")
	err = hyperscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hyperscan_bin_url := "https://github.com/intel/hyperscan/archive/refs/tags/v5.4.2.bin"
	hyperscan_cmd_bin := exec.Command("curl", "-L", hyperscan_bin_url, "-o", "binary.bin")
	err = hyperscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hyperscan_src_url := "https://github.com/intel/hyperscan/archive/refs/tags/v5.4.2.src.tar.gz"
	hyperscan_cmd_src := exec.Command("curl", "-L", hyperscan_src_url, "-o", "source.tar.gz")
	err = hyperscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hyperscan_cmd_direct := exec.Command("./binary")
	err = hyperscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ragel")
exec.Command("latte", "install", "ragel")
	fmt.Println("Instalando dependencia: pcre")
exec.Command("latte", "install", "pcre")
}
