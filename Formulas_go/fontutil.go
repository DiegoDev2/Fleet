package main

// FontUtil - X.Org: Font package creation/installation utilities
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installFontUtil() {
	// Método 1: Descargar y extraer .tar.gz
	fontutil_tar_url := "https://www.x.org/archive/individual/font/font-util-1.4.1.tar.xz"
	fontutil_cmd_tar := exec.Command("curl", "-L", fontutil_tar_url, "-o", "package.tar.gz")
	err := fontutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fontutil_zip_url := "https://www.x.org/archive/individual/font/font-util-1.4.1.tar.xz"
	fontutil_cmd_zip := exec.Command("curl", "-L", fontutil_zip_url, "-o", "package.zip")
	err = fontutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fontutil_bin_url := "https://www.x.org/archive/individual/font/font-util-1.4.1.tar.xz"
	fontutil_cmd_bin := exec.Command("curl", "-L", fontutil_bin_url, "-o", "binary.bin")
	err = fontutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fontutil_src_url := "https://www.x.org/archive/individual/font/font-util-1.4.1.tar.xz"
	fontutil_cmd_src := exec.Command("curl", "-L", fontutil_src_url, "-o", "source.tar.gz")
	err = fontutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fontutil_cmd_direct := exec.Command("./binary")
	err = fontutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
}
