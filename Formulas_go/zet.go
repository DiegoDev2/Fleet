package main

// Zet - CLI utility to find the union, intersection, and set difference of files
// Homepage: https://github.com/yarrow/zet

import (
	"fmt"
	
	"os/exec"
)

func installZet() {
	// Método 1: Descargar y extraer .tar.gz
	zet_tar_url := "https://github.com/yarrow/zet/archive/refs/tags/v1.0.0.tar.gz"
	zet_cmd_tar := exec.Command("curl", "-L", zet_tar_url, "-o", "package.tar.gz")
	err := zet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zet_zip_url := "https://github.com/yarrow/zet/archive/refs/tags/v1.0.0.zip"
	zet_cmd_zip := exec.Command("curl", "-L", zet_zip_url, "-o", "package.zip")
	err = zet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zet_bin_url := "https://github.com/yarrow/zet/archive/refs/tags/v1.0.0.bin"
	zet_cmd_bin := exec.Command("curl", "-L", zet_bin_url, "-o", "binary.bin")
	err = zet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zet_src_url := "https://github.com/yarrow/zet/archive/refs/tags/v1.0.0.src.tar.gz"
	zet_cmd_src := exec.Command("curl", "-L", zet_src_url, "-o", "source.tar.gz")
	err = zet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zet_cmd_direct := exec.Command("./binary")
	err = zet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
