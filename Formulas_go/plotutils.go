package main

// Plotutils - C/C++ function library for exporting 2-D vector graphics
// Homepage: https://www.gnu.org/software/plotutils/

import (
	"fmt"
	
	"os/exec"
)

func installPlotutils() {
	// Método 1: Descargar y extraer .tar.gz
	plotutils_tar_url := "https://ftp.gnu.org/gnu/plotutils/plotutils-2.6.tar.gz"
	plotutils_cmd_tar := exec.Command("curl", "-L", plotutils_tar_url, "-o", "package.tar.gz")
	err := plotutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	plotutils_zip_url := "https://ftp.gnu.org/gnu/plotutils/plotutils-2.6.zip"
	plotutils_cmd_zip := exec.Command("curl", "-L", plotutils_zip_url, "-o", "package.zip")
	err = plotutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	plotutils_bin_url := "https://ftp.gnu.org/gnu/plotutils/plotutils-2.6.bin"
	plotutils_cmd_bin := exec.Command("curl", "-L", plotutils_bin_url, "-o", "binary.bin")
	err = plotutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	plotutils_src_url := "https://ftp.gnu.org/gnu/plotutils/plotutils-2.6.src.tar.gz"
	plotutils_cmd_src := exec.Command("curl", "-L", plotutils_src_url, "-o", "source.tar.gz")
	err = plotutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	plotutils_cmd_direct := exec.Command("./binary")
	err = plotutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libxaw")
exec.Command("latte", "install", "libxaw")
}
