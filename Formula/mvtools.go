package main

// Mvtools - Filters for motion estimation and compensation
// Homepage: https://github.com/dubhater/vapoursynth-mvtools

import (
	"fmt"
	
	"os/exec"
)

func installMvtools() {
	// Método 1: Descargar y extraer .tar.gz
	mvtools_tar_url := "https://github.com/dubhater/vapoursynth-mvtools/archive/refs/tags/v24.tar.gz"
	mvtools_cmd_tar := exec.Command("curl", "-L", mvtools_tar_url, "-o", "package.tar.gz")
	err := mvtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mvtools_zip_url := "https://github.com/dubhater/vapoursynth-mvtools/archive/refs/tags/v24.zip"
	mvtools_cmd_zip := exec.Command("curl", "-L", mvtools_zip_url, "-o", "package.zip")
	err = mvtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mvtools_bin_url := "https://github.com/dubhater/vapoursynth-mvtools/archive/refs/tags/v24.bin"
	mvtools_cmd_bin := exec.Command("curl", "-L", mvtools_bin_url, "-o", "binary.bin")
	err = mvtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mvtools_src_url := "https://github.com/dubhater/vapoursynth-mvtools/archive/refs/tags/v24.src.tar.gz"
	mvtools_cmd_src := exec.Command("curl", "-L", mvtools_src_url, "-o", "source.tar.gz")
	err = mvtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mvtools_cmd_direct := exec.Command("./binary")
	err = mvtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: vapoursynth")
	exec.Command("latte", "install", "vapoursynth").Run()
}
