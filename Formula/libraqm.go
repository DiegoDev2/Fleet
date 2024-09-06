package main

// Libraqm - Library for complex text layout
// Homepage: https://github.com/HOST-Oman/libraqm

import (
	"fmt"
	
	"os/exec"
)

func installLibraqm() {
	// Método 1: Descargar y extraer .tar.gz
	libraqm_tar_url := "https://github.com/HOST-Oman/libraqm/archive/refs/tags/v0.10.1.tar.gz"
	libraqm_cmd_tar := exec.Command("curl", "-L", libraqm_tar_url, "-o", "package.tar.gz")
	err := libraqm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libraqm_zip_url := "https://github.com/HOST-Oman/libraqm/archive/refs/tags/v0.10.1.zip"
	libraqm_cmd_zip := exec.Command("curl", "-L", libraqm_zip_url, "-o", "package.zip")
	err = libraqm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libraqm_bin_url := "https://github.com/HOST-Oman/libraqm/archive/refs/tags/v0.10.1.bin"
	libraqm_cmd_bin := exec.Command("curl", "-L", libraqm_bin_url, "-o", "binary.bin")
	err = libraqm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libraqm_src_url := "https://github.com/HOST-Oman/libraqm/archive/refs/tags/v0.10.1.src.tar.gz"
	libraqm_cmd_src := exec.Command("curl", "-L", libraqm_src_url, "-o", "source.tar.gz")
	err = libraqm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libraqm_cmd_direct := exec.Command("./binary")
	err = libraqm_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: fribidi")
	exec.Command("latte", "install", "fribidi").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
}
