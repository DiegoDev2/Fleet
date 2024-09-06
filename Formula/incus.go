package main

// Incus - CLI client for interacting with Incus
// Homepage: https://linuxcontainers.org/incus

import (
	"fmt"
	
	"os/exec"
)

func installIncus() {
	// Método 1: Descargar y extraer .tar.gz
	incus_tar_url := "https://linuxcontainers.org/downloads/incus/incus-6.5.tar.xz"
	incus_cmd_tar := exec.Command("curl", "-L", incus_tar_url, "-o", "package.tar.gz")
	err := incus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	incus_zip_url := "https://linuxcontainers.org/downloads/incus/incus-6.5.tar.xz"
	incus_cmd_zip := exec.Command("curl", "-L", incus_zip_url, "-o", "package.zip")
	err = incus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	incus_bin_url := "https://linuxcontainers.org/downloads/incus/incus-6.5.tar.xz"
	incus_cmd_bin := exec.Command("curl", "-L", incus_bin_url, "-o", "binary.bin")
	err = incus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	incus_src_url := "https://linuxcontainers.org/downloads/incus/incus-6.5.tar.xz"
	incus_cmd_src := exec.Command("curl", "-L", incus_src_url, "-o", "source.tar.gz")
	err = incus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	incus_cmd_direct := exec.Command("./binary")
	err = incus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
