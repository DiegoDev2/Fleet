package main

// Pcapplusplus - C++ network sniffing, packet parsing and crafting framework
// Homepage: https://pcapplusplus.github.io

import (
	"fmt"
	
	"os/exec"
)

func installPcapplusplus() {
	// Método 1: Descargar y extraer .tar.gz
	pcapplusplus_tar_url := "https://github.com/seladb/PcapPlusPlus/archive/refs/tags/v24.09.tar.gz"
	pcapplusplus_cmd_tar := exec.Command("curl", "-L", pcapplusplus_tar_url, "-o", "package.tar.gz")
	err := pcapplusplus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcapplusplus_zip_url := "https://github.com/seladb/PcapPlusPlus/archive/refs/tags/v24.09.zip"
	pcapplusplus_cmd_zip := exec.Command("curl", "-L", pcapplusplus_zip_url, "-o", "package.zip")
	err = pcapplusplus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcapplusplus_bin_url := "https://github.com/seladb/PcapPlusPlus/archive/refs/tags/v24.09.bin"
	pcapplusplus_cmd_bin := exec.Command("curl", "-L", pcapplusplus_bin_url, "-o", "binary.bin")
	err = pcapplusplus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcapplusplus_src_url := "https://github.com/seladb/PcapPlusPlus/archive/refs/tags/v24.09.src.tar.gz"
	pcapplusplus_cmd_src := exec.Command("curl", "-L", pcapplusplus_src_url, "-o", "source.tar.gz")
	err = pcapplusplus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcapplusplus_cmd_direct := exec.Command("./binary")
	err = pcapplusplus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
