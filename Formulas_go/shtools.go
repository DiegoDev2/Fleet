package main

// Shtools - Spherical Harmonic Tools
// Homepage: https://shtools.github.io/SHTOOLS/

import (
	"fmt"
	
	"os/exec"
)

func installShtools() {
	// Método 1: Descargar y extraer .tar.gz
	shtools_tar_url := "https://github.com/SHTOOLS/SHTOOLS/archive/refs/tags/v4.13.1.tar.gz"
	shtools_cmd_tar := exec.Command("curl", "-L", shtools_tar_url, "-o", "package.tar.gz")
	err := shtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shtools_zip_url := "https://github.com/SHTOOLS/SHTOOLS/archive/refs/tags/v4.13.1.zip"
	shtools_cmd_zip := exec.Command("curl", "-L", shtools_zip_url, "-o", "package.zip")
	err = shtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shtools_bin_url := "https://github.com/SHTOOLS/SHTOOLS/archive/refs/tags/v4.13.1.bin"
	shtools_cmd_bin := exec.Command("curl", "-L", shtools_bin_url, "-o", "binary.bin")
	err = shtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shtools_src_url := "https://github.com/SHTOOLS/SHTOOLS/archive/refs/tags/v4.13.1.src.tar.gz"
	shtools_cmd_src := exec.Command("curl", "-L", shtools_src_url, "-o", "source.tar.gz")
	err = shtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shtools_cmd_direct := exec.Command("./binary")
	err = shtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
