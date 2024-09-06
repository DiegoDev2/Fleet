package main

// Isl - Integer Set Library for the polyhedral model
// Homepage: https://libisl.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installIsl() {
	// Método 1: Descargar y extraer .tar.gz
	isl_tar_url := "https://libisl.sourceforge.io/isl-0.27.tar.xz"
	isl_cmd_tar := exec.Command("curl", "-L", isl_tar_url, "-o", "package.tar.gz")
	err := isl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	isl_zip_url := "https://libisl.sourceforge.io/isl-0.27.tar.xz"
	isl_cmd_zip := exec.Command("curl", "-L", isl_zip_url, "-o", "package.zip")
	err = isl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	isl_bin_url := "https://libisl.sourceforge.io/isl-0.27.tar.xz"
	isl_cmd_bin := exec.Command("curl", "-L", isl_bin_url, "-o", "binary.bin")
	err = isl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	isl_src_url := "https://libisl.sourceforge.io/isl-0.27.tar.xz"
	isl_cmd_src := exec.Command("curl", "-L", isl_src_url, "-o", "source.tar.gz")
	err = isl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	isl_cmd_direct := exec.Command("./binary")
	err = isl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
