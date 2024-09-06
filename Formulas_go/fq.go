package main

// Fq - Brokered message queue optimized for performance
// Homepage: https://github.com/circonus-labs/fq

import (
	"fmt"
	
	"os/exec"
)

func installFq() {
	// Método 1: Descargar y extraer .tar.gz
	fq_tar_url := "https://github.com/circonus-labs/fq/archive/refs/tags/v0.13.11.tar.gz"
	fq_cmd_tar := exec.Command("curl", "-L", fq_tar_url, "-o", "package.tar.gz")
	err := fq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fq_zip_url := "https://github.com/circonus-labs/fq/archive/refs/tags/v0.13.11.zip"
	fq_cmd_zip := exec.Command("curl", "-L", fq_zip_url, "-o", "package.zip")
	err = fq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fq_bin_url := "https://github.com/circonus-labs/fq/archive/refs/tags/v0.13.11.bin"
	fq_cmd_bin := exec.Command("curl", "-L", fq_bin_url, "-o", "binary.bin")
	err = fq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fq_src_url := "https://github.com/circonus-labs/fq/archive/refs/tags/v0.13.11.src.tar.gz"
	fq_cmd_src := exec.Command("curl", "-L", fq_src_url, "-o", "source.tar.gz")
	err = fq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fq_cmd_direct := exec.Command("./binary")
	err = fq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: concurrencykit")
exec.Command("latte", "install", "concurrencykit")
	fmt.Println("Instalando dependencia: jlog")
exec.Command("latte", "install", "jlog")
	fmt.Println("Instalando dependencia: bind")
exec.Command("latte", "install", "bind")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
