package main

// VapoursynthImwri - VapourSynth filters - ImageMagick HDRI writer/reader
// Homepage: https://github.com/vapoursynth/vs-imwri

import (
	"fmt"
	
	"os/exec"
)

func installVapoursynthImwri() {
	// Método 1: Descargar y extraer .tar.gz
	vapoursynthimwri_tar_url := "https://github.com/vapoursynth/vs-imwri/archive/refs/tags/R2.tar.gz"
	vapoursynthimwri_cmd_tar := exec.Command("curl", "-L", vapoursynthimwri_tar_url, "-o", "package.tar.gz")
	err := vapoursynthimwri_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vapoursynthimwri_zip_url := "https://github.com/vapoursynth/vs-imwri/archive/refs/tags/R2.zip"
	vapoursynthimwri_cmd_zip := exec.Command("curl", "-L", vapoursynthimwri_zip_url, "-o", "package.zip")
	err = vapoursynthimwri_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vapoursynthimwri_bin_url := "https://github.com/vapoursynth/vs-imwri/archive/refs/tags/R2.bin"
	vapoursynthimwri_cmd_bin := exec.Command("curl", "-L", vapoursynthimwri_bin_url, "-o", "binary.bin")
	err = vapoursynthimwri_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vapoursynthimwri_src_url := "https://github.com/vapoursynth/vs-imwri/archive/refs/tags/R2.src.tar.gz"
	vapoursynthimwri_cmd_src := exec.Command("curl", "-L", vapoursynthimwri_src_url, "-o", "source.tar.gz")
	err = vapoursynthimwri_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vapoursynthimwri_cmd_direct := exec.Command("./binary")
	err = vapoursynthimwri_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: imagemagick")
	exec.Command("latte", "install", "imagemagick").Run()
	fmt.Println("Instalando dependencia: vapoursynth")
	exec.Command("latte", "install", "vapoursynth").Run()
	fmt.Println("Instalando dependencia: jpeg-xl")
	exec.Command("latte", "install", "jpeg-xl").Run()
	fmt.Println("Instalando dependencia: libheif")
	exec.Command("latte", "install", "libheif").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
}
