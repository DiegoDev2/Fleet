package main

// Pngxx - C++ wrapper for libpng library
// Homepage: https://www.nongnu.org/pngpp/

import (
	"fmt"
	
	"os/exec"
)

func installPngxx() {
	// Método 1: Descargar y extraer .tar.gz
	pngxx_tar_url := "https://download.savannah.gnu.org/releases/pngpp/png++-0.2.10.tar.gz"
	pngxx_cmd_tar := exec.Command("curl", "-L", pngxx_tar_url, "-o", "package.tar.gz")
	err := pngxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pngxx_zip_url := "https://download.savannah.gnu.org/releases/pngpp/png++-0.2.10.zip"
	pngxx_cmd_zip := exec.Command("curl", "-L", pngxx_zip_url, "-o", "package.zip")
	err = pngxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pngxx_bin_url := "https://download.savannah.gnu.org/releases/pngpp/png++-0.2.10.bin"
	pngxx_cmd_bin := exec.Command("curl", "-L", pngxx_bin_url, "-o", "binary.bin")
	err = pngxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pngxx_src_url := "https://download.savannah.gnu.org/releases/pngpp/png++-0.2.10.src.tar.gz"
	pngxx_cmd_src := exec.Command("curl", "-L", pngxx_src_url, "-o", "source.tar.gz")
	err = pngxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pngxx_cmd_direct := exec.Command("./binary")
	err = pngxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
