package main

// Toxcore - C library implementing the Tox peer to peer network protocol
// Homepage: https://tox.chat/

import (
	"fmt"
	
	"os/exec"
)

func installToxcore() {
	// Método 1: Descargar y extraer .tar.gz
	toxcore_tar_url := "https://github.com/TokTok/c-toxcore/releases/download/v0.2.19/c-toxcore-0.2.19.tar.gz"
	toxcore_cmd_tar := exec.Command("curl", "-L", toxcore_tar_url, "-o", "package.tar.gz")
	err := toxcore_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toxcore_zip_url := "https://github.com/TokTok/c-toxcore/releases/download/v0.2.19/c-toxcore-0.2.19.zip"
	toxcore_cmd_zip := exec.Command("curl", "-L", toxcore_zip_url, "-o", "package.zip")
	err = toxcore_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toxcore_bin_url := "https://github.com/TokTok/c-toxcore/releases/download/v0.2.19/c-toxcore-0.2.19.bin"
	toxcore_cmd_bin := exec.Command("curl", "-L", toxcore_bin_url, "-o", "binary.bin")
	err = toxcore_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toxcore_src_url := "https://github.com/TokTok/c-toxcore/releases/download/v0.2.19/c-toxcore-0.2.19.src.tar.gz"
	toxcore_cmd_src := exec.Command("curl", "-L", toxcore_src_url, "-o", "source.tar.gz")
	err = toxcore_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toxcore_cmd_direct := exec.Command("./binary")
	err = toxcore_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libconfig")
	exec.Command("latte", "install", "libconfig").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: libvpx")
	exec.Command("latte", "install", "libvpx").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
}
