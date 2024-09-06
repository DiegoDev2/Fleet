package main

// Libsigcxx - Callback framework for C++
// Homepage: https://libsigcplusplus.github.io/libsigcplusplus/

import (
	"fmt"
	
	"os/exec"
)

func installLibsigcxx() {
	// Método 1: Descargar y extraer .tar.gz
	libsigcxx_tar_url := "https://download.gnome.org/sources/libsigc++/3.6/libsigc++-3.6.0.tar.xz"
	libsigcxx_cmd_tar := exec.Command("curl", "-L", libsigcxx_tar_url, "-o", "package.tar.gz")
	err := libsigcxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsigcxx_zip_url := "https://download.gnome.org/sources/libsigc++/3.6/libsigc++-3.6.0.tar.xz"
	libsigcxx_cmd_zip := exec.Command("curl", "-L", libsigcxx_zip_url, "-o", "package.zip")
	err = libsigcxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsigcxx_bin_url := "https://download.gnome.org/sources/libsigc++/3.6/libsigc++-3.6.0.tar.xz"
	libsigcxx_cmd_bin := exec.Command("curl", "-L", libsigcxx_bin_url, "-o", "binary.bin")
	err = libsigcxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsigcxx_src_url := "https://download.gnome.org/sources/libsigc++/3.6/libsigc++-3.6.0.tar.xz"
	libsigcxx_cmd_src := exec.Command("curl", "-L", libsigcxx_src_url, "-o", "source.tar.gz")
	err = libsigcxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsigcxx_cmd_direct := exec.Command("./binary")
	err = libsigcxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
