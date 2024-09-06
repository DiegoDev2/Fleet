package main

// OpenjdkAT8 - Development kit for the Java programming language
// Homepage: https://openjdk.java.net/

import (
	"fmt"
	
	"os/exec"
)

func installOpenjdkAT8() {
	// Método 1: Descargar y extraer .tar.gz
	openjdkat8_tar_url := "https://github.com/openjdk/jdk8u/archive/refs/tags/jdk8u422-ga.tar.gz"
	openjdkat8_cmd_tar := exec.Command("curl", "-L", openjdkat8_tar_url, "-o", "package.tar.gz")
	err := openjdkat8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openjdkat8_zip_url := "https://github.com/openjdk/jdk8u/archive/refs/tags/jdk8u422-ga.zip"
	openjdkat8_cmd_zip := exec.Command("curl", "-L", openjdkat8_zip_url, "-o", "package.zip")
	err = openjdkat8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openjdkat8_bin_url := "https://github.com/openjdk/jdk8u/archive/refs/tags/jdk8u422-ga.bin"
	openjdkat8_cmd_bin := exec.Command("curl", "-L", openjdkat8_bin_url, "-o", "binary.bin")
	err = openjdkat8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openjdkat8_src_url := "https://github.com/openjdk/jdk8u/archive/refs/tags/jdk8u422-ga.src.tar.gz"
	openjdkat8_cmd_src := exec.Command("curl", "-L", openjdkat8_src_url, "-o", "source.tar.gz")
	err = openjdkat8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openjdkat8_cmd_direct := exec.Command("./binary")
	err = openjdkat8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
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
