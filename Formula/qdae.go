package main

// Qdae - Quick and Dirty Apricot Emulator
// Homepage: https://www.seasip.info/Unix/QDAE/

import (
	"fmt"
	
	"os/exec"
)

func installQdae() {
	// Método 1: Descargar y extraer .tar.gz
	qdae_tar_url := "https://www.seasip.info/Unix/QDAE/qdae-0.0.10.tar.gz"
	qdae_cmd_tar := exec.Command("curl", "-L", qdae_tar_url, "-o", "package.tar.gz")
	err := qdae_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qdae_zip_url := "https://www.seasip.info/Unix/QDAE/qdae-0.0.10.zip"
	qdae_cmd_zip := exec.Command("curl", "-L", qdae_zip_url, "-o", "package.zip")
	err = qdae_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qdae_bin_url := "https://www.seasip.info/Unix/QDAE/qdae-0.0.10.bin"
	qdae_cmd_bin := exec.Command("curl", "-L", qdae_bin_url, "-o", "binary.bin")
	err = qdae_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qdae_src_url := "https://www.seasip.info/Unix/QDAE/qdae-0.0.10.src.tar.gz"
	qdae_cmd_src := exec.Command("curl", "-L", qdae_src_url, "-o", "source.tar.gz")
	err = qdae_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qdae_cmd_direct := exec.Command("./binary")
	err = qdae_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
