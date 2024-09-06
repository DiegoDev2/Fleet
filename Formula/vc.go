package main

// Vc - SIMD Vector Classes for C++
// Homepage: https://github.com/VcDevel/Vc

import (
	"fmt"
	
	"os/exec"
)

func installVc() {
	// Método 1: Descargar y extraer .tar.gz
	vc_tar_url := "https://github.com/VcDevel/Vc/archive/refs/tags/1.4.5.tar.gz"
	vc_cmd_tar := exec.Command("curl", "-L", vc_tar_url, "-o", "package.tar.gz")
	err := vc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vc_zip_url := "https://github.com/VcDevel/Vc/archive/refs/tags/1.4.5.zip"
	vc_cmd_zip := exec.Command("curl", "-L", vc_zip_url, "-o", "package.zip")
	err = vc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vc_bin_url := "https://github.com/VcDevel/Vc/archive/refs/tags/1.4.5.bin"
	vc_cmd_bin := exec.Command("curl", "-L", vc_bin_url, "-o", "binary.bin")
	err = vc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vc_src_url := "https://github.com/VcDevel/Vc/archive/refs/tags/1.4.5.src.tar.gz"
	vc_cmd_src := exec.Command("curl", "-L", vc_src_url, "-o", "source.tar.gz")
	err = vc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vc_cmd_direct := exec.Command("./binary")
	err = vc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
