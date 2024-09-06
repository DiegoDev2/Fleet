package main

// Tailwindcss - Utility-first CSS framework
// Homepage: https://tailwindcss.com

import (
	"fmt"
	
	"os/exec"
)

func installTailwindcss() {
	// Método 1: Descargar y extraer .tar.gz
	tailwindcss_tar_url := "https://github.com/tailwindlabs/tailwindcss/archive/refs/tags/v3.4.10.tar.gz"
	tailwindcss_cmd_tar := exec.Command("curl", "-L", tailwindcss_tar_url, "-o", "package.tar.gz")
	err := tailwindcss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tailwindcss_zip_url := "https://github.com/tailwindlabs/tailwindcss/archive/refs/tags/v3.4.10.zip"
	tailwindcss_cmd_zip := exec.Command("curl", "-L", tailwindcss_zip_url, "-o", "package.zip")
	err = tailwindcss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tailwindcss_bin_url := "https://github.com/tailwindlabs/tailwindcss/archive/refs/tags/v3.4.10.bin"
	tailwindcss_cmd_bin := exec.Command("curl", "-L", tailwindcss_bin_url, "-o", "binary.bin")
	err = tailwindcss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tailwindcss_src_url := "https://github.com/tailwindlabs/tailwindcss/archive/refs/tags/v3.4.10.src.tar.gz"
	tailwindcss_cmd_src := exec.Command("curl", "-L", tailwindcss_src_url, "-o", "source.tar.gz")
	err = tailwindcss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tailwindcss_cmd_direct := exec.Command("./binary")
	err = tailwindcss_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
