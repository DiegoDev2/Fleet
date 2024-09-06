package main

// Kalign - Fast multiple sequence alignment program for biological sequences
// Homepage: https://github.com/TimoLassmann/kalign

import (
	"fmt"
	
	"os/exec"
)

func installKalign() {
	// Método 1: Descargar y extraer .tar.gz
	kalign_tar_url := "https://github.com/TimoLassmann/kalign/archive/refs/tags/v3.4.0.tar.gz"
	kalign_cmd_tar := exec.Command("curl", "-L", kalign_tar_url, "-o", "package.tar.gz")
	err := kalign_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kalign_zip_url := "https://github.com/TimoLassmann/kalign/archive/refs/tags/v3.4.0.zip"
	kalign_cmd_zip := exec.Command("curl", "-L", kalign_zip_url, "-o", "package.zip")
	err = kalign_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kalign_bin_url := "https://github.com/TimoLassmann/kalign/archive/refs/tags/v3.4.0.bin"
	kalign_cmd_bin := exec.Command("curl", "-L", kalign_bin_url, "-o", "binary.bin")
	err = kalign_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kalign_src_url := "https://github.com/TimoLassmann/kalign/archive/refs/tags/v3.4.0.src.tar.gz"
	kalign_cmd_src := exec.Command("curl", "-L", kalign_src_url, "-o", "source.tar.gz")
	err = kalign_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kalign_cmd_direct := exec.Command("./binary")
	err = kalign_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
