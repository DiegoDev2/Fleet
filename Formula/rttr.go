package main

// Rttr - C++ Reflection Library
// Homepage: https://www.rttr.org

import (
	"fmt"
	
	"os/exec"
)

func installRttr() {
	// Método 1: Descargar y extraer .tar.gz
	rttr_tar_url := "https://github.com/rttrorg/rttr/archive/refs/tags/v0.9.6.tar.gz"
	rttr_cmd_tar := exec.Command("curl", "-L", rttr_tar_url, "-o", "package.tar.gz")
	err := rttr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rttr_zip_url := "https://github.com/rttrorg/rttr/archive/refs/tags/v0.9.6.zip"
	rttr_cmd_zip := exec.Command("curl", "-L", rttr_zip_url, "-o", "package.zip")
	err = rttr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rttr_bin_url := "https://github.com/rttrorg/rttr/archive/refs/tags/v0.9.6.bin"
	rttr_cmd_bin := exec.Command("curl", "-L", rttr_bin_url, "-o", "binary.bin")
	err = rttr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rttr_src_url := "https://github.com/rttrorg/rttr/archive/refs/tags/v0.9.6.src.tar.gz"
	rttr_cmd_src := exec.Command("curl", "-L", rttr_src_url, "-o", "source.tar.gz")
	err = rttr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rttr_cmd_direct := exec.Command("./binary")
	err = rttr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
}
