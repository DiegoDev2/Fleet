package main

// Cddlib - Double description method for general polyhedral cones
// Homepage: https://www.inf.ethz.ch/personal/fukudak/cdd_home/

import (
	"fmt"
	
	"os/exec"
)

func installCddlib() {
	// Método 1: Descargar y extraer .tar.gz
	cddlib_tar_url := "https://github.com/cddlib/cddlib/releases/download/0.94m/cddlib-0.94m.tar.gz"
	cddlib_cmd_tar := exec.Command("curl", "-L", cddlib_tar_url, "-o", "package.tar.gz")
	err := cddlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cddlib_zip_url := "https://github.com/cddlib/cddlib/releases/download/0.94m/cddlib-0.94m.zip"
	cddlib_cmd_zip := exec.Command("curl", "-L", cddlib_zip_url, "-o", "package.zip")
	err = cddlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cddlib_bin_url := "https://github.com/cddlib/cddlib/releases/download/0.94m/cddlib-0.94m.bin"
	cddlib_cmd_bin := exec.Command("curl", "-L", cddlib_bin_url, "-o", "binary.bin")
	err = cddlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cddlib_src_url := "https://github.com/cddlib/cddlib/releases/download/0.94m/cddlib-0.94m.src.tar.gz"
	cddlib_cmd_src := exec.Command("curl", "-L", cddlib_src_url, "-o", "source.tar.gz")
	err = cddlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cddlib_cmd_direct := exec.Command("./binary")
	err = cddlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
}
