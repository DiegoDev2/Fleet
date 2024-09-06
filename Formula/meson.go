package main

// Meson - Fast and user friendly build system
// Homepage: https://mesonbuild.com/

import (
	"fmt"
	
	"os/exec"
)

func installMeson() {
	// Método 1: Descargar y extraer .tar.gz
	meson_tar_url := "https://github.com/mesonbuild/meson/releases/download/1.5.1/meson-1.5.1.tar.gz"
	meson_cmd_tar := exec.Command("curl", "-L", meson_tar_url, "-o", "package.tar.gz")
	err := meson_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	meson_zip_url := "https://github.com/mesonbuild/meson/releases/download/1.5.1/meson-1.5.1.zip"
	meson_cmd_zip := exec.Command("curl", "-L", meson_zip_url, "-o", "package.zip")
	err = meson_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	meson_bin_url := "https://github.com/mesonbuild/meson/releases/download/1.5.1/meson-1.5.1.bin"
	meson_cmd_bin := exec.Command("curl", "-L", meson_bin_url, "-o", "binary.bin")
	err = meson_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	meson_src_url := "https://github.com/mesonbuild/meson/releases/download/1.5.1/meson-1.5.1.src.tar.gz"
	meson_cmd_src := exec.Command("curl", "-L", meson_src_url, "-o", "source.tar.gz")
	err = meson_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	meson_cmd_direct := exec.Command("./binary")
	err = meson_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
