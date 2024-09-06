package main

// Opusfile - API for decoding and seeking in .opus files
// Homepage: https://www.opus-codec.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpusfile() {
	// Método 1: Descargar y extraer .tar.gz
	opusfile_tar_url := "https://downloads.xiph.org/releases/opus/opusfile-0.12.tar.gz"
	opusfile_cmd_tar := exec.Command("curl", "-L", opusfile_tar_url, "-o", "package.tar.gz")
	err := opusfile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opusfile_zip_url := "https://downloads.xiph.org/releases/opus/opusfile-0.12.zip"
	opusfile_cmd_zip := exec.Command("curl", "-L", opusfile_zip_url, "-o", "package.zip")
	err = opusfile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opusfile_bin_url := "https://downloads.xiph.org/releases/opus/opusfile-0.12.bin"
	opusfile_cmd_bin := exec.Command("curl", "-L", opusfile_bin_url, "-o", "binary.bin")
	err = opusfile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opusfile_src_url := "https://downloads.xiph.org/releases/opus/opusfile-0.12.src.tar.gz"
	opusfile_cmd_src := exec.Command("curl", "-L", opusfile_src_url, "-o", "source.tar.gz")
	err = opusfile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opusfile_cmd_direct := exec.Command("./binary")
	err = opusfile_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
}
