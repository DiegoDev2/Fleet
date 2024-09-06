package main

// Lzlib - Data compression library
// Homepage: https://www.nongnu.org/lzip/lzlib.html

import (
	"fmt"
	
	"os/exec"
)

func installLzlib() {
	// Método 1: Descargar y extraer .tar.gz
	lzlib_tar_url := "https://download.savannah.gnu.org/releases/lzip/lzlib/lzlib-1.14.tar.gz"
	lzlib_cmd_tar := exec.Command("curl", "-L", lzlib_tar_url, "-o", "package.tar.gz")
	err := lzlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lzlib_zip_url := "https://download.savannah.gnu.org/releases/lzip/lzlib/lzlib-1.14.zip"
	lzlib_cmd_zip := exec.Command("curl", "-L", lzlib_zip_url, "-o", "package.zip")
	err = lzlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lzlib_bin_url := "https://download.savannah.gnu.org/releases/lzip/lzlib/lzlib-1.14.bin"
	lzlib_cmd_bin := exec.Command("curl", "-L", lzlib_bin_url, "-o", "binary.bin")
	err = lzlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lzlib_src_url := "https://download.savannah.gnu.org/releases/lzip/lzlib/lzlib-1.14.src.tar.gz"
	lzlib_cmd_src := exec.Command("curl", "-L", lzlib_src_url, "-o", "source.tar.gz")
	err = lzlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lzlib_cmd_direct := exec.Command("./binary")
	err = lzlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
