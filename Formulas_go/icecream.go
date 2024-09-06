package main

// Icecream - Distributed compiler with a central scheduler to share build load
// Homepage: https://en.opensuse.org/Icecream

import (
	"fmt"
	
	"os/exec"
)

func installIcecream() {
	// Método 1: Descargar y extraer .tar.gz
	icecream_tar_url := "https://github.com/icecc/icecream/archive/refs/tags/1.4.tar.gz"
	icecream_cmd_tar := exec.Command("curl", "-L", icecream_tar_url, "-o", "package.tar.gz")
	err := icecream_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icecream_zip_url := "https://github.com/icecc/icecream/archive/refs/tags/1.4.zip"
	icecream_cmd_zip := exec.Command("curl", "-L", icecream_zip_url, "-o", "package.zip")
	err = icecream_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icecream_bin_url := "https://github.com/icecc/icecream/archive/refs/tags/1.4.bin"
	icecream_cmd_bin := exec.Command("curl", "-L", icecream_bin_url, "-o", "binary.bin")
	err = icecream_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icecream_src_url := "https://github.com/icecc/icecream/archive/refs/tags/1.4.src.tar.gz"
	icecream_cmd_src := exec.Command("curl", "-L", icecream_src_url, "-o", "source.tar.gz")
	err = icecream_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icecream_cmd_direct := exec.Command("./binary")
	err = icecream_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: docbook2x")
exec.Command("latte", "install", "docbook2x")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: lzo")
exec.Command("latte", "install", "lzo")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: libcap-ng")
exec.Command("latte", "install", "libcap-ng")
}
