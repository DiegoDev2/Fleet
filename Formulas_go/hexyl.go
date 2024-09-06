package main

// Hexyl - Command-line hex viewer
// Homepage: https://github.com/sharkdp/hexyl

import (
	"fmt"
	
	"os/exec"
)

func installHexyl() {
	// Método 1: Descargar y extraer .tar.gz
	hexyl_tar_url := "https://github.com/sharkdp/hexyl/archive/refs/tags/v0.14.0.tar.gz"
	hexyl_cmd_tar := exec.Command("curl", "-L", hexyl_tar_url, "-o", "package.tar.gz")
	err := hexyl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hexyl_zip_url := "https://github.com/sharkdp/hexyl/archive/refs/tags/v0.14.0.zip"
	hexyl_cmd_zip := exec.Command("curl", "-L", hexyl_zip_url, "-o", "package.zip")
	err = hexyl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hexyl_bin_url := "https://github.com/sharkdp/hexyl/archive/refs/tags/v0.14.0.bin"
	hexyl_cmd_bin := exec.Command("curl", "-L", hexyl_bin_url, "-o", "binary.bin")
	err = hexyl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hexyl_src_url := "https://github.com/sharkdp/hexyl/archive/refs/tags/v0.14.0.src.tar.gz"
	hexyl_cmd_src := exec.Command("curl", "-L", hexyl_src_url, "-o", "source.tar.gz")
	err = hexyl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hexyl_cmd_direct := exec.Command("./binary")
	err = hexyl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
