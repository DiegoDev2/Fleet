package main

// Logswan - Fast Web log analyzer using probabilistic data structures
// Homepage: https://www.logswan.org

import (
	"fmt"
	
	"os/exec"
)

func installLogswan() {
	// Método 1: Descargar y extraer .tar.gz
	logswan_tar_url := "https://github.com/fcambus/logswan/archive/refs/tags/2.1.14.tar.gz"
	logswan_cmd_tar := exec.Command("curl", "-L", logswan_tar_url, "-o", "package.tar.gz")
	err := logswan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	logswan_zip_url := "https://github.com/fcambus/logswan/archive/refs/tags/2.1.14.zip"
	logswan_cmd_zip := exec.Command("curl", "-L", logswan_zip_url, "-o", "package.zip")
	err = logswan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	logswan_bin_url := "https://github.com/fcambus/logswan/archive/refs/tags/2.1.14.bin"
	logswan_cmd_bin := exec.Command("curl", "-L", logswan_bin_url, "-o", "binary.bin")
	err = logswan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	logswan_src_url := "https://github.com/fcambus/logswan/archive/refs/tags/2.1.14.src.tar.gz"
	logswan_cmd_src := exec.Command("curl", "-L", logswan_src_url, "-o", "source.tar.gz")
	err = logswan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	logswan_cmd_direct := exec.Command("./binary")
	err = logswan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: libmaxminddb")
	exec.Command("latte", "install", "libmaxminddb").Run()
}
