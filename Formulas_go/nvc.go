package main

// Nvc - VHDL compiler and simulator
// Homepage: https://github.com/nickg/nvc

import (
	"fmt"
	
	"os/exec"
)

func installNvc() {
	// Método 1: Descargar y extraer .tar.gz
	nvc_tar_url := "https://github.com/nickg/nvc/releases/download/r1.13.3/nvc-1.13.3.tar.gz"
	nvc_cmd_tar := exec.Command("curl", "-L", nvc_tar_url, "-o", "package.tar.gz")
	err := nvc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nvc_zip_url := "https://github.com/nickg/nvc/releases/download/r1.13.3/nvc-1.13.3.zip"
	nvc_cmd_zip := exec.Command("curl", "-L", nvc_zip_url, "-o", "package.zip")
	err = nvc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nvc_bin_url := "https://github.com/nickg/nvc/releases/download/r1.13.3/nvc-1.13.3.bin"
	nvc_cmd_bin := exec.Command("curl", "-L", nvc_bin_url, "-o", "binary.bin")
	err = nvc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nvc_src_url := "https://github.com/nickg/nvc/releases/download/r1.13.3/nvc-1.13.3.src.tar.gz"
	nvc_cmd_src := exec.Command("curl", "-L", nvc_src_url, "-o", "source.tar.gz")
	err = nvc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nvc_cmd_direct := exec.Command("./binary")
	err = nvc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: check")
exec.Command("latte", "install", "check")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
}
