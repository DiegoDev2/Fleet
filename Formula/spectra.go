package main

// Spectra - Header-only C++ library for large scale eigenvalue problems
// Homepage: https://spectralib.org

import (
	"fmt"
	
	"os/exec"
)

func installSpectra() {
	// Método 1: Descargar y extraer .tar.gz
	spectra_tar_url := "https://github.com/yixuan/spectra/archive/refs/tags/v1.0.1.tar.gz"
	spectra_cmd_tar := exec.Command("curl", "-L", spectra_tar_url, "-o", "package.tar.gz")
	err := spectra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spectra_zip_url := "https://github.com/yixuan/spectra/archive/refs/tags/v1.0.1.zip"
	spectra_cmd_zip := exec.Command("curl", "-L", spectra_zip_url, "-o", "package.zip")
	err = spectra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spectra_bin_url := "https://github.com/yixuan/spectra/archive/refs/tags/v1.0.1.bin"
	spectra_cmd_bin := exec.Command("curl", "-L", spectra_bin_url, "-o", "binary.bin")
	err = spectra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spectra_src_url := "https://github.com/yixuan/spectra/archive/refs/tags/v1.0.1.src.tar.gz"
	spectra_cmd_src := exec.Command("curl", "-L", spectra_src_url, "-o", "source.tar.gz")
	err = spectra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spectra_cmd_direct := exec.Command("./binary")
	err = spectra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
}
