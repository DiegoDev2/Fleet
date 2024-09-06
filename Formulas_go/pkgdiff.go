package main

// Pkgdiff - Tool for analyzing changes in software packages (e.g. RPM, DEB, TAR.GZ)
// Homepage: https://lvc.github.io/pkgdiff/

import (
	"fmt"
	
	"os/exec"
)

func installPkgdiff() {
	// Método 1: Descargar y extraer .tar.gz
	pkgdiff_tar_url := "https://github.com/lvc/pkgdiff/archive/refs/tags/1.7.2.tar.gz"
	pkgdiff_cmd_tar := exec.Command("curl", "-L", pkgdiff_tar_url, "-o", "package.tar.gz")
	err := pkgdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pkgdiff_zip_url := "https://github.com/lvc/pkgdiff/archive/refs/tags/1.7.2.zip"
	pkgdiff_cmd_zip := exec.Command("curl", "-L", pkgdiff_zip_url, "-o", "package.zip")
	err = pkgdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pkgdiff_bin_url := "https://github.com/lvc/pkgdiff/archive/refs/tags/1.7.2.bin"
	pkgdiff_cmd_bin := exec.Command("curl", "-L", pkgdiff_bin_url, "-o", "binary.bin")
	err = pkgdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pkgdiff_src_url := "https://github.com/lvc/pkgdiff/archive/refs/tags/1.7.2.src.tar.gz"
	pkgdiff_cmd_src := exec.Command("curl", "-L", pkgdiff_src_url, "-o", "source.tar.gz")
	err = pkgdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pkgdiff_cmd_direct := exec.Command("./binary")
	err = pkgdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: binutils")
exec.Command("latte", "install", "binutils")
	fmt.Println("Instalando dependencia: gawk")
exec.Command("latte", "install", "gawk")
	fmt.Println("Instalando dependencia: wdiff")
exec.Command("latte", "install", "wdiff")
}
