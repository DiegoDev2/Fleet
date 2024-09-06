package main

// Aria2 - Download with resuming and segmented downloading
// Homepage: https://aria2.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installAria2() {
	// Método 1: Descargar y extraer .tar.gz
	aria2_tar_url := "https://github.com/aria2/aria2/releases/download/release-1.37.0/aria2-1.37.0.tar.xz"
	aria2_cmd_tar := exec.Command("curl", "-L", aria2_tar_url, "-o", "package.tar.gz")
	err := aria2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aria2_zip_url := "https://github.com/aria2/aria2/releases/download/release-1.37.0/aria2-1.37.0.tar.xz"
	aria2_cmd_zip := exec.Command("curl", "-L", aria2_zip_url, "-o", "package.zip")
	err = aria2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aria2_bin_url := "https://github.com/aria2/aria2/releases/download/release-1.37.0/aria2-1.37.0.tar.xz"
	aria2_cmd_bin := exec.Command("curl", "-L", aria2_bin_url, "-o", "binary.bin")
	err = aria2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aria2_src_url := "https://github.com/aria2/aria2/releases/download/release-1.37.0/aria2-1.37.0.tar.xz"
	aria2_cmd_src := exec.Command("curl", "-L", aria2_src_url, "-o", "source.tar.gz")
	err = aria2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aria2_cmd_direct := exec.Command("./binary")
	err = aria2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
