package main

// Libarchive - Multi-format archive and compression library
// Homepage: https://www.libarchive.org

import (
	"fmt"
	
	"os/exec"
)

func installLibarchive() {
	// Método 1: Descargar y extraer .tar.gz
	libarchive_tar_url := "https://www.libarchive.org/downloads/libarchive-3.7.4.tar.xz"
	libarchive_cmd_tar := exec.Command("curl", "-L", libarchive_tar_url, "-o", "package.tar.gz")
	err := libarchive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libarchive_zip_url := "https://www.libarchive.org/downloads/libarchive-3.7.4.tar.xz"
	libarchive_cmd_zip := exec.Command("curl", "-L", libarchive_zip_url, "-o", "package.zip")
	err = libarchive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libarchive_bin_url := "https://www.libarchive.org/downloads/libarchive-3.7.4.tar.xz"
	libarchive_cmd_bin := exec.Command("curl", "-L", libarchive_bin_url, "-o", "binary.bin")
	err = libarchive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libarchive_src_url := "https://www.libarchive.org/downloads/libarchive-3.7.4.tar.xz"
	libarchive_cmd_src := exec.Command("curl", "-L", libarchive_src_url, "-o", "source.tar.gz")
	err = libarchive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libarchive_cmd_direct := exec.Command("./binary")
	err = libarchive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libb2")
exec.Command("latte", "install", "libb2")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
