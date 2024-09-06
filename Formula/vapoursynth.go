package main

// Vapoursynth - Video processing framework with simplicity in mind
// Homepage: https://www.vapoursynth.com

import (
	"fmt"
	
	"os/exec"
)

func installVapoursynth() {
	// Método 1: Descargar y extraer .tar.gz
	vapoursynth_tar_url := "https://github.com/vapoursynth/vapoursynth/archive/refs/tags/R69.tar.gz"
	vapoursynth_cmd_tar := exec.Command("curl", "-L", vapoursynth_tar_url, "-o", "package.tar.gz")
	err := vapoursynth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vapoursynth_zip_url := "https://github.com/vapoursynth/vapoursynth/archive/refs/tags/R69.zip"
	vapoursynth_cmd_zip := exec.Command("curl", "-L", vapoursynth_zip_url, "-o", "package.zip")
	err = vapoursynth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vapoursynth_bin_url := "https://github.com/vapoursynth/vapoursynth/archive/refs/tags/R69.bin"
	vapoursynth_cmd_bin := exec.Command("curl", "-L", vapoursynth_bin_url, "-o", "binary.bin")
	err = vapoursynth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vapoursynth_src_url := "https://github.com/vapoursynth/vapoursynth/archive/refs/tags/R69.src.tar.gz"
	vapoursynth_cmd_src := exec.Command("curl", "-L", vapoursynth_src_url, "-o", "source.tar.gz")
	err = vapoursynth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vapoursynth_cmd_direct := exec.Command("./binary")
	err = vapoursynth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cython")
	exec.Command("latte", "install", "cython").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: zimg")
	exec.Command("latte", "install", "zimg").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
