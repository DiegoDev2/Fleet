package main

// Libxkbcommon - Keyboard handling library
// Homepage: https://xkbcommon.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxkbcommon() {
	// Método 1: Descargar y extraer .tar.gz
	libxkbcommon_tar_url := "https://xkbcommon.org/download/libxkbcommon-1.7.0.tar.xz"
	libxkbcommon_cmd_tar := exec.Command("curl", "-L", libxkbcommon_tar_url, "-o", "package.tar.gz")
	err := libxkbcommon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxkbcommon_zip_url := "https://xkbcommon.org/download/libxkbcommon-1.7.0.tar.xz"
	libxkbcommon_cmd_zip := exec.Command("curl", "-L", libxkbcommon_zip_url, "-o", "package.zip")
	err = libxkbcommon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxkbcommon_bin_url := "https://xkbcommon.org/download/libxkbcommon-1.7.0.tar.xz"
	libxkbcommon_cmd_bin := exec.Command("curl", "-L", libxkbcommon_bin_url, "-o", "binary.bin")
	err = libxkbcommon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxkbcommon_src_url := "https://xkbcommon.org/download/libxkbcommon-1.7.0.tar.xz"
	libxkbcommon_cmd_src := exec.Command("curl", "-L", libxkbcommon_src_url, "-o", "source.tar.gz")
	err = libxkbcommon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxkbcommon_cmd_direct := exec.Command("./binary")
	err = libxkbcommon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: xkeyboardconfig")
	exec.Command("latte", "install", "xkeyboardconfig").Run()
	fmt.Println("Instalando dependencia: xorg-server")
	exec.Command("latte", "install", "xorg-server").Run()
}
