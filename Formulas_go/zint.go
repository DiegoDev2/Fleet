package main

// Zint - Barcode encoding library supporting over 50 symbologies
// Homepage: https://www.zint.org.uk/

import (
	"fmt"
	
	"os/exec"
)

func installZint() {
	// Método 1: Descargar y extraer .tar.gz
	zint_tar_url := "https://downloads.sourceforge.net/project/zint/zint/2.13.0/zint-2.13.0-src.tar.gz"
	zint_cmd_tar := exec.Command("curl", "-L", zint_tar_url, "-o", "package.tar.gz")
	err := zint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zint_zip_url := "https://downloads.sourceforge.net/project/zint/zint/2.13.0/zint-2.13.0-src.zip"
	zint_cmd_zip := exec.Command("curl", "-L", zint_zip_url, "-o", "package.zip")
	err = zint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zint_bin_url := "https://downloads.sourceforge.net/project/zint/zint/2.13.0/zint-2.13.0-src.bin"
	zint_cmd_bin := exec.Command("curl", "-L", zint_bin_url, "-o", "binary.bin")
	err = zint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zint_src_url := "https://downloads.sourceforge.net/project/zint/zint/2.13.0/zint-2.13.0-src.src.tar.gz"
	zint_cmd_src := exec.Command("curl", "-L", zint_src_url, "-o", "source.tar.gz")
	err = zint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zint_cmd_direct := exec.Command("./binary")
	err = zint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
