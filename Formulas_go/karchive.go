package main

// Karchive - Reading, creating, and manipulating file archives
// Homepage: https://api.kde.org/frameworks/karchive/html/index.html

import (
	"fmt"
	
	"os/exec"
)

func installKarchive() {
	// Método 1: Descargar y extraer .tar.gz
	karchive_tar_url := "https://download.kde.org/stable/frameworks/6.5/karchive-6.5.0.tar.xz"
	karchive_cmd_tar := exec.Command("curl", "-L", karchive_tar_url, "-o", "package.tar.gz")
	err := karchive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	karchive_zip_url := "https://download.kde.org/stable/frameworks/6.5/karchive-6.5.0.tar.xz"
	karchive_cmd_zip := exec.Command("curl", "-L", karchive_zip_url, "-o", "package.zip")
	err = karchive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	karchive_bin_url := "https://download.kde.org/stable/frameworks/6.5/karchive-6.5.0.tar.xz"
	karchive_cmd_bin := exec.Command("curl", "-L", karchive_bin_url, "-o", "binary.bin")
	err = karchive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	karchive_src_url := "https://download.kde.org/stable/frameworks/6.5/karchive-6.5.0.tar.xz"
	karchive_cmd_src := exec.Command("curl", "-L", karchive_src_url, "-o", "source.tar.gz")
	err = karchive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	karchive_cmd_direct := exec.Command("./binary")
	err = karchive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: extra-cmake-modules")
exec.Command("latte", "install", "extra-cmake-modules")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
