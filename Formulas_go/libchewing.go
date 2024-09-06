package main

// Libchewing - Intelligent phonetic input method library
// Homepage: https://chewing.im/

import (
	"fmt"
	
	"os/exec"
)

func installLibchewing() {
	// Método 1: Descargar y extraer .tar.gz
	libchewing_tar_url := "https://github.com/chewing/libchewing/releases/download/v0.9.0/libchewing-0.9.0.tar.zst"
	libchewing_cmd_tar := exec.Command("curl", "-L", libchewing_tar_url, "-o", "package.tar.gz")
	err := libchewing_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libchewing_zip_url := "https://github.com/chewing/libchewing/releases/download/v0.9.0/libchewing-0.9.0.tar.zst"
	libchewing_cmd_zip := exec.Command("curl", "-L", libchewing_zip_url, "-o", "package.zip")
	err = libchewing_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libchewing_bin_url := "https://github.com/chewing/libchewing/releases/download/v0.9.0/libchewing-0.9.0.tar.zst"
	libchewing_cmd_bin := exec.Command("curl", "-L", libchewing_bin_url, "-o", "binary.bin")
	err = libchewing_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libchewing_src_url := "https://github.com/chewing/libchewing/releases/download/v0.9.0/libchewing-0.9.0.tar.zst"
	libchewing_cmd_src := exec.Command("curl", "-L", libchewing_src_url, "-o", "source.tar.gz")
	err = libchewing_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libchewing_cmd_direct := exec.Command("./binary")
	err = libchewing_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: corrosion")
exec.Command("latte", "install", "corrosion")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
