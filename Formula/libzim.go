package main

// Libzim - Reference implementation of the ZIM specification
// Homepage: https://github.com/openzim/libzim

import (
	"fmt"
	
	"os/exec"
)

func installLibzim() {
	// Método 1: Descargar y extraer .tar.gz
	libzim_tar_url := "https://github.com/openzim/libzim/archive/refs/tags/9.2.3.tar.gz"
	libzim_cmd_tar := exec.Command("curl", "-L", libzim_tar_url, "-o", "package.tar.gz")
	err := libzim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libzim_zip_url := "https://github.com/openzim/libzim/archive/refs/tags/9.2.3.zip"
	libzim_cmd_zip := exec.Command("curl", "-L", libzim_zip_url, "-o", "package.zip")
	err = libzim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libzim_bin_url := "https://github.com/openzim/libzim/archive/refs/tags/9.2.3.bin"
	libzim_cmd_bin := exec.Command("curl", "-L", libzim_bin_url, "-o", "binary.bin")
	err = libzim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libzim_src_url := "https://github.com/openzim/libzim/archive/refs/tags/9.2.3.src.tar.gz"
	libzim_cmd_src := exec.Command("curl", "-L", libzim_src_url, "-o", "source.tar.gz")
	err = libzim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libzim_cmd_direct := exec.Command("./binary")
	err = libzim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: xapian")
	exec.Command("latte", "install", "xapian").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
