package main

// OpusTools - Utilities to encode, inspect, and decode .opus files
// Homepage: https://www.opus-codec.org

import (
	"fmt"
	
	"os/exec"
)

func installOpusTools() {
	// Método 1: Descargar y extraer .tar.gz
	opustools_tar_url := "https://archive.mozilla.org/pub/opus/opus-tools-0.2.tar.gz"
	opustools_cmd_tar := exec.Command("curl", "-L", opustools_tar_url, "-o", "package.tar.gz")
	err := opustools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opustools_zip_url := "https://archive.mozilla.org/pub/opus/opus-tools-0.2.zip"
	opustools_cmd_zip := exec.Command("curl", "-L", opustools_zip_url, "-o", "package.zip")
	err = opustools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opustools_bin_url := "https://archive.mozilla.org/pub/opus/opus-tools-0.2.bin"
	opustools_cmd_bin := exec.Command("curl", "-L", opustools_bin_url, "-o", "binary.bin")
	err = opustools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opustools_src_url := "https://archive.mozilla.org/pub/opus/opus-tools-0.2.src.tar.gz"
	opustools_cmd_src := exec.Command("curl", "-L", opustools_src_url, "-o", "source.tar.gz")
	err = opustools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opustools_cmd_direct := exec.Command("./binary")
	err = opustools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libopusenc")
exec.Command("latte", "install", "libopusenc")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: opusfile")
exec.Command("latte", "install", "opusfile")
}
