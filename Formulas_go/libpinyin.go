package main

// Libpinyin - Library to deal with pinyin
// Homepage: https://github.com/libpinyin/libpinyin

import (
	"fmt"
	
	"os/exec"
)

func installLibpinyin() {
	// Método 1: Descargar y extraer .tar.gz
	libpinyin_tar_url := "https://github.com/libpinyin/libpinyin/archive/refs/tags/2.8.1.tar.gz"
	libpinyin_cmd_tar := exec.Command("curl", "-L", libpinyin_tar_url, "-o", "package.tar.gz")
	err := libpinyin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpinyin_zip_url := "https://github.com/libpinyin/libpinyin/archive/refs/tags/2.8.1.zip"
	libpinyin_cmd_zip := exec.Command("curl", "-L", libpinyin_zip_url, "-o", "package.zip")
	err = libpinyin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpinyin_bin_url := "https://github.com/libpinyin/libpinyin/archive/refs/tags/2.8.1.bin"
	libpinyin_cmd_bin := exec.Command("curl", "-L", libpinyin_bin_url, "-o", "binary.bin")
	err = libpinyin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpinyin_src_url := "https://github.com/libpinyin/libpinyin/archive/refs/tags/2.8.1.src.tar.gz"
	libpinyin_cmd_src := exec.Command("curl", "-L", libpinyin_src_url, "-o", "source.tar.gz")
	err = libpinyin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpinyin_cmd_direct := exec.Command("./binary")
	err = libpinyin_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: berkeley-db")
exec.Command("latte", "install", "berkeley-db")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
}
