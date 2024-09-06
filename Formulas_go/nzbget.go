package main

// Nzbget - Binary newsgrabber for nzb files
// Homepage: https://nzbget.com

import (
	"fmt"
	
	"os/exec"
)

func installNzbget() {
	// Método 1: Descargar y extraer .tar.gz
	nzbget_tar_url := "https://github.com/nzbgetcom/nzbget/archive/refs/tags/v24.2.tar.gz"
	nzbget_cmd_tar := exec.Command("curl", "-L", nzbget_tar_url, "-o", "package.tar.gz")
	err := nzbget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nzbget_zip_url := "https://github.com/nzbgetcom/nzbget/archive/refs/tags/v24.2.zip"
	nzbget_cmd_zip := exec.Command("curl", "-L", nzbget_zip_url, "-o", "package.zip")
	err = nzbget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nzbget_bin_url := "https://github.com/nzbgetcom/nzbget/archive/refs/tags/v24.2.bin"
	nzbget_cmd_bin := exec.Command("curl", "-L", nzbget_bin_url, "-o", "binary.bin")
	err = nzbget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nzbget_src_url := "https://github.com/nzbgetcom/nzbget/archive/refs/tags/v24.2.src.tar.gz"
	nzbget_cmd_src := exec.Command("curl", "-L", nzbget_src_url, "-o", "source.tar.gz")
	err = nzbget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nzbget_cmd_direct := exec.Command("./binary")
	err = nzbget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: sevenzip")
exec.Command("latte", "install", "sevenzip")
}
