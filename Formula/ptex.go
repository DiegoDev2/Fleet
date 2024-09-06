package main

// Ptex - Texture mapping system
// Homepage: https://ptex.us/

import (
	"fmt"
	
	"os/exec"
)

func installPtex() {
	// Método 1: Descargar y extraer .tar.gz
	ptex_tar_url := "https://github.com/wdas/ptex/archive/refs/tags/v2.4.3.tar.gz"
	ptex_cmd_tar := exec.Command("curl", "-L", ptex_tar_url, "-o", "package.tar.gz")
	err := ptex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ptex_zip_url := "https://github.com/wdas/ptex/archive/refs/tags/v2.4.3.zip"
	ptex_cmd_zip := exec.Command("curl", "-L", ptex_zip_url, "-o", "package.zip")
	err = ptex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ptex_bin_url := "https://github.com/wdas/ptex/archive/refs/tags/v2.4.3.bin"
	ptex_cmd_bin := exec.Command("curl", "-L", ptex_bin_url, "-o", "binary.bin")
	err = ptex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ptex_src_url := "https://github.com/wdas/ptex/archive/refs/tags/v2.4.3.src.tar.gz"
	ptex_cmd_src := exec.Command("curl", "-L", ptex_src_url, "-o", "source.tar.gz")
	err = ptex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ptex_cmd_direct := exec.Command("./binary")
	err = ptex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
