package main

// Openjdk - Development kit for the Java programming language
// Homepage: https://openjdk.java.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpenjdk() {
	// Método 1: Descargar y extraer .tar.gz
	openjdk_tar_url := "https://github.com/openjdk/jdk22u/archive/refs/tags/jdk-22.0.2-ga.tar.gz"
	openjdk_cmd_tar := exec.Command("curl", "-L", openjdk_tar_url, "-o", "package.tar.gz")
	err := openjdk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjdk_zip_url := "https://github.com/openjdk/jdk22u/archive/refs/tags/jdk-22.0.2-ga.zip"
	openjdk_cmd_zip := exec.Command("curl", "-L", openjdk_zip_url, "-o", "package.zip")
	err = openjdk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjdk_bin_url := "https://github.com/openjdk/jdk22u/archive/refs/tags/jdk-22.0.2-ga.bin"
	openjdk_cmd_bin := exec.Command("curl", "-L", openjdk_bin_url, "-o", "binary.bin")
	err = openjdk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjdk_src_url := "https://github.com/openjdk/jdk22u/archive/refs/tags/jdk-22.0.2-ga.src.tar.gz"
	openjdk_cmd_src := exec.Command("curl", "-L", openjdk_src_url, "-o", "source.tar.gz")
	err = openjdk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjdk_cmd_direct := exec.Command("./binary")
	err = openjdk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
	fmt.Println("Instalando dependencia: libxtst")
exec.Command("latte", "install", "libxtst")
}
