package main

// Lsyncd - Synchronize local directories with remote targets
// Homepage: https://github.com/lsyncd/lsyncd

import (
	"fmt"
	
	"os/exec"
)

func installLsyncd() {
	// Método 1: Descargar y extraer .tar.gz
	lsyncd_tar_url := "https://github.com/lsyncd/lsyncd/archive/refs/tags/release-2.3.1.tar.gz"
	lsyncd_cmd_tar := exec.Command("curl", "-L", lsyncd_tar_url, "-o", "package.tar.gz")
	err := lsyncd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lsyncd_zip_url := "https://github.com/lsyncd/lsyncd/archive/refs/tags/release-2.3.1.zip"
	lsyncd_cmd_zip := exec.Command("curl", "-L", lsyncd_zip_url, "-o", "package.zip")
	err = lsyncd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lsyncd_bin_url := "https://github.com/lsyncd/lsyncd/archive/refs/tags/release-2.3.1.bin"
	lsyncd_cmd_bin := exec.Command("curl", "-L", lsyncd_bin_url, "-o", "binary.bin")
	err = lsyncd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lsyncd_src_url := "https://github.com/lsyncd/lsyncd/archive/refs/tags/release-2.3.1.src.tar.gz"
	lsyncd_cmd_src := exec.Command("curl", "-L", lsyncd_src_url, "-o", "source.tar.gz")
	err = lsyncd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lsyncd_cmd_direct := exec.Command("./binary")
	err = lsyncd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: lua")
	exec.Command("latte", "install", "lua").Run()
}
