package main

// MypaintBrushes - Brushes used by MyPaint and other software using libmypaint
// Homepage: https://github.com/mypaint/mypaint-brushes

import (
	"fmt"
	
	"os/exec"
)

func installMypaintBrushes() {
	// Método 1: Descargar y extraer .tar.gz
	mypaintbrushes_tar_url := "https://github.com/mypaint/mypaint-brushes/archive/refs/tags/v2.0.2.tar.gz"
	mypaintbrushes_cmd_tar := exec.Command("curl", "-L", mypaintbrushes_tar_url, "-o", "package.tar.gz")
	err := mypaintbrushes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mypaintbrushes_zip_url := "https://github.com/mypaint/mypaint-brushes/archive/refs/tags/v2.0.2.zip"
	mypaintbrushes_cmd_zip := exec.Command("curl", "-L", mypaintbrushes_zip_url, "-o", "package.zip")
	err = mypaintbrushes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mypaintbrushes_bin_url := "https://github.com/mypaint/mypaint-brushes/archive/refs/tags/v2.0.2.bin"
	mypaintbrushes_cmd_bin := exec.Command("curl", "-L", mypaintbrushes_bin_url, "-o", "binary.bin")
	err = mypaintbrushes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mypaintbrushes_src_url := "https://github.com/mypaint/mypaint-brushes/archive/refs/tags/v2.0.2.src.tar.gz"
	mypaintbrushes_cmd_src := exec.Command("curl", "-L", mypaintbrushes_src_url, "-o", "source.tar.gz")
	err = mypaintbrushes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mypaintbrushes_cmd_direct := exec.Command("./binary")
	err = mypaintbrushes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libmypaint")
exec.Command("latte", "install", "libmypaint")
}
