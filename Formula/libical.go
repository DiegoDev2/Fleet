package main

// Libical - Implementation of iCalendar protocols and data formats
// Homepage: https://libical.github.io/libical/

import (
	"fmt"
	
	"os/exec"
)

func installLibical() {
	// Método 1: Descargar y extraer .tar.gz
	libical_tar_url := "https://github.com/libical/libical/releases/download/v3.0.18/libical-3.0.18.tar.gz"
	libical_cmd_tar := exec.Command("curl", "-L", libical_tar_url, "-o", "package.tar.gz")
	err := libical_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libical_zip_url := "https://github.com/libical/libical/releases/download/v3.0.18/libical-3.0.18.zip"
	libical_cmd_zip := exec.Command("curl", "-L", libical_zip_url, "-o", "package.zip")
	err = libical_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libical_bin_url := "https://github.com/libical/libical/releases/download/v3.0.18/libical-3.0.18.bin"
	libical_cmd_bin := exec.Command("curl", "-L", libical_bin_url, "-o", "binary.bin")
	err = libical_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libical_src_url := "https://github.com/libical/libical/releases/download/v3.0.18/libical-3.0.18.src.tar.gz"
	libical_cmd_src := exec.Command("curl", "-L", libical_src_url, "-o", "source.tar.gz")
	err = libical_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libical_cmd_direct := exec.Command("./binary")
	err = libical_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
}
