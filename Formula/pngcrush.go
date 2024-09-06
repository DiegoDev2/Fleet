package main

// Pngcrush - Optimizer for PNG files
// Homepage: https://pmt.sourceforge.io/pngcrush/

import (
	"fmt"
	
	"os/exec"
)

func installPngcrush() {
	// Método 1: Descargar y extraer .tar.gz
	pngcrush_tar_url := "https://downloads.sourceforge.net/project/pmt/pngcrush/1.8.13/pngcrush-1.8.13-nolib.tar.xz"
	pngcrush_cmd_tar := exec.Command("curl", "-L", pngcrush_tar_url, "-o", "package.tar.gz")
	err := pngcrush_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pngcrush_zip_url := "https://downloads.sourceforge.net/project/pmt/pngcrush/1.8.13/pngcrush-1.8.13-nolib.tar.xz"
	pngcrush_cmd_zip := exec.Command("curl", "-L", pngcrush_zip_url, "-o", "package.zip")
	err = pngcrush_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pngcrush_bin_url := "https://downloads.sourceforge.net/project/pmt/pngcrush/1.8.13/pngcrush-1.8.13-nolib.tar.xz"
	pngcrush_cmd_bin := exec.Command("curl", "-L", pngcrush_bin_url, "-o", "binary.bin")
	err = pngcrush_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pngcrush_src_url := "https://downloads.sourceforge.net/project/pmt/pngcrush/1.8.13/pngcrush-1.8.13-nolib.tar.xz"
	pngcrush_cmd_src := exec.Command("curl", "-L", pngcrush_src_url, "-o", "source.tar.gz")
	err = pngcrush_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pngcrush_cmd_direct := exec.Command("./binary")
	err = pngcrush_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
