package main

// Libvterm - C99 library which implements a VT220 or xterm terminal emulator
// Homepage: https://www.leonerd.org.uk/code/libvterm/

import (
	"fmt"
	
	"os/exec"
)

func installLibvterm() {
	// Método 1: Descargar y extraer .tar.gz
	libvterm_tar_url := "https://launchpad.net/libvterm/trunk/v0.3/+download/libvterm-0.3.3.tar.gz"
	libvterm_cmd_tar := exec.Command("curl", "-L", libvterm_tar_url, "-o", "package.tar.gz")
	err := libvterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvterm_zip_url := "https://launchpad.net/libvterm/trunk/v0.3/+download/libvterm-0.3.3.zip"
	libvterm_cmd_zip := exec.Command("curl", "-L", libvterm_zip_url, "-o", "package.zip")
	err = libvterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvterm_bin_url := "https://launchpad.net/libvterm/trunk/v0.3/+download/libvterm-0.3.3.bin"
	libvterm_cmd_bin := exec.Command("curl", "-L", libvterm_bin_url, "-o", "binary.bin")
	err = libvterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvterm_src_url := "https://launchpad.net/libvterm/trunk/v0.3/+download/libvterm-0.3.3.src.tar.gz"
	libvterm_cmd_src := exec.Command("curl", "-L", libvterm_src_url, "-o", "source.tar.gz")
	err = libvterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvterm_cmd_direct := exec.Command("./binary")
	err = libvterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
