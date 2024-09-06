package main

// UtilMacros - X.Org: Set of autoconf macros used to build other xorg packages
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installUtilMacros() {
	// Método 1: Descargar y extraer .tar.gz
	utilmacros_tar_url := "https://www.x.org/archive/individual/util/util-macros-1.20.1.tar.xz"
	utilmacros_cmd_tar := exec.Command("curl", "-L", utilmacros_tar_url, "-o", "package.tar.gz")
	err := utilmacros_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	utilmacros_zip_url := "https://www.x.org/archive/individual/util/util-macros-1.20.1.tar.xz"
	utilmacros_cmd_zip := exec.Command("curl", "-L", utilmacros_zip_url, "-o", "package.zip")
	err = utilmacros_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	utilmacros_bin_url := "https://www.x.org/archive/individual/util/util-macros-1.20.1.tar.xz"
	utilmacros_cmd_bin := exec.Command("curl", "-L", utilmacros_bin_url, "-o", "binary.bin")
	err = utilmacros_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	utilmacros_src_url := "https://www.x.org/archive/individual/util/util-macros-1.20.1.tar.xz"
	utilmacros_cmd_src := exec.Command("curl", "-L", utilmacros_src_url, "-o", "source.tar.gz")
	err = utilmacros_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	utilmacros_cmd_direct := exec.Command("./binary")
	err = utilmacros_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
