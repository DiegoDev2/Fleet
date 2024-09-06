package main

// Nemu - Ncurses UI for QEMU
// Homepage: https://github.com/nemuTUI/nemu

import (
	"fmt"
	
	"os/exec"
)

func installNemu() {
	// Método 1: Descargar y extraer .tar.gz
	nemu_tar_url := "https://github.com/nemuTUI/nemu/archive/refs/tags/v3.3.1.tar.gz"
	nemu_cmd_tar := exec.Command("curl", "-L", nemu_tar_url, "-o", "package.tar.gz")
	err := nemu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nemu_zip_url := "https://github.com/nemuTUI/nemu/archive/refs/tags/v3.3.1.zip"
	nemu_cmd_zip := exec.Command("curl", "-L", nemu_zip_url, "-o", "package.zip")
	err = nemu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nemu_bin_url := "https://github.com/nemuTUI/nemu/archive/refs/tags/v3.3.1.bin"
	nemu_cmd_bin := exec.Command("curl", "-L", nemu_bin_url, "-o", "binary.bin")
	err = nemu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nemu_src_url := "https://github.com/nemuTUI/nemu/archive/refs/tags/v3.3.1.src.tar.gz"
	nemu_cmd_src := exec.Command("curl", "-L", nemu_src_url, "-o", "source.tar.gz")
	err = nemu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nemu_cmd_direct := exec.Command("./binary")
	err = nemu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
