package main

// OpenjdkAT11 - Development kit for the Java programming language
// Homepage: https://openjdk.java.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpenjdkAT11() {
	// Método 1: Descargar y extraer .tar.gz
	openjdkat11_tar_url := "https://github.com/openjdk/jdk11u/archive/refs/tags/jdk-11.0.24-ga.tar.gz"
	openjdkat11_cmd_tar := exec.Command("curl", "-L", openjdkat11_tar_url, "-o", "package.tar.gz")
	err := openjdkat11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjdkat11_zip_url := "https://github.com/openjdk/jdk11u/archive/refs/tags/jdk-11.0.24-ga.zip"
	openjdkat11_cmd_zip := exec.Command("curl", "-L", openjdkat11_zip_url, "-o", "package.zip")
	err = openjdkat11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjdkat11_bin_url := "https://github.com/openjdk/jdk11u/archive/refs/tags/jdk-11.0.24-ga.bin"
	openjdkat11_cmd_bin := exec.Command("curl", "-L", openjdkat11_bin_url, "-o", "binary.bin")
	err = openjdkat11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjdkat11_src_url := "https://github.com/openjdk/jdk11u/archive/refs/tags/jdk-11.0.24-ga.src.tar.gz"
	openjdkat11_cmd_src := exec.Command("curl", "-L", openjdkat11_src_url, "-o", "source.tar.gz")
	err = openjdkat11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjdkat11_cmd_direct := exec.Command("./binary")
	err = openjdkat11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: libxrandr")
	exec.Command("latte", "install", "libxrandr").Run()
	fmt.Println("Instalando dependencia: libxrender")
	exec.Command("latte", "install", "libxrender").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: libxtst")
	exec.Command("latte", "install", "libxtst").Run()
}
