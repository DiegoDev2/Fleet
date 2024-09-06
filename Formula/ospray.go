package main

// Ospray - Ray-tracing-based rendering engine for high-fidelity visualization
// Homepage: https://www.ospray.org/

import (
	"fmt"
	
	"os/exec"
)

func installOspray() {
	// Método 1: Descargar y extraer .tar.gz
	ospray_tar_url := "https://github.com/ospray/ospray/archive/refs/tags/v3.2.0.tar.gz"
	ospray_cmd_tar := exec.Command("curl", "-L", ospray_tar_url, "-o", "package.tar.gz")
	err := ospray_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ospray_zip_url := "https://github.com/ospray/ospray/archive/refs/tags/v3.2.0.zip"
	ospray_cmd_zip := exec.Command("curl", "-L", ospray_zip_url, "-o", "package.zip")
	err = ospray_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ospray_bin_url := "https://github.com/ospray/ospray/archive/refs/tags/v3.2.0.bin"
	ospray_cmd_bin := exec.Command("curl", "-L", ospray_bin_url, "-o", "binary.bin")
	err = ospray_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ospray_src_url := "https://github.com/ospray/ospray/archive/refs/tags/v3.2.0.src.tar.gz"
	ospray_cmd_src := exec.Command("curl", "-L", ospray_src_url, "-o", "source.tar.gz")
	err = ospray_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ospray_cmd_direct := exec.Command("./binary")
	err = ospray_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: embree")
	exec.Command("latte", "install", "embree").Run()
	fmt.Println("Instalando dependencia: ispc")
	exec.Command("latte", "install", "ispc").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
}
