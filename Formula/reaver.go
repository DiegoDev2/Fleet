package main

// Reaver - Implements brute force attack to recover WPA/WPA2 passkeys
// Homepage: https://github.com/t6x/reaver-wps-fork-t6x

import (
	"fmt"
	
	"os/exec"
)

func installReaver() {
	// Método 1: Descargar y extraer .tar.gz
	reaver_tar_url := "https://github.com/t6x/reaver-wps-fork-t6x/releases/download/v1.6.6/reaver-1.6.6.tar.xz"
	reaver_cmd_tar := exec.Command("curl", "-L", reaver_tar_url, "-o", "package.tar.gz")
	err := reaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reaver_zip_url := "https://github.com/t6x/reaver-wps-fork-t6x/releases/download/v1.6.6/reaver-1.6.6.tar.xz"
	reaver_cmd_zip := exec.Command("curl", "-L", reaver_zip_url, "-o", "package.zip")
	err = reaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reaver_bin_url := "https://github.com/t6x/reaver-wps-fork-t6x/releases/download/v1.6.6/reaver-1.6.6.tar.xz"
	reaver_cmd_bin := exec.Command("curl", "-L", reaver_bin_url, "-o", "binary.bin")
	err = reaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reaver_src_url := "https://github.com/t6x/reaver-wps-fork-t6x/releases/download/v1.6.6/reaver-1.6.6.tar.xz"
	reaver_cmd_src := exec.Command("curl", "-L", reaver_src_url, "-o", "source.tar.gz")
	err = reaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reaver_cmd_direct := exec.Command("./binary")
	err = reaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pixiewps")
	exec.Command("latte", "install", "pixiewps").Run()
}
