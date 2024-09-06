package main

// Diskus - Minimal, fast alternative to 'du -sh'
// Homepage: https://github.com/sharkdp/diskus

import (
	"fmt"
	
	"os/exec"
)

func installDiskus() {
	// Método 1: Descargar y extraer .tar.gz
	diskus_tar_url := "https://github.com/sharkdp/diskus/archive/refs/tags/v0.7.0.tar.gz"
	diskus_cmd_tar := exec.Command("curl", "-L", diskus_tar_url, "-o", "package.tar.gz")
	err := diskus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	diskus_zip_url := "https://github.com/sharkdp/diskus/archive/refs/tags/v0.7.0.zip"
	diskus_cmd_zip := exec.Command("curl", "-L", diskus_zip_url, "-o", "package.zip")
	err = diskus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	diskus_bin_url := "https://github.com/sharkdp/diskus/archive/refs/tags/v0.7.0.bin"
	diskus_cmd_bin := exec.Command("curl", "-L", diskus_bin_url, "-o", "binary.bin")
	err = diskus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	diskus_src_url := "https://github.com/sharkdp/diskus/archive/refs/tags/v0.7.0.src.tar.gz"
	diskus_cmd_src := exec.Command("curl", "-L", diskus_src_url, "-o", "source.tar.gz")
	err = diskus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	diskus_cmd_direct := exec.Command("./binary")
	err = diskus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
