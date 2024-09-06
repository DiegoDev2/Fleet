package main

// Nift - Cross-platform open source framework for managing and generating websites
// Homepage: https://nift.dev/

import (
	"fmt"
	
	"os/exec"
)

func installNift() {
	// Método 1: Descargar y extraer .tar.gz
	nift_tar_url := "https://github.com/nifty-site-manager/nsm/archive/refs/tags/v3.0.3.tar.gz"
	nift_cmd_tar := exec.Command("curl", "-L", nift_tar_url, "-o", "package.tar.gz")
	err := nift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nift_zip_url := "https://github.com/nifty-site-manager/nsm/archive/refs/tags/v3.0.3.zip"
	nift_cmd_zip := exec.Command("curl", "-L", nift_zip_url, "-o", "package.zip")
	err = nift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nift_bin_url := "https://github.com/nifty-site-manager/nsm/archive/refs/tags/v3.0.3.bin"
	nift_cmd_bin := exec.Command("curl", "-L", nift_bin_url, "-o", "binary.bin")
	err = nift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nift_src_url := "https://github.com/nifty-site-manager/nsm/archive/refs/tags/v3.0.3.src.tar.gz"
	nift_cmd_src := exec.Command("curl", "-L", nift_src_url, "-o", "source.tar.gz")
	err = nift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nift_cmd_direct := exec.Command("./binary")
	err = nift_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
}
